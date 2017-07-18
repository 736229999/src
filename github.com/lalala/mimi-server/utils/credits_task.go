package utils

import (
	"context"

	"github.com/caojunxyz/mimi-api/proto"
	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
)

func AddCredits(db dbproto.DbUsercenterAgentClient, accountId int64, reason apiproto.CreditsTask_TaskType, detail string, args ...interface{}) {
	dbArg := &dbproto.ChangeVirtualFundArg{
		AccountId: accountId,
		Reason:    int32(reason),
		Detail:    detail,
		Var:       CreditsTaskTable[reason].Awards,
	}
	if reason == apiproto.CreditsTask_Buycai && len(args) == 1 {
		money, ok := args[0].(float64)
		if ok {
			n := int32(money / 2)
			dbArg.Var *= n
		}
	}
	db.ChangeCredits(context.Background(), dbArg)
}

type Task struct {
	Awards int32
	Title  string
	Desc   string
	IsOnce bool
}

var CreditsTaskTable = map[apiproto.CreditsTask_TaskType]Task{
	apiproto.CreditsTask_Buycai:       Task{Awards: 1, Title: "购彩消费", Desc: "每消费2元得1积分"},
	apiproto.CreditsTask_InviteFriend: Task{Awards: 5, Title: "邀请好友", Desc: "每邀请1位好友, 双方各得5积分"},
	apiproto.CreditsTask_FirstWin:     Task{Awards: 10, Title: "首次中奖", Desc: "首次中奖任意彩种", IsOnce: true},
	// apiproto.CreditsTask_FirstBuycai:   Task{Awards: 10, Title: "首次购彩", Desc: "第一次完成购彩", IsOnce: true},
	apiproto.CreditsTask_AuthRealname:  Task{Awards: 20, Title: "实名认证", Desc: "通过实名认证", IsOnce: true},
	apiproto.CreditsTask_BindPhone:     Task{Awards: 10, Title: "绑定手机", Desc: "绑定手机号", IsOnce: true},
	apiproto.CreditsTask_RegistAccount: Task{Awards: 10, Title: "新用户注册", Desc: "注册账户成功", IsOnce: true},
}

var CreditsTaskList = []apiproto.CreditsTask_TaskType{
	apiproto.CreditsTask_Buycai,
	apiproto.CreditsTask_InviteFriend,
	apiproto.CreditsTask_FirstWin,
	// apiproto.CreditsTask_FirstBuycai,
	apiproto.CreditsTask_AuthRealname,
	apiproto.CreditsTask_BindPhone,
	apiproto.CreditsTask_RegistAccount,
}
