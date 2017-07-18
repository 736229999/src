package dlt

import (
	"fmt"
	"log"

	"github.com/caojunxyz/mimi-api/proto"
	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/lottery/site/zhongfu"
	"github.com/caojunxyz/mimi-server/lottery/validator"
)

func convert_pt(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	qianBalls := validator.GetStringBalls(sch, K_QIAN, 2)
	houBalls := validator.GetStringBalls(sch, K_HOU, 2)
	return zhongfu.Scheme{
		Type:   "普通投注",
		TypeId: fmt.Sprint(zfTypeId),
		Money:  fmt.Sprint(sch.GetMoney()),
		Num:    fmt.Sprint(sch.GetNum()),
		Balls:  zhongfu.MakeBalls(ballsFormat, qianBalls, houBalls),
	}
}

func validate_pt(sch *apiproto.BuycaiScheme) bool {
	qianBalls, ok := validator.CheckBalls(sch, K_QIAN, 5, 18, MIN_RED, MAX_RED)
	if !ok {
		return false
	}
	houBalls, ok := validator.CheckBalls(sch, K_HOU, 2, 12, MIN_BLUE, MAX_BLUE)
	if !ok {
		return false
	}

	numQian := len(qianBalls)
	numHou := len(houBalls)
	num := int32(validator.Comb(numQian, 5) * validator.Comb(numHou, 2))
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
