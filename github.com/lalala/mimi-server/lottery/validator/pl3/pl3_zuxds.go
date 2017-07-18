package pl3

import (
	"fmt"
	"log"

	"github.com/caojunxyz/mimi-api/proto"
	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/lottery/site/zhongfu"
	"github.com/caojunxyz/mimi-server/lottery/validator"
)

func convert_zuxds(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	baiBalls := validator.GetStringBalls(sch, K_BAI, 1)
	shiBalls := validator.GetStringBalls(sch, K_SHI, 1)
	geBalls := validator.GetStringBalls(sch, K_GE, 1)

	return zhongfu.Scheme{
		Type:   "单式",
		TypeId: fmt.Sprint(zfTypeId),
		Money:  fmt.Sprint(sch.GetMoney()),
		Num:    fmt.Sprint(sch.GetNum()),
		Balls:  zhongfu.MakeBalls(ballsFormat, baiBalls, shiBalls, geBalls),
	}
}

func validate_zuxds(sch *apiproto.BuycaiScheme) bool {
	_, ok := validator.CheckBalls(sch, K_BAI, 1, 1, MIN, MAX)
	if !ok {
		return false
	}
	_, ok = validator.CheckBalls(sch, K_SHI, 1, 1, MIN, MAX)
	if !ok {
		return false
	}
	_, ok = validator.CheckBalls(sch, K_GE, 1, 1, MIN, MAX)
	if !ok {
		return false
	}
	if sch.GetNum() != 1 {
		log.Println("注数计算不一致:", 1, sch.GetNum())
		return false
	}

	sumMoney := float64(2)
	if !validator.IsEqualMoney(sumMoney, sch.GetMoney()) {
		log.Println("金额计算不一致:", sumMoney, sch.GetMoney())
		return false
	}

	return true
}
