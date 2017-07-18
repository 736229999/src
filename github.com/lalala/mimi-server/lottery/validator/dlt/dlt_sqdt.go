package dlt

import (
	"fmt"
	"log"

	"github.com/caojunxyz/mimi-api/proto"
	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/lottery/site/zhongfu"
	"github.com/caojunxyz/mimi-server/lottery/validator"
)

func convert_sqdt(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	qdBalls := validator.GetStringBalls(sch, K_QD, 2)
	qtBalls := validator.GetStringBalls(sch, K_QT, 2)
	hdBalls := validator.GetStringBalls(sch, K_HD, 2)
	htBalls := validator.GetStringBalls(sch, K_HT, 2)
	return zhongfu.Scheme{
		Type:   "双区胆拖",
		TypeId: fmt.Sprint(zfTypeId),
		Money:  fmt.Sprint(sch.GetMoney()),
		Num:    fmt.Sprint(sch.GetNum()),
		Balls:  zhongfu.MakeBalls(ballsFormat, qdBalls, qtBalls, hdBalls, htBalls),
	}
}

func validate_sqdt(sch *apiproto.BuycaiScheme) bool {
	qdBalls, ok := validator.CheckBalls(sch, K_QD, 1, 4, 1, 35)
	if !ok {
		return false
	}
	qtBalls, ok := validator.CheckBalls(sch, K_QT, 2, 14, 1, 35)
	if !ok {
		return false
	}
	if !validator.CheckDuplicate(qdBalls, qtBalls) {
		return false
	}

	hdBalls, ok := validator.CheckBalls(sch, K_HD, 1, 1, 1, 12)
	if !ok {
		return false
	}
	htBalls, ok := validator.CheckBalls(sch, K_HT, 2, 12, 1, 12)
	if !ok {
		return false
	}
	if !validator.CheckDuplicate(hdBalls, htBalls) {
		return false
	}

	numQd := len(qdBalls)
	numQt := len(qtBalls)
	numHd := len(hdBalls)
	numHt := len(htBalls)
	num := int32(validator.Comb(numQt, 5-numQd) * validator.Comb(numHt, 2-numHd))
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
