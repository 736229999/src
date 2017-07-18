package bjpk10

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/caojunxyz/mimi-server/lottery/saleissue"
	"github.com/caojunxyz/mimi-server/lottery/validator"
)

const DAY_MAX_NO = 179

var PlayList = map[string]validator.PlayType{
	"pt-x1":  validator.PlayType{9401, validate_ptx1, convert_ptx1, "%s"},       // 普通-选1
	"pt-x2":  validator.PlayType{9402, validate_ptx2, convert_ptx2, "%s,%s"},    // 普通-选2
	"pt-x3":  validator.PlayType{9403, validate_ptx3, convert_ptx3, "%s,%s,%s"}, // 普通-选3
	"jq-x2":  validator.PlayType{9411, validate_jqx2, convert_jqx2, "%s,%s"},    // 精确-选2
	"jq-zx2": validator.PlayType{9414, validate_jqzx2, convert_jqzx2, "%s"},     // 精确-组选2
	"jq-zx3": validator.PlayType{9415, validate_jqzx3, convert_jqzx3, "%s"},     // 精确-组选3
	"jq-zx4": validator.PlayType{9416, validate_jqzx4, convert_jqzx4, "%s"},     // 精确-组选4
	"wz-1":   validator.PlayType{9417, validate_wz1, convert_wz1, "%s"},         // 位置-位置1
	"wz-2":   validator.PlayType{9418, validate_wz2, convert_wz2, "%s"},         // 位置-位置2
	"tlj":    validator.PlayType{9419, validate_tlj, convert_tlj, ""},           // 拖拉机
	"cjo":    validator.PlayType{9420, validate_cjo, convert_cjo, ""},           // 猜奇偶
	"cdx":    validator.PlayType{9421, validate_cdx, convert_cdx, ""},           // 猜大小
	"hs":     validator.PlayType{9422, validate_hs, convert_hs, "%s"},           // 和数
}

const (
	K1  = "k1"
	K2  = "k2"
	K3  = "k3"
	K4  = "k4"
	K5  = "k5"
	K6  = "k6"
	K7  = "k7"
	K8  = "k8"
	K9  = "k9"
	K10 = "k10"
)

const (
	MIN = 1
	MAX = 10
)

// 生成销售期号列表
// prevIssue: 头一天最后一期
// days: 生成天数
func MakeSaleIssueList(prevIssue saleissue.SaleIssue, dayNum int) []saleissue.SaleIssue {
	if dayNum <= 0 {
		log.Println("天数无效:", dayNum)
		return nil
	}

	prevIssueNo, err := strconv.Atoi(prevIssue.Issue)
	if err != nil {
		log.Println(err)
		return nil
	}

	result := []saleissue.SaleIssue{}
	interval := 5 * time.Minute
	loc := validator.DefaultLocal
	var startTime, endTime, openTime time.Time
	startTime = time.Unix(prevIssue.StartTime, 0)
	for i := 0; i < dayNum; i++ {
		year, month, day := startTime.AddDate(0, 0, 1).Date()
		startTime = time.Date(year, month, day, 9, 2, 0, 0, loc)
		endTime = startTime.Add(interval).Add(-time.Minute)
		openTime = startTime.Add(interval).Add(time.Second * 50)
		for no := 1; no <= DAY_MAX_NO; no++ {
			si := saleissue.SaleIssue{
				Issue:     fmt.Sprint(prevIssueNo + i*DAY_MAX_NO + no),
				StartTime: startTime.Unix(),
				EndTime:   endTime.Unix(),
				OpenTime:  openTime.Unix(),
			}
			result = append(result, si)
			startTime = startTime.Add(interval)
			endTime = endTime.Add(interval)
			openTime = openTime.Add(interval)
		}
	}
	return result
}
