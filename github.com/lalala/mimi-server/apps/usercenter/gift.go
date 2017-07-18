package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"encoding/json"

	apiproto "github.com/caojunxyz/mimi-api/proto"
	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/utils"
)

// TODO: 配置指定类型的礼包, 根据设备类型（小米/华为/苹果/Oppo等）选择不同礼包
var _giftConfig = [...][3]int32{
	[3]int32{300, 30, int32(apiproto.LotteryId_Cqssc)},
	[3]int32{200, 15, int32(apiproto.LotteryId_Gd11x5)},
	[3]int32{200, 15, int32(apiproto.LotteryId_Bjpk10)},
	[3]int32{30, 2, int32(apiproto.LotteryId_Dlt)},
	[3]int32{30, 2, int32(apiproto.LotteryId_Ssq)},
	[3]int32{30, 2, int32(apiproto.LotteryId_Fc3d)},
	[3]int32{30, 2, int32(apiproto.LotteryId_Pl3)},
	[3]int32{30, 2, int32(apiproto.LotteryId_Pl5)},
	[3]int32{2, 1, int32(apiproto.LotteryId_AllId)},
}

// 查看手机号注册礼包
func (srv *UcServer) HandleOpenPhoneRegistGift(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleOpenPhoneRegistGift")
	//result := &apiproto.RegistGiftPackage{Title: "运气爆棚，赶快领取礼包去买彩票吧！"}
	//for _, v := range _giftConfig {
	//	ticket := &dbproto.BuycaiTicket{UseBase: v[0], UseSub: v[1], RestrictId: v[2]}
	//	desc := utils.GetBuycaiTicketRestrictDesc(ticket)
	//	item := &apiproto.GiftItem{Type: 0, Desc: desc, Value: ticket.GetUseSub()}
	//	result.List = append(result.List, item)
	//}
	//utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", result) Credits
	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

	giftJson, err := srv.dbClient.QueryPhoneUserRegisterGift(context.Background(), &dbproto.Nil{})
	if err != nil {
		log.Printf("%+v\n", err)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "获取失败", nil)
		return
	}

	if giftJson.GetValue() == "" {
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "获取失败", nil)
		return
	}

	gift := &dbproto.GiftPackageArg{}
	if err = json.Unmarshal([]byte(giftJson.GetValue()), &gift); err != nil {
		log.Printf("%+v\n", err)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "获取失败", nil)
		return
	}

	result := &apiproto.RegistGiftPackage{Title: "运气爆棚，赶快领取礼包去买彩票吧！"}
	if len(gift.Tickets) > 0 {
		for k, v := range gift.Tickets {
			ticket := &dbproto.BuycaiTicket{UseBase: v.UseBase, UseSub: v.UseSub, RestrictId: v.RestrictId}
			desc := utils.GetBuycaiTicketRestrictDesc(ticket)
			item := &apiproto.GiftItem{Type: 0, Desc: desc, Value: ticket.GetUseSub()}
			result.List = append(result.List, item)
			log.Println(k, "desc:", desc, "item:", item)
		}
	}
	// result.Credits = int32(gift.Credits)

	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", result)
}

func (srv *UcServer) HandleReceivePhoneRegistGift(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleReceivePhoneRegistGift")
	accountId, _, err := utils.ParseHttpRequest(w, r, nil)
	if err != nil {
		return
	}
	ret, err := srv.dbClient.QueryPhoneRegistGiftReceived(context.Background(), &dbproto.IntValue{Value: accountId})
	if err != nil {
		log.Println(err)
		http.Error(w, "db: QueryPhoneRegistGiftReceived", http.StatusInternalServerError)
		return
	}
	if ret.GetValue() {
		log.Println("已经领取过:", accountId)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "已经领取过!", nil)
		return
	}

	giftJson, err := srv.dbClient.QueryPhoneUserRegisterGift(context.Background(), &dbproto.Nil{})
	if err != nil {
		log.Printf("%+v\n", err)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "获取失败", nil)
		return
	}

	if giftJson.GetValue() == "" {
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "获取失败", nil)
		return
	}

	gift := &dbproto.GiftPackageArg{}
	if err = json.Unmarshal([]byte(giftJson.GetValue()), &gift); err != nil {
		log.Printf("%+v\n", err)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "获取失败", nil)
		return
	}

	now := time.Now()
	validStart := now
	validEnd := now.AddDate(0, 0, 7)

	//组装数据.
	result := &apiproto.RegistGiftPackage{Title: "运气爆棚，赶快领取礼包去买彩票吧！"}
	dbArg := &dbproto.InsertBuycaiTicketArg{}
	if len(gift.Tickets) > 0 {
		for _, v := range gift.Tickets {
			ticket := &dbproto.BuycaiTicket{
				AccountId: accountId, UseBase: v.GetUseBase(), UseSub: v.GetUseSub(), RestrictId: v.GetRestrictId(),
				Title:   "新用户绑定手机礼包",
				AddTime: now.Unix(), ValidStart: validStart.Unix(), ValidEnd: validEnd.Unix(),
			}
			dbArg.List = append(dbArg.List, ticket)

			desc := utils.GetBuycaiTicketRestrictDesc(ticket)
			item := &apiproto.GiftItem{Type: 0, Desc: desc, Value: ticket.GetUseSub()}
			result.List = append(result.List, item)
		}
	}

	phoneResisterGift := &dbproto.PhoneRegistGift{
		AccountId: accountId,
		Credits:   int32(gift.GetCredits()),
		List:      gift.GetTickets(),
	}
	_, err = srv.dbClient.InsertPhoneRegistGiftReceived(context.Background(), phoneResisterGift)
	if err != nil {
		log.Printf("%+v\n", err)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "领取失败", nil)
		return
	}

	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", result)
}
