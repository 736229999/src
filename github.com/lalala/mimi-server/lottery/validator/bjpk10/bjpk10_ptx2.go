package bjpk10

import (
	"fmt"
	"log"

	"github.com/caojunxyz/mimi-api/proto"
	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/lottery/site/zhongfu"
	"github.com/caojunxyz/mimi-server/lottery/validator"
)

func convert_ptx2(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	balls1 := validator.GetStringBalls(sch, K1, 2)
	balls2 := validator.GetStringBalls(sch, K2, 2)
	return zhongfu.Scheme{
		Type:   "猜冠亚军",
		TypeId: fmt.Sprint(zfTypeId),
		Money:  fmt.Sprint(sch.GetMoney()),
		Num:    fmt.Sprint(sch.GetNum()),
		Balls:  zhongfu.MakeBalls(ballsFormat, balls1, balls2),
	}
}

func validate_ptx2(sch *apiproto.BuycaiScheme) bool {
	balls1, ok := validator.CheckBalls(sch, K1, 1, 10, MIN, MAX)
	if !ok {
		return false
	}
	balls2, ok := validator.CheckBalls(sch, K2, 1, 10, MIN, MAX)
	if !ok {
		return false
	}

	var num int32 = 0
	for _, v1 := range balls1 {
		for _, v2 := range balls2 {
			if v1 != v2 {
				num++
			}
		}
	}

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
