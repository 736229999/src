package cqssc

import (
	"fmt"
	"log"

	"github.com/caojunxyz/mimi-api/proto"
	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/lottery/site/zhongfu"
	"github.com/caojunxyz/mimi-server/lottery/validator"
)

func convert(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	balls := validator.GetStringBalls(sch, K1, 1)

	var typeName string = "单式"
	if len(balls) > 1 {
		typeName = "复式"

	}

	return zhongfu.Scheme{
		Type:   typeName,
		TypeId: fmt.Sprint(zfTypeId),
		Money:  fmt.Sprint(sch.GetMoney()),
		Num:    fmt.Sprint(sch.GetNum()),
		Balls:  zhongfu.MakeBalls2(ballsFormat, ",", balls),
	}
}

func validate(sch *apiproto.BuycaiScheme) bool {
	balls, ok := validator.CheckBalls(sch, K1, 1, 10, MIN, MAX)
	if !ok {
		return false
	}

	num := len(balls)
	if int32(num) != sch.GetNum() {
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

// 一帆风顺
func convert_1ffs(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	result := convert(sch, zfTypeId, ballsFormat)
	return result
}

func validate_1ffs(sch *apiproto.BuycaiScheme) bool {
	ok := validate(sch)
	return ok
}

// 好事成双
func convert_hscs(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	result := convert(sch, zfTypeId, ballsFormat)
	return result
}

func validate_hscs(sch *apiproto.BuycaiScheme) bool {
	ok := validate(sch)
	return ok
}

// 三星报喜
func convert_3xbx(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	result := convert(sch, zfTypeId, ballsFormat)
	return result
}

func validate_3xbx(sch *apiproto.BuycaiScheme) bool {
	ok := validate(sch)
	return ok
}

// 四季发财
func convert_4jfc(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	result := convert(sch, zfTypeId, ballsFormat)
	return result
}

func validate_4jfc(sch *apiproto.BuycaiScheme) bool {
	ok := validate(sch)
	return ok
}
