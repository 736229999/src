package main

import (
	"log"
	"net/http"
	"time"

	apiproto "github.com/caojunxyz/mimi-api/proto"
	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	ucproto "github.com/caojunxyz/mimi-server/apps/usercenter/proto"
	"github.com/caojunxyz/mimi-server/utils"

	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net"
	"strings"
)

type UnifiedOrderReq struct {
	Appid            string `xml:"appid"`
	Mch_id           string `xml:"mch_id"`
	Nonce_str        string `xml:"nonce_str"`
	Sign             string `xml:"sign"`
	Body             string `xml:"body"`
	Out_trade_no     string `xml:"out_trade_no"`
	Total_fee        int64  `xml:"total_fee"`
	Spbill_create_ip net.IP `xml:"spbill_create_ip"`
	Notify_url       string `xml:"notify_url"`
	Trade_type       string `xml:"trade_type"`
}

type UnifiedOrderResp struct {
	Return_code string `xml:"return_code"`
	Return_msg  string `xml:"return_msg"`
	Appid       string `xml:"appid"`
	Mch_id      string `xml:"mch_id"`
	Nonce_str   string `xml:"nonce_str"`
	Sign        string `xml:"sign"`
	Result_code string `xml:"result_code"`
	Prepay_id   string `xml:"prepay_id"`
	Trade_type  string `xml:"trade_type"`
}

type PayCallbackRequest struct {
	Return_code    string `xml:"return_code"`
	Return_msg     string `xml:"return_msg"`
	Appid          string `xml:"appid"`
	Mch_id         string `xml:"mch_id"`
	Nonce          string `xml:"nonce_str"`
	Sign           string `xml:"sign"`
	Result_code    string `xml:"result_code"`
	Openid         string `xml:"openid"`
	Is_subscribe   string `xml:"is_subscribe"`
	Trade_type     string `xml:"trade_type"`
	Bank_type      string `xml:"bank_type"`
	Total_fee      int64  `xml:"total_fee"`
	Fee_type       string `xml:"fee_type"`
	Cash_fee       int    `xml:"cash_fee"`
	Cash_fee_Type  string `xml:"cash_fee_type"`
	Transaction_id string `xml:"transaction_id"`
	Out_trade_no   string `xml:"out_trade_no"`
	Attach         string `xml:"attach"`
	Time_end       string `xml:"time_end"`
}

type PayCallbackResponse struct {
	Return_code string `xml:"return_code"`
	Return_msg  string `xml:"return_msg"`
}

