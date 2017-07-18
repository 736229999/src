package dlt

import (
	"fmt"
	"log"

	"github.com/caojunxyz/mimi-api/proto"
	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/lottery/site/zhongfu"
	"github.com/caojunxyz/mimi-server/lottery/validator"
)

func convert_hqdt(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	qianBalls := validator.GetStringBalls(sch, K_QIAN, 2)
	hdBalls := validator.GetStringBalls(sch, K_HD, 2)
	htBalls := validator.GetStringBalls(sch, K_HT, 2)
	return zhongfu.Scheme{
		Type:   "后区胆拖",
		TypeId: fmt.Sprint(zfTypeId),
		Money:  fmt.Sprint(sch.GetMoney()),
		Num:    fmt.Sprint(sch.GetNum()),
		Balls:  zhongfu.MakeBalls(ballsFormat, qianBalls, hdBalls, htBalls),
	}
}

func validate_hqdt(sch *apiproto.BuycaiScheme) bool {
	_, ok := validator.CheckBalls(sch, K_QIAN, 5, 5, MIN_RED, MAX_RED)
	if !ok {
		return false
	}
	hdBalls, ok := validator.CheckBalls(sch, K_HD, 0, 1, MIN_BLUE, MAX_BLUE)
	if !ok {
		return false
	}
	htBalls, ok := validator.CheckBalls(sch, K_HT, 2, 12, MIN_BLUE, MAX_BLUE)
	if !ok {
		return false
	}

	if !validator.CheckDuplicate(hdBalls, htBalls) {
		return false
	}

	numHd := len(hdBalls)
	numHt := len(htBalls)
	num := int32(validator.Comb(numHt, 2-numHd))
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
