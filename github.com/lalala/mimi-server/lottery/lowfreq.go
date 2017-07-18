package lottery

import (
	apiproto "github.com/caojunxyz/mimi-api/proto"
	"github.com/caojunxyz/mimi-server/lottery/validator/dlt"
	"github.com/caojunxyz/mimi-server/lottery/validator/fc3d"
	"github.com/caojunxyz/mimi-server/lottery/validator/pl3"
	"github.com/caojunxyz/mimi-server/lottery/validator/pl5"
	"github.com/caojunxyz/mimi-server/lottery/validator/ssq"
)

//-------------------------------------------------------------------------------------------------------------------------------------------
// 每周一、三、六 20:30开奖
var dltConfig = Config{
	Name:        "大乐透",
	Code:        Dlt,
	Id:          apiproto.LotteryId_Dlt,
	Type:        apiproto.LotteryType_LowFreq,
	ZfId:        "39",
	BlueNum:     2,
	MaxMultiple: 99,
	BonusNames: map[int]string{
		1: "一等奖", 2: "追加一等奖",
		3: "二等奖", 4: "追加二等奖",
		5: "三等奖", 6: "追加三等奖",
		7: "四等奖", 8: "追加四等奖",
		9: "五等奖", 10: "追加五等奖",
		11: "六等奖",
	},
	PlayList:        dlt.PlayList,
	IssuesGenerator: dlt.MakeSaleIssueList,
}

//-------------------------------------------------------------------------------------------------------------------------------------------
// 每日 20:30开奖
var fc3dConfig = Config{
	Name:        "福彩3D",
	Code:        Fc3d,
	Id:          apiproto.LotteryId_Fc3d,
	ZfId:        "6",
	Type:        apiproto.LotteryType_LowFreq,
	BlueNum:     0,
	MaxMultiple: 99,
	BonusNames: map[int]string{
		1: "单选", 2: "组选3", 3: "组选6",
	},
	PlayList:        fc3d.PlayList,
	IssuesGenerator: fc3d.MakeSaleIssueList,
}

//-------------------------------------------------------------------------------------------------------------------------------------------
// 每周二、四、日晚上21:15开奖,
var ssqConfig = Config{
	Name:        "双色球",
	Code:        Ssq,
	Id:          apiproto.LotteryId_Ssq,
	Type:        apiproto.LotteryType_LowFreq,
	ZfId:        "5",
	BlueNum:     1,
	MaxMultiple: 99,
	BonusNames: map[int]string{
		1: "一等奖", 2: "二等奖", 3: "三等奖",
		4: "四等奖", 5: "五等奖", 6: "六等奖",
	},
	PlayList:        ssq.PlayList,
	IssuesGenerator: ssq.MakeSaleIssueList,
}

//-------------------------------------------------------------------------------------------------------------------------------------------
// 每日20:30开奖
var pl3Config = Config{
	Name:            "排列3",
	Code:            Pl3,
	Id:              apiproto.LotteryId_Pl3,
	Type:            apiproto.LotteryType_LowFreq,
	ZfId:            "63",
	MaxMultiple:     99,
	PlayList:        pl3.PlayList,
	IssuesGenerator: pl3.MakeSaleIssueList,
}

//-------------------------------------------------------------------------------------------------------------------------------------------
// 每日20:30开奖
var pl5Config = Config{
	Name:            "排列5",
	Code:            Pl5,
	Id:              apiproto.LotteryId_Pl5,
	Type:            apiproto.LotteryType_LowFreq,
	ZfId:            "64",
	MaxMultiple:     99,
	PlayList:        pl5.PlayList,
	IssuesGenerator: pl5.MakeSaleIssueList,
}
