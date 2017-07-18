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

func (srv *FootballService) HandleTeamListGet(c *gin.Context) {
	log.Println("Handle GamesList Get")
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pagesize", "10")
	teamName := c.DefaultQuery("name", "")

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

	msg := dbproto.QueryFbTeamArg{
		Name:     teamName,
		Page:     pageDb,
		PageSize: pageSizeDb,
	}
	log.Printf("msg: %+v", msg)
	res, err := srv.Db().QueryFbTeamList(context.Background(), &msg)
	if err != nil {
		log.Println("QueryGamesList error", err)
		srv.Json("服务器异常", http.StatusInternalServerError, c)
		return
	}
	srv.Json(res, http.StatusOK, c)
}

func (srv *FootballService) GetTeamById(c *gin.Context) {
	log.Println("Handle Team Get one")
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}
	res, err := srv.Db().QueryFbTeamById(context.Background(), &dbproto.IntValue{Value: id})
	if err != nil {
		log.Println("QueryNewsById error", err)
		srv.Json("服务器异常", http.StatusInternalServerError, c)
		return
	}
	srv.Log(c, id, "获取一条比赛信息", core.QUERY_OPERATION)
	srv.Json(res, http.StatusOK, c)
}

func (srv *FootballService) HandleTeamAdd(c *gin.Context) {
	log.Println("Handle Tean Add")
	msg := dbproto.FbTeamInfo{}
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
	rid, err := srv.Db().CreateFbTeam(context.Background(), &msg)
	if err != nil {
		log.Println("InsertFbTeam ", err)
		srv.Json(err, http.StatusInternalServerError, c)
		return
	}
	log.Printf("Id is %v\n", rid.GetValue())
	srv.Log(c, msg, "添加球队信息", core.ADD_OPERATION)
	srv.Json(rid, http.StatusOK, c)
}

func (srv *FootballService) HandleTeamUpdate(c *gin.Context) {
	log.Println("Handle Team Update")
	msg := dbproto.FbTeamInfo{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Println("BindJSON", err)
		srv.Json("数据格式错误", http.StatusBadRequest, c)
		return
	}
	log.Printf("msg is %+v\n", msg)
	rid, err := srv.Db().CreateFbTeam(context.Background(), &msg)
	if err != nil {
		log.Println("Update Fbteam error", err)
		srv.Json("服务器异常", http.StatusInternalServerError, c)
		return
	}
	srv.Log(c, msg, "更新球队信息", core.UPDATE_OPERATION)
	log.Printf("Fbteam affect is %v\n", rid.GetValue())
	srv.Json("", http.StatusOK, c)
}
