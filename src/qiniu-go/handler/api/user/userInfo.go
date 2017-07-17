package user

import (
	"github.com/henrylee2cn/faygo"
	"qiniu-go/model"
	"qiniu-go/util"
)

//个人信息

type GetUserInfo struct {
	Uid	 int	    `param:"<in:query><required> <desc:用户id>"`
}
/**
	获取个人信息的api
 */
func (t *GetUserInfo) Serve(ctx *faygo.Context) error {

	//

	var userInfo *model.Userinfo
	if userInfo == nil {
		userInfo = new(model.Userinfo)
	}
	flag,err := util.Engine.Where("user_id=?",t.Uid).Get(userInfo)
	util.CheckError(err)
	if flag {
		return ctx.JSON(200,userInfo,true)
	}
	return ctx.JSON(400,"该用户不存在",true)

}
func (t *GetUserInfo) Doc() faygo.Doc {
	return faygo.Doc{
		Note:   "获取个人信息的api",
		Return: "返回个人信息的json",
	}
}





type GetUserLevel struct {
	Uid	 int	    `param:"<in:query><required> <desc:用户id>"`
}
/**
	获取个人等级的api
 */
func (t *GetUserLevel) Serve(ctx *faygo.Context) error {

	//
	var userLevel  *model.Userlevel
	if userLevel == nil {
		userLevel = new(model.Userlevel)
	}

	flag,err := util.Engine.Where("user_id=?",t.Uid).Get(userLevel)
	util.CheckError(err)

	if flag {
		return ctx.JSON(200,userLevel,true)
	}
	return ctx.JSON(400,"该用户不存在",true)
}
func (t *GetUserLevel) Doc() faygo.Doc {
	return faygo.Doc{
		Note:   "获取个人等级的api",
		Return: "返回个人等级和经验值的json",
	}
}