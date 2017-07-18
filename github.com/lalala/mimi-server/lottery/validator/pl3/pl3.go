package pl3

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/caojunxyz/mimi-server/lottery/saleissue"
	"github.com/caojunxyz/mimi-server/lottery/validator"
)

var PlayList = map[string]validator.PlayType{
	"zhixtz": validator.PlayType{6301, validate_zhixtz, convert_zhixtz, ""},   //直选投注
	"zuxds":  validator.PlayType{6302, validate_zuxds, convert_zuxds, "%s"},   //组选单式
	"zu3fs":  validator.PlayType{6304, validate_zu3fs, convert_zu3fs, "%s"},   //组三复式
	"zu6fs":  validator.PlayType{6303, validate_zu6fs, convert_zu6fs, "%s"},   //组六复式
	"zuxhz":  validator.PlayType{6306, validate_zuxhz, convert_zuxhz, "%s"},   //组选和值
	"zhixhz": validator.PlayType{6305, validate_zhixhz, convert_zhixhz, "%s"}, //直选和值
}

const (
	K     = "k"
	K1    = "k1"
	K2    = "k2"
	K_BAI = "bai"
	K_SHI = "shi"
	K_GE  = "ge"
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
