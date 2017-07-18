package statistics

import (
	"github.com/caojunxyz/mimi-admin/dbadminagent/proto"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"log"
	"net/http"
)

func (srv *StatisticsService) HandleGetRechargeList(c *gin.Context) {

	msg := &dbproto.RechargeOrderList{}
	if err := c.BindJSON(&msg); err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}
	recharge, err := srv.Db().QueryRechargeList(context.Background(), msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return

	}
	recharge.StartTime = msg.GetStartTime()
	recharge.EndTime = msg.GetEndTime()
	recharge.Size = msg.GetSize()
	recharge.Page = msg.GetPage()

	srv.Json(recharge, http.StatusOK, c)
}

//根据月份获取充值记录.
func (srv *StatisticsService) HandleGetRechargeByMonth(c *gin.Context) {

	recharge, err := srv.Db().QueryRechargeListByMonth(context.Background(), &dbproto.Nil{})
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json(" 获取失败", http.StatusForbidden, c)
		return
	}

	srv.Json(recharge, http.StatusOK, c)
}

//根据年月获取充值数据.
func (srv *StatisticsService) HandleGetRechargeByYear(c *gin.Context) {

	recharge, err := srv.Db().QueryRechangeListByYear(context.Background(), &dbproto.Nil{})
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	srv.Json(recharge, http.StatusOK, c)
}
