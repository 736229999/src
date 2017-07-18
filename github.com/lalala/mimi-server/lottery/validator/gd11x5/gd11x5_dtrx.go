package gd11x5

import (
	"fmt"
	"log"

	"github.com/caojunxyz/mimi-api/proto"
	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/lottery/site/zhongfu"
	"github.com/caojunxyz/mimi-server/lottery/validator"
)

// n: 胆拖任选n（1 ～ 8）
func _convert_dtrx(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string, n int) zhongfu.Scheme {
	danBalls := validator.GetStringBalls(sch, K_DAN, 2)
	tuoBalls := validator.GetStringBalls(sch, K_TUO, 2)
	return zhongfu.Scheme{
		Type:   "胆拖",
		TypeId: fmt.Sprint(zfTypeId),
		Money:  fmt.Sprint(sch.GetMoney()),
		Num:    fmt.Sprint(sch.GetNum()),
		Balls:  zhongfu.MakeBalls(ballsFormat, danBalls, tuoBalls),
	}
}

func _validate_dtrx(sch *apiproto.BuycaiScheme, n int) bool {
	danBalls, ok := validator.CheckBalls(sch, K_DAN, 1, int32(n-1), MIN, MAX)
	if !ok {
		return false
	}
	tuoBalls, ok := validator.CheckBalls(sch, K_TUO, 2, 10, MIN, MAX)
	if !ok {
		return false
	}

	numDan := len(danBalls)
	numTuo := len(tuoBalls)
	if numDan+numTuo > 11 {
		log.Println("数量错误:", numDan, numTuo)
		return false
	}

	num := validator.Comb(numTuo, n-numDan)
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

//--------------------------------------------------------------------------------------------------------
func convert_dtrx2(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	return _convert_dtrx(sch, zfTypeId, ballsFormat, 2)
}

func validate_dtrx2(sch *apiproto.BuycaiScheme) bool {
	return _validate_dtrx(sch, 2)
}

//--------------------------------------------------------------------------------------------------------
func convert_dtrx3(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	return _convert_dtrx(sch, zfTypeId, ballsFormat, 3)
}

func validate_dtrx3(sch *apiproto.BuycaiScheme) bool {
	return _validate_dtrx(sch, 3)
}

//--------------------------------------------------------------------------------------------------------
func convert_dtrx4(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	return _convert_dtrx(sch, zfTypeId, ballsFormat, 4)
}

func validate_dtrx4(sch *apiproto.BuycaiScheme) bool {
	return _validate_dtrx(sch, 4)
}

//--------------------------------------------------------------------------------------------------------
func convert_dtrx5(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	return _convert_dtrx(sch, zfTypeId, ballsFormat, 5)
}

func validate_dtrx5(sch *apiproto.BuycaiScheme) bool {
	return _validate_dtrx(sch, 5)
}

//--------------------------------------------------------------------------------------------------------
func convert_dtrx6(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	return _convert_dtrx(sch, zfTypeId, ballsFormat, 6)
}

func validate_dtrx6(sch *apiproto.BuycaiScheme) bool {
	return _validate_dtrx(sch, 6)
}

//--------------------------------------------------------------------------------------------------------
func convert_dtrx7(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	return _convert_dtrx(sch, zfTypeId, ballsFormat, 7)
}

func validate_dtrx7(sch *apiproto.BuycaiScheme) bool {
	return _validate_dtrx(sch, 7)
}

//--------------------------------------------------------------------------------------------------------
func convert_dtrx8(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	return _convert_dtrx(sch, zfTypeId, ballsFormat, 8)
}

func validate_dtrx8(sch *apiproto.BuycaiScheme) bool {
	return _validate_dtrx(sch, 8)
}
