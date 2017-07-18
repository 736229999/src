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

func (srv *FootballService) HandleLeagueListGet(c *gin.Context) {
	log.Println("Handle LeagueList Get")
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pagesize", "10")
	leagueName := c.DefaultQuery("name", "")
	log.Println(page, pageSize)

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

	msg := dbproto.QueryFbLeagueArg{
		Name:     leagueName,
		Page:     pageDb,
		PageSize: pageSizeDb,
	}
	log.Printf("msg: %+v", msg)
	res, err := srv.Db().QueryFbLeague(context.Background(), &msg)
	if err != nil {
		log.Println("QueryGamesList error", err)
		srv.Json("服务器异常", http.StatusInternalServerError, c)
		return
	}
	srv.Json(res, http.StatusOK, c)
}

func (srv *FootballService) HandleLeagueAdd(c *gin.Context) {
	log.Println("Handle League Add")
	msg := dbproto.FbLeagueInfo{}
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
	rid, err := srv.Db().CreateFbLeague(context.Background(), &msg)
	if err != nil {
		log.Println("Insert FbLeague ", err)
		srv.Json(err, http.StatusInternalServerError, c)
		return
	}
	log.Printf("Id is %v\n", rid.GetValue())
	srv.Log(c, msg, "添加球队信息", core.ADD_OPERATION)
	srv.Json(rid, http.StatusOK, c)
}

func (srv *FootballService) HandleLeagueUpdate(c *gin.Context) {
	log.Println("Handle League Update")
	msg := dbproto.FbLeagueInfo{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Println("BindJSON", err)
		srv.Json("数据格式错误", http.StatusBadRequest, c)
		return
	}
	log.Printf("msg is %+v\n", msg)
	rid, err := srv.Db().CreateFbLeague(context.Background(), &msg)
	if err != nil {
		log.Println("Update Fbleague error", err)
		srv.Json("服务器异常", http.StatusInternalServerError, c)
		return
	}
	log.Printf("Fbleague affect is %v\n", rid.GetValue())
	srv.Log(c, msg, "更新球队信息", core.UPDATE_OPERATION)
	srv.Json("", http.StatusOK, c)
}
