package football

import (
	"github.com/caojunxyz/mimi-admin/backend/core"
	"github.com/caojunxyz/mimi-admin/dbadminagent/proto"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"strconv"
)

func (srv *FootballService) HandleOpencaiListGet(c *gin.Context) {
	log.Println("Handle OpencaiList Get")
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pagesize", "10")
	date := c.DefaultQuery("date", "0")

	dateDb, err := strconv.ParseInt(date, 10, 0)
	if err != nil {
		log.Println("error", err)
		srv.Json("时间参数错误", http.StatusForbidden, c)
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
	msg := dbproto.QueryOpencaiArg{
		Date:     dateDb,
		Page:     pageDb,
		PageSize: pageSizeDb,
	}
	log.Printf("msg: %+v", msg)
	res, err := srv.Db().QueryFbResult(context.Background(), &msg)
	if err != nil {
		log.Println("QueryGamesList error", err)
		srv.Json("服务器异常", http.StatusInternalServerError, c)
		return
	}
	srv.Log(c, "", "获取一条赔率信息", core.QUERY_OPERATION)
	srv.Json(res, http.StatusOK, c)
}

func (srv *FootballService) GetOpencaiById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}
	res, err := srv.Db().QueryFbResultById(context.Background(), &dbproto.IntValue{Value: id})
	if err != nil {
		log.Println("QueryNewsById error", err)
		srv.Json("服务器异常", http.StatusInternalServerError, c)
		return
	}
	srv.Log(c, id, "获取一条比赛信息", core.QUERY_OPERATION)
	srv.Json(res, http.StatusOK, c)
}

func (srv *FootballService) HandleOpencaiAdd(c *gin.Context) {
	log.Println("Handle Game Add")
	msg := dbproto.FbGameresult{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Println("BindJSON", err)
		srv.Json("数据格式错误", http.StatusBadRequest, c)
		return
	}

	log.Printf("msg is %+v\n", msg)
	rid, err := srv.Db().CreateFbResult(context.Background(), &msg)
	if err != nil {
		log.Println("InsertGame ", err)
		srv.Json(err, http.StatusInternalServerError, c)
		return
	}
	log.Printf("Opencai Id is %v\n", rid.GetValue())
	srv.Log(c, msg, "添加足彩开奖信息", core.ADD_OPERATION)
	srv.Json(rid, http.StatusOK, c)
}
func (srv *FootballService) HandleOpencaiUpdate(c *gin.Context) {
	log.Println("Handle Odds Update")
	msg := dbproto.FbGameresult{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Println("BindJSON", err)
		srv.Json("数据格式错误", http.StatusBadRequest, c)
		return
	}
	log.Printf("msg is %+v\n", msg)
	//共用一个db方法
	rid, err := srv.Db().CreateFbResult(context.Background(), &msg)
	if err != nil {
		log.Println("UpdateNews error", err)
		srv.Json("服务器异常", http.StatusInternalServerError, c)
		return
	}
	srv.Log(c, msg, "更新竞彩开奖消息", core.UPDATE_OPERATION)
	log.Printf("opencai affect is %v\n", rid.GetValue())
	srv.Json("", http.StatusOK, c)
}
