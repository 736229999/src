package cqssc

import (
	"fmt"
	"log"
	"strings"

	"github.com/caojunxyz/mimi-api/proto"
	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/lottery/site/zhongfu"
	"github.com/caojunxyz/mimi-server/lottery/validator"
)

// 一星定位胆
func convert_1x_dwd(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	wanBalls := validator.GetStringBalls(sch, K_WAN, 1)
	qianBalls := validator.GetStringBalls(sch, K_QIAN, 1)
	baiBalls := validator.GetStringBalls(sch, K_BAI, 1)
	shiBalls := validator.GetStringBalls(sch, K_SHI, 1)
	geBalls := validator.GetStringBalls(sch, K_GE, 1)

	var typeName string
	wannum := len(wanBalls)
	qiannum := len(qianBalls)
	bainum := len(baiBalls)
	shinum := len(shiBalls)
	genum := len(geBalls)
	wan := "-"
	qian := "-"
	bai := "-"
	shi := "-"
	ge := "-"
	if wannum+qiannum+bainum+shinum+genum == 1 {
		typeName = "单式"
	} else {
		typeName = "复式"
	}

	if wannum != 0 {
		wan = fmt.Sprintf("(%s)", strings.Join(wanBalls, ""))
	}
	if qiannum != 0 {
		qian = fmt.Sprintf("(%s)", strings.Join(qianBalls, ""))
	}
	if bainum != 0 {
		bai = fmt.Sprintf("(%s)", strings.Join(baiBalls, ""))
	}
	if shinum != 0 {
		shi = fmt.Sprintf("(%s)", strings.Join(shiBalls, ""))
	}
	if genum != 0 {
		ge = fmt.Sprintf("(%s)", strings.Join(geBalls, ""))
	}

	return zhongfu.Scheme{
		Type:   typeName,
		TypeId: fmt.Sprint(zfTypeId),
		Money:  fmt.Sprint(sch.GetMoney()),
		Num:    fmt.Sprint(sch.GetNum()),
		Balls:  fmt.Sprintf("%s%s%s%s%s", wan, qian, bai, shi, ge),
	}
}

func validate_1x_dwd(sch *apiproto.BuycaiScheme) bool {
	wanBalls, _ := validator.CheckBalls(sch, K_WAN, 0, 10, MIN, MAX)
	qianBalls, _ := validator.CheckBalls(sch, K_QIAN, 0, 10, MIN, MAX)
	baiBalls, _ := validator.CheckBalls(sch, K_BAI, 0, 10, MIN, MAX)
	shiBalls, _ := validator.CheckBalls(sch, K_SHI, 0, 10, MIN, MAX)
	geBalls, _ := validator.CheckBalls(sch, K_GE, 0, 10, MIN, MAX)

	numWan := len(wanBalls)
	numQian := len(qianBalls)
	numBai := len(baiBalls)
	numShi := len(shiBalls)
	numGe := len(geBalls)

	num := int32(numWan + numQian + numBai + numShi + numGe)
	if num == 0 {
		log.Println(num)
		return false
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

// 一星直选
func convert_1x_zhixuan(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	geBalls := validator.GetStringBalls(sch, K_GE, 1)

	var typename string = "单式"
	if len(geBalls) > 1 {
		typename = "复式"

	}

	return zhongfu.Scheme{
		Type:   typename,
		TypeId: fmt.Sprint(zfTypeId),
		Money:  fmt.Sprint(sch.GetMoney()),
		Num:    fmt.Sprint(sch.GetNum()),
		Balls:  zhongfu.MakeBalls2(ballsFormat, "", geBalls),
	}
}

func validate_1x_zhixuan(sch *apiproto.BuycaiScheme) bool {
	balls, ok := validator.CheckBalls(sch, K_GE, 1, 10, MIN, MAX)
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
