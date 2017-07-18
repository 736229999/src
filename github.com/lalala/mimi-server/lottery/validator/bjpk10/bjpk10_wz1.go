package bjpk10

import (
	"fmt"
	"log"

	"github.com/caojunxyz/mimi-api/proto"
	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/lottery/site/zhongfu"
	"github.com/caojunxyz/mimi-server/lottery/validator"
)

func convert_wz1(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	balls := validator.GetStringBalls(sch, K1, 2)
	return zhongfu.Scheme{
		Type:   "位置一",
		TypeId: fmt.Sprint(zfTypeId),
		Money:  fmt.Sprint(sch.GetMoney()),
		Num:    fmt.Sprint(sch.GetNum()),
		Balls:  zhongfu.MakeBalls(ballsFormat, balls),
	}
}

func validate_wz1(sch *apiproto.BuycaiScheme) bool {
	balls, ok := validator.CheckBalls(sch, K1, 1, 10, MIN, MAX)
	if !ok {
		return false
	}

	numBalls := len(balls)
	if int32(numBalls) != sch.GetNum() {
		log.Println("注数计算不一致:", numBalls, sch.GetNum())
		return false
	}

	sumMoney := float64(numBalls * 2)
	if !validator.IsEqualMoney(sumMoney, sch.GetMoney()) {
		log.Println("金额计算不一致:", sumMoney, sch.GetMoney())
		return false
	}

	return true
}
