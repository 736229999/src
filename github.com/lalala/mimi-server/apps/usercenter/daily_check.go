package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/caojunxyz/gotu"
	apiproto "github.com/caojunxyz/mimi-api/proto"
	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/utils"
)

// TODO: 配置表
var levelTable = [...]int32{0, 13, 120, 600, 1200, 12000, 36000, 120000, 1200000, 120000000}

func CalUserLevel(exp int32) int32 {
	level := 0
	for i, v := range levelTable {
		if exp >= v {
			level = i + 1
		}
	}
	return int32(level)
}

// 连续签到exp+2, 否则+1
func (srv *UcServer) HandleDailyCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleDailyCheck")
	accountId, _, err := utils.ParseHttpRequest(w, r, nil)
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}
	userInfo, err := srv.dbClient.QueryUserInfo(context.Background(), &dbproto.IntValue{Value: accountId})
	if err != nil {
		log.Println(err)
		http.Error(w, "错误：系统错误!", http.StatusInternalServerError)
		return
	}

	contDays := userInfo.GetContCheckDays()
	now := time.Now()
	lastCheckTime := time.Unix(userInfo.GetDailyCheckTime(), 0)
	if gotu.IsSameDay(now, lastCheckTime) {
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "今日已签到", nil)
		return
	}

	addExp := int32(0)
	if gotu.DaysInterval(now, lastCheckTime) > 1 {
		addExp = 2
		contDays += 1
	} else {
		addExp = 1
		contDays = 1
	}

	exp := userInfo.GetExp() + addExp
	level := CalUserLevel(exp)
	checkArg := &dbproto.DailyCheckArg{AccountId: accountId, Exp: exp, Level: level, ContCheckDays: contDays}
	_, err = srv.dbClient.SetDailyCheck(context.Background(), checkArg)
	if err != nil {
		log.Println(err)
		http.Error(w, "错误：系统错误!", http.StatusInternalServerError)
		return
	}

	result := &apiproto.DailyCheckReply{ContCheckDays: contDays, Exp: addExp}
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", result)
}
