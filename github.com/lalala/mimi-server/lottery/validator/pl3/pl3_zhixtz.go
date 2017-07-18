package pl3

import (
	"fmt"
	"log"
	"strings"

	"github.com/caojunxyz/mimi-api/proto"
	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/lottery/site/zhongfu"
	"github.com/caojunxyz/mimi-server/lottery/validator"
)

func convert_zhixtz(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	baiBalls := validator.GetStringBalls(sch, K_BAI, 1)
	shiBalls := validator.GetStringBalls(sch, K_SHI, 1)
	geBalls := validator.GetStringBalls(sch, K_GE, 1)

	ballsList := [...][]string{baiBalls, shiBalls, geBalls}
	typeName := "单式"
	var balls string
	for _, v := range ballsList {
		if len(v) > 1 {
			typeName = "复式"
			balls += fmt.Sprintf("(%s)", strings.Join(v, ""))
		} else {
			balls += fmt.Sprintf("%s", strings.Join(v, ""))
		}
	}

	return zhongfu.Scheme{
		Type:   typeName,
		TypeId: fmt.Sprint(zfTypeId),
		Money:  fmt.Sprint(sch.GetMoney()),
		Num:    fmt.Sprint(sch.GetNum()),
		Balls:  balls,
	}
}

func validate_zhixtz(sch *apiproto.BuycaiScheme) bool {
	baiBalls, ok := validator.CheckBalls(sch, K_BAI, 1, 10, MIN, MAX)
	if !ok {
		return false
	}
	shiBalls, ok := validator.CheckBalls(sch, K_SHI, 1, 10, MIN, MAX)
	if !ok {
		return false
	}
	geBalls, ok := validator.CheckBalls(sch, K_GE, 1, 10, MIN, MAX)
	if !ok {
		return false
	}

	numBai := len(baiBalls)
	numShi := len(shiBalls)
	numGe := len(geBalls)
	num := int32(numBai * numShi * numGe)
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
