package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
)

//创建订单号.
func (srv *RechargeServer) createOrderNo(w http.ResponseWriter, r *http.Request, accountId int64) (string, error) {

	//获取当前这个用户的所有订单数量.
	num, err := srv.DbAgentClient().QueryOrderQuantityById(context.Background(), &dbproto.IntValue{Value: accountId})
	if err != nil {
		log.Printf("%+v\n", err)
		return "", err
	}
	log.Println("num:", num.GetValue())
	count := num.GetValue()
	//获取当前的日期和时间.
	//格式化的时候去掉 "-"，" ",":".
	datetime := time.Now().Format("20060102150405")
	log.Println("date:", datetime)
	//生成订单号.
	//订单号组成.
	//日期时间+用户id+充值订单数.
	orderNo := datetime + fmt.Sprintf("%d", accountId) + fmt.Sprintf("%d", count)

	return orderNo, nil
}

//生成一个随机数.
func (srv *RechargeServer) randomString() string {

	t := time.Now().UnixNano()
	h := md5.New()
	h.Write([]byte(string(t)))
	b := h.Sum(nil)
	str := hex.EncodeToString(b)
	return str
}
