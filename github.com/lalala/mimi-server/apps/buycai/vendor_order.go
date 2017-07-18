package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/caojunxyz/gotu"
	apiproto "github.com/caojunxyz/mimi-api/proto"
	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/lottery"
	"github.com/caojunxyz/mimi-server/utils"
)

func (srv *BuycaiServer) HandleVendorOrderHistory(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleVendorOrderHistory")
	var msg apiproto.VendorOrderHistoryRequest
	accountId, _, err := utils.ParseHttpRequest(w, r, &msg)
	if err != nil {
		return
	}

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
	dbArg := &dbproto.QueryVendorOrderHistoryArg{
		AccountId:  accountId,
		LotteryId:  int32(msg.GetLotteryId()),
		StartTime:  startTime,
		EndTime:    now.Unix(),
		Limit:      pageSize,
		Offset:     offset,
		StatusList: convertToServerVendorOrderStatusList(msg.GetStatus()),
	}

	stream, err := srv.dbUc.QueryVendorOrderHistory(context.Background(), dbArg)
	if err != nil {
		http.Error(w, "db：QueryVendorOrderHistory!", http.StatusInternalServerError)
		return
	}

	result := &apiproto.VendorOrderHistory{PageSize: pageSize}
	for {
		v, err := stream.Recv()
		if err == nil {
			log.Printf("%+v\n", v)
			record := &apiproto.VendorOrderRecord{
				Id:        v.GetId(),
				LotteryId: apiproto.LotteryId(v.GetLotteryId()),
				Status:    convertToClientVendorOrderStatus(v.GetStatus()),
				Issue:     v.GetIssue(),
				Money:     v.GetMoney(),
				IsChase:   v.GetIssueNum() > 1,
				WinMoney:  v.GetWinMoney(),
				AddTime:   v.GetAddTime(),
			}
			// log.Printf("%+v\n", record)
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

func (srv *BuycaiServer) HandleVendorOrderDetail(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleVendorOrderDetail")
	var msg apiproto.IntValue
	_, _, err := utils.ParseHttpRequest(w, r, &msg)
	if err != nil {
		return
	}

	vendorOrder, err := srv.dbUc.QueryBuycaiVendorOrderById(context.Background(), &dbproto.IntValue{Value: msg.GetValue()})
	if err != nil {
		http.Error(w, "db：QueryBuycaiVendorOrderById!", http.StatusInternalServerError)
		return
	}

	lotteryId := apiproto.LotteryId(vendorOrder.GetLotteryId())
	cfg := lottery.GetConfig(lotteryId)
	agt := srv.getAgent(lotteryId)
	issue := vendorOrder.GetIssue()
	saleIssue := agt.GetSaleIssue(issue)
	if saleIssue == nil {
		log.Println(issue, lotteryId)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "无效期号", nil)
		return
	}

	orderTime := vendorOrder.GetAddTime()
	result := &apiproto.BuycaiVendorOrder{
		Id:          vendorOrder.GetId(),
		LotteryId:   lotteryId,
		Status:      convertToClientVendorOrderStatus(dbproto.VendorOrderStatus(vendorOrder.GetStatus())),
		OpenTime:    saleIssue.OpenTime,
		Multiple:    vendorOrder.GetMultiple(),
		Money:       vendorOrder.GetMoney(),
		SumNum:      vendorOrder.GetSumNum(),
		WinMoney:    vendorOrder.GetWinMoney(),
		SchemeList:  convertToClientSchemeList(vendorOrder.GetSchemeList()),
		UserOrderNo: lottery.MakeOrderNo(cfg, time.Unix(orderTime, 0).Format("20060102150405"), vendorOrder.GetUserOrderId()),
		UserOrderId: vendorOrder.GetUserOrderId(),
		OrderTime:   orderTime,
		ReqTime:     vendorOrder.GetVendorReqTime(),
		RespTime:    vendorOrder.GetVendorRespTime(),
		CurIssue:    vendorOrder.GetIssue(),
	}
	balls := strings.Replace(saleIssue.OpenBalls, "+", ",", -1)
	ballsList := strings.Split(balls, ",")
	if len(ballsList) > 0 && ballsList[0] != "" {
		log.Println(balls, ballsList)
		result.OpenBalls = ballsList
	}
	// log.Printf("%+v\n", result)
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", result)
}
