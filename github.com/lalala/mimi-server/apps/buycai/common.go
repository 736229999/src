package main

import (
	"log"

	apiproto "github.com/caojunxyz/mimi-api/proto"
	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
)

func convertToClientVendorOrderStatus(status dbproto.VendorOrderStatus) apiproto.VendorOrderStatus {
	switch status {
	case dbproto.VendorOrderStatus_VO_NotBet:
		return apiproto.VendorOrderStatus_VOS_WaitTicket
	case dbproto.VendorOrderStatus_VO_BetFail:
		return apiproto.VendorOrderStatus_VOS_TicketFail
	case dbproto.VendorOrderStatus_VO_BetSuccess:
		return apiproto.VendorOrderStatus_VOS_WaitOpen
	case dbproto.VendorOrderStatus_VO_Win:
		return apiproto.VendorOrderStatus_VOS_Win
	case dbproto.VendorOrderStatus_VO_NotWin:
		return apiproto.VendorOrderStatus_VOS_NotWin
	}
	return apiproto.VendorOrderStatus_VOS_All
}

func convertToServerVendorOrderStatusList(status apiproto.VendorOrderStatus) []int32 {
	result := []int32{}
	switch status {
	case apiproto.VendorOrderStatus_VOS_WaitTicket:
		result = append(result, int32(dbproto.VendorOrderStatus_VO_NotBet))
	case apiproto.VendorOrderStatus_VOS_TicketFail:
		result = append(result, int32(dbproto.VendorOrderStatus_VO_BetFail))
	case apiproto.VendorOrderStatus_VOS_WaitOpen:
		result = append(result, int32(dbproto.VendorOrderStatus_VO_BetSuccess))
	case apiproto.VendorOrderStatus_VOS_Win:
		result = append(result, int32(dbproto.VendorOrderStatus_VO_Win))
	case apiproto.VendorOrderStatus_VOS_NotWin:
		result = append(result, int32(dbproto.VendorOrderStatus_VO_NotWin))
	}
	return result
}

func convertToClientUserOrderStatus(status dbproto.UserOrderStatus) apiproto.UserOrderStatus {
	switch status {
	case dbproto.UserOrderStatus_UO_Doing:
		return apiproto.UserOrderStatus_UOS_Doing
	case dbproto.UserOrderStatus_UO_FinishStop:
		return apiproto.UserOrderStatus_UOS_Finish
	case dbproto.UserOrderStatus_UO_WinStop:
		return apiproto.UserOrderStatus_UOS_Stop
	case dbproto.UserOrderStatus_UO_FailStop:
		return apiproto.UserOrderStatus_UOS_Stop
	case dbproto.UserOrderStatus_UO_CancelStop:
		return apiproto.UserOrderStatus_UOS_Stop
	}
	log.Panic(status)
	return apiproto.UserOrderStatus_UOS_All
}

func convertToServerUserOrderStatusList(status apiproto.UserOrderStatus) []int32 {
	result := []int32{}
	switch status {
	case apiproto.UserOrderStatus_UOS_Doing:
		result = append(result, int32(dbproto.UserOrderStatus_UO_Doing))
	case apiproto.UserOrderStatus_UOS_Finish:
		result = append(result, int32(dbproto.UserOrderStatus_UO_FinishStop))
	case apiproto.UserOrderStatus_UOS_Stop:
		result = append(result, int32(dbproto.UserOrderStatus_UO_WinStop))
		result = append(result, int32(dbproto.UserOrderStatus_UO_FailStop))
	}
	return result
}

func convertToClientSchemeList(list []*dbproto.BuycaiScheme) []*apiproto.BuycaiScheme {
	result := []*apiproto.BuycaiScheme{}
	for _, v := range list {
		sch := &apiproto.BuycaiScheme{Type: v.GetType(), Num: v.GetNum(), Money: v.GetMoney()}
		selectBalls := v.GetSelectBalls()
		sch.SelectBalls = make(map[string]*apiproto.Balls)
		for k, ball := range selectBalls {
			ball0 := &apiproto.Balls{}
			for _, no := range ball.List {
				ball0.List = append(ball0.List, no)
			}
			sch.SelectBalls[k] = ball0
		}
		result = append(result, sch)
	}
	return result
}
