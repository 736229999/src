package main

import (
	"net/http"
	"log"

	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"golang.org/x/net/context"
	"github.com/caojunxyz/mimi-server/utils"
	"github.com/caojunxyz/mimi-api/proto"
	"strconv"
)

//活动列表
func (srv *ActivityServer) HandleActivityList(w http.ResponseWriter, r *http.Request)  {
	log.Println("活动列表api........")
	query := r.URL.Query()
	if len(query) != 1 {
		http.Error(w, "参数错误", http.StatusBadRequest)
		return
	}
	account_id,err := strconv.Atoi(query["account_id"][0])
	if err != nil {
		log.Printf("%+v\n",err)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "转换失败", nil)
		return
	}
	activityList ,err := srv.dbClient.QueryActivityList(context.Background(),&dbproto.IntValue{Value:int64(account_id)})
	if err != nil {
		log.Printf("%+v\n",err)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "获取活动列表失败", nil)
		return
	}
	log.Println(activityList)
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", activityList)

}
//活动详情
func (srv *ActivityServer) HandleActivityDetail(w http.ResponseWriter, r *http.Request)  {
	log.Println("活动详情api.........")
	activityAccount := &dbproto.ActivityAccount{}
	activityAccount.AccountId = 1
	activityAccount.ActivityId = 1
	activityDetail ,err := srv.dbClient.ActivityDetail(context.Background(),activityAccount)
	if err != nil {
		log.Printf("%+v\n",err)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "获取活动详情失败", nil)
		return
	}
	log.Println(activityDetail)
	//utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", activityList)
}
