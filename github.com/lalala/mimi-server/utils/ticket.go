package utils

import (
	"fmt"
	"time"

	apiproto "github.com/caojunxyz/mimi-api/proto"
	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/lottery"
)

func GetBuycaiTicketStatus(tkt *dbproto.BuycaiTicket) apiproto.TicketStatus {
	validStart := time.Unix(tkt.GetValidStart(), 0)
	validEnd := time.Unix(tkt.GetValidEnd(), 0)
	// 优先判断是否已使用
	if tkt.GetOrderId() > 0 {
		return apiproto.TicketStatus_Used
	}
	if time.Now().After(validEnd) {
		return apiproto.TicketStatus_Expired
	}
	if time.Now().Before(validStart) {
		return apiproto.TicketStatus_WaitValid
	}
	return apiproto.TicketStatus_WaitUse
}

func GetBuycaiTicketRestrictDesc(tkt *dbproto.BuycaiTicket) string {
	restrictType := tkt.GetRestrictType()
	restrictId := tkt.GetRestrictId()
	if restrictId > 0 {
		name := lottery.GetConfig(apiproto.LotteryId(restrictId)).Name
		return fmt.Sprintf("仅可用于购买%s", name)
	}
	if restrictId < 0 {
		name := lottery.GetConfig(apiproto.LotteryId(-restrictId)).Name
		return fmt.Sprintf("不可用于购买%s", name)
	}
	if restrictType > 0 {
		name := lottery.GetLotteryTypeName(apiproto.LotteryType(restrictType))
		return fmt.Sprintf("仅可用于购买%s", name)
	}
	if restrictType < 0 {
		name := lottery.GetLotteryTypeName(apiproto.LotteryType(-restrictType))
		return fmt.Sprintf("不可用于购买%s", name)
	}
	return "全场通用"
}

func IsBuycaiTicketCanUse(accountId int64, lotteryId int32, pay float64, tkt *dbproto.BuycaiTicket) bool {
	if tkt.GetOrderId() > 0 {
		return false
	}
	if tkt.GetAccountId() != accountId {
		return false
	}
	if tkt.GetId() == 0 {
		return false
	}

	cf := lottery.GetConfig(apiproto.LotteryId(lotteryId))
	if cf == nil {
		return false
	}
	restrictId := tkt.GetRestrictId()
	if restrictId > 0 && restrictId != int32(lotteryId) {
		return false
	}
	if restrictId < 0 && -restrictId == int32(lotteryId) {
		return false
	}
	restrictType := tkt.GetRestrictType()
	if restrictType > 0 && restrictType != int32(cf.Type) {
		return false
	}
	if restrictType < 0 && restrictType == int32(cf.Type) {
		return false
	}
	now := time.Now().Unix()
	if tkt.GetValidStart() > now || tkt.GetValidEnd() <= now {
		return false
	}
	if int32(pay) < tkt.GetUseBase() {
		return false
	}
	return true
}
