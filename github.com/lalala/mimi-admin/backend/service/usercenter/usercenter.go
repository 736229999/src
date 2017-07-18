package usercenter

import (
	"github.com/caojunxyz/mimi-admin/backend/core"
	"github.com/caojunxyz/mimi-admin/dbadminagent/proto"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"strconv"
)

type UsercenterService struct {
	core.Service
}

//获取usercenter的用户列表.
func (srv *UsercenterService) HandleUserList(c *gin.Context) {

	msg := &dbproto.UsercenterList{}
	if err := c.BindJSON(&msg); err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	ucList, err := srv.Db().QueryUsercenterList(context.Background(), msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	ucList.Size = msg.GetSize()
	ucList.Page = msg.GetPage()
	ucList.Nickname = msg.GetNickname()
	ucList.Phone = msg.GetPhone()
	srv.Json(ucList, http.StatusOK, c)
}

//获取user的详细信息.
func (srv *UsercenterService) HandleUserDetail(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	userInfo, err := srv.Db().QueryUsercenterDetail(context.Background(), &dbproto.IntValue{Value: int64(id)})
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}
	srv.Json(userInfo, http.StatusOK, c)
}

//获取用户的金额统计.
func (srv *UsercenterService) HandleUserFund(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	fund, err := srv.Db().QueryUsercenterFundById(context.Background(), &dbproto.IntValue{Value: int64(id)})
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	//获取充值列表数据.
	rechargeList, err := srv.Db().QueryUsercenterRechargeById(context.Background(), &dbproto.IntValue{Value: int64(id)})
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}
	fund.RechargeList = rechargeList.RechargeList

	srv.Json(fund, http.StatusOK, c)
}

//获取用户的提现记录.
func (srv *UsercenterService) HandleUserWithdraw(c *gin.Context) {

	// id, err := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// 	log.Printf("%+v\n", err)
	// 	srv.Json("获取失败", http.StatusForbidden, c)
	// 	return
	// }

	// withdrawList, err := srv.Db().QueryUsercenterWithdrawById(context.Background(), &dbproto.IntValue{Value:int64(id)})
	// if err != nil {
	// 	log.Printf("%+v\n", err)
	// 	srv.Json("获取失败", http.StatusForbidden, c)
	// 	return
	// }

	// srv.Json(withdrawList, http.StatusOK, c)
}
