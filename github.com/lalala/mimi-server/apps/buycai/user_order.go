package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/caojunxyz/gotu"
	"github.com/caojunxyz/mimi-api/proto"
	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/lottery"
	"github.com/caojunxyz/mimi-server/utils"
)

func (srv *BuycaiServer) HandleUserOrderHistory(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleUserOrderHistory")
	var msg apiproto.UserOrderHistoryRequest
	accountId, _, err := utils.ParseHttpRequest(w, r, &msg)
	if err != nil {
		return
	}

	log.Printf("%d --> %+v\n", accountId, msg)
	pageSize := int32(20)
	offset := msg.GetPage() * pageSize
	now := time.Now()
	var startTime int64
	switch msg.GetRange() {
	case apiproto.TimeRange_TR_Today:
		startTime = gotu.BeginningOfDay(now).Unix()
	case apiproto.TimeRange_TR_Week:
		startTime = gotu.BeginningOfDay(now).AddDate(0, 0, -7).Unix()
	case apiproto.TimeRange_TR_Month:
		startTime = gotu.BeginningOfMonth(now).Unix()
	case apiproto.TimeRange_TR_ThreeMonth:
		startTime = gotu.BeginningOfMonth(now).AddDate(0, -2, 0).Unix()
	}
	dbArg := &dbproto.QueryUserOrderHistoryArg{
		AccountId:  accountId,
		LotteryId:  int32(msg.GetLotteryId()),
		StartTime:  startTime,
		EndTime:    now.Unix(),
		Limit:      pageSize,
		Offset:     offset,
		IsChase:    msg.GetType() == apiproto.BuycaiOrderType_Chase,
		StatusList: convertToServerUserOrderStatusList(msg.GetStatus()),
	}
	log.Println(msg.GetStatus(), dbArg.StatusList)

	stream, err := srv.dbUc.QueryUserOrderHistory(context.Background(), dbArg)
	if err != nil {
		http.Error(w, "db：QueryUserOrderHistory!", http.StatusInternalServerError)
		return
	}

	result := &apiproto.UserOrderHistory{PageSize: pageSize}
	for {
		v, err := stream.Recv()
		if err == nil {
			record := &apiproto.UserOrderRecord{
				Id:        v.GetId(),
				LotteryId: apiproto.LotteryId(v.GetLotteryId()),
				Status:    convertToClientUserOrderStatus(v.GetStatus()),
				Money:     v.GetMoney(),
				IssueNum:  v.GetIssueNum(),
				ChaseNo:   v.GetChaseNo(),
				WinMoney:  v.GetWinMoney(),
				AddTime:   v.GetAddTime(),
				CurIssue:  v.GetCurIssue(),
			}
			log.Printf("%+v\n", record)
			result.List = append(result.List, record)
			continue
		}
		if err != io.EOF {
			log.Println(err)
		}
		break
	}
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", result)
}

func (srv *BuycaiServer) HandleUserOrderDetail(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleUserOrderDetail")
	var msg apiproto.IntValue
	_, _, err := utils.ParseHttpRequest(w, r, &msg)
	if err != nil {
		return
	}

	userOrderId := msg.GetValue()
	userOrder, err := srv.dbUc.QueryBuycaiUserOrderById(context.Background(), &dbproto.IntValue{Value: userOrderId})
	if err != nil {
		http.Error(w, "db：QueryBuycaiUserOrderById!", http.StatusInternalServerError)
		return
	}

	lotteryId := apiproto.LotteryId(userOrder.GetLotteryId())
	cfg := lottery.GetConfig(lotteryId)
	orderTime := userOrder.GetOrderTime()
	schemeList := userOrder.GetSchemeList()
	result := &apiproto.BuycaiUserOrder{
		Id:             userOrder.GetId(),
		LotteryId:      lotteryId,
		SumMoney:       userOrder.GetSumMoney(),
		Status:         convertToClientUserOrderStatus(dbproto.UserOrderStatus(userOrder.GetStatus())),
		OrderTime:      userOrder.GetOrderTime(),
		TicketSubMoney: userOrder.GetTicketSubMoney(),
		TotalWinMoney:  userOrder.GetTotalWinMoney(),
		SumNum:         getSchemeListSumNum(schemeList),
		IssueNum:       userOrder.GetIssueNum(),
		ChaseNo:        userOrder.GetChaseNo(),
		SchemeList:     convertToClientSchemeList(schemeList),
		OrderType:      apiproto.BuycaiOrderType_Normal,
		UserOrderNo:    lottery.MakeOrderNo(cfg, time.Unix(orderTime, 0).Format("20060102150405"), userOrder.GetId()),
	}

	if result.IssueNum > 1 {
		result.OrderType = apiproto.BuycaiOrderType_Chase
	}
	allIssues, err := srv.dbUc.BuycaiQueryUserOrderAllIssues(context.Background(), &dbproto.IntValue{Value: userOrderId})
	if err != nil {
		http.Error(w, "db：QueryBuycaiUserOrderById!", http.StatusInternalServerError)
		return
	}
	for _, v := range allIssues.List {
		issueInfo := &apiproto.BuycaiIssueInfo{
			Issue:         v.GetIssue(),
			Multiple:      v.GetMultiple(),
			Money:         v.GetMoney(),
			WinMoney:      v.GetWinMoney(),
			Status:        convertToClientVendorOrderStatus(v.GetStatus()),
			ChaseNo:       v.GetChaseNo(),
			VendorOrderId: v.GetVendorOrderId(),
		}
		result.IssueList = append(result.IssueList, issueInfo)
	}
	result.CurIssue = result.IssueList[0].Issue
	// log.Printf("%+v\n", result)
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", result)
}

func (srv *BuycaiServer) HandleStopChase(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleStopChase")
	var msg apiproto.IntValue
	accountId, _, err := utils.ParseHttpRequest(w, r, &msg)
	if err != nil {
		return
	}

	userOrderId := msg.GetValue()
	_, err = srv.dbUc.BuycaiUserCancelStopChase(context.Background(), &dbproto.IntValue{Value: userOrderId})
	if err != nil {
		log.Println(accountId, err)
		http.Error(w, "db：BuycaiUserCancelStopChase!", http.StatusInternalServerError)
		return
	}
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", nil)
}
