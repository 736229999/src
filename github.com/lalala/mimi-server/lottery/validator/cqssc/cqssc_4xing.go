package cqssc

import (
	"fmt"
	"log"

	"github.com/caojunxyz/mimi-api/proto"
	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/lottery/site/zhongfu"
	"github.com/caojunxyz/mimi-server/lottery/validator"
)

// 四星-前四直选
func convert_4x_q4zhixuan(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	wanBalls := validator.GetStringBalls(sch, K_WAN, 1)
	qianBalls := validator.GetStringBalls(sch, K_QIAN, 1)
	baiBalls := validator.GetStringBalls(sch, K_BAI, 1)
	shiBalls := validator.GetStringBalls(sch, K_SHI, 1)
	//
	// typeName := "单式"
	// if len(wanBalls) > 1 || len(qianBalls) > 1 || len(baiBalls) > 1 || len(shiBalls) > 1 {
	// 	typeName = "复式"
	// }

	return zhongfu.Scheme{
		Type:   "",
		TypeId: fmt.Sprint(zfTypeId),
		Money:  fmt.Sprint(sch.GetMoney()),
		Num:    fmt.Sprint(sch.GetNum()),
		Balls:  zhongfu.MakeBalls2(ballsFormat, "", wanBalls, qianBalls, baiBalls, shiBalls),
	}
}

