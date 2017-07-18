package options

import (
	"github.com/caojunxyz/mimi-admin/backend/core"
	"github.com/caojunxyz/mimi-admin/dbadminagent/proto"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"log"
	"net/http"
)

//更新客服信息.
func (srv *OptionsService) HandleOptionsUpdateContact(c *gin.Context) {

	msg := &dbproto.Contact{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("更新失败", http.StatusForbidden, c)
		return
	}

	_, err = srv.Db().UpdateContact(context.Background(), msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("更新失败", http.StatusForbidden, c)
		return
	}

	srv.Log(c, msg, "更新客服信息", core.UPDATE_OPERATION)
	srv.Json("更新成功", http.StatusOK, c)
}

//获取客服信息.
func (srv *OptionsService) HandleOptionsQueryContact(c *gin.Context) {

	contact, err := srv.Db().QueryContact(context.Background(), &dbproto.Nil{})
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	srv.Log(c, nil, "获取客服信息", core.QUERY_OPERATION)
	srv.Json(contact, http.StatusOK, c)
}
