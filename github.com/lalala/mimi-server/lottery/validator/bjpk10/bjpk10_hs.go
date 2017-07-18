package bjpk10

import (
	"fmt"
	"log"
	"strings"

	"github.com/caojunxyz/mimi-api/proto"
	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/lottery/site/zhongfu"
	"github.com/caojunxyz/mimi-server/lottery/validator"
)

func convert_hs(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	balls := validator.GetStringBalls(sch, K1, 2)
	var v string
	if len(balls) > 1 {
		v = strings.Join(balls, ",")
	} else if len(balls) == 1 {
		v = fmt.Sprint(balls[0])
	}

	return zhongfu.Scheme{
		Type:   "和值",
		TypeId: fmt.Sprint(zfTypeId),
		Money:  fmt.Sprint(sch.GetMoney()),
		Num:    fmt.Sprint(sch.GetNum()),
		Balls:  v,
	}
}

func validate_hs(sch *apiproto.BuycaiScheme) bool {
	balls, ok := validator.CheckBalls(sch, K1, 1, 22, 6, 27)
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
