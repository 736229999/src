package logs

import (
	"github.com/caojunxyz/mimi-admin/backend/core"
	"github.com/caojunxyz/mimi-admin/dbadminagent/proto"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"log"
	"net/http"
)

type LogsService struct {
	core.Service
}

//日志列表.
func (srv *LogsService) HandleLogList(c *gin.Context) {

	msg := dbproto.LogReply{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	log.Printf("msg:%+v\n", msg)

	msg.UserId = int64(srv.GetUserInfo(c, "id").(float64))

	logReply, err := srv.Db().QueryLog(context.Background(), &msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	logReply.Account = msg.Account
	logReply.Size = msg.Size
	logReply.Page = msg.Page

	srv.Log(c, msg, "日志列表", core.QUERY_OPERATION)
	srv.Json(logReply, http.StatusOK, c)
}
