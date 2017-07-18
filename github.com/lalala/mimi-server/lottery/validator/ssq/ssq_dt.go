package ssq

import (
	"fmt"
	"log"

	"github.com/caojunxyz/mimi-api/proto"
	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/lottery/site/zhongfu"
	"github.com/caojunxyz/mimi-server/lottery/validator"
)

func convert_dt(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	danBalls := validator.GetStringBalls(sch, K_DAN, 2)
	tuoBalls := validator.GetStringBalls(sch, K_TUO, 2)
	blueBalls := validator.GetStringBalls(sch, K_BLUE, 2)
	return zhongfu.Scheme{
		Type:   "胆拖",
		TypeId: fmt.Sprint(zfTypeId),
		Money:  fmt.Sprint(sch.GetMoney()),
		Num:    fmt.Sprint(sch.GetNum()),
		Balls:  zhongfu.MakeBalls(ballsFormat, danBalls, tuoBalls, blueBalls),
	}
}

func validate_dt(sch *apiproto.BuycaiScheme) bool {
	danBalls, ok := validator.CheckBalls(sch, K_DAN, 1, 5, MIN_RED, MAX_RED)
	if !ok {
		return false
	}
	tuoBalls, ok := validator.CheckBalls(sch, K_TUO, 2, 33, MIN_RED, MAX_RED)
	if !ok {
		return false
	}
	blueBalls, ok := validator.CheckBalls(sch, K_BLUE, 1, 16, MIN_BLUE, MAX_BLUE)
	if !ok {
		return false
	}
	if !validator.CheckDuplicate(danBalls, tuoBalls) {
		return false
	}

	numDan := len(danBalls)
	numTuo := len(tuoBalls)
	numBlue := len(blueBalls)
	num := int32(validator.Comb(numTuo, 6-numDan) * validator.Comb(numBlue, 1))
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
