package gd11x5

import (
	"fmt"
	"log"

	"github.com/caojunxyz/mimi-api/proto"
	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/lottery/site/zhongfu"
	"github.com/caojunxyz/mimi-server/lottery/validator"
)

// n: 普通任选n（1 ～ 8）
func _convert_ptrx(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string, n int) zhongfu.Scheme {
	balls := validator.GetStringBalls(sch, K1, 2)
	typeName := "单式"
	if len(balls) > n {
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

func _validate_ptrx(sch *apiproto.BuycaiScheme, n int) bool {
	balls, ok := validator.CheckBalls(sch, K1, int32(n), 11, MIN, MAX)
	if !ok {
		return false
	}

	num := validator.Comb(len(balls), n)
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
func convert_ptrx2(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	return _convert_ptrx(sch, zfTypeId, ballsFormat, 2)
}

func validate_ptrx2(sch *apiproto.BuycaiScheme) bool {
	return _validate_ptrx(sch, 2)
}

//--------------------------------------------------------------------------------------------------------
func convert_ptrx3(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	return _convert_ptrx(sch, zfTypeId, ballsFormat, 3)
}

func validate_ptrx3(sch *apiproto.BuycaiScheme) bool {
	return _validate_ptrx(sch, 3)
}

//--------------------------------------------------------------------------------------------------------
func convert_ptrx4(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	return _convert_ptrx(sch, zfTypeId, ballsFormat, 4)
}

func validate_ptrx4(sch *apiproto.BuycaiScheme) bool {
	return _validate_ptrx(sch, 4)
}

//--------------------------------------------------------------------------------------------------------
func convert_ptrx5(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	return _convert_ptrx(sch, zfTypeId, ballsFormat, 5)
}

func validate_ptrx5(sch *apiproto.BuycaiScheme) bool {
	return _validate_ptrx(sch, 5)
}

//--------------------------------------------------------------------------------------------------------
func convert_ptrx6(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	return _convert_ptrx(sch, zfTypeId, ballsFormat, 6)
}

func validate_ptrx6(sch *apiproto.BuycaiScheme) bool {
	return _validate_ptrx(sch, 6)
}

//--------------------------------------------------------------------------------------------------------
func convert_ptrx7(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	return _convert_ptrx(sch, zfTypeId, ballsFormat, 7)
}

func validate_ptrx7(sch *apiproto.BuycaiScheme) bool {
	return _validate_ptrx(sch, 7)
}

//--------------------------------------------------------------------------------------------------------
func convert_ptrx8(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	return _convert_ptrx(sch, zfTypeId, ballsFormat, 8)
}

func validate_ptrx8(sch *apiproto.BuycaiScheme) bool {
	return _validate_ptrx(sch, 8)
}
