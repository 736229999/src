package statistics

import (
	"github.com/caojunxyz/mimi-admin/backend/core"
	"github.com/caojunxyz/mimi-admin/dbadminagent/proto"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"log"
	"net/http"
)

//控制台获取新增用户和用户总数.
func (srv *StatisticsService) HandleUserStatistics(c *gin.Context) {

	userStatistics, err := srv.Db().QueryUserStatisticsNum(context.Background(), &dbproto.Nil{})
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	srv.Log(c, nil, "获取新增的用户和用户总数", core.QUERY_OPERATION)
	srv.Json(userStatistics, http.StatusOK, c)
}
