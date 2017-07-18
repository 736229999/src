package cqssc

import (
	"fmt"
	"log"
	"time"

	"github.com/caojunxyz/mimi-server/lottery/saleissue"
	"github.com/caojunxyz/mimi-server/lottery/validator"
)

const DAY_MAX_NO = 120

const (
	K_WAN  = "wan"
	K_QIAN = "qian"
	K_BAI  = "bai"
	K_SHI  = "shi"
	K_GE   = "ge"
	K_DAN  = "dan"
	K_TUO  = "tuo"
	K1     = "k1"
	K2     = "k2"
)

const (
	MIN = 0
	MAX = 9
)

var _zuxuan_hz = map[int32]int32{
	1:  1,
	2:  2,
	3:  2,
	4:  4,
	5:  5,
	6:  6,
	7:  8,
	8:  10,
	9:  11,
	10: 13,
	11: 14,
	12: 14,
	13: 15,
	14: 15,
	15: 14,
	16: 14,
	17: 13,
	18: 11,
	19: 10,
	20: 8,
	21: 6,
	22: 5,
	23: 4,
	24: 2,
	25: 2,
	26: 1,
}

var _zhixuan_hz = map[int32]int32{
	0:  1,
	1:  3,
	2:  6,
	3:  10,
	4:  15,
	5:  21,
	6:  28,
	7:  36,
	8:  45,
	9:  55,
	10: 63,
	11: 69,
	12: 73,
	13: 75,
	14: 75,
	15: 73,
	16: 69,
	17: 63,
	18: 55,
	19: 45,
	20: 36,
	21: 28,
	22: 21,
	23: 15,
	24: 10,
	25: 6,
	26: 3,
	27: 1,
}
var PlayList = map[string]validator.PlayType{
	"5x-zhixuan":     validator.PlayType{2803, validate_5x_zhixuan, convert_5x_zhixuan, ""},                      //五星直选
	"5x-tongxuan":    validator.PlayType{2805, validate_5x_tongxuan, convert_5x_tongxuan, ""},                    //五星通选
	"5x-zuxuan120":   validator.PlayType{2860, validate_5x_zuxuan120, convert_5x_zuxuan120, "%s"},                //五星组选120
	"5x-zuxuan60":    validator.PlayType{2861, validate_5x_zuxuan60, convert_5x_zuxuan60, "%s,%s"},               //五星组选60
	"5x-zuxuan30":    validator.PlayType{2862, validate_5x_zuxuan30, convert_5x_zuxuan30, "%s,%s"},               //五星组选30
	"5x-zuxuan20":    validator.PlayType{2863, validate_5x_zuxuan20, convert_5x_zuxuan20, "%s,%s"},               //五星组选20
	"5x-zuxuan10":    validator.PlayType{2864, validate_5x_zuxuan10, convert_5x_zuxuan10, "%s,%s"},               //五星组选10
	"5x-zuxuan5":     validator.PlayType{2865, validate_5x_zuxuan5, convert_5x_zuxuan5, "%s,%s"},                 //五星组选5
	"4x-q4zhixuan":   validator.PlayType{2819, validate_4x_q4zhixuan, convert_4x_q4zhixuan, "(%s)(%s)(%s)(%s)-"}, //前四直选
	"4x-h4zhixuan":   validator.PlayType{0, validate_4x_h4zhixuan, convert_4x_h4zhixuan, ""},                     //后四直选
	"4x-h4zuxuan24":  validator.PlayType{2822, validate_4x_h4zuxuan24, convert_4x_h4zuxuan24, "%s"},              //后四组选24
	"4x-h4zuxuan12":  validator.PlayType{2823, validate_4x_h4zuxuan12, convert_4x_h4zuxuan12, "%s,%s"},           //后四组选12
	"4x-h4zuxuan6":   validator.PlayType{2824, validate_4x_h4zuxuan6, convert_4x_h4zuxuan6, "%s,%s"},             //后四组选6
	"4x-h4zuxuan4":   validator.PlayType{2825, validate_4x_h4zuxuan4, convert_4x_h4zuxuan4, "%s,%s"},             //后四组选4
	"q3-q3zhixuan":   validator.PlayType{0, validate_q3zhixuan, convert_q3zhixuan, "%s%s%s--"},                   //前三直选
	"q3-q3zu3":       validator.PlayType{2829, validate_q3zu3, convert_q3zu3, "%s"},                              //前三组三
	"q3-q3zu6":       validator.PlayType{2830, validate_q3zu6, convert_q3zu6, "%s"},                              //前三组六
	"q3-q3zuxuanhz":  validator.PlayType{2832, validate_q3zuxuan_hz, convert_q3zuxuan_hz, "%s"},                  //前三组选和值
	"q3-q3zhixuanhz": validator.PlayType{2828, validate_q3zhixuan_hz, convert_q3zhixuan_hz, "%s"},                //前三直选和值
	"z3-z3zhixuan":   validator.PlayType{0, validate_z3zhixuan, convert_z3zhixuan, ""},                           //中三直选
	"z3-z3zu3":       validator.PlayType{2836, validate_z3zu3, convert_z3zu3, "%s"},                              //中三组三
	"z3-z3zu6":       validator.PlayType{2837, validate_z3zu6, convert_z3zu6, "%s"},                              //中三组六
	"z3-z3zuxuanhz":  validator.PlayType{2839, validate_z3zuxuan_hz, convert_z3zuxuan_hz, "%s"},                  //中三组选和值
	"z3-z3zhixuanhz": validator.PlayType{2835, validate_z3zhixuan_hz, convert_z3zhixuan_hz, "%s"},                //中三直选和值
	"h3-h3zhixuan":   validator.PlayType{2803, validate_h3zhixuan, convert_h3zhixuan, ""},                        //后三直选
	"h3-h3zu3":       validator.PlayType{2811, validate_h3zu3, convert_h3zu3, "%s"},                              //后三组三
	"h3-h3zu6":       validator.PlayType{2812, validate_h3zu6, convert_h3zu6, "%s"},                              //后三组六
	"h3-h3zuxuanhz":  validator.PlayType{2815, validate_h3zuxuan_hz, convert_h3zuxuan_hz, "%s"},                  //后三组选和值
	"h3-h3zhixuanhz": validator.PlayType{2810, validate_h3zhixuan_hz, convert_h3zhixuan_hz, "%s"},                //后三直选和值
	"bdw-q31m":       validator.PlayType{2842, validate_bdwq31m, convert_bdwq31m, "%s"},                          //前三一码
	"bdw-q32m":       validator.PlayType{2843, validate_bdwq32m, convert_bdwq32m, "%s"},                          //前三二码
	"bdw-z31m":       validator.PlayType{2844, validate_bdwz31m, convert_bdwz31m, "%s"},                          //中三一码
	"bdw-z32m":       validator.PlayType{2845, validate_bdwz32m, convert_bdwz32m, "%s"},                          //中三二码
	"bdw-h31m":       validator.PlayType{2846, validate_bdwh31m, convert_bdwh31m, "%s"},                          //后三一码
	"bdw-h32m":       validator.PlayType{2847, validate_bdwh32m, convert_bdwh32m, "%s"},                          //后三二码
	"bdw-4x1m":       validator.PlayType{2870, validate_bdw4x1m, convert_bdw4x1m, "%s"},                          //四星一码
	"bdw-4x2m":       validator.PlayType{2871, validate_bdw4x2m, convert_bdw4x2m, "%s"},                          //四星二码
	"bdw-5x2m":       validator.PlayType{2872, validate_bdw5x2m, convert_bdw5x2m, "%s"},                          //五星二码
	"bdw-5x3m":       validator.PlayType{2873, validate_bdw5x3m, convert_bdw5x3m, "%s"},                          //五星三码
	"2x-q2zhixuan":   validator.PlayType{2849, validate_q2zhixuan, convert_q2zhixuan, ""},                        //前二直选
	"2x-q2zuxuan":    validator.PlayType{0, validate_q2zuxuan, convert_q2zuxuan, ""},                             //前二组选
	"2x-h2zhixuan":   validator.PlayType{2803, validate_h2zhixuan, convert_h2zhixuan, ""},                        //后二直选
	"2x-h2zuxuan":    validator.PlayType{0, validate_h2zuxuan, convert_h2zuxuan, ""},                             //后二组选
	"2x-h2zuxuandt":  validator.PlayType{2818, validate_h2zuxuan_dt, convert_h2zuxuan_dt, "(%s)%s"},              //后二组选胆拖
	"1x-dwd":         validator.PlayType{2866, validate_1x_dwd, convert_1x_dwd, "%s%s%s%s%s"},                    //定位胆
	"1x-1xzhixuan":   validator.PlayType{2803, validate_1x_zhixuan, convert_1x_zhixuan, "----(%s)"},              //一星直选
	"q2ddxds":        validator.PlayType{2853, validate_q2ddxds, convert_q2ddxds, "%s%s"},                        //前二大大小单双
	"h2ddxds":        validator.PlayType{2804, validate_h2ddxds, convert_h2ddxds, "%s%s"},                        //后二大大小单双
	"rx-rx3":         validator.PlayType{2854, validate_rx3, convert_rx3, ""},                                    //任选三
	"rx-rx2":         validator.PlayType{2855, validate_rx2, convert_rx2, ""},                                    //任选二
	"qw-1ffs":        validator.PlayType{2856, validate_1ffs, convert_1ffs, "%s"},                                //一帆风顺
	"qw-hscs":        validator.PlayType{2857, validate_hscs, convert_hscs, "%s"},                                //好事成双
	"qw-3xbx":        validator.PlayType{2858, validate_3xbx, convert_3xbx, "%s"},                                //三星报喜
	"qw-4jfc":        validator.PlayType{2859, validate_4jfc, convert_4jfc, "%s"},                                //四季发财
}

