package main

import (
	"net/http"
	"github.com/caojunxyz/mimi-api/proto"
	"log"
	"github.com/caojunxyz/mimi-server/utils"
	"golang.org/x/net/context"
	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
)

//领取礼包.
func (srv *UcServer) HandleReceiveGift (w http.ResponseWriter, r *http.Request) {

	msg := &apiproto.IntValue{}
	accountId, _, err := utils.ParseHttpRequest(w, r, msg)
	if err != nil {
		log.Printf("%+v\n", err)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "领取失败", nil)
		return
	}

	//根据活动id获取礼包模板.
	template, err := srv.dbClient.QueryGiftTemplateById(context.Background(), &dbproto.IntValue{Value:msg.GetValue()})
	if err != nil {
		log.Printf("%+v\n", err)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "领取失败", nil)
	}

	//根据礼包模板生成对应的礼包.
	userGift := &dbproto.UserGiftPackage{Gift:new(dbproto.UserGiftPackageContent)}
	credits := template.GetContent().GetCredits()
	//验证当前礼包模板中赠送积分是否随机.
	if credits.GetRandomCredits() {
		//积分随机.
		//获取积分随机的最大上限和最少下限.
		upper_limit := credits.GetUpperLimit()
		lower_limit := credits.GetLowerLimit()
		userGift.Gift.Credits = RandNum(lower_limit, upper_limit)

	} else {
		//积分不随机.
		userGift.Gift.Credits = credits.GetCredits()
	}

	//验证当前礼包模板中赠送的购彩券是否随机.
	ticket := template.GetContent().GetTickets()
	if ticket.GetRandomTickets() {
		//购彩券随机.
		upper_limit := ticket.GetUpperLimit()
		lower_limit  := ticket.GetLowerLimit()
		randNum := RandNum(lower_limit, upper_limit)
		log.Println("upper_limit:", upper_limit, "lower_limit:", lower_limit, "随机数：", randNum)
		for k, v := range ticket.GetTickets() {
			if k + 1 <= int(randNum) {
				userGift.Gift.Tickets = append(userGift.Gift.Tickets , v)
			}
		}

	} else {
		//购彩券不随机.
		userGift.Gift.Tickets = ticket.GetTickets()
	}
	userGift.AccountId = accountId

	//将礼包添加到user_gift表中.
	srv.dbClient.ReceiveGift(context.Background(), userGift)

}
