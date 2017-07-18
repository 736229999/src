package validator

import (
	"fmt"
	"log"
	"time"

	apiproto "github.com/caojunxyz/mimi-api/proto"
	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/lottery/site/zhongfu"
)

var DefaultLocal *time.Location
var TIME_FORMAT = "2006-01-02 15:04:05"

func init() {
	loc, err := time.LoadLocation("Asia/Chongqing")
	if err != nil {
		log.Panic(err)
	}
	DefaultLocal = loc
}

type SchemeValidator func(*apiproto.BuycaiScheme) bool
type ZfSchemeConverter func(*dbproto.BuycaiScheme, int, string) zhongfu.Scheme

type PlayType struct {
	ZfId        int
	Validator   SchemeValidator
	ZfConverter ZfSchemeConverter
	BallsFormat string
}

func CheckBalls(scheme *apiproto.BuycaiScheme, key string, minNum, maxNum, minVal, maxVal int32) ([]int32, bool) {
	balls, ok := scheme.GetSelectBalls()[key]
	if !ok {
		if minNum > 0 {
			log.Printf("没有'%s'的号码\n", key)
		}
		return nil, false
	}

	num := int32(len(balls.List))
	if num < minNum || num > maxNum {
		log.Printf("'%s'数量超出范围: %d [%d, %d]\n", key, num, minNum, maxNum)
		return nil, false
	}
	for _, v := range balls.List {
		if v < minVal || v > maxVal {
			log.Printf("'%s'值超出范围: %d [%d, %d]\n", key, v, minVal, maxVal)
			return nil, false
		}
	}
	return balls.List, true
}

// 有重复号码返回false，否则返回true
func CheckDuplicate(balls1, balls2 []int32) bool {
	for _, v1 := range balls1 {
		for _, v2 := range balls2 {
			if v1 == v2 {
				log.Printf("重复号码: %d\n", v1)
				return false
			}
		}
	}
	return true
}

func GetStringBalls(scheme *dbproto.BuycaiScheme, key string, width int) []string {
	balls, ok := scheme.GetSelectBalls()[key]
	if !ok {
		return nil
	}
	ret := []string{}
	format := fmt.Sprintf("%%0%dd", width)
	for _, v := range balls.List {
		ret = append(ret, fmt.Sprintf(format, v))
	}
	return ret
}