// 生成销售期号列表
// prevIssue: 头一天最后一期
// days: 生成天数
func MakeSaleIssueList(prevIssue saleissue.SaleIssue, dayNum int) []saleissue.SaleIssue {
	if dayNum <= 0 {
		log.Println("天数无效:", dayNum)
		return nil
	}

	result := []saleissue.SaleIssue{}
	var interval time.Duration
	interval5M := time.Minute * 5
	interval10M := time.Minute * 10
	loc := validator.DefaultLocal
	var startTime, endTime, openTime, beginTime time.Time
	startTime = time.Unix(prevIssue.StartTime, 0)
	for i := 0; i < dayNum; i++ {
		year, month, day := startTime.AddDate(0, 0, 1).Date()
		for no := 1; no <= DAY_MAX_NO; no++ {
			if no <= 24 {
				interval = interval5M
				beginTime = time.Date(year, month, day, 0, 0, 0, 0, loc)
				startTime = beginTime.Add(interval * time.Duration(no-1))
			} else if no > 24 && no <= 96 {
				interval = interval10M
				beginTime = time.Date(year, month, day, 10, 0, 0, 0, loc)
				startTime = beginTime.Add(interval * time.Duration(no-24-1))
			} else if no > 96 {
				interval = interval5M
				beginTime = time.Date(year, month, day, 22, 0, 0, 0, loc)
				startTime = beginTime.Add(interval * time.Duration(no-96-1))
			}
			endTime = startTime.Add(interval).Add(-time.Minute)
			openTime = startTime.Add(interval).Add(time.Second * 50)
			if no == 24 {
				endTime = time.Date(year, month, day, 10, 0, 0, 0, loc).Add(-time.Minute)
				openTime = time.Date(year, month, day, 10, 0, 0, 0, loc).Add(time.Second * 50)
			}

			si := saleissue.SaleIssue{
				Issue:     fmt.Sprintf("%04d%02d%02d-%03d", year, month, day, no),
				StartTime: startTime.Unix(),
				EndTime:   endTime.Unix(),
				OpenTime:  openTime.Unix(),
			}
			result = append(result, si)
		}
	}
	return result
}
