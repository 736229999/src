package activity

import (
	"github.com/gin-gonic/gin"
	"github.com/caojunxyz/mimi-admin/dbadminagent/proto"
	"log"
	"net/http"
	"golang.org/x/net/context"
	"github.com/caojunxyz/mimi-admin/backend/core"
	"strconv"
	"fmt"
)

//添加礼包模板.
func (srv *ActivityService) HandleGiftTemplateAdd (c *gin.Context) {

	msg := &dbproto.GiftTemplate{}
	if err := c.BindJSON(&msg); err != nil {
		log.Printf("%+v\n", err)
		srv.Json("添加失败", http.StatusForbidden, c)
		return
	}

	msg.Creator = fmt.Sprintf("%s", srv.GetUserInfo(c, "username"))
	_, err := srv.Db().InsertGiftTemplate(context.Background(), msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("添加失败", http.StatusForbidden, c)
		return
	}

	srv.Log(c, msg, "添加礼包模板", core.ADD_OPERATION)
	srv.Json("添加成功", http.StatusOK, c)

}

//获取礼包模板的列表.
func (srv *ActivityService) HandleGiftTemplateList (c *gin.Context) {

	msg := &dbproto.GiftTemplateList{}
	if err := c.BindJSON(&msg); err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	giftTemplate, err := srv.Db().QueryGiftTemplateList(context.Background(), msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}
	giftTemplate.Page  = msg.GetPage()
	giftTemplate.Size  = msg.GetSize()
	giftTemplate.Title = msg.GetTitle()

	srv.Log(c, msg, "获取礼包模板列表", core.QUERY_OPERATION)
	srv.Json(giftTemplate, http.StatusOK, c)
}

//根据id获取礼包模板.
func (srv *ActivityService) HandleGiftTemplateById (c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	giftTemplate, err := srv.Db().QueryGiftTemplateById(context.Background(), &dbproto.IntValue{Value:int64(id)})
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	srv.Log(c, id, "根据id获取礼包模板", core.QUERY_OPERATION)
	srv.Json(giftTemplate, http.StatusOK, c)
}

//更改.
func (srv *ActivityService) HandleGiftTemplateUpdate (c *gin.Context) {

	msg := &dbproto.GiftTemplate{}
	if err := c.BindJSON(&msg); err != nil {
		log.Printf("%+v\n", err)
		srv.Json("更新失败", http.StatusForbidden, c)
		return
	}

	_, err := srv.Db().UpdateGiftTemplateById(context.Background(), msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("更新失败", http.StatusForbidden, c)
		return
	}

	srv.Log(c, msg, "添加礼包模板", core.UPDATE_OPERATION)
	srv.Json("添加成功", http.StatusOK, c)
}

//根据id删除.
func (srv *ActivityService) HandleGiftTemplateDelete (c *gin.Context) {

	msg := &dbproto.GiftTemplate{}
	if err := c.BindJSON(&msg); err != nil {
		log.Printf("%+v\n", err)
		srv.Json("删除失败", http.StatusForbidden, c)
		return
	}

	_, err := srv.Db().DeleteGiftTemplateById(context.Background(), &dbproto.IntValue{Value:msg.GetId()})
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("删除失败", http.StatusForbidden, c)
		return
	}

	srv.Log(c, msg, "根据id删除礼包", core.DELETE_OPERATION)
	srv.Json("删除成功", http.StatusOK, c)
}