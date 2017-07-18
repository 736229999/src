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

type FootballService struct {
	core.Service
}

func (srv *FootballService) HandleGamesListGet(c *gin.Context) {
	log.Println("Handle GamesList Get")
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pagesize", "10")
	team := c.DefaultQuery("teamname", "0")
	start := c.DefaultQuery("start", "0")
	end := c.DefaultQuery("end", "0")

	if team == "" {
		team = "0"
	}
	teamDb, err := strconv.ParseInt(team, 10, 0)
	if err != nil {
		log.Println("error", err)
		srv.Json("球队参数错误", http.StatusForbidden, c)
		return
	}
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
	msg := dbproto.QueryFbGameArg{
		Team:     teamDb,
		Start:    startDb,
		End:      endDb,
		Page:     pageDb,
		PageSize: pageSizeDb,
	}
	log.Printf("msg: %+v", msg)
	res, err := srv.Db().QueryFbGame(context.Background(), &msg)
	if err != nil {
		log.Println("QueryGamesList error", err)
		srv.Json("服务器异常", http.StatusInternalServerError, c)
		return
	}
	srv.Json(res, http.StatusOK, c)
}

func (srv *FootballService) GetFbGameById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}
	res, err := srv.Db().QueryFbGameById(context.Background(), &dbproto.IntValue{Value: id})
	if err != nil {
		log.Println("QueryNewsById error", err)
		srv.Json("服务器异常", http.StatusInternalServerError, c)
		return
	}
	srv.Log(c, id, "获取一条比赛信息", core.QUERY_OPERATION)
	srv.Json(res, http.StatusOK, c)
}

func (srv *FootballService) HandleGameAdd(c *gin.Context) {
	log.Println("Handle Game Add")
	msg := dbproto.GameInfo{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Println("BindJSON", err)
		srv.Json("数据格式错误", http.StatusBadRequest, c)
		return
	}
	if msg.Id != 0 {
		srv.Json("id必须为0", http.StatusForbidden, c)
		return
	}

	log.Printf("msg is %+v\n", msg)
	rid, err := srv.Db().CreateFbGame(context.Background(), &msg)
	if err != nil {
		log.Println("InsertGame ", err)
		srv.Json(err, http.StatusInternalServerError, c)
		return
	}
	log.Printf("Id is %v\n", rid.GetValue())
	srv.Log(c, msg, "添加比赛信息", core.ADD_OPERATION)
	srv.Json(rid, http.StatusOK, c)
}

func (srv *FootballService) HandleGameUpdate(c *gin.Context) {
	log.Println("Handle Game Update")
	msg := dbproto.GameInfo{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Println("BindJSON", err)
		srv.Json("数据格式错误", http.StatusBadRequest, c)
		return
	}
	log.Printf("msg is %+v\n", msg)
	//共用一条dbagent
	rid, err := srv.Db().CreateFbGame(context.Background(), &msg)
	if err != nil {
		log.Println("UpdateNews error", err)
		srv.Json("服务器异常", http.StatusInternalServerError, c)
		return
	}
	srv.Log(c, msg, "更新比赛消息", core.UPDATE_OPERATION)
	log.Printf("affect is %v\n", rid.GetValue())
	srv.Json("", http.StatusOK, c)
}
