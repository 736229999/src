package user

import (
	"github.com/caojunxyz/mimi-admin/backend/core"
	"github.com/gin-gonic/gin"
	"log"

	"fmt"
	"github.com/caojunxyz/mimi-admin/backend/auth"
	"github.com/caojunxyz/mimi-admin/backend/utils"
	"github.com/caojunxyz/mimi-admin/dbadminagent/proto"
	"golang.org/x/net/context"
	"net/http"
	"strconv"
)

type UserService struct {
	core.Service
}

func (srv *UserService) HandleLogin(c *gin.Context) {

	email := c.PostForm("email")
	password := c.PostForm("password")

	if len(email) < 1 {
		srv.Json("邮箱不能为空", http.StatusForbidden, c)
		return
	}

	if len(password) < 1 {
		srv.Json("密码不能为空", http.StatusForbidden, c)
		return
	}

	if !utils.IsEmail(email) {
		srv.Json("邮箱格式错误", http.StatusForbidden, c)
		return
	}

	if !utils.IsPassword(password) {
		srv.Json("密码格式不对", http.StatusForbidden, c)
		return
	}

	userInfo, err := srv.Db().QueryUserInfoByEmail(context.Background(), &dbproto.StringValue{Value: email})
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("用户名或者密码错误", http.StatusForbidden, c)
		return
	}

	//加盐加密.
	if userInfo.GetPassword() != srv.MD5(password+userInfo.GetSalt()) {
		srv.Json("用户名或者密码不正确", http.StatusForbidden, c)
		return
	}

	//将用户信息记录session.
	dist := make(map[string]interface{})
	dist["id"] = userInfo.GetId()
	dist["email"] = userInfo.GetEmail()
	dist["username"] = userInfo.GetUsername()
	dist["pasword"] = userInfo.GetPassword()
	dist["salt"] = userInfo.GetSalt()
	dist["status"] = userInfo.GetStatus()
	dist["mobile"] = userInfo.GetMobile()
	dist["create_time"] = userInfo.GetCreateTime()
	dist["creator"] = userInfo.GetCreator()
	dist["register_ip"] = userInfo.GetRegisterIp()

	token, err := auth.CreateToekn(dist)
	if err != nil {
		log.Println("%+v\n", err)
		srv.Json("登录失败", http.StatusForbidden, c)
		return
	}
	c.Request.Header.Set("Authorization", token)

	srv.Log(c, dist, "用户登录", core.QUERY_OPERATION)
	srv.Json(token, http.StatusOK, c)
	return
}

//更新用户.
func (srv *UserService) HandleUserUpdate(c *gin.Context) {

	mobile := c.PostForm("mobile")
	username := c.PostForm("username")
	//userId := srv.SessionGet(c, "id")

	if len(mobile) < 1 {
		srv.Json("手机号码不能为空", http.StatusForbidden, c)
		return
	}

	if len(username) < 1 {
		srv.Json("昵称不能为空", http.StatusForbidden, c)
		return
	}

	if !utils.IsMobile(mobile) {
		srv.Json("手机号格式不正确", http.StatusForbidden, c)
		return
	}

	if !utils.IsName(username) {
		srv.Json("昵称不正确", http.StatusForbidden, c)
		return
	}

	userInfo := &dbproto.AdminUserInfoArg{
		Id:       0,
		Username: username,
		Mobile:   mobile,
	}
	//更新.
	userInfo, err := srv.Db().SetUserInfoById(context.Background(), userInfo)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("服务器异常", http.StatusForbidden, c)
		return
	}

	//修改session.
	dist := make(map[string]interface{})
	dist["id"] = userInfo.GetId()
	dist["email"] = userInfo.GetEmail()
	dist["username"] = userInfo.GetUsername()
	dist["pasword"] = userInfo.GetPassword()
	dist["salt"] = userInfo.GetSalt()
	dist["status"] = userInfo.GetStatus()
	dist["mobile"] = userInfo.GetMobile()
	dist["create_time"] = userInfo.GetCreateTime()
	dist["creator"] = userInfo.GetCreator()
	dist["register_ip"] = userInfo.GetRegisterIp()
	//srv.SessionSet(c, dist)

	srv.Json(dist, http.StatusOK, c)
}

