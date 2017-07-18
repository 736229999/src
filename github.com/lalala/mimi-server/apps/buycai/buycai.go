package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	apiproto "github.com/caojunxyz/mimi-api/proto"
	"github.com/caojunxyz/mimi-server/lottery"
	"github.com/caojunxyz/mimi-server/utils"
)

// 在售期号列表
func (srv *BuycaiServer) HandleBuycaiInfo(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleBuycaiInfo")
	id, err := utils.ParseLotteryIdArg(r)
	if err != nil {
		log.Println(err)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "无效彩种", nil)
	}
	agt := srv.getAgent(id)
	if agt == nil {
		log.Println("id:", id)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "无效彩种", nil)
		return
	}
	result := &apiproto.BuycaiInfo{
		MaxMultiple: lottery.GetConfig(id).MaxMultiple,
	}
	if v := agt.GetLastIssue(); v != nil {
		result.Last = &apiproto.SaleIssue{Issue: v.Issue, StartTime: v.StartTime, EndTime: v.EndTime, OpenTime: v.OpenTime}
	}
	if v := agt.GetCurIssue(); v != nil {
		result.Current = &apiproto.SaleIssue{Issue: v.Issue, StartTime: v.StartTime, EndTime: v.EndTime, OpenTime: v.OpenTime}
	}
	if v := agt.GetNextIssue(); v != nil {
		result.Next = &apiproto.SaleIssue{Issue: v.Issue, StartTime: v.StartTime, EndTime: v.EndTime, OpenTime: v.OpenTime}
	}
	saleList := agt.GetOnSaleList()
	for _, v := range saleList {
		result.SaleList = append(result.SaleList, v.Issue)
	}
	// data, _ := json.Marshal(result)
	// log.Printf("%v --> %v\n", id, string(data))
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", result)
}

func (srv *BuycaiServer) HandleCommitOrder(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleCommitOrder")
	var order apiproto.BuycaiOrder
	accountId, _, err := utils.ParseHttpRequest(w, r, &order)
	if err != nil {
		return
	}

	id := order.GetLotteryId()
	data, _ := json.Marshal(order)
	log.Printf("%d --> (%v) %s\n", accountId, id, string(data))
	agt := srv.getAgent(id)
	if agt == nil {
		log.Println("id:", id)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "无效彩种", nil)
		return
	}

	userOrder, err := agt.AddOrder(accountId, &order)
	if err != nil {
		log.Println("id:", id, err)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, err.Error(), nil)
		return
	}

	cfg := lottery.GetConfig(id)
	orderTime := userOrder.GetOrderTime()
	result := &apiproto.BuycaiUserOrder{
		Id:          userOrder.GetId(),
		LotteryId:   id,
		SumMoney:    order.GetCai() + order.GetBalance(),
		UserOrderNo: lottery.MakeOrderNo(cfg, time.Unix(orderTime, 0).Format("20060102150405"), userOrder.GetId()),
	}
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", result)
}
