package dlt

import (
	"log"
	"strconv"
	"time"

	"fmt"

	"github.com/caojunxyz/mimi-server/lottery/saleissue"
	"github.com/caojunxyz/mimi-server/lottery/validator"
)

var PlayList = map[string]validator.PlayType{
	"pt":   validator.PlayType{3901, validate_pt, convert_pt, "%s # %s"},
	"qqdt": validator.PlayType{3903, validate_qqdt, convert_qqdt, "%s , %s # %s"},
	"hqdt": validator.PlayType{3906, validate_hqdt, convert_hqdt, "%s # %s , %s"},
	"sqdt": validator.PlayType{3907, validate_sqdt, convert_sqdt, "%s , %s # %s , %s"},
}

const (
	K_QD   = "qian-dan"
	K_QT   = "qian-tuo"
	K_HD   = "hou-dan"
	K_HT   = "hou-tuo"
	K_QIAN = "qian"
	K_HOU  = "hou"
)

const (
	MIN_RED  = 1
	MAX_RED  = 35
	MIN_BLUE = 1
	MAX_BLUE = 12
)

var weekDaysTable = map[time.Weekday]int{
	time.Monday:    2,
	time.Saturday:  2,
	time.Wednesday: 3,
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