//用户列表.
func (srv *UserService) HandleUserList(c *gin.Context) {

	userList, err := srv.Db().QueryUserList(context.Background(), &dbproto.Nil{})
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("服务器异常", http.StatusForbidden, c)
		return
	}

	log.Printf("%+v\n", userList.UserList)
	srv.Json(userList.GetUserList(), http.StatusOK, c)
}

//权限列表.
func (srv *UserService) HandleUserPrivilegeList(c *gin.Context) {

	privilegeList, err := srv.Db().QueryPrivilegeList(context.Background(), &dbproto.Nil{})
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("服务器异常", http.StatusForbidden, c)
		return
	}

	srv.Json(privilegeList, http.StatusOK, c)
}

//添加权限.
func (srv *UserService) HandleUserPrivilegeAdd(c *gin.Context) {

	msg := &dbproto.Privilege{}
	if err := c.BindJSON(&msg); err != nil {
		log.Printf("%+v\n", err)
		srv.Json("添加失败", http.StatusForbidden, c)
		return
	}

	msg.Creator = fmt.Sprintf("%s", srv.GetUserInfo(c, "username"))

	_, err := srv.Db().InsertPrivilege(context.Background(), msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("添加失败", http.StatusForbidden, c)
		return
	}

	srv.Json("添加成功", http.StatusOK, c)

	//name := c.PostForm("name")
	//key := c.PostForm("key")
	//path := c.PostForm("path")
	//
	//if len(name) < 1 {
	//	srv.Json("权限名称不能为空", http.StatusForbidden, c)
	//	return
	//}
	//if len(key) < 1 {
	//	srv.Json("权限的key不能为空", http.StatusForbidden, c)
	//	return
	//}
	//if len(path) < 1 {
	//	srv.Json("权限路径不能为空", http.StatusForbidden, c)
	//	return
	//}
	//
	//privileges := &dbproto.AdminPrivileges{
	//	Name:name,
	//	Key:key,
	//	Path:path,
	//	Creator:0,
	//	CreateTime:time.Now().Unix(),
	//}
	//
	//_, err := srv.Db().InsertPrivileges(context.Background(), privileges)
	//if err != nil {
	//	log.Printf("%+v\n", err)
	//	srv.Json("添加失败", http.StatusForbidden, c)
	//	return
	//}
	//
	//srv.Json("添加成功", http.StatusOK, c)
}

//权限编辑.
func (srv *UserService) HandleUserPrivilegesEdit(c *gin.Context) {

	idStr := c.PostForm("id")
	name := c.PostForm("name")
	key := c.PostForm("key")
	path := c.PostForm("path")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("数据错误", http.StatusForbidden, c)
		return
	}

	if id < 1 {
		log.Printf("%+v\n", err)
		srv.Json("数据错误", http.StatusForbidden, c)
		return
	}

	if len(name) < 1 {
		srv.Json("权限名称不能为空", http.StatusForbidden, c)
		return
	}

	if len(key) < 1 {
		srv.Json("权限key不能为空", http.StatusForbidden, c)
		return
	}

	if len(path) < 1 {
		srv.Json("权限路径不能为空", http.StatusForbidden, c)
		return
	}

	privileges := &dbproto.AdminPrivileges{
		Id:   id,
		Name: name,
		Key:  key,
		Path: path,
	}

	_, err = srv.Db().SetPrivileges(context.Background(), privileges)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("服务器异常", http.StatusForbidden, c)
		return
	}

	srv.Json("服务器异常", http.StatusOK, c)
}

//删除权限.
func (srv *UserService) HandleUserPrivilegesDelete(c *gin.Context) {

	idStr := c.PostForm("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("数据错误", http.StatusForbidden, c)
		return
	}

	if id < 1 {
		srv.Json("数据错误", http.StatusForbidden, c)
		return
	}

	_, err = srv.Db().DeletePrivileges(context.Background(), &dbproto.IntValue{Value: id})

	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("服务器异常", http.StatusForbidden, c)
		return
	}

	srv.Json("删除成功", http.StatusOK, c)
}

//角色列表.
func (srv *UserService) HandleUserRoleList(c *gin.Context) {

	roleList, err := srv.Db().QueryRoleList(context.Background(), &dbproto.Nil{})
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("服务器异常", http.StatusForbidden, c)
		return
	}

	srv.Json(roleList, http.StatusOK, c)
}

//添加角色.
func (srv *UserService) HandleUserRolesAdd(c *gin.Context) {

	//userId := srv.SessionGet(c,"id")
	//
	//privi_ids := c.PostForm("privi_ids")

}
