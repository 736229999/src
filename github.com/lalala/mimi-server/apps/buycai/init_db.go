package main

import (
	"context"
	"log"
	"time"

	"github.com/caojunxyz/mimi-api/proto"
	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/lottery"
	"github.com/caojunxyz/mimi-server/lottery/saleissue"
)

var prevIssues = map[apiproto.LotteryId]saleissue.SaleIssue{
	apiproto.LotteryId_Bjpk10: saleissue.SaleIssue{
		Issue: "625213", StartTime: time.Date(2017, time.Month(6), 24, 0, 0, 0, 0, time.Local).Unix(),
	},
	apiproto.LotteryId_Cqssc: saleissue.SaleIssue{
		Issue: "20170624-120", StartTime: time.Date(2017, time.Month(6), 24, 0, 0, 0, 0, time.Local).Unix(),
	},
	apiproto.LotteryId_Gd11x5: saleissue.SaleIssue{
		Issue: "17062484", StartTime: time.Date(2017, time.Month(6), 24, 0, 0, 0, 0, time.Local).Unix(),
	},
	apiproto.LotteryId_Dlt: saleissue.SaleIssue{
		Issue:     "2017068",
		StartTime: time.Date(2017, time.Month(6), 12, 20, 10, 0, 0, time.Local).Unix(),
		EndTime:   time.Date(2017, time.Month(6), 14, 20, 0, 0, 0, time.Local).Unix(),
		OpenTime:  time.Date(2017, time.Month(6), 14, 20, 30, 0, 0, time.Local).Unix(),
	},
	apiproto.LotteryId_Fc3d: saleissue.SaleIssue{
		Issue:     "2017158",
		StartTime: time.Date(2017, time.Month(6), 13, 20, 40, 0, 0, time.Local).Unix(),
		EndTime:   time.Date(2017, time.Month(6), 14, 20, 30, 0, 0, time.Local).Unix(),
		OpenTime:  time.Date(2017, time.Month(6), 14, 21, 15, 0, 0, time.Local).Unix(),
	},
	apiproto.LotteryId_Pl3: saleissue.SaleIssue{
		Issue:     "2017158",
		StartTime: time.Date(2017, time.Month(6), 13, 20, 00, 0, 0, time.Local).Unix(),
		EndTime:   time.Date(2017, time.Month(6), 14, 19, 45, 0, 0, time.Local).Unix(),
		OpenTime:  time.Date(2017, time.Month(6), 14, 20, 30, 0, 0, time.Local).Unix(),
	},
	apiproto.LotteryId_Pl5: saleissue.SaleIssue{
		Issue:     "2017158",
		StartTime: time.Date(2017, time.Month(6), 13, 20, 00, 0, 0, time.Local).Unix(),
		EndTime:   time.Date(2017, time.Month(6), 14, 19, 45, 0, 0, time.Local).Unix(),
		OpenTime:  time.Date(2017, time.Month(6), 14, 20, 30, 0, 0, time.Local).Unix(),
	},
	apiproto.LotteryId_Ssq: saleissue.SaleIssue{
		Issue:     "2017068",
		StartTime: time.Date(2017, time.Month(6), 11, 20, 40, 0, 0, time.Local).Unix(),
		EndTime:   time.Date(2017, time.Month(6), 13, 20, 30, 0, 0, time.Local).Unix(),
		OpenTime:  time.Date(2017, time.Month(6), 13, 21, 15, 0, 0, time.Local).Unix(),
	},
}

func (agt *BuycaiAgent) initSaleIssues() {
	cfg := lottery.GetConfig(agt.id)
	log.Printf("%s initSaleIssues...\n", cfg.Code)
	// openInfo, err := agt.dbOpencai.OpencaiQueryLatestIssue(context.Background(), &dbproto.StringValue{Value: cfg.Code})
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	var list []saleissue.SaleIssue
	if cfg.Type == apiproto.LotteryType_LowFreq {
		list = cfg.IssuesGenerator(prevIssues[cfg.Id], 60)
	} else if cfg.Type == apiproto.LotteryType_HighFreq {
		list = cfg.IssuesGenerator(prevIssues[cfg.Id], 60)
	}

	// TODO: 批量插入
	for _, v := range list {
		insertArg := &dbproto.BuycaiUpsertIssueArg{
			Code:      cfg.Code,
			SaleIssue: &dbproto.BuycaiSaleIssue{Issue: v.Issue, StartTime: v.StartTime, EndTime: v.EndTime, OpenTime: v.OpenTime},
		}
		_, err := agt.dbBuycai.BuycaiUpsertIssue(context.Background(), insertArg)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
