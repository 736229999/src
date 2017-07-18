package bjpk10

import (
	"fmt"
	"log"

	"github.com/caojunxyz/mimi-api/proto"
	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/lottery/site/zhongfu"
	"github.com/caojunxyz/mimi-server/lottery/validator"
)

func convert_cjo(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	balls := validator.GetStringBalls(sch, K1, 1)
	var v string
	if balls[0] == "0" { // 与客户端显示顺序对应
		v = "全奇"
	} else {
		v = "全偶"
	}
	return zhongfu.Scheme{
		Type:   "猜奇偶",
		TypeId: fmt.Sprint(zfTypeId),
		Money:  fmt.Sprint(sch.GetMoney()),
		Num:    fmt.Sprint(sch.GetNum()),
		Balls:  v,
	}
}

func validate_cjo(sch *apiproto.BuycaiScheme) bool {
	_, ok := validator.CheckBalls(sch, K1, 1, 1, 0, 1)
	if !ok {
		return false
	}

	if sch.GetNum() != 1 {
		log.Println("注数计算不一致:", sch.GetNum())
		return false
	}

	sumMoney := float64(2)
	if !validator.IsEqualMoney(sumMoney, sch.GetMoney()) {
		log.Println("金额计算不一致:", sumMoney, sch.GetMoney())
		return false
	}
	return true
}
