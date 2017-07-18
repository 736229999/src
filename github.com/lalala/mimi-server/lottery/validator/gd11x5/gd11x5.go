package gd11x5

import (
	"fmt"
	"log"
	"time"

	"github.com/caojunxyz/mimi-server/lottery/saleissue"
	"github.com/caojunxyz/mimi-server/lottery/validator"
)

const DAY_MAX_NO = 84

var PlayList = map[string]validator.PlayType{
	"pt-rx2": validator.PlayType{7802, validate_ptrx2, convert_ptrx2, "%s"},      //任选二-普通
	"pt-rx3": validator.PlayType{7803, validate_ptrx3, convert_ptrx3, "%s"},      //任选三-普通
	"pt-rx4": validator.PlayType{7804, validate_ptrx4, convert_ptrx4, "%s"},      //任选四-普通
	"pt-rx5": validator.PlayType{7805, validate_ptrx5, convert_ptrx5, "%s"},      //任选五-普通
	"pt-rx6": validator.PlayType{7806, validate_ptrx6, convert_ptrx6, "%s"},      //任选六-普通
	"pt-rx7": validator.PlayType{7807, validate_ptrx7, convert_ptrx7, "%s"},      //任选七-普通
	"pt-rx8": validator.PlayType{7808, validate_ptrx8, convert_ptrx8, "%s"},      //任选八-普通
	"pt-q1":  validator.PlayType{7801, validate_ptq1, convert_ptq1, "%s"},        //前一-普通
	"pt-q2":  validator.PlayType{7809, validate_ptq2, convert_ptq2, "%s,%s"},     //前二-普通
	"pt-q3":  validator.PlayType{7810, validate_ptq3, convert_ptq3, "%s,%s,%s"},  //前三-普通
	"zx-q2":  validator.PlayType{7811, validate_zxq2, convert_zxq2, "%s"},        //前二-组选
	"zx-q3":  validator.PlayType{7812, validate_zxq3, convert_zxq3, "%s"},        // 前三-组选
	"dt-q2":  validator.PlayType{7813, validate_dtq2, convert_dtq2, "%s , %s"},   //前二-胆拖
	"dt-q3":  validator.PlayType{7814, validate_dtq3, convert_dtq3, "%s , %s"},   //前三-胆拖
	"dt-rx2": validator.PlayType{7815, validate_dtrx2, convert_dtrx2, "%s , %s"}, //任选二-胆拖
	"dt-rx3": validator.PlayType{7816, validate_dtrx3, convert_dtrx3, "%s , %s"}, //任选三-胆拖
	"dt-rx4": validator.PlayType{7817, validate_dtrx4, convert_dtrx4, "%s , %s"}, //任选四-胆拖
	"dt-rx5": validator.PlayType{7818, validate_dtrx5, convert_dtrx5, "%s , %s"}, //任选五-胆拖
	"dt-rx6": validator.PlayType{7819, validate_dtrx6, convert_dtrx6, "%s , %s"}, //任选六-胆拖
	"dt-rx7": validator.PlayType{7820, validate_dtrx7, convert_dtrx7, "%s , %s"}, //任选七-胆拖
	"dt-rx8": validator.PlayType{7821, validate_dtrx8, convert_dtrx8, "%s , %s"}, //任选八-胆拖
}

const (
	K1    = "k1"
	K2    = "k2"
	K3    = "k3"
	K_DAN = "dan"
	K_TUO = "tuo"
)

const (
	MIN = 1
	MAX = 11
)

// 生成销售期号列表
// prevIssue: 头一天最后一期
// days: 生成天数
func MakeSaleIssueList(prevIssue saleissue.SaleIssue, dayNum int) []saleissue.SaleIssue {
	if dayNum <= 0 {
		log.Println("天数无效:", dayNum)
		return nil
	}

	result := []saleissue.SaleIssue{}
	interval := 10 * time.Minute
	loc := validator.DefaultLocal
	var startTime, endTime, openTime time.Time
	startTime = time.Unix(prevIssue.StartTime, 0)
	for i := 0; i < dayNum; i++ {
		year, month, day := startTime.AddDate(0, 0, 1).Date()
		startTime = time.Date(year, month, day, 9, 0, 0, 0, loc)
		endTime = startTime.Add(interval).Add(-time.Minute)
		openTime = startTime.Add(interval).Add(time.Second * 50)
		for no := 1; no <= DAY_MAX_NO; no++ {
			si := saleissue.SaleIssue{
				Issue:     fmt.Sprintf("%02d%02d%02d%02d", year-2000, month, day, no),
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
