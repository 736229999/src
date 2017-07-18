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

// 任选三
func convert_rx3(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
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
	if wannum+qiannum+bainum+shinum+genum == 3 {
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

func validate_rx3(sch *apiproto.BuycaiScheme) bool {
	wanBalls, _ := validator.CheckBalls(sch, K_WAN, 0, 10, MIN, MAX)
	qianBalls, _ := validator.CheckBalls(sch, K_QIAN, 0, 10, MIN, MAX)
	baiBalls, _ := validator.CheckBalls(sch, K_BAI, 0, 10, MIN, MAX)
	shiBalls, _ := validator.CheckBalls(sch, K_SHI, 0, 10, MIN, MAX)
	geBalls, _ := validator.CheckBalls(sch, K_GE, 0, 10, MIN, MAX)

	pos := 0
	if len(wanBalls) > 0 {
		pos++
	}
	if len(qianBalls) > 0 {
		pos++
	}
	if len(baiBalls) > 0 {
		pos++
	}
	if len(shiBalls) > 0 {
		pos++
	}
	if len(geBalls) > 0 {
		pos++
	}

	if pos < 3 {
		log.Println(len(wanBalls), len(qianBalls), len(baiBalls), len(shiBalls), len(geBalls))
		return false
	}

	//验证
	num := 0
	tmp := [5][]int32{wanBalls, qianBalls, baiBalls, shiBalls, geBalls}

	for i := 0; i < 3; i++ {
		for j := i + 1; j < 4; j++ {
			for k := j + 1; k < 5; k++ {
				for range tmp[i] {
					for range tmp[j] {
						for range tmp[k] {
							num++
						}
					}
				}
			}
		}
	}

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

// 任选二
func convert_rx2(sch *dbproto.BuycaiScheme, zfTypeId int, ballsFormat string) zhongfu.Scheme {
	wanBalls := validator.GetStringBalls(sch, K_WAN, 1)
	qianBalls := validator.GetStringBalls(sch, K_QIAN, 1)
	baiBalls := validator.GetStringBalls(sch, K_BAI, 1)
	shiBalls := validator.GetStringBalls(sch, K_SHI, 1)
	geBalls := validator.GetStringBalls(sch, K_GE, 1)

	var typeName string
	var Bvalue string
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
	if wannum+qiannum+bainum+shinum+genum == 3 {
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

	Bvalue = fmt.Sprintf("%s%s%s%s%s", wan, qian, bai, shi, ge)

	return zhongfu.Scheme{
		Type:   typeName,
		TypeId: fmt.Sprint(zfTypeId),
		Money:  fmt.Sprint(sch.GetMoney()),
		Num:    fmt.Sprint(sch.GetNum()),
		Balls:  Bvalue,
	}
}

func validate_rx2(sch *apiproto.BuycaiScheme) bool {
	wanBalls, _ := validator.CheckBalls(sch, K_WAN, 0, 10, MIN, MAX)
	qianBalls, _ := validator.CheckBalls(sch, K_QIAN, 0, 10, MIN, MAX)
	baiBalls, _ := validator.CheckBalls(sch, K_BAI, 0, 10, MIN, MAX)
	shiBalls, _ := validator.CheckBalls(sch, K_SHI, 0, 10, MIN, MAX)
	geBalls, _ := validator.CheckBalls(sch, K_GE, 0, 10, MIN, MAX)

	pos := 0
	if len(wanBalls) > 0 {
		pos++
	}
	if len(qianBalls) > 0 {
		pos++
	}
	if len(baiBalls) > 0 {
		pos++
	}
	if len(shiBalls) > 0 {
		pos++
	}
	if len(geBalls) > 0 {
		pos++
	}

	if pos < 2 {
		log.Println(len(wanBalls), len(qianBalls), len(baiBalls), len(shiBalls), len(geBalls))
		return false
	}

	//验证
	num := 0
	tmp := [5][]int32{wanBalls, qianBalls, baiBalls, shiBalls, geBalls}

	for i := 0; i < 4; i++ {
		for j := i + 1; j < 5; j++ {
			for range tmp[i] {
				for range tmp[j] {
					num++
				}
			}
		}
	}

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
