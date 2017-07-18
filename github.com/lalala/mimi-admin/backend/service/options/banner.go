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

// HandleBannerAdd 添加一条Banner信息
func (srv *OptionsService) HandleBannerAdd(c *gin.Context) {
	log.Println("Handle Banner Add")
	msg := dbproto.Banner{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Println("BindJSON", err)
		srv.Json("数据格式错误", http.StatusBadRequest, c)
		return
	}
	log.Printf("msg is %+v\n", msg)
	rid, err := srv.Db().CreateBanner(context.Background(), &msg)
	if err != nil {
		log.Println("CreateBanner error", err)
		srv.Json(err, http.StatusInternalServerError, c)
		return
	}
	log.Printf("Id is %v\n", rid)
	srv.Log(c, msg, "添加轮播图", core.ADD_OPERATION)
	srv.Json("", http.StatusOK, c)
}

// HandleBannerDetail 获取一条Banner
func (srv *OptionsService) HandleBannerDetail(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}
	res, err := srv.Db().QueryBannerById(context.Background(), &dbproto.BannerId{Id: id})
	if err != nil {
		log.Println("QueryBannerById error", err)
		srv.Json(err, http.StatusInternalServerError, c)
		return
	}
	log.Printf("res is %+v", res)
	srv.Log(c, id, "获取轮播图详细信息", core.QUERY_OPERATION)
	srv.Json(res, http.StatusOK, c)
}

// HandleBannerList 获取Banner列表
func (srv *OptionsService) HandleBannerList(c *gin.Context) {
	log.Println("Handle Banner List")
	// msg := dbproto.QueryBannerArg{}
	// err := c.BindJSON(&msg)
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

	locationStr := c.DefaultQuery("location", strconv.Itoa(int(dbproto.Banner_Location_Home)))
	location, err := strconv.Atoi(locationStr)
	if err != nil {
		log.Println("locationStr", err)
		srv.Json("数据格式错误", http.StatusBadRequest, c)
		return
	}

	log.Println("location", dbproto.Banner_LocationTypeBackend(location))
	msg := dbproto.QueryBannerArg{
		Location: dbproto.Banner_LocationTypeBackend(location),
		Page:     int32(pageDb),
		PageSize: int32(pageSizeDb),
	}
	res, err := srv.Db().QueryBannerList(context.Background(), &msg)
	if err != nil {
		log.Println("QueryBannerList ", err)
		srv.Json("服务器异常", http.StatusInternalServerError, c)
		return
	}
	// log.Printf("res is %+v", res)
	srv.Log(c, msg, "查询轮播图列表", core.QUERY_OPERATION)
	srv.Json(res, http.StatusOK, c)
}

// HandleUpdateBanner 更新Banner信息
func (srv *OptionsService) HandleUpdateBanner(c *gin.Context) {
	msg := dbproto.Banner{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Println("BindJSON", err)
		srv.Json("数据格式错误", http.StatusBadRequest, c)
		return
	}
	log.Printf("msg is %+v\n", msg)
	rid, err := srv.Db().UpdateBanner(context.Background(), &msg)
	if err != nil {
		log.Println("UpdateBanner error", err)
		srv.Json("服务器异常", http.StatusInternalServerError, c)
		return
	}
	log.Printf("Id is %v\n", rid.GetValue())
	srv.Log(c, msg, "更新轮播图", core.UPDATE_OPERATION)
	srv.Json("", http.StatusOK, c)
}

// HandlePreBanner 获取预览的Banner
func (srv *OptionsService) HandlePreBanner(c *gin.Context) {
	log.Println("Handle Banner Preview")

	locationStr := c.DefaultQuery("location", strconv.Itoa(int(dbproto.Banner_Location_Home)))
	location, err := strconv.Atoi(locationStr)
	if err != nil {
		log.Println("locationStr", err)
		srv.Json("数据格式错误", http.StatusBadRequest, c)
		return
	}

	log.Println("location", dbproto.Banner_LocationTypeBackend(location))
	msg := dbproto.QueryBannerArg{
		Location: dbproto.Banner_LocationTypeBackend(location),
	}
	res, err := srv.Db().QueryPreBanner(context.Background(), &msg)
	if err != nil {
		log.Println("QueryPreBanner ", err)
		srv.Json("服务器异常", http.StatusInternalServerError, c)
		return
	}
	// log.Printf("res is %+v", res)
	srv.Log(c, "", "获取预览轮播图", core.QUERY_OPERATION)
	srv.Json(res, http.StatusOK, c)
}
