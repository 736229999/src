package options

import (
	"github.com/caojunxyz/mimi-admin/backend/core"
	dbproto "github.com/caojunxyz/mimi-admin/dbadminagent/proto"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"strconv"
)

//获取用户反馈的信息.
func (srv *OptionsService) HandleOptionsFeedbackList(c *gin.Context) {

	msg := &dbproto.FeedbackList{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	list, err := srv.Db().QueryFeedbackList(context.Background(), msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	list.Page = msg.GetPage()
	list.Size = msg.GetSize()
	srv.Log(c, msg, "获取用户的反馈信息", core.QUERY_OPERATION)
	srv.Json(list, http.StatusOK, c)
}

//删除用户反馈的信息.
func (srv *OptionsService) HandleOptionsFeedbackDel(c *gin.Context) {

	msg := make(map[string]int64)
	err := c.BindJSON(&msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("删除失败", http.StatusForbidden, c)
		return
	}

	_, err = srv.Db().DeleteFeedbackById(context.Background(), &dbproto.IntValue{Value: msg["id"]})
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("删除失败", http.StatusForbidden, c)
		return
	}
	srv.Log(c, msg, "删除用户反馈的信息", core.DELETE_OPERATION)
	srv.Json("删除成功", http.StatusOK, c)
}

//获取用户反馈详细.
func (srv *OptionsService) HandleOptionsFeedbackDetail(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}
	log.Println(id)

	feedback, err := srv.Db().QueryFeedbackById(context.Background(), &dbproto.IntValue{Value: int64(id)})
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	srv.Log(c, id, "看出反馈的详细内容", core.QUERY_OPERATION)
	srv.Json(feedback, http.StatusOK, c)
}

//处理反馈状态.
func (srv *OptionsService) HandleOptionsFeedbackUpdate(c *gin.Context) {

	msg := &dbproto.Feedback{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("处理失败", http.StatusForbidden, c)
		return
	}

	_, err = srv.Db().UpdateFeedbackById(context.Background(), msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("处理失败", http.StatusForbidden, c)
		return
	}

	srv.Json("处理成功", http.StatusOK, c)
}
