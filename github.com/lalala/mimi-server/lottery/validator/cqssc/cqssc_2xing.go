package cqssc

import (
	"fmt"
	"log"

	"github.com/caojunxyz/mimi-api/proto"
	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/lottery/site/zhongfu"
	"github.com/caojunxyz/mimi-server/lottery/validator"
)

// 前二直选
func convert_q2zhixuan(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	wanBalls := validator.GetStringBalls(sch, K_WAN, 1)
	qianBalls := validator.GetStringBalls(sch, K_QIAN, 1)

	var typeName string
	if len(wanBalls) == 1 && len(qianBalls) == 1 {
		ballsFormat = "%s%s---"
		typeName = "单式"
	} else {
		ballsFormat = "(%s)(%s)---"
		typeName = "复式"
	}

	return zhongfu.Scheme{
		Type:   typeName,
		TypeId: fmt.Sprint(zfTypeId),
		Money:  fmt.Sprint(sch.GetMoney()),
		Num:    fmt.Sprint(sch.GetNum()),
		Balls:  zhongfu.MakeBalls2(ballsFormat, "", wanBalls, qianBalls),
	}
}

func validate_q2zhixuan(sch *apiproto.BuycaiScheme) bool {
	wanBalls, ok := validator.CheckBalls(sch, K_WAN, 1, 10, MIN, MAX)
	if !ok {
		return false
	}
	qianBalls, ok := validator.CheckBalls(sch, K_QIAN, 1, 10, MIN, MAX)
	if !ok {
		return false
	}
	numWan := len(wanBalls)
	numQian := len(qianBalls)

	num := int32(numWan * numQian)

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

// 前二组选
func convert_q2zuxuan(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	balls := validator.GetStringBalls(sch, K1, 1)

	typeName := "单式"
	sep := ""
	zfTypeId = 2851
	if len(balls) > 2 {
		typeName = "复式"
		sep = " "
		zfTypeId = 2850
	}
	return zhongfu.Scheme{
		Type:   typeName,
		TypeId: fmt.Sprint(zfTypeId),
		Money:  fmt.Sprint(sch.GetMoney()),
		Num:    fmt.Sprint(sch.GetNum()),
		Balls:  zhongfu.MakeBalls2("%s", sep, balls),
	}
}

func validate_q2zuxuan(sch *apiproto.BuycaiScheme) bool {
	balls, ok := validator.CheckBalls(sch, K1, 2, 10, MIN, MAX)
	if !ok {
		return false
	}

	num := validator.Comb(len(balls), 2)
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

// 后二直选
func convert_h2zhixuan(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	shiBalls := validator.GetStringBalls(sch, K_SHI, 1)
	geBalls := validator.GetStringBalls(sch, K_GE, 1)

	var typeName string
	if len(shiBalls) == 1 && len(geBalls) == 1 {
		ballsFormat = "---%s%s"
		typeName = "单式"
	} else {
		ballsFormat = "---(%s)(%s)"
		typeName = "复式"
	}

	return zhongfu.Scheme{
		Type:   typeName,
		TypeId: fmt.Sprint(zfTypeId),
		Money:  fmt.Sprint(sch.GetMoney()),
		Num:    fmt.Sprint(sch.GetNum()),
		Balls:  zhongfu.MakeBalls2(ballsFormat, "", shiBalls, geBalls),
	}
}

func validate_h2zhixuan(sch *apiproto.BuycaiScheme) bool {
	shiBalls, ok := validator.CheckBalls(sch, K_SHI, 1, 10, MIN, MAX)
	if !ok {
		return false
	}
	geBalls, ok := validator.CheckBalls(sch, K_GE, 1, 10, MIN, MAX)
	if !ok {
		return false
	}
	numshi := len(shiBalls)
	numge := len(geBalls)
	num := int32(numshi * numge)

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

// 后二组选
func convert_h2zuxuan(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	balls := validator.GetStringBalls(sch, K1, 1)
	typeName := "单式"
	sep := ""
	zfTypeId = 2869
	if len(balls) > 2 {
		typeName = "复式"
		sep = " "
		zfTypeId = 2806
	}

	return zhongfu.Scheme{
		Type:   typeName,
		TypeId: fmt.Sprint(zfTypeId),
		Money:  fmt.Sprint(sch.GetMoney()),
		Num:    fmt.Sprint(sch.GetNum()),
		Balls:  zhongfu.MakeBalls2("%s", sep, balls),
	}
}

func validate_h2zuxuan(sch *apiproto.BuycaiScheme) bool {
	balls, ok := validator.CheckBalls(sch, K1, 2, 10, MIN, MAX)
	if !ok {
		return false
	}

	num := validator.Comb(len(balls), 2)
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

// 后二组选胆拖
func convert_h2zuxuan_dt(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	danBalls := validator.GetStringBalls(sch, K_DAN, 2)
	tuoBalls := validator.GetStringBalls(sch, K_TUO, 2)

	return zhongfu.Scheme{
		Type:   "胆拖",
		TypeId: fmt.Sprint(zfTypeId),
		Money:  fmt.Sprint(sch.GetMoney()),
		Num:    fmt.Sprint(sch.GetNum()),
		Balls:  zhongfu.MakeBalls2(ballsFormat, "", danBalls, tuoBalls),
	}
}

func validate_h2zuxuan_dt(sch *apiproto.BuycaiScheme) bool {
	danBalls, ok := validator.CheckBalls(sch, K_DAN, 1, 5, MIN, MAX)
	if !ok {
		return false
	}
	tuoBalls, ok := validator.CheckBalls(sch, K_TUO, 2, 33, MIN, MAX)
	if !ok {
		return false
	}
	numTuo := len(tuoBalls)

	if !validator.CheckDuplicate(danBalls, tuoBalls) {
		return false
	}

	if int32(numTuo) != sch.GetNum() {
		log.Println("注数计算不一致:", numTuo, sch.GetNum())
		return false
	}

	sumMoney := float64(numTuo * 2)
	if !validator.IsEqualMoney(sumMoney, sch.GetMoney()) {
		log.Println("金额计算不一致:", sumMoney, sch.GetMoney())
		return false
	}

	return true
}
