package ssq

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/caojunxyz/mimi-server/lottery/saleissue"
	"github.com/caojunxyz/mimi-server/lottery/validator"
)

var PlayList = map[string]validator.PlayType{
	"pt": validator.PlayType{501, validate_pt, convert_pt, "%s # %s"},      // 普通投注
	"dt": validator.PlayType{502, validate_dt, convert_dt, "%s , %s # %s"}, // 胆拖投注
}

const (
	K_RED  = "red"
	K_BLUE = "blue"
	K_DAN  = "dan"
	K_TUO  = "tuo"
)

const (
	MIN_RED  = 1
	MAX_RED  = 33
	MIN_BLUE = 1
	MAX_BLUE = 16
)

var weekDaysTable = map[time.Weekday]int{
	time.Tuesday:  2,
	time.Sunday:   2,
	time.Thursday: 3,
}

// 生成销售期号列表
// 期号格式: 2017011
func MakeSaleIssueList(prevIssue saleissue.SaleIssue, num int) []saleissue.SaleIssue {
	result := []saleissue.SaleIssue{}
	no, err := strconv.Atoi(prevIssue.Issue[4:])
	if err != nil {
		log.Println(err)
		return nil
	}

	var weekDay time.Weekday
	startTime := time.Unix(prevIssue.StartTime, 0)
	endTime := time.Unix(prevIssue.EndTime, 0)
	openTime := time.Unix(prevIssue.OpenTime, 0)
	days := 0
	for i := 0; i < num; i++ {
		no++
		weekDay = startTime.Weekday()
		days = weekDaysTable[weekDay]
		if days == 0 {
			log.Panicf("无效时间:%v", startTime)
		}
		startTime = startTime.AddDate(0, 0, days)

		weekDay = endTime.Weekday()
		days = weekDaysTable[weekDay]
		endTime = endTime.AddDate(0, 0, days)

		prevYear, _, _ := openTime.Date()
		openTime = openTime.AddDate(0, 0, days)
		year, _, _ := openTime.Date()
		if year > prevYear {
			no = 1
		}
		si := saleissue.SaleIssue{
			Issue:     fmt.Sprintf("%d%03d", year, no),
			StartTime: startTime.Unix(),
			EndTime:   endTime.Unix(),
			OpenTime:  openTime.Unix(),
		}
		result = append(result, si)
	}
	return result
}
