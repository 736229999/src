package order

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"github.com/caojunxyz/mimi-admin/dbadminagent/proto"
	"golang.org/x/net/context"
	"github.com/caojunxyz/mimi-admin/backend/core"
)


type OrderService struct {
	core.Service
}
//查询所有用户订单
func (srv *OrderService) QueryUserOrderList(c *gin.Context) {
	msg := dbproto.UserOrderList{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	list,err := srv.Db().QueryUserOrderList(context.Background(),&msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}
	list.Size = msg.GetSize()
	list.Page = msg.GetPage()
	srv.Log(c, msg, "用户订单列表", core.QUERY_OPERATION)
	srv.Json(list,200,c)
}

//查询所有用户订单带条件
func (srv *OrderService) QueryUserOrderListWithCondition(c *gin.Context) {
	msg := dbproto.UserOrderList{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}
	log.Println(msg)
	list,err := srv.Db().QueryUserOrderListWithCondition(context.Background(),&msg)

	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}
	list.Size = msg.GetSize()
	list.Page = msg.GetPage()
	list.Condition = msg.GetCondition()
	srv.Log(c, msg, "用户订单列表", core.QUERY_OPERATION)
	srv.Json(list,http.StatusOK,c)
}
