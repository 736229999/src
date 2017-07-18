package cqssc

import (
	"fmt"
	"log"

	"github.com/caojunxyz/mimi-api/proto"
	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/lottery/site/zhongfu"
	"github.com/caojunxyz/mimi-server/lottery/validator"
)

var _dxds_values = map[string]string{
	"0": "大",
	"1": "小",
	"2": "单",
	"3": "双",
}

// 前二大大小单双
func convert_q2ddxds(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	wanBalls := validator.GetStringBalls(sch, K_WAN, 1)
	qianBalls := validator.GetStringBalls(sch, K_QIAN, 1)

	wan := _dxds_values[wanBalls[0]]
	qian := _dxds_values[qianBalls[0]]

	return zhongfu.Scheme{
		Type:   "单式",
		TypeId: fmt.Sprint(zfTypeId),
		Money:  fmt.Sprint(sch.GetMoney()),
		Num:    fmt.Sprint(sch.GetNum()),
		Balls:  fmt.Sprintf("%s%s", wan, qian),
	}
}

func validate_q2ddxds(sch *apiproto.BuycaiScheme) bool {
	_, ok := validator.CheckBalls(sch, K_WAN, 1, 1, MIN, MAX)
	if !ok {
		return false
	}
	_, ok = validator.CheckBalls(sch, K_QIAN, 1, 1, MIN, MAX)
	if !ok {
		return false
	}

	num := int32(1)
	if num != sch.GetNum() {
		log.Println("注数计算不一致:", num, sch.GetNum())
		return false
	}

	sumMoney := float64(2)
	if !validator.IsEqualMoney(sumMoney, sch.GetMoney()) {
		log.Println("金额计算不一致:", sumMoney, sch.GetMoney())
		return false
	}

	return true
}

// 后二大大小单双
func convert_h2ddxds(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	shiBalls := validator.GetStringBalls(sch, K_SHI, 1)
	geBalls := validator.GetStringBalls(sch, K_GE, 1)

	shi := _dxds_values[shiBalls[0]]
	ge := _dxds_values[geBalls[0]]

	return zhongfu.Scheme{
		Type:   "单式",
		TypeId: fmt.Sprint(zfTypeId),
		Money:  fmt.Sprint(sch.GetMoney()),
		Num:    fmt.Sprint(sch.GetNum()),
		Balls:  fmt.Sprintf("%s%s", shi, ge),
	}
}

func validate_h2ddxds(sch *apiproto.BuycaiScheme) bool {
	_, ok := validator.CheckBalls(sch, K_SHI, 1, 1, MIN, MAX)
	if !ok {
		return false
	}
	_, ok = validator.CheckBalls(sch, K_GE, 1, 1, MIN, MAX)
	if !ok {
		return false
	}

	num := int32(1)
	if num != sch.GetNum() {
		log.Println("注数计算不一致:", num, sch.GetNum())
		return false
	}

	sumMoney := float64(2)
	if !validator.IsEqualMoney(sumMoney, sch.GetMoney()) {
		log.Println("金额计算不一致:", sumMoney, sch.GetMoney())
		return false
	}

	return true
}
