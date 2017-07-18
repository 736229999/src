package activity

import (
	"github.com/caojunxyz/mimi-admin/backend/core"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	dbproto "github.com/caojunxyz/mimi-admin/dbadminagent/proto"
	"golang.org/x/net/context"
	"strconv"
)

type ActivityService struct {
	core.Service
}
//添加活动
func (srv *ActivityService) AddActivity (c *gin.Context) {
	msg := dbproto.Activity{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}
	res ,err := srv.Db().InsertActivity(context.Background(),&msg)
	if err != nil {
		log.Println(err)
		srv.Json("添加失败", http.StatusForbidden, c)
		return
	}
	log.Println(res)
	srv.Log(c, msg, "添加活动", core.QUERY_OPERATION)
	srv.Json("添加成功", http.StatusOK, c)
}

//查询所有活动
func (srv *ActivityService) QueryActivityList (c *gin.Context) {
	msg := dbproto.ActivityReplyList{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	list,err := srv.Db().QueryActivityList(context.Background(),&msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}
	list.Size = msg.GetSize()
	list.Page = msg.GetPage()
	srv.Log(c, msg, "日志列表", core.QUERY_OPERATION)
	srv.Json(list,200,c)

}

//删除活动
func (srv *ActivityService) DeleteActivity (c *gin.Context) {
	msg := make(map[string]int64)
	defer func() {
		msg = nil
	}()
	err := c.BindJSON(&msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}
	_,err = srv.Db().DeleteActivity(context.Background(),&dbproto.IntValue{Value:msg["id"]})
	if err != nil {
		log.Println(err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}
	srv.Log(c, msg, "日志列表", core.QUERY_OPERATION)
	srv.Json("删除成功",200,c)
}
//获取单个活动详情
func (srv *ActivityService) QueryActivityById (c *gin.Context) {
	msg := make(map[string]string)
	defer func() {
		msg = nil
	}()
	err := c.BindJSON(&msg)
	log.Println("前台传过来的id：",msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}
	id,err := strconv.Atoi(msg["id"])
	if err != nil {
		log.Println("字符串转换失败：",err)
		return
	}
	activity,err := srv.Db().QueryActivityById(context.Background(),&dbproto.IntValue{Value:int64(id)})
	if err != nil {
		log.Println(err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}
	log.Println("获取到的rpc活动：",activity)
	srv.Log(c, msg, "日志列表", core.QUERY_OPERATION)
	srv.Json(activity,200,c)
}
//更新活动
func (srv *ActivityService) UpdateActivity (c *gin.Context) {
	msg := dbproto.Activity{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}
	_,err = srv.Db().UpdateActivity(context.Background(),&msg)
	if err !=nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}
	srv.Log(c, msg, "日志列表", core.QUERY_OPERATION)
	srv.Json("更新成功",200,c)
}