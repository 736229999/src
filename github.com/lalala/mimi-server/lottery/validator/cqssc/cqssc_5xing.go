package cqssc

import (
	"fmt"
	"log"

	"github.com/caojunxyz/mimi-api/proto"
	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/lottery/site/zhongfu"
	"github.com/caojunxyz/mimi-server/lottery/validator"
)

func _convert_5x(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	wanBalls := validator.GetStringBalls(sch, K_WAN, 1)
	qianBalls := validator.GetStringBalls(sch, K_QIAN, 1)
	baiBalls := validator.GetStringBalls(sch, K_BAI, 1)
	shiBalls := validator.GetStringBalls(sch, K_SHI, 1)
	geBalls := validator.GetStringBalls(sch, K_GE, 1)

	typeName := "单式"
	if len(wanBalls) > 1 || len(qianBalls) > 1 || len(baiBalls) > 1 || len(shiBalls) > 1 || len(geBalls) > 1 {
		typeName = "复式"
	}

	if len(wanBalls) == 1 && len(qianBalls) == 1 && len(baiBalls) == 1 && len(shiBalls) == 1 && len(geBalls) == 1 {
		ballsFormat = "%s%s%s%s%s"
	} else {
		ballsFormat = "(%s)(%s)(%s)(%s)(%s)"
	}

	return zhongfu.Scheme{
		Type:   typeName,
		TypeId: fmt.Sprint(zfTypeId),
		Money:  fmt.Sprint(sch.GetMoney()),
		Num:    fmt.Sprint(sch.GetNum()),
		Balls:  zhongfu.MakeBalls2(ballsFormat, "", wanBalls, qianBalls, baiBalls, shiBalls, geBalls),
	}
}

func _validate_5x(sch *apiproto.BuycaiScheme) bool {
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
	geBalls, ok := validator.CheckBalls(sch, K_GE, 1, 10, MIN, MAX)
	if !ok {
		return false
	}

	numWan := len(wanBalls)
	numQian := len(qianBalls)
	numBai := len(baiBalls)
	numShi := len(shiBalls)
	numGe := len(geBalls)
	num := int32(numWan * numQian * numBai * numShi * numGe)

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

// 五星直选
func convert_5x_zhixuan(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	return _convert_5x(sch, zfTypeId, ballsFormat)
}

func validate_5x_zhixuan(sch *apiproto.BuycaiScheme) bool {
	return _validate_5x(sch)
}

// 五星通选
func convert_5x_tongxuan(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	return _convert_5x(sch, zfTypeId, ballsFormat)
}

func validate_5x_tongxuan(sch *apiproto.BuycaiScheme) bool {
	return _validate_5x(sch)
}

// 五星组选120
func convert_5x_zuxuan120(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	balls := validator.GetStringBalls(sch, K1, 1)

	typeName := "单式"
	if len(balls) > 5 {
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

func validate_5x_zuxuan120(sch *apiproto.BuycaiScheme) bool {
	balls, ok := validator.CheckBalls(sch, K1, 5, 10, MIN, MAX)
	if !ok {
		return false
	}
	numballs := len(balls)

	num := int32(validator.Comb(numballs, 5))
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

// 五星组选60
func convert_5x_zuxuan60(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	k1balls := validator.GetStringBalls(sch, K1, 1)
	k2balls := validator.GetStringBalls(sch, K2, 1)

	typeName := "单式"
	if len(k1balls) > 1 || len(k2balls) > 3 {
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

func validate_5x_zuxuan60(sch *apiproto.BuycaiScheme) bool {
	k1balls, ok := validator.CheckBalls(sch, K1, 1, 10, MIN, MAX)
	if !ok {
		return false
	}
	k2balls, ok := validator.CheckBalls(sch, K2, 3, 10, MIN, MAX)
	if !ok {
		return false
	}

	interNum := validator.CalInter2Num(k1balls, k2balls)
	x := (len(k1balls) - interNum) * validator.Comb(len(k2balls), 3)
	y := interNum * validator.Comb(len(k2balls)-1, 3)
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

// 五星组选30
func convert_5x_zuxuan30(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	k1balls := validator.GetStringBalls(sch, K1, 1)
	k2balls := validator.GetStringBalls(sch, K2, 1)

	typeName := "单式"
	if len(k1balls) > 2 || len(k2balls) > 1 {
		typeName = "复式"
	}

	return zhongfu.Scheme{
		Type:   typeName,
		TypeId: fmt.Sprint(zfTypeId),
		Money:  fmt.Sprint(sch.GetMoney()),
		Num:    fmt.Sprint(sch.GetNum()),
		Balls:  zhongfu.MakeBalls2(ballsFormat, "", k1balls, k2balls),
	}
}

func validate_5x_zuxuan30(sch *apiproto.BuycaiScheme) bool {
	k1balls, ok := validator.CheckBalls(sch, K1, 2, 10, MIN, MAX)
	if !ok {
		return false
	}
	k2balls, ok := validator.CheckBalls(sch, K2, 1, 10, MIN, MAX)
	if !ok {
		return false
	}

	interNum := validator.CalInter2Num(k1balls, k2balls)
	x := (len(k2balls) - interNum) * validator.Comb(len(k1balls), 2)
	y := interNum * validator.Comb(len(k1balls)-1, 2)
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

// 五星组选20
func convert_5x_zuxuan20(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
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

func validate_5x_zuxuan20(sch *apiproto.BuycaiScheme) bool {
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

// 五星组选10
func convert_5x_zuxuan10(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
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

func validate_5x_zuxuan10(sch *apiproto.BuycaiScheme) bool {
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

// 五星组选5
func convert_5x_zuxuan5(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
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

func validate_5x_zuxuan5(sch *apiproto.BuycaiScheme) bool {
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
