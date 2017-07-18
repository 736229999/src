package pl3

import (
	"fmt"
	"log"

	"github.com/caojunxyz/mimi-api/proto"
	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/lottery/site/zhongfu"
	"github.com/caojunxyz/mimi-server/lottery/validator"
)

var zuxhz = map[int32]int32{
	1:  1,
	2:  2,
	3:  2,
	4:  4,
	5:  5,
	6:  6,
	7:  8,
	8:  10,
	9:  11,
	10: 13,
	11: 14,
	12: 14,
	13: 15,
	14: 15,
	15: 14,
	16: 14,
	17: 13,
	18: 11,
	19: 10,
	20: 8,
	21: 6,
	22: 5,
	23: 4,
	24: 2,
	25: 2,
	26: 1,
}

func convert_zuxhz(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
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

func validate_zuxhz(sch *apiproto.BuycaiScheme) bool {
	balls, ok := validator.CheckBalls(sch, K, 1, 3*MAX-1, MIN, 3*MAX-1)
	if !ok {
		return false
	}

	num := int32(0)
	for _, v := range balls {
		num += zuxhz[v]
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
