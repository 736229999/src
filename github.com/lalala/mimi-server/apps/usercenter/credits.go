package main

import (
	"context"
	"log"
	"net/http"

	apiproto "github.com/caojunxyz/mimi-api/proto"
	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/utils"
)

func (srv *UcServer) HandleCreditsTaskInfo(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleCreditsTaskInfo")
	accountId, _, err := utils.ParseHttpRequest(w, r, nil)
	if err != nil {
		return
	}

	taskInfo, err := srv.dbClient.QueryAccountCreditsTaskInfo(context.Background(), &dbproto.IntValue{Value: accountId})
	if err != nil {
		http.Error(w, "错误：数据库查询错误!", http.StatusInternalServerError)
		return
	}

	result := &apiproto.CreditsTaskInfo{Credits: taskInfo.GetCredits()}
	for _, taskType := range utils.CreditsTaskList {
		cf := utils.CreditsTaskTable[taskType]
		task := &apiproto.CreditsTask{
			AwardCredits: cf.Awards,
			Type:         taskType,
			Title:        cf.Title,
			Desc:         cf.Desc,
			IsFinish:     false,
		}

		for _, v := range taskInfo.TaskList {
			if v.GetReason() == int32(taskType) {
				task.IsFinish = (cf.IsOnce && v.GetSumVar() > 0)
				break
			}
		}
		result.TaskList = append(result.TaskList, task)
	}
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", result)
}
