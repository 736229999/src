package fc3d

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/caojunxyz/mimi-server/lottery/saleissue"
	"github.com/caojunxyz/mimi-server/lottery/validator"
)

var PlayList = map[string]validator.PlayType{
	"zx": validator.PlayType{601, validate_zx, convert_zx, "%s,%s,%s"}, //直选
	"z3": validator.PlayType{602, validate_z3, convert_z3, "%s"},       //组三
	"z6": validator.PlayType{603, validate_z6, convert_z6, "%s"},       //组六
}

const (
	K     = "k"
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