//接受充值的金额.
//生成充值订单.
func (srv *RechargeServer) HandleWechatRechargeOrderCommit(w http.ResponseWriter, r *http.Request) {

	//获取请求的时间.
	now := time.Now()

	log.Println("收到请求")
	var msg apiproto.RechargeRequest

	account_id, ip, err := utils.ParseHttpRequest(w, r, &msg)

	log.Printf("%+v", msg)
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}

	osStr := strings.ToLower(msg.Os)

	if osStr != "ios" && osStr != "android" {
		log.Printf("数据传递错误，只能传递iOS %+v\n", msg)
		http.Error(w, "请求解析错误!", http.StatusForbidden)
		return
	}

	os := msg.Os

	//生成一个随机字符串.
	nonce_str := srv.randomString()

	if msg.Money == 0 {
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "充值金额不能为0", nil)
		return
	}

	if !*debug {

		//注意数据库存储的单位为分，分，分.
		if msg.Money%100 != 0 {
			utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "充值的金额必须为整数", nil)
			return
		}
	}

	//获取充值给哪个用户.
	//生成一个订单号码.
	log.Printf("开始创建订单%+v\n", &msg)
	order_no, err := srv.createOrderNo(w, r, account_id)
	log.Println("order:", order_no)
	if err != nil {
		log.Printf("%+v\n", err)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "充值失败", nil)
		return
	}

	req := UnifiedOrderReq{}
	req.Appid = GetWechat(os).AppId
	req.Body = "会员充值"
	req.Mch_id = GetWechat(os).MchId
	req.Nonce_str = srv.randomString()
	req.Notify_url = GetWechat(os).GetNotifyUrl(*gw)
	req.Out_trade_no = order_no
	req.Spbill_create_ip = ip
	req.Total_fee = msg.Money
	req.Trade_type = "APP"

	//组装数据，计算签名 .
	m := make(map[string]interface{})
	m["appid"] = req.Appid
	m["body"] = req.Body
	m["mch_id"] = req.Mch_id
	m["nonce_str"] = req.Nonce_str
	m["notify_url"] = req.Notify_url
	m["out_trade_no"] = req.Out_trade_no
	m["spbill_create_ip"] = req.Spbill_create_ip
	m["total_fee"] = req.Total_fee
	m["trade_type"] = req.Trade_type
	log.Println("第一次签名sign")
	req.Sign = Sign(m, GetWechat(os).ApiKey)
	log.Println(req.Sign)

	bytes_req, err := xml.Marshal(req)
	if err != nil {
		log.Panicf("%+v\n", err)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "充值失败", nil)
		return
	}

	str_req := string(bytes_req)
	//wxpay的unifiedorder接口需要http body中xmldoc的根节点是<xml></xml>这种，所以这里需要replace一下
	str_req = strings.Replace(str_req, "UnifiedOrder", "xml", -1)
	bytes_req = []byte(str_req)

	//发送unified order请求.
	requ, err := http.NewRequest("POST", "https://api.mch.weixin.qq.com/pay/unifiedorder", bytes.NewReader(bytes_req))
	if err != nil {
		log.Printf("%+v\n", err)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "充值失败", nil)
		return
	}
	requ.Header.Set("Accept", "application/xml")
	//这里的http header的设置是必须设置的.
	requ.Header.Set("Content-Type", "application/xml;charset=utf-8")

	client := http.Client{}
	resp, _err := client.Do(requ)
	if _err != nil {
		log.Printf("%+v\n", err)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "充值失败", nil)
		return
	}

	xmlResp := UnifiedOrderResp{}

	defer resp.Body.Close()
	by, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("%+v\n", err)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "充值失败", nil)
		return
	}

	_err = xml.Unmarshal(by, &xmlResp)
	//处理return code.
	if xmlResp.Return_code == "FAIL" {
		log.Println("微信支付统一下单不成功", xmlResp.Return_msg)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "充值失败", nil)
		return
	}

	//这里微信支付统一下单成功以后在将充值订单信息记录到数据库.
	//获取充值给哪个用户.
	rechargeOrderArg := &dbproto.RechargeOrder{
		OrderNo:       order_no,
		AccountId:     account_id,
		Money:         msg.Money,
		Os:            msg.Os,
		ClientReqTime: now.Unix(),
		CreateTime:    now.Unix(),
		UcRespTime:    0,
		PaymentMethod: WECHAT_PAY,
	}

	//存入充值订单表.
	_, err = srv.DbAgentClient().RechargeCreateOrder(context.Background(), rechargeOrderArg)
	if err != nil {
		log.Printf("记录充值订单失败 %v\n", err)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "充值失败", nil)
		return
	}

	//这里已经得到微信支付的prepay id，需要返给客户端，由客户端继续完成支付流程
	response := &apiproto.RechargeResponse{}
	response.AppId = xmlResp.Appid
	response.PartnerId = GetWechat(os).MchId
	response.PrepayId = xmlResp.Prepay_id
	response.NonceStr = nonce_str
	response.Timestamp = now.Unix()
	response.Package = GetWechat(os).Package

	dist := make(map[string]interface{})
	dist["appid"] = response.AppId
	dist["partnerid"] = response.PartnerId
	dist["prepayid"] = xmlResp.Prepay_id
	dist["noncestr"] = response.NonceStr
	dist["timestamp"] = response.Timestamp
	dist["package"] = response.Package

	log.Println("第二次签名")
	response.Sign = Sign(dist, GetWechat(os).ApiKey)

	fmt.Println("response:", response)
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", response)
}

