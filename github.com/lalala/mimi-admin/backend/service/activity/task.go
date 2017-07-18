package activity

import (
	"github.com/gin-gonic/gin"
	"net/http"
	//
	"log"
	"github.com/caojunxyz/mimi-admin/dbadminagent/proto"
	"golang.org/x/net/context"
	"github.com/caojunxyz/mimi-admin/backend/core"
	"strconv"
)

//添加任务
func (srv *ActivityService) InserTask(c *gin.Context) {
	msg := dbproto.Task{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}
	log.Println("前台传过来的数据：",msg)
	_,err = srv.Db().InsertTask(context.Background(),&msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}
	srv.Log(c, msg, "日志列表", core.QUERY_OPERATION)
	srv.Json("添加成功", http.StatusOK, c)
}
//查询所有的任务列表
func (srv *ActivityService) QueryTaskList(c *gin.Context) {
	msg := dbproto.TaskReplyList{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}
	list,err := srv.Db().QueryTaskList(context.Background(),&msg)
	log.Println(list)
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
//查询所有任务不带分页
func (srv *ActivityService) QueryAllTask(c *gin.Context) {
	list,err := srv.Db().QueryAllTaskList(context.Background(),&dbproto.Nil{})
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}
	srv.Log(c, list, "日志列表", core.QUERY_OPERATION)
	srv.Json(list,200,c)
}
//查询所有礼包模板不带分页
func (srv *ActivityService) QueryAllGiftTemplate(c *gin.Context) {
	list,err := srv.Db().QueryAllGiftTemplateList(context.Background(),&dbproto.Nil{})
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}
	srv.Log(c, list, "日志列表", core.QUERY_OPERATION)
	srv.Json(list,200,c)
}
//删除任务
func (srv *ActivityService) DeleteTask(c *gin.Context) {
	msg := make(map[string]int64)
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

	_,err = srv.Db().DeleteTask(context.Background(),&dbproto.IntValue{Value:msg["id"]})
	if err != nil {
		log.Println(err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}
	srv.Log(c, msg, "日志列表", core.QUERY_OPERATION)
	srv.Json("删除任务成功",200,c)
}
//获取单个任务详情
func (srv *ActivityService) QueryTaskById(c *gin.Context) {
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
	task,err := srv.Db().QueryTaskById(context.Background(),&dbproto.IntValue{Value:int64(id)})
	if err != nil {
		log.Println(err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}
	srv.Log(c, msg, "日志列表", core.QUERY_OPERATION)
	srv.Json(task,200,c)
}
//更新任务
func (srv *ActivityService) UpdateTask(c *gin.Context) {
	msg := dbproto.Task{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}
	_,err = srv.Db().UpdateTask(context.Background(),&msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}
	srv.Log(c, msg, "日志列表", core.QUERY_OPERATION)
	srv.Json("更新成功", http.StatusOK, c)
}
//获取所有的任务类型
func (srv *ActivityService) QueryAllTaskType(c *gin.Context) {
	types,err := srv.Db().QueryAllTaskType(context.Background(),&dbproto.Nil{})
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}
	srv.Log(c, types, "日志列表", core.QUERY_OPERATION)
	srv.Json(types, http.StatusOK, c)
}