package bjpk10

import (
	"fmt"
	"log"

	"github.com/caojunxyz/mimi-api/proto"
	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/lottery/site/zhongfu"
	"github.com/caojunxyz/mimi-server/lottery/validator"
)

func convert_tlj(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	name := "拖拉机"
	return zhongfu.Scheme{
		Type:   name,
		TypeId: fmt.Sprint(zfTypeId),
		Money:  fmt.Sprint(sch.GetMoney()),
		Num:    fmt.Sprint(sch.GetNum()),
		Balls:  name,
	}
}

func validate_tlj(sch *apiproto.BuycaiScheme) bool {
	valTlj := int32(0)
	_, ok := validator.CheckBalls(sch, K1, 1, 1, valTlj, valTlj) // 与客户端显示顺序对应
	if !ok {
		return false
	}
	if sch.GetNum() != 1 {
		log.Println("注数计算不一致:", 1, sch.GetNum())
		return false
	}

	sumMoney := float64(2)
	if !validator.IsEqualMoney(sumMoney, sch.GetMoney()) {
		log.Println("金额计算不一致:", sumMoney, sch.GetMoney())
		return false
	}
	return true
}
