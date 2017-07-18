package dlt

import (
	"fmt"
	"log"

	"github.com/caojunxyz/mimi-api/proto"
	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/lottery/site/zhongfu"
	"github.com/caojunxyz/mimi-server/lottery/validator"
)

func convert_qqdt(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	qdBalls := validator.GetStringBalls(sch, K_QD, 2)
	qtBalls := validator.GetStringBalls(sch, K_QT, 2)
	houBalls := validator.GetStringBalls(sch, K_HOU, 2)
	return zhongfu.Scheme{
		Type:   "前区胆拖",
		TypeId: fmt.Sprint(zfTypeId),
		Money:  fmt.Sprint(sch.GetMoney()),
		Num:    fmt.Sprint(sch.GetNum()),
		Balls:  zhongfu.MakeBalls(ballsFormat, qdBalls, qtBalls, houBalls),
	}
}

func validate_qqdt(sch *apiproto.BuycaiScheme) bool {
	qdBalls, ok := validator.CheckBalls(sch, K_QD, 1, 4, MIN_RED, MAX_RED)
	if !ok {
		return false
	}
	qtBalls, ok := validator.CheckBalls(sch, K_QT, 2, 14, MIN_RED, MAX_RED)
	if !ok {
		return false
	}
	houBalls, ok := validator.CheckBalls(sch, K_HOU, 2, 2, MIN_BLUE, MAX_BLUE)
	if !ok {
		return false
	}

	if !validator.CheckDuplicate(qdBalls, qtBalls) {
		return false
	}

	numQt := len(qtBalls)
	numQd := len(qdBalls)
	numHou := len(houBalls)
	num := int32(validator.Comb(numQt, 5-numQd) * validator.Comb(numHou, 2))
	if num != sch.GetNum() {
		log.Println("注数计算不一致:", num, sch.GetNum())
		return false
	}

	sumMoney := float64(num * 2)
	if !validator.IsEqualMoney(sumMoney, sch.GetMoney()) {
		log.Println("金额计算不一致:", sumMoney, sch.GetMoney())
		return false
	}
	return true
}