func (srv *RechargeServer) HandleWechatRechargeNotify(w http.ResponseWriter, r *http.Request) {

	log.Println("开始执行微信支付回调...")

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("读取http body失败，原因!", err)
		http.Error(w.(http.ResponseWriter), http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	log.Println("微信支付异步通知，HTTP Body:", string(body))
	var mr PayCallbackRequest
	err = xml.Unmarshal(body, &mr)
	if err != nil {
		log.Println("解析HTTP Body格式到xml失败，原因!", err)
		return
	}

	var reqMap map[string]interface{}
	reqMap = make(map[string]interface{})

	reqMap["appid"] = mr.Appid
	reqMap["mch_id"] = mr.Mch_id
	reqMap["nonce_str"] = mr.Nonce
	reqMap["return_code"] = mr.Return_code
	reqMap["return_msg"] = mr.Return_msg
	reqMap["result_code"] = mr.Result_code
	reqMap["openid"] = mr.Openid
	reqMap["is_subscribe"] = mr.Is_subscribe
	reqMap["trade_type"] = mr.Trade_type
	reqMap["bank_type"] = mr.Bank_type
	reqMap["total_fee"] = mr.Total_fee
	reqMap["fee_type"] = mr.Fee_type
	reqMap["cash_fee"] = mr.Cash_fee
	reqMap["cash_fee_type"] = mr.Cash_fee_Type
	reqMap["transaction_id"] = mr.Transaction_id
	reqMap["out_trade_no"] = mr.Out_trade_no
	reqMap["attach"] = mr.Attach
	reqMap["time_end"] = mr.Time_end

	log.Println("获取微信返回的订单号：", mr.Out_trade_no)
	data, err := srv.DbAgentClient().QueryRechargeInfoByOrderNo(context.Background(), &dbproto.StringValue{Value: mr.Out_trade_no})
	if err != nil {
		log.Println("根据订单号调用查询数据的方法失败:", err)
		return
	}

	var resp PayCallbackResponse

	//判断数据是否为空.
	if data == nil {

		resp.Return_code = "SUCCESS"
		resp.Return_msg = "OK"
		bytes, _err := xml.Marshal(resp)
		strResp := strings.Replace(string(bytes), "PayCallbackResponse", "xml", -1)
		if _err != nil {
			log.Println("xml编码失败：")
			log.Printf("%+v", _err)
			return
		}
		fmt.Fprintf(w, strResp)
		return
	}

	//todo 暂时模拟成ios
	api_key := GetWechat(data.Os).ApiKey

	//进行签名校验
	if mr.Sign == Sign(reqMap, api_key) {
		//根据返回来数据中的订单号和订单金额.
		//来判断是否一致.
		//判断订单是否未支付.clea
		if data.Status == WAIT_PAY {

			//判断数据库中存储的金额是否与接口传递过来的金额一致.
			//如果金额一致则修改数据库中的订单状态为已支付支付.
			//修改失败则返回给支付接口提示失败.
			if data.Money == mr.Total_fee {
				//修改订单状态为支付.
				_, err = srv.DbAgentClient().SetRechargeSuccess(context.Background(), &dbproto.StringValue{Value: mr.Out_trade_no})
				if err == nil {

					//回调请求地址.
					uc := ucproto.RechargeResult{
						AccountId: data.AccountId,
						Money:     mr.Total_fee,
						OrderNo:   data.OrderNo,
						Method:    "微信充值",
					}

					log.Printf("%+v", uc)
					_, err = srv.UcAgentClient().NotifyRecharged(context.Background(), &uc)
					if err != nil {
						log.Println("充值添加金额失败：", err)
						log.Println("参数：", uc)

						resp.Return_code = "FAIL"
						resp.Return_msg = "failed to verify sign, please retry!"

						bytes, _err := xml.Marshal(resp)
						strResp := strings.Replace(string(bytes), "PayCallbackResponse", "xml", -1)
						if _err != nil {
							log.Println("xml编码失败：")
							log.Printf("%+v", _err)
							return
						}
						fmt.Fprintf(w, strResp)

					} else {
						//修改uc_resp_time时间.
						arg := dbproto.RechargeOrder{
							OrderNo:    data.OrderNo,
							UcRespTime: time.Now().Unix(),
						}

						srv.DbAgentClient().RechargeSetUcRespTime(context.Background(), &arg)

						resp.Return_code = "SUCCESS"
						resp.Return_msg = "OK"

						bytes, _err := xml.Marshal(resp)
						strResp := strings.Replace(string(bytes), "PayCallbackResponse", "xml", -1)
						if _err != nil {
							log.Println("xml编码失败：")
							log.Printf("%+v", _err)
							return
						}
						fmt.Fprintf(w, strResp)
					}

				} else {
					resp.Return_code = "FAIL"
					resp.Return_msg = "failed to verify sign, please retry!"
					bytes, _err := xml.Marshal(resp)
					strResp := strings.Replace(string(bytes), "PayCallbackResponse", "xml", -1)
					if _err != nil {
						log.Println("xml编码失败：")
						log.Printf("%+v", _err)
						return
					}

					fmt.Fprintf(w, strResp)
				}
			} else {
				resp.Return_code = "FAIL"
				resp.Return_msg = "failed to verify sign, please retry!"
				bytes, _err := xml.Marshal(resp)
				strResp := strings.Replace(string(bytes), "PayCallbackResponse", "xml", -1)
				if _err != nil {
					log.Println("xml编码失败：")
					log.Printf("%+v", _err)
					return
				}
				fmt.Fprintf(w, strResp)
			}
		} else {
			resp.Return_code = "SUCCESS"
			resp.Return_msg = "OK"
			bytes, _err := xml.Marshal(resp)
			strResp := strings.Replace(string(bytes), "PayCallbackResponse", "xml", -1)
			if _err != nil {
				log.Println("xml编码失败：")
				log.Printf("%+v", _err)
				return
			}
			fmt.Fprintf(w, strResp)
		}

	} else {
		resp.Return_code = "FAIL"
		resp.Return_msg = "failed to verify sign, please retry!"
		bytes, _err := xml.Marshal(resp)
		strResp := strings.Replace(string(bytes), "PayCallbackResponse", "xml", -1)
		if _err != nil {
			log.Println("xml编码失败：")
			log.Printf("%+v", _err)
			return
		}
		fmt.Fprintf(w, strResp)
	}

}
