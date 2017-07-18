package options

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"strconv"

	"github.com/caojunxyz/mimi-admin/backend/core"
	dbproto "github.com/caojunxyz/mimi-admin/dbadminagent/proto"
)

// HandleNewsAdd 用于添加一条新闻
func (srv *OptionsService) HandleNewsAdd(c *gin.Context) {
	log.Println("Handle News Add")
	msg := dbproto.News{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Println("BindJSON", err)
		srv.Json("数据格式错误", http.StatusBadRequest, c)
		return
	}
	if len(msg.Title) < 1 {
		srv.Json("标题不能为空", http.StatusForbidden, c)
		return
	}
	if len(msg.Content) < 1 {
		srv.Json("内容不能为空", http.StatusForbidden, c)
		return
	}
	if len(msg.Description) < 1 {
		srv.Json("封面描述不能为空", http.StatusForbidden, c)
		return
	}
	log.Printf("msg is %+v\n", msg)
	rid, err := srv.Db().CreateNews(context.Background(), &msg)
	if err != nil {
		log.Println("InsertNews ", err)
		srv.Json(err, http.StatusInternalServerError, c)
		return
	}
	log.Printf("Id is %v\n", rid)
	srv.Log(c, msg, "添加新闻", core.ADD_OPERATION)
	srv.Json("", http.StatusOK, c)
	return
}

// HandleNewsListGet 获取新闻列表
func (srv *OptionsService) HandleNewsListGet(c *gin.Context) {
	log.Println("Handle NewsList Get")
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pagesize", "10")
	start := c.DefaultQuery("start", "0")
	end := c.DefaultQuery("end", "0")
	log.Println(page, pageSize, start, end)
	startDb, err := strconv.ParseInt(start, 10, 0)
	if err != nil {
		log.Println("error", err)
		srv.Json("开始时间参数错误", http.StatusForbidden, c)
		return
	}
	endDb, err := strconv.ParseInt(end, 10, 0)
	if err != nil {
		log.Println("error", err)
		srv.Json("结束时间参数错误", http.StatusForbidden, c)
		return
	}
	pageDb, err := strconv.ParseInt(page, 10, 0)
	if err != nil {
		log.Println("error", err)
		srv.Json("页码参数错误", http.StatusForbidden, c)
		return
	}
	pageSizeDb, err := strconv.ParseInt(pageSize, 10, 0)
	if err != nil {
		log.Println("error", err)
		srv.Json("每页尺寸参数错误", http.StatusForbidden, c)
		return
	}
	msg := dbproto.QueryNewsArg{
		Title:    c.Query("title"),
		Author:   c.Query("author"),
		Start:    startDb,
		End:      endDb,
		Page:     pageDb,
		PageSize: pageSizeDb,
	}
	log.Printf("msg: %+v", msg)
	log.Printf("msg title: %+v", c.Query("title"))
	res, err := srv.Db().QueryNewsList(context.Background(), &msg)
	if err != nil {
		log.Println("QueryNewsList error", err)
		srv.Json("服务器异常", http.StatusInternalServerError, c)
		return
	}
	// log.Printf("res %+v\n", res)
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
	srv.Log(c, msg, "查询新闻列表", core.QUERY_OPERATION)
	c.JSON(http.StatusOK, res)
	// srv.Json(res, http.StatusOK, c)
}

// GetNewsById 获取一条新闻
func (srv *OptionsService) GetNewsById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}
	res, err := srv.Db().QueryNewsById(context.Background(), &dbproto.NewsId{Id: id})
	if err != nil {
		log.Println("QueryNewsById error", err)
		srv.Json("服务器异常", http.StatusInternalServerError, c)
		return
	}
	// log.Printf("res is %+v", res)
	srv.Log(c, id, "获取新闻详细信息", core.QUERY_OPERATION)
	srv.Json(res, http.StatusOK, c)
}

// QueryNewsOfSelect banner管理中下拉框搜索接口
func (srv *OptionsService) QueryNewsOfSelect(c *gin.Context) {
	keyword := c.DefaultQuery("keyword", "")
	log.Println("keyword", keyword)
	msg := dbproto.QueryNewsOfSelect{
		KeyWord: keyword,
	}
	res, err := srv.Db().QueryBakendSelectOfNews(context.Background(), &msg)
	if err != nil {
		log.Println("QueryBakendSelectOfNews error", err)
		srv.Json("服务器异常", http.StatusInternalServerError, c)
		return
	}
	// log.Printf("res %+v\n", res)
	srv.Json(res, http.StatusOK, c)
}

// HandleUpdateNews 修改新闻
func (srv *OptionsService) HandleUpdateNews(c *gin.Context) {
	msg := dbproto.News{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Println("BindJSON", err)
		srv.Json("数据格式错误", http.StatusBadRequest, c)
		return
	}
	log.Printf("msg is %+v\n", msg)
	rid, err := srv.Db().UpdateNews(context.Background(), &msg)
	if err != nil {
		log.Println("UpdateNews error", err)
		srv.Json("服务器异常", http.StatusInternalServerError, c)
		return
	}
	srv.Log(c, msg, "更新新闻", core.UPDATE_OPERATION)
	log.Printf("affect is %v\n", rid.GetValue())
	srv.Json("", http.StatusOK, c)
}
