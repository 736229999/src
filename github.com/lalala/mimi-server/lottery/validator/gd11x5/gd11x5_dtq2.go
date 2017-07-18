package gd11x5

import (
	"fmt"
	"log"

	"github.com/caojunxyz/mimi-api/proto"
	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/lottery/site/zhongfu"
	"github.com/caojunxyz/mimi-server/lottery/validator"
)

func convert_dtq2(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	danBalls := validator.GetStringBalls(sch, K_DAN, 2)
	tuoBalls := validator.GetStringBalls(sch, K_TUO, 2)
	return zhongfu.Scheme{
		Type:   "胆拖",
		TypeId: fmt.Sprint(zfTypeId),
		Money:  fmt.Sprint(sch.GetMoney()),
		Num:    fmt.Sprint(sch.GetNum()),
		Balls:  zhongfu.MakeBalls(ballsFormat, danBalls, tuoBalls),
	}
}

func validate_dtq2(sch *apiproto.BuycaiScheme) bool {
	danBalls, ok := validator.CheckBalls(sch, K_DAN, 1, 1, MIN, MAX)
	if !ok {
		return false
	}
	tuoBalls, ok := validator.CheckBalls(sch, K_TUO, 2, 10, MIN, MAX)
	if !ok {
		return false
	}

	if !validator.CheckDuplicate(danBalls, tuoBalls) {
		return false
	}

	numDan := len(danBalls)
	numTuo := len(tuoBalls)
	num := validator.Comb(numTuo, 2-numDan)
	if int32(num) != sch.GetNum() {
		log.Println("注数计算不一致:", numTuo, sch.GetNum())
		return false
	}

	sumMoney := float64(num * 2)
	if !validator.IsEqualMoney(sumMoney, sch.GetMoney()) {
		log.Println("金额计算不一致:", sumMoney, sch.GetMoney())
		return false
	}

	return true
}