func validate_4x_q4zhixuan(sch *apiproto.BuycaiScheme) bool {
	wanBalls, ok := validator.CheckBalls(sch, K_WAN, 1, 10, MIN, MAX)
	if !ok {
		return false
	}
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

	numWan := len(wanBalls)
	numQian := len(qianBalls)
	numBai := len(baiBalls)
	numShi := len(shiBalls)

	num := int32(numWan * numQian * numBai * numShi)

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

// 四星-后四直选
func convert_4x_h4zhixuan(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	qianBalls := validator.GetStringBalls(sch, K_QIAN, 1)
	baiBalls := validator.GetStringBalls(sch, K_BAI, 1)
	shiBalls := validator.GetStringBalls(sch, K_SHI, 1)
	geBalls := validator.GetStringBalls(sch, K_GE, 1)

	ballsFormat = "-%s%s%s%s"
	typeName := "单式"
	zfTypeId = 2821
	if len(qianBalls) > 1 || len(baiBalls) > 1 || len(shiBalls) > 1 || len(geBalls) > 1 {
		typeName = "复式"
		ballsFormat = "-(%s)(%s)(%s)(%s)"
		zfTypeId = 2820
	}
	return zhongfu.Scheme{
		Type:   typeName,
		TypeId: fmt.Sprint(zfTypeId),
		Money:  fmt.Sprint(sch.GetMoney()),
		Num:    fmt.Sprint(sch.GetNum()),
		Balls:  zhongfu.MakeBalls2(ballsFormat, "", qianBalls, baiBalls, shiBalls, geBalls),
	}
}

func validate_4x_h4zhixuan(sch *apiproto.BuycaiScheme) bool {
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
	geBalls, ok := validator.CheckBalls(sch, K_GE, 1, 10, MIN, MAX)
	if !ok {
		return false
	}

	numQian := len(qianBalls)
	numBai := len(baiBalls)
	numShi := len(shiBalls)
	numGe := len(geBalls)
	num := int32(numQian * numBai * numShi * numGe)

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

// 四星-后四组选24
func convert_4x_h4zuxuan24(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	balls := validator.GetStringBalls(sch, K1, 1)

	typeName := "单式"
	if len(balls) > 4 {
		typeName = "复式"
	}

	return zhongfu.Scheme{
		Type:   typeName,
		TypeId: fmt.Sprint(zfTypeId),
		Money:  fmt.Sprint(sch.GetMoney()),
		Num:    fmt.Sprint(sch.GetNum()),
		Balls:  zhongfu.MakeBalls2(ballsFormat, " ", balls),
	}
}

func validate_4x_h4zuxuan24(sch *apiproto.BuycaiScheme) bool {
	balls, ok := validator.CheckBalls(sch, K1, 4, 10, MIN, MAX)
	if !ok {
		return false
	}
	numBalls := len(balls)

	num := validator.Comb(numBalls, 4)
	if num != int(sch.GetNum()) {
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

// 四星-后四组选12
func convert_4x_h4zuxuan12(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	k1balls := validator.GetStringBalls(sch, K1, 1)
	k2balls := validator.GetStringBalls(sch, K2, 1)

	typeName := "单式"
	if len(k1balls) > 1 || len(k2balls) > 2 {
		typeName = "复式"
	}

	return zhongfu.Scheme{
		Type:   typeName,
		TypeId: fmt.Sprint(zfTypeId),
		Money:  fmt.Sprint(sch.GetMoney()),
		Num:    fmt.Sprint(sch.GetNum()),
		Balls:  zhongfu.MakeBalls2(ballsFormat, " ", k1balls, k2balls),
	}
}

func validate_4x_h4zuxuan12(sch *apiproto.BuycaiScheme) bool {
	k1balls, ok := validator.CheckBalls(sch, K1, 1, 10, MIN, MAX)
	if !ok {
		return false
	}
	k2balls, ok := validator.CheckBalls(sch, K2, 2, 10, MIN, MAX)
	if !ok {
		return false
	}

	interNum := validator.CalInter2Num(k1balls, k2balls)
	x := (len(k1balls) - interNum) * validator.Comb(len(k2balls), 2)
	y := interNum * validator.Comb(len(k2balls)-1, 2)
	num := x + y
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

// 四星-后四组选6
func convert_4x_h4zuxuan6(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	k1balls := validator.GetStringBalls(sch, K1, 1)
	k2balls := validator.GetStringBalls(sch, K2, 1)

	typeName := "单式"
	if len(k1balls) > 1 || len(k2balls) > 1 {
		typeName = "复式"
	}

	return zhongfu.Scheme{
		Type:   typeName,
		TypeId: fmt.Sprint(zfTypeId),
		Money:  fmt.Sprint(sch.GetMoney()),
		Num:    fmt.Sprint(sch.GetNum()),
		Balls:  zhongfu.MakeBalls2(ballsFormat, " ", k1balls, k2balls),
	}
}

func validate_4x_h4zuxuan6(sch *apiproto.BuycaiScheme) bool {
	k1balls, ok := validator.CheckBalls(sch, K1, 1, 10, MIN, MAX)
	if !ok {
		return false
	}
	k2balls, ok := validator.CheckBalls(sch, K2, 1, 10, MIN, MAX)
	if !ok {
		return false
	}

	interNum := validator.CalInter2Num(k1balls, k2balls)
	x := (len(k1balls) - interNum) * len(k2balls)
	y := interNum * (len(k2balls) - 1)
	num := x + y
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

// 四星-后四组选4
func convert_4x_h4zuxuan4(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	k1balls := validator.GetStringBalls(sch, K1, 1)
	k2balls := validator.GetStringBalls(sch, K2, 1)

	typeName := "单式"
	if len(k1balls) > 1 || len(k2balls) > 1 {
		typeName = "复式"
	}

	return zhongfu.Scheme{
		Type:   typeName,
		TypeId: fmt.Sprint(zfTypeId),
		Money:  fmt.Sprint(sch.GetMoney()),
		Num:    fmt.Sprint(sch.GetNum()),
		Balls:  zhongfu.MakeBalls2(ballsFormat, " ", k1balls, k2balls),
	}
}

func validate_4x_h4zuxuan4(sch *apiproto.BuycaiScheme) bool {
	k1balls, ok := validator.CheckBalls(sch, K1, 1, 10, MIN, MAX)
	if !ok {
		return false
	}
	k2balls, ok := validator.CheckBalls(sch, K2, 1, 10, MIN, MAX)
	if !ok {
		return false
	}

	interNum := validator.CalInter2Num(k1balls, k2balls)
	x := (len(k1balls) - interNum) * len(k2balls)
	y := interNum * (len(k2balls) - 1)
	num := x + y
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
