package cqssc

import (
	"fmt"
	"log"

	"github.com/caojunxyz/mimi-api/proto"
	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/lottery/site/zhongfu"
	"github.com/caojunxyz/mimi-server/lottery/validator"
)

// 中三直选
func convert_z3zhixuan(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	qianBalls := validator.GetStringBalls(sch, K_QIAN, 1)
	baiBalls := validator.GetStringBalls(sch, K_BAI, 1)
	shiBalls := validator.GetStringBalls(sch, K_SHI, 1)

	typeName := "单式"
	zfTypeId = 2834
	ballsFormat = "-%s%s%s-"
	if len(qianBalls) > 1 || len(baiBalls) > 1 || len(shiBalls) > 1 {
		typeName = "复式"
		zfTypeId = 2833
		ballsFormat = "-(%s)(%s)(%s)-"
	}
	return zhongfu.Scheme{
		Type:   typeName,
		TypeId: fmt.Sprint(zfTypeId),
		Money:  fmt.Sprint(sch.GetMoney()),
		Num:    fmt.Sprint(sch.GetNum()),
		Balls:  zhongfu.MakeBalls2(ballsFormat, "", qianBalls, baiBalls, shiBalls),
	}
}

func validate_z3zhixuan(sch *apiproto.BuycaiScheme) bool {
	qianBalls, ok := validator.CheckBalls(sch, K_QIAN, 1, 10, MIN, MAX)
	if !ok {
		return false
	}
	baiBalls, ok := validator.CheckBalls(sch, K_BAI, 1, 10, MIN, MAX)
	if !ok {
		return false
	}
	shiBalls, ok := validator.CheckBalls(sch, K_SHI, 1, 10, MIN, MAX)
	if !ok {
		return false
	}
	numQian := len(qianBalls)
	numBai := len(baiBalls)
	numShi := len(shiBalls)
	num := int32(numQian * numBai * numShi)

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

// 中三组三
func convert_z3zu3(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	balls := validator.GetStringBalls(sch, K1, 1)

	return zhongfu.Scheme{
		Type:   "",
		TypeId: fmt.Sprint(zfTypeId),
		Money:  fmt.Sprint(sch.GetMoney()),
		Num:    fmt.Sprint(sch.GetNum()),
		Balls:  zhongfu.MakeBalls2(ballsFormat, " ", balls),
	}
}

func validate_z3zu3(sch *apiproto.BuycaiScheme) bool {
	balls, ok := validator.CheckBalls(sch, K1, 2, 10, MIN, MAX)
	if !ok {
		return false
	}
	numBalls := len(balls)

	num := validator.Comb(numBalls, 2) * 2
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

// 中三组六
func convert_z3zu6(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	balls := validator.GetStringBalls(sch, K1, 1)

	return zhongfu.Scheme{
		Type:   "",
		TypeId: fmt.Sprint(zfTypeId),
		Money:  fmt.Sprint(sch.GetMoney()),
		Num:    fmt.Sprint(sch.GetNum()),
		Balls:  zhongfu.MakeBalls2(ballsFormat, " ", balls),
	}
}

func validate_z3zu6(sch *apiproto.BuycaiScheme) bool {
	balls, ok := validator.CheckBalls(sch, K1, 3, 10, MIN, MAX)
	if !ok {
		return false
	}
	numBalls := len(balls)

	num := validator.Comb(numBalls, 3)
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

// 中三组选和值
func convert_z3zuxuan_hz(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
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

func validate_z3zuxuan_hz(sch *apiproto.BuycaiScheme) bool {
	balls, ok := validator.CheckBalls(sch, K1, 1, 26, 1, 26)
	if !ok {
		return false
	}

	var num int32 = 0
	for _, v := range balls {
		num += _zuxuan_hz[v]
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

// 中三直选和值
func convert_z3zhixuan_hz(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
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

func validate_z3zhixuan_hz(sch *apiproto.BuycaiScheme) bool {
	balls, ok := validator.CheckBalls(sch, K1, 1, 28, 0, 27)
	if !ok {
		return false
	}

	var num int32 = 0
	for _, v := range balls {
		num += _zhixuan_hz[v]
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
