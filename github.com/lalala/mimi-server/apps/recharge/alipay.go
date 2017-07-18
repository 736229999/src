package main

import (
	"context"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/caojunxyz/mimi-api/proto"
	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	ucproto "github.com/caojunxyz/mimi-server/apps/usercenter/proto"
	"github.com/caojunxyz/mimi-server/utils"
)

//支付宝充值提交订单.
func (srv *RechargeServer) HandleAlipayRechargeCommitOrder(w http.ResponseWriter, r *http.Request) {

	var params apiproto.RechargeRequest

	account_id, _, err := utils.ParseHttpRequest(w, r, &params)

	if err != nil {
		log.Printf("%+v\n", err)
		return
	}

	osStr := strings.ToLower(params.Os)

	if osStr != "ios" && osStr != "android" {
		log.Printf("数据传递错误，只能传递iOS %+v\n", params)
		http.Error(w, "请求解析错误!", http.StatusForbidden)
		return
	}

	now := time.Now()
	order_no, err := srv.createOrderNo(w, r, account_id)
	if err != nil {
		log.Printf("%+v\n", err)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "充值失败", nil)
		return
	}
	log.Println("生成的订单号：", order_no)
	arg := &dbproto.RechargeOrder{
		OrderNo:       order_no,
		AccountId:     account_id,
		Money:         params.Money,
		Status:        WAIT_PAY,
		ClientReqTime: now.Unix(),
		CreateTime:    now.Unix(),
		UcRespTime:    0,
		PaymentMethod: ALIPAY_PAY,
	}

	_, err = srv.DbAgentClient().RechargeCreateOrder(context.Background(), arg)
	if err != nil {
		log.Printf("支付宝充值订单失败", err)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "充值失败", nil)
		return
	}

	aliConf := GetAlipay()
	alipayReply := &apiproto.AlipayRechargeReply{
		AppId:     aliConf.AppId,
		Method:    aliConf.Method,
		SignType:  aliConf.SignType,
		Charset:   aliConf.Charset,
		Version:   aliConf.Version,
		NotifyUrl: aliConf.GetNotifyUrl(*gw),
		Timestamp: now.Format("2006-01-02 15:04:05"),
	}

	money := float64(params.Money) / 100

	dist := make(map[string]interface{})
	dist["subject"] = "支付宝充值"
	dist["out_trade_no"] = order_no
	dist["total_amount"] = money
	dist["product_code"] = "QUICK_MSECURITY_PAY"
	dist["passback_params"] = alipayReply.Timestamp

	jsonByte, err := json.Marshal(dist)
	if err != nil {
		log.Println(err)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "充值失败", nil)
		return
	}

	alipayReply.BizContent = string(jsonByte)

	//生成签名.
	signMap := make(map[string]interface{})
	signMap["app_id"] = alipayReply.AppId
	signMap["method"] = alipayReply.Method
	signMap["sign_type"] = alipayReply.SignType
	signMap["charset"] = alipayReply.Charset
	signMap["version"] = alipayReply.Version
	signMap["notify_url"] = aliConf.GetNotifyUrl(*gw)
	signMap["timestamp"] = alipayReply.Timestamp
	signMap["biz_content"] = alipayReply.BizContent

	signStr := AlipaySignString(signMap)

	AppPrivatekey := aliConf.PrivateKey

	PrivateKey := "-----BEGIN RSA PRIVATE KEY-----\n" + AppPrivatekey + "\n-----END RSA PRIVATE KEY-----"

	block, _ := pem.Decode([]byte(PrivateKey))

	if block == nil {
		log.Println("RSA私钥错误")
		return
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Println("加密失败", err)
		return
	}

	sign, _ := RsaSign(signStr, privateKey)

	alipayReply.Sign = url.QueryEscape(sign)

	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", alipayReply)
}

func (srv *RechargeServer) HandleAlipayRechargeNotify(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)
	log.Println("回调的内容：", string(body))

	//调用url.ParseQuery来获取到参数列表，url.ParseQuery还会自动做url safe decode
	values_m, _err := url.ParseQuery(string(body))
	if _err != nil {
		log.Println(_err)
		return
	}

	var m map[string]interface{}
	m = make(map[string]interface{})

	for k, v := range values_m {
		if k == "sign" || k == "sign_type" { //不要'sign'和'sign_type'
			continue
		}
		m[k] = v[0]
	}

	sign := values_m["sign"][0]

	//待签名字符串.
	strPreSign := AlipaySignString(m)

	//加密.
	status, er := RSAVerify([]byte(strPreSign), []byte(sign))

	if er != nil {
		log.Println(er)
		return
	}

	//对比支付宝返回来的sign与返回来的参数进行加密的sign.
	if status {
		//.
		log.Println("签名相同")
		//根据传递过来的订单号，来修改支付订单的状态.
		order_no := m["out_trade_no"].(string)

		log.Println("order_no:", order_no)
		//根据订单号去查询订单状态.
		data, err := srv.DbAgentClient().QueryRechargeInfoByOrderNo(context.Background(), &dbproto.StringValue{Value: order_no})
		if err != nil {
			log.Println("根据订单号调用查询数据的方法失败:", err)
			return
		}
		log.Printf("%+v\n", data)

		//判断支付宝传递过来的订单是否是有效订单（自己平台生成的订单）.
		//判断当前这个订单是否待支付.
		if data.Status == WAIT_PAY {

			//判断data.Money 与 支付宝传递过来的金额是否相同
			//获取支付宝传递过来的total_amount.
			total_amount := m["total_amount"].(string)

			money_float, err := strconv.ParseFloat(total_amount, 64)
			log.Println(money_float)

			//数据库中存储的金额单位为分.
			//支付宝返回的金额单位为元，这里需要乘以100.
			money := money_float * 100

			if err != nil {
				log.Println(err)
				return
			}
			if data.Money == int64(money) {

				log.Println("金额相同")
				//判断支付宝传递过来的seller_id 和 支付宝本来的seller_id是否相同.
				//获取支付宝传递过来的Seller_id.
				seller_id := m["seller_id"].(string)
				if GetAlipay().SellerId == seller_id {
					log.Println("seller_id相同")
					//修改订单状态.
					_, err = srv.DbAgentClient().SetRechargeSuccess(context.Background(), &dbproto.StringValue{Value: order_no})
					if err != nil {
						log.Println("支付宝充值订单，状态修改失败", err)
						return
					}
					log.Println("支付宝充值订单，状态修改成功", err)

					//回调请求地址.
					uc := ucproto.RechargeResult{
						AccountId: data.AccountId,
						Money:     int64(money),
						OrderNo:   data.OrderNo,
						Method:    "支付宝充值",
					}

					//添加账户金额.
					_, err = srv.UcAgentClient().NotifyRecharged(context.Background(), &uc)

					if err != nil {

						log.Printf("充值添加金额失败：", err, "参数:", uc)

					} else {

						//修改uc_resp_time时间.
						arg := dbproto.RechargeOrder{
							OrderNo:    data.OrderNo,
							UcRespTime: time.Now().Unix(),
						}
						_, err = srv.DbAgentClient().RechargeSetUcRespTime(context.Background(), &arg)
						if err != nil {
							log.Println("设置充值反馈的时间失败:", err)

						}

						//通知支付宝，回调成功.
						w.Write([]byte("success"))
					}
				} else {
					log.Println("Seller_id不相同")
				}
			} else {
				log.Println("金额不相同", data.Money, money)
			}
		}

	} else {
		log.Println("签名不相同")
		fmt.Fprintf(w, "success")
	}
}
