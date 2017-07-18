package statistics

import (
	"github.com/caojunxyz/mimi-admin/backend/core"
	"github.com/caojunxyz/mimi-admin/dbadminagent/proto"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"log"
	"net/http"
)

type StatisticsService struct {
	core.Service
}

//统计今日收入，历史总收入，今日订单数，历史总订单数.
func (srv *StatisticsService) HandleOrderAndIncome(c *gin.Context) {

	data, err := srv.Db().QueryOrderAndIncome(context.Background(), &dbproto.Nil{})
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	srv.Log(c, nil, "获取控制台的订单和收入统计", core.QUERY_OPERATION)
	srv.Json(data, http.StatusOK, c)
}
