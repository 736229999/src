package activity

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
	"encoding/json"
	dbproto "github.com/caojunxyz/mimi-admin/dbadminagent/proto"
	"golang.org/x/net/context"

	"github.com/caojunxyz/mimi-server/utils"
	"github.com/caojunxyz/mimi-admin/backend/core"
)

//添加cdk.
func (srv *ActivityService) HandleActivityCdKeyAdd(c *gin.Context) {

	msg := &dbproto.CdkeyBatch{}
	if err := c.BindJSON(&msg); err != nil {
		log.Printf("%+v\n", err)
		srv.Json("添加失败", http.StatusForbidden, c)
		return
	}

	_, err := srv.Db().InsertCdkeyBatch(context.Background(), msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("添加失败", http.StatusForbidden, c)
		return
	}

	srv.Log(c, msg, "添加cdkey", core.ADD_OPERATION)
	srv.Json("添加成功", http.StatusOK, c)
}

//获取cdk列表.
func (srv *ActivityService) HandleActivityCdKeyList(c *gin.Context) {

	msg := dbproto.CdkeyReply{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	cdkey, err := srv.Db().QueryCdkeyList(context.Background(), &msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	cdkey.Title = msg.GetTitle()
	cdkey.Size = msg.GetSize()
	cdkey.Page = msg.GetPage()

	srv.Log(c, msg, "获取cdkey列表", core.QUERY_OPERATION)
	srv.Json(cdkey, http.StatusOK, c)
}

//查询所有的礼包模板.
func (srv *ActivityService) HandleActivityGiftTemplateList(c *gin.Context) {

	list, err := srv.Db().QueryGiftTemplateListAll (context.Background(), &dbproto.Nil{})
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	srv.Log(c, nil, "查询所有的礼包模板", core.QUERY_OPERATION)
	srv.Json(list, http.StatusOK, c)
}

//获取cdkey的详细信息.
func (srv *ActivityService) HandleActivityCdkeyDetail(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	detail, err := srv.Db().QueryCdkeyDetailById(context.Background(), &dbproto.IntValue{Value: int64(id)})
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	if err = json.Unmarshal([]byte(detail.GetContent()), &detail.Gift); err != nil {
		log.Printf("%+v\n")
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	srv.Log(c, id, "根据id获取cdkey的详细信息", core.QUERY_OPERATION)
	srv.Json(detail, http.StatusOK, c)
}

//更新cdkey.
func (srv *ActivityService) HandleActivityCdKeyUpdate(c *gin.Context) {

	msg := &dbproto.CdkeyDetail{}
	if err := c.BindJSON(&msg); err != nil {
		log.Printf("%+v\n", err)
		srv.Json("更新失败", http.StatusForbidden, c)
		return
	}

	_, err := srv.Db().UpdateCdkeyById(context.Background(), msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("更新失败", http.StatusForbidden, c)
		return
	}

	srv.Log(c, msg, "更新cdkey", core.UPDATE_OPERATION)
	srv.Json("更新成功", http.StatusOK, c)
}

//删除cdkey.
func (srv *ActivityService) HandleActivityCdKeyDelete(c *gin.Context) {

	params := make(map[string]int64)
	err := c.BindJSON(&params)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("删除失败", http.StatusForbidden, c)
		return
	}

	_, err = srv.Db().DeleteCdkeyById(context.Background(), &dbproto.IntValue{Value: params["id"]})
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("删除失败", http.StatusForbidden, c)
		return
	}

	srv.Log(c, params, "删除cdkey", core.DELETE_OPERATION)
	srv.Json("删除成功", http.StatusOK, c)
}

//导出csv.
func (srv *ActivityService) HandleActivityCdKeyExport(c *gin.Context) {

	msg := make(map[string]int64)
	err := c.BindJSON(&msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("导出失败", http.StatusForbidden, c)
		return
	}
	detail, err := srv.Db().QueryCdkeyDetailById(context.Background(), &dbproto.IntValue{Value: msg["id"]})
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	log.Println(msg)
	record := utils.GenerateCdkey(int(msg["id"]), 10000)
	log.Printf("%+v\n", record)

	dist := make(map[string]interface{})
	dist["cdkey"] = record
	dist["title"] = detail.GetTitle()

	srv.Log(c, msg, "导出csv", core.QUERY_OPERATION)
	srv.Json(dist, http.StatusOK, c)
}
