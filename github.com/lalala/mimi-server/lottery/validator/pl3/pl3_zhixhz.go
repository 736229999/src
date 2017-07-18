package pl3

import (
	"fmt"
	"log"

	"github.com/caojunxyz/mimi-api/proto"
	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/lottery/site/zhongfu"
	"github.com/caojunxyz/mimi-server/lottery/validator"
)

var zhixhz = map[int32]int32{
	0:  1,
	1:  3,
	2:  6,
	3:  10,
	4:  15,
	5:  21,
	6:  28,
	7:  36,
	8:  45,
	9:  55,
	10: 63,
	11: 69,
	12: 73,
	13: 75,
	14: 75,
	15: 73,
	16: 69,
	17: 63,
	18: 55,
	19: 45,
	20: 36,
	21: 28,
	22: 21,
	23: 15,
	24: 10,
	25: 6,
	26: 3,
	27: 1,
}

func convert_zhixhz(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	balls := validator.GetStringBalls(sch, K, 1)
	typeName := "单式"
	if len(balls) > 1 {
		typeName = "复式"
	}

	return zhongfu.Scheme{
		Type:   typeName,
		TypeId: fmt.Sprint(zfTypeId),
		Money:  fmt.Sprint(sch.GetMoney()),
		Num:    fmt.Sprint(sch.GetNum()),
		Balls:  zhongfu.MakeBalls(ballsFormat, balls),
	}
}

func validate_zhixhz(sch *apiproto.BuycaiScheme) bool {
	balls, ok := validator.CheckBalls(sch, K, 1, 3*MAX+1, MIN, 3*MAX)
	if !ok {
		return false
	}

	num := int32(0)
	for _, v := range balls {
		num += zhixhz[v]
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
