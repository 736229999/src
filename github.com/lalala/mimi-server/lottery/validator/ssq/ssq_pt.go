package ssq

import (
	"fmt"
	"log"

	"github.com/caojunxyz/mimi-api/proto"
	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/lottery/site/zhongfu"
	"github.com/caojunxyz/mimi-server/lottery/validator"
)

func convert_pt(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	redBalls := validator.GetStringBalls(sch, K_RED, 2)
	blueBalls := validator.GetStringBalls(sch, K_BLUE, 2)

	typeName := "单式"
	if len(redBalls) > 6 || len(blueBalls) > 1 {
		typeName = "复式"
	}
	return zhongfu.Scheme{
		Type:   typeName,
		TypeId: fmt.Sprint(zfTypeId),
		Money:  fmt.Sprint(sch.GetMoney()),
		Num:    fmt.Sprint(sch.GetNum()),
		Balls:  zhongfu.MakeBalls(ballsFormat, redBalls, blueBalls),
	}
}

func validate_pt(sch *apiproto.BuycaiScheme) bool {
	redBalls, ok := validator.CheckBalls(sch, K_RED, 6, 33, MIN_RED, MAX_RED)
	if !ok {
		return false
	}
	blueBalls, ok := validator.CheckBalls(sch, K_BLUE, 1, 16, MIN_BLUE, MAX_BLUE)
	if !ok {
		return false
	}

	numRed := len(redBalls)
	numBlue := len(blueBalls)
	num := int32(validator.Comb(numRed, 6) * validator.Comb(numBlue, 1))
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
