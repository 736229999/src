package lottery

import (
	"github.com/caojunxyz/mimi-api/proto"
	"github.com/caojunxyz/mimi-server/lottery/validator/bjpk10"
	"github.com/caojunxyz/mimi-server/lottery/validator/cqssc"
	"github.com/caojunxyz/mimi-server/lottery/validator/gd11x5"
)

//-------------------------------------------------------------------------------------------------------------------------------------------
// 10:00-22:00（72期）10分钟一期，22:00-02:00（48期）5分钟一期
var cqsscConfig = Config{
	Name:            "重庆时时彩",
	Code:            Cqssc,
	Id:              apiproto.LotteryId_Cqssc,
	Type:            apiproto.LotteryType_HighFreq,
	ZfId:            "28",
	DayMaxNo:        cqssc.DAY_MAX_NO,
	MaxMultiple:     9999,
	PlayList:        cqssc.PlayList,
	IssuesGenerator: cqssc.MakeSaleIssueList,
}

//-------------------------------------------------------------------------------------------------------------------------------------------
// 09:02 至 23:57 每期间隔5分钟，等待开奖时间为1分钟，全天共179期
var bjpk10Config = Config{
	Name:            "北京PK拾",
	Code:            Bjpk10,
	Id:              apiproto.LotteryId_Bjpk10,
	ZfId:            "94",
	Type:            apiproto.LotteryType_HighFreq,
	DayMaxNo:        bjpk10.DAY_MAX_NO,
	MaxMultiple:     9999,
	PlayList:        bjpk10.PlayList,
	IssuesGenerator: bjpk10.MakeSaleIssueList,
}

//-------------------------------------------------------------------------------------------------------------------------------------------
// 9:00 ~ 23:00 销售, 10分钟一期, 9:10分第一期开奖
var gd11x5Config = Config{
	Name:            "广东11选5",
	Code:            Gd11x5,
	Id:              apiproto.LotteryId_Gd11x5,
	ZfId:            "78",
	Type:            apiproto.LotteryType_HighFreq,
	DayMaxNo:        gd11x5.DAY_MAX_NO,
	MaxMultiple:     9999,
	PlayList:        gd11x5.PlayList,
	IssuesGenerator: gd11x5.MakeSaleIssueList,
}

//-------------------------------------------------------------------------------------------------------------------------------------------
