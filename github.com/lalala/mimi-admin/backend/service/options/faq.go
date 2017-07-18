package options

import (
	"log"
	// dbproto "github.com/caojunxyz/mimi-server/dbagent/proto"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"net/http"

	"github.com/caojunxyz/mimi-admin/backend/core"
	dbproto "github.com/caojunxyz/mimi-admin/dbadminagent/proto"
	"strconv"
)

// HandleFaqAdd 添加一条Faq信息
func (srv *OptionsService) HandleFaqAdd(c *gin.Context) {
	log.Println("Handle Faq Add")
	msg := dbproto.Faq{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Println("BindJSON", err)
		srv.Json("数据格式错误", http.StatusBadRequest, c)
		return
	}
	log.Printf("msg is %+v\n", msg)
	rid, err := srv.Db().CreateFaq(context.Background(), &msg)
	if err != nil {
		log.Println("CreateFaq error", err)
		srv.Json("服务器异常", http.StatusInternalServerError, c)
		return
	}
	log.Printf("Id is %v\n", rid)
	srv.Log(c, msg, "添加常见问题", core.ADD_OPERATION)
	srv.Json("", http.StatusOK, c)
}

// HandleFaqList 获取Faq列表
func (srv *OptionsService) HandleFaqList(c *gin.Context) {
	log.Println("Handle Faq List")

	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pagesize", "10")

	pageDb, err := strconv.Atoi(page)
	if err != nil {
		log.Println("error", err)
		srv.Json("页码参数错误", http.StatusForbidden, c)
		return
	}
	pageSizeDb, err := strconv.Atoi(pageSize)
	if err != nil {
		log.Println("error", err)
		srv.Json("每页尺寸参数错误", http.StatusForbidden, c)
		return
	}

	title := c.DefaultQuery("title", "")

	msg := dbproto.QueryFaqArg{
		Title:    title,
		Page:     int32(pageDb),
		PageSize: int32(pageSizeDb),
	}
	res, err := srv.Db().QueryFaqList(context.Background(), &msg)
	if err != nil {
		log.Println("QueryFaqList ", err)
		srv.Json("服务器异常", http.StatusInternalServerError, c)
		return
	}
	// log.Printf("res is %+v", res)
	srv.Log(c, msg, "查询Faq列表", core.QUERY_OPERATION)
	srv.Json(res, http.StatusOK, c)
}

// HandleFaqDetail 获取一条Banner
func (srv *OptionsService) HandleFaqDetail(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}
	res, err := srv.Db().QueryFaqById(context.Background(), &dbproto.FaqId{Id: id})
	if err != nil {
		log.Println("QueryFaqById error", err)
		srv.Json("服务器异常", http.StatusInternalServerError, c)
		return
	}
	log.Printf("res is %+v", res)
	srv.Log(c, id, "获取某个常见问题", core.QUERY_OPERATION)
	srv.Json(res, http.StatusOK, c)
}

// HandleUpdateFaq 更新Banner信息
func (srv *OptionsService) HandleUpdateFaq(c *gin.Context) {
	msg := dbproto.Faq{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Println("BindJSON", err)
		srv.Json("数据格式错误", http.StatusBadRequest, c)
		return
	}
	log.Printf("msg is %+v\n", msg)
	rid, err := srv.Db().UpdateFaq(context.Background(), &msg)
	if err != nil {
		log.Println("UpdateFaq error", err)
		srv.Json("服务器异常", http.StatusInternalServerError, c)
		return
	}
	log.Printf("Id is %v\n", rid.GetValue())
	srv.Log(c, msg, "更新常见问题", core.UPDATE_OPERATION)
	srv.Json("", http.StatusOK, c)
}
