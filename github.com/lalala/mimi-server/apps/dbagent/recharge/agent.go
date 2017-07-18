package recharge

import (
	"fmt"
	"log"

	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"golang.org/x/net/context"
)

//定义几个充值订单的状态.
const (
	WAIT_PAY    = 0 //待支付.
	PAY_SUCCESS = 1 //已支付.
	PAY_CLOSE   = 2 //支付取消.

	WECHAT_PAY = 0 //微信支付.
	ALIPAY_PAY = 1 //支付宝支付.
)

//创建充值订单.
func (srv *DbRechargeAgent) RechargeCreateOrder(ctx context.Context, arg *dbproto.RechargeOrder) (*dbproto.Nil, error) {

	fmt.Println("进到了db")
	log.Println(arg)
	//获取数据.
	money := arg.GetMoney()
	order_no := arg.GetOrderNo()
	account_id := arg.GetAccountId()
	os := arg.GetOs()
	status := WAIT_PAY
	client_req_time := arg.GetClientReqTime()
	create_time := arg.GetCreateTime()
	uc_resp_time := arg.GetUcRespTime()
	payment_method := arg.GetPaymentMethod()

	//预处理sql.
	st, err := srv.dbConn.Prepare("INSERT INTO recharge_order(order_no, account_id, money, status, os, client_req_time, create_time, uc_resp_time, payment_method) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)")
	if err != nil {
		log.Println("预处理sql报错:", err)
		return nil, err
	}
	_, err = st.Exec(order_no, account_id, money, status, os, client_req_time, create_time, uc_resp_time, payment_method)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &dbproto.Nil{}, nil
}

//根据充值订单号来查询充值订单的数据.
func (srv *DbRechargeAgent) QueryRechargeInfoByOrderNo(tx context.Context, arg *dbproto.StringValue) (*dbproto.RechargeOrder, error) {

	//获取订单号.
	order_no := arg.GetValue()
	st, err := srv.dbConn.Prepare("SELECT order_no,account_id, money,status,create_time,os FROM recharge_order WHERE order_no = $1")
	if err != nil {
		log.Println("预处理sql失败:", err)
		return nil, err
	}
	data := &dbproto.RechargeOrder{}
	err = st.QueryRow(order_no).Scan(&data.OrderNo, &data.AccountId, &data.Money, &data.Status, &data.CreateTime, &data.Os)

	//判断是否查询到数据.
	if data.OrderNo != "" {

		//判断是否scan获取数据错误.
		if err != nil {
			log.Println("scan获取数据失败：", err)
			log.Println("order_no:", order_no)
			return nil, err
		}
	}

	return data, nil
}

//根据充值订单号来修改充值订单状态.
func (srv *DbRechargeAgent) SetRechargeSuccess(ctx context.Context, arg *dbproto.StringValue) (*dbproto.Nil, error) {

	//获取订单号.
	order_no := arg.GetValue()

	st, err := srv.dbConn.Prepare("UPDATE recharge_order SET status = $1 WHERE order_no = $2")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	_, err = st.Exec(PAY_SUCCESS, order_no)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &dbproto.Nil{}, nil
}

//查询充值没有回调的订单.
//func (srv *DbRechargeAgent) QueryRechargeNotCallbackOrder (arg *dbproto.StringValue, stream dbproto.DbAgent_QueryRechargeNotCallbackOrderServer) error {
//
//	st ,err := srv.dbConn.Prepare("SELECT order_no FROM recharge_order WHERE status = $1 AND callback_success_time = 0")
//	if err != nil {
//		log.Println("充值,查询未回调给账户增加余额的方法，预处理sql失败:", err)
//	}
//	rows, err := st.Query(PAY_SUCCESS)
//
//	if err != nil {
//		log.Println("充值，查询没有回调的充值订单，查询失败：", err)
//	}
//
//	for rows.Next() {
//		var order string
//		err = rows.Scan(&order)
//		if err != nil {
//			log.Println("取出数据失败：", err)
//			break
//		}
//		orderInfo := &dbproto.CallbackOrderInfo{
//			OrderNo:order,
//		}
//		stream.Send(orderInfo)
//	}
//
//	return nil
//}

//充值设置用户中心反馈的时间.
func (srv *DbRechargeAgent) RechargeSetUcRespTime(ctx context.Context, arg *dbproto.RechargeOrder) (*dbproto.Nil, error) {

	uc_resp_time := arg.GetUcRespTime()
	order_no := arg.GetOrderNo()

	st, err := srv.dbConn.Prepare("UPDATE recharge_order SET uc_resp_time = $1 WHERE order_no = $2")
	if err != nil {
		log.Printf("充值设置用户中心反馈的时间预处理sql失败：", err)
	}
	_, err = st.Exec(uc_resp_time, order_no)
	if err != nil {
		log.Printf("充值设置用户中心反馈的时间，执行sql失败：", err)
	}
	return &dbproto.Nil{}, nil
}

//根据用户id查询订单数量.
func (srv *DbRechargeAgent) QueryOrderQuantityById(ctx context.Context, arg *dbproto.IntValue) (*dbproto.IntValue, error) {

	accountId := arg.GetValue()
	log.Println("accountId:", accountId)

	st, err := srv.dbConn.Prepare("SELECT COUNT(*) AS num FROM recharge_order WHERE account_id = $1")
	if err != nil {
		log.Printf("%+v\n", err)
	}
	var num int64 = 0

	err = st.QueryRow(accountId).Scan(&num)

	log.Println("这个用户的总订单数：", num)
	if err != nil {
		log.Printf("%+v", err)
	}

	return &dbproto.IntValue{Value: num}, err
}
