package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"strings"

	"sort"

	apiproto "github.com/caojunxyz/mimi-api/proto"
	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/lottery"
	"github.com/caojunxyz/mimi-server/utils"
	"google.golang.org/grpc"
)

var dbc dbproto.DbThirdApiAgentClient

func connectDb(addr string) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Panic(err)
	}
	dbc = dbproto.NewDbThirdApiAgentClient(conn)
}

func dbQueryIssuesByOpendate(code, date string) []*dbproto.OpenInfo {
	result := make([]*dbproto.OpenInfo, 0)
	arg := &dbproto.OpencaiQueryArg{
		Code: code,
		Args: []string{date},
	}
	stream, err := dbc.OpencaiQueryByOpendate(context.Background(), arg)
	if err != nil {
		log.Println(err)
		return nil
	}

	for {
		v, err := stream.Recv()
		if err == nil {
			result = append(result, v)
			continue
		}
		if err != io.EOF {
			log.Println(err)
		}
		break
	}
	return result
}

func dbQueryIssuesByLatestNum(code string, num int) []*dbproto.OpenInfo {
	result := make([]*dbproto.OpenInfo, 0)
	arg := &dbproto.OpencaiQueryArg{
		Code: code,
		Args: []string{fmt.Sprint(num)},
	}
	stream, err := dbc.OpencaiQueryByLatestNum(context.Background(), arg)
	if err != nil {
		log.Println(err)
		return nil
	}

	for {
		v, err := stream.Recv()
		if err == nil {
			result = append(result, v)
			continue
		}
		if err != io.EOF {
			log.Println(err)
		}
		break
	}
	return result
}

func dbInfoToApiInfo(conf *lottery.Config, v *dbproto.OpenInfo) *apiproto.OpenInfo {
	ret := &apiproto.OpenInfo{
		Type:     conf.Type,
		Id:       conf.Id,
		Name:     conf.Name,
		BlueNum:  conf.BlueNum,
		Issue:    v.Issue,
		OpenTime: v.OpenTime,
	}

	balls := strings.Replace(v.Balls, "+", ",", -1)
	ret.Balls = strings.Split(balls, ",")

	if detail := v.GetDetail(); detail != nil {
		apiDetail := &apiproto.OpenDetail{Sale: detail.GetSale(), Pool: detail.GetPool()}
		for _, v0 := range detail.BonusList {
			name := conf.BonusNames[int(v0.GetId())]
			bd := &apiproto.BonusDetail{Name: name, Num: v0.GetNum(), Money: v0.GetMoney()}
			apiDetail.BonusList = append(apiDetail.BonusList, bd)
		}
		ret.Detail = apiDetail
	}
	return ret
}

func getLatestOpen() []*apiproto.OpenInfo {
	confList := lottery.GetLottery()
	ret := make([]*apiproto.OpenInfo, 0, len(confList))
	for _, conf := range confList {
		if conf.Type != apiproto.LotteryType_LowFreq && conf.Type != apiproto.LotteryType_HighFreq {
			continue
		}
		if list := dbQueryIssuesByLatestNum(conf.Code, 1); len(list) == 1 {
			latest := list[0]
			ret = append(ret, dbInfoToApiInfo(&conf, latest))
		}
	}
	sort.SliceStable(ret, func(i, j int) bool {
		return ret[i].GetId() < ret[j].GetId()
	})
	return ret
}

var dayName = map[int]string{2: "前天", 1: "昨天", 0: "今天"}

func getHistory(id apiproto.LotteryId) *apiproto.History {
	history := &apiproto.History{}
	if conf := lottery.GetConfig(id); conf != nil {
		if conf.Type == apiproto.LotteryType_HighFreq {
			for n := 2; n >= 0; n-- {
				dayTime := utils.TimeBeforeDays(n)
				date := dayTime.Format("2006-01-02")
				dayHistory := &apiproto.DayHistory{
					Name: dayName[n],
					Date: date,
				}
				list := dbQueryIssuesByOpendate(conf.Code, date)
				for _, v := range list {
					dayHistory.List = append(dayHistory.List, dbInfoToApiInfo(conf, v))
				}
				history.Days = append(history.Days, dayHistory)
			}
		} else if conf.Type == apiproto.LotteryType_LowFreq {
			dayHistory := &apiproto.DayHistory{}
			list := dbQueryIssuesByLatestNum(conf.Code, 30)
			for _, v := range list {
				dayHistory.List = append(dayHistory.List, dbInfoToApiInfo(conf, v))
			}
			history.Days = append(history.Days, dayHistory)
		}

	}
	return history
}

func getDigestHistory(id apiproto.LotteryId) *apiproto.DigestHistory {
	result := &apiproto.DigestHistory{}
	if conf := lottery.GetConfig(id); conf != nil {
		list := dbQueryIssuesByLatestNum(conf.Code, 10)
		for _, v := range list {
			info := dbInfoToApiInfo(conf, v)
			dig := &apiproto.OpenInfoDigest{
				Issue:   info.Issue,
				Balls:   info.Balls,
				BlueNum: info.BlueNum,
			}
			result.List = append(result.List, dig)
		}
	}
	return result
}
