package activity

import (
	"github.com/gin-gonic/gin"
	dbproto "github.com/caojunxyz/mimi-admin/dbadminagent/proto"
	"log"
	"net/http"
	"golang.org/x/net/context"
	"encoding/json"
	"strconv"
	"github.com/caojunxyz/mimi-admin/backend/core"
)


//获取礼包列表.
func (srv *ActivityService) HandleGiftLit (c *gin.Context) {

	msg := dbproto.GiftList{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	log.Printf("%+v\n",msg)

	giftList, err := srv.Db().QueryGiftList(context.Background(), &msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	giftList.Page  = msg.GetPage()
	giftList.Title = msg.GetTitle()
	giftList.Size  = msg.GetSize()

	srv.Log(c, msg, "获取礼包列表", core.QUERY_OPERATION)
	srv.Json(giftList, http.StatusOK, c)
}

//获取礼包的类型列表.
//func (srv *ActivityService) HandleGiftTypeList (c *gin.Context) {
//
//	giftTypeList := admin.GiftTypes
//
//	srv.Log(c, nil, "获取礼包的类型列表", core.QUERY_OPERATION)
//	srv.Json(giftTypeList, http.StatusOK, c)
//}

//添加礼包.
func (srv *ActivityService) HandleGiftAdd (c *gin.Context) {

	msg := dbproto.GiftPackageRequest{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("添加失败", http.StatusForbidden, c)
		return
	}

	msg.Content, err =  srv.ToJson(msg.GetGift())
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("失败", http.StatusForbidden, c)
		return
	}

	_, err = srv.Db().InsertGift(context.Background(), &msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("添加失败", http.StatusForbidden, c)
		return
	}

	srv.Log(c, msg, "添加礼包", core.ADD_OPERATION)
	srv.Json("ok", http.StatusOK, c)
}

//获取查看礼包详细数据.
func (srv *ActivityService) HandleDetail (c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	giftDetail, err := srv.Db().QueryGiftDetailById(context.Background(), &dbproto.IntValue{Value:int64(id)})
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	if err := json.Unmarshal([]byte(giftDetail.GetContent()), &giftDetail.Gift); err != nil {
		log.Printf("%+v\v", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	srv.Log(c, id, "获取超礼包详细数据", core.QUERY_OPERATION)
	srv.Json(giftDetail, http.StatusOK, c)
}

//删除礼包.
func (srv *ActivityService) HandleDeleteGift (c *gin.Context) {

	var msg map[string]int64
	err := c.BindJSON(&msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("删除失败", http.StatusForbidden, c)
		return
	}

	_, err = srv.Db().DeleteGiftById(context.Background(), &dbproto.IntValue{Value:msg["id"]})
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("删除失败", http.StatusForbidden, c)
		return
	}

	srv.Log(c, msg, "删除礼包", core.DELETE_OPERATION)
	srv.Json("删除成功", http.StatusOK, c)
}

//更新礼包.
func (srv *ActivityService) HandleGiftUpdate (c *gin.Context) {

	msg := dbproto.GiftPackageRequest{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("更新失败", http.StatusForbidden, c)
		return
	}

	msg.Content, err = srv.ToJson(msg.GetGift())
	if err != nil {
		log.Printf("%+v]\n", err)
		srv.Json("更新失败", http.StatusForbidden, c)
		return
	}

	_, err = srv.Db().UpdateGiftById(context.Background(), &msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("更新失败", http.StatusForbidden, c)
		return
	}

	srv.Log(c, msg, "更新礼包", core.UPDATE_OPERATION)
	srv.Json("更新成功", http.StatusOK, c)
}
