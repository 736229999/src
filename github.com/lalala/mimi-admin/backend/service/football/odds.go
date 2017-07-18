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

func (srv *FootballService) GetOddsById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}
	res, err := srv.Db().QueryOddsById(context.Background(), &dbproto.IntValue{Value: id})
	if err != nil {
		log.Println("QueryOddsById error", err)
		srv.Json("服务器异常", http.StatusInternalServerError, c)
		return
	}
	srv.Log(c, id, "获取一条赔率信息", core.QUERY_OPERATION)
	srv.Json(res, http.StatusOK, c)
}

func (srv *FootballService) HandleOddsUpdate(c *gin.Context) {
	log.Println("Handle Odds Update")
	msg := dbproto.PlayOdds{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Println("BindJSON", err)
		srv.Json("数据格式错误", http.StatusBadRequest, c)
		return
	}
	log.Printf("msg is %+v\n", msg)

	rid, err := srv.Db().UpdatePlayOdds(context.Background(), &msg)
	if err != nil {
		log.Println("UpdateNews error", err)
		srv.Json("服务器异常", http.StatusInternalServerError, c)
		return
	}
	log.Printf("affect is %v\n", rid.GetValue())
	srv.Log(c, msg, "更新赔率信息", core.UPDATE_OPERATION)
	srv.Json("", http.StatusOK, c)
}

func (srv *FootballService) HandleOddsAdd(c *gin.Context) {
	log.Println("Handle Game Add")
	msg := dbproto.PlayOdds{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Println("BindJSON", err)
		srv.Json("数据格式错误", http.StatusBadRequest, c)
		return
	}
	if len(msg.Spf) == 0 || len(msg.Rqspf) == 0 || len(msg.Zjqs) == 0 || len(msg.Bf) == 0 || len(msg.Bqc) == 0 {
		srv.Json("赔率参数不能为空", http.StatusForbidden, c)
		return
	}

	log.Printf("msg is %+v\n", msg)
	rid, err := srv.Db().CreatePlayOdds(context.Background(), &msg)
	if err != nil {
		log.Println("InsertOdds ", err)
		srv.Json(err, http.StatusInternalServerError, c)
		return
	}
	log.Printf("Id is %v\n", rid.GetValue())
	srv.Log(c, msg, "添加赔率信息", core.ADD_OPERATION)
	srv.Json(rid, http.StatusOK, c)
	return
}
