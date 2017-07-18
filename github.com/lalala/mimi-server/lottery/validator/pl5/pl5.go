package pl5

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/caojunxyz/mimi-server/lottery/saleissue"
	"github.com/caojunxyz/mimi-server/lottery/validator"
)

var PlayList = map[string]validator.PlayType{
	"pt": validator.PlayType{6401, validate_pt, convert_pt, ""}, // 普通投注
}

const (
	K_WAN  = "wan"
	K_QIAN = "qian"
	K_BAI  = "bai"
	K_SHI  = "shi"
	K_GE   = "ge"
)

const (
	MIN = 0
	MAX = 9
)

// 生成销售期号列表
// 期号格式: 2017011
func MakeSaleIssueList(prevIssue saleissue.SaleIssue, num int) []saleissue.SaleIssue {
	result := []saleissue.SaleIssue{}
	no, err := strconv.Atoi(prevIssue.Issue[4:])
	if err != nil {
		log.Println(err)
		return nil
	}

	startTime := time.Unix(prevIssue.StartTime, 0)
	endTime := time.Unix(prevIssue.EndTime, 0)
	openTime := time.Unix(prevIssue.OpenTime, 0)
	for i := 0; i < num; i++ {
		no++
		startTime = startTime.AddDate(0, 0, 1)
		endTime = endTime.AddDate(0, 0, 1)
		prevYear, _, _ := openTime.Date()
		openTime = openTime.AddDate(0, 0, 1)
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
