package user

import (
	"github.com/henrylee2cn/faygo"
	"qiniu-go/util"
	"qiniu-go/model/apiModel"
)

//个人信息

type WxLogin struct {
	UnionId	 string	    `param:"<in:query><required> <desc:微信unionid>"`

}
/**
	获取个人信息的api
 */
func (t *WxLogin) Serve(ctx *faygo.Context) error {

	//如果unionid存在就直接返回，不存在就插入数据库，再返回
	user  := new(apiModel.User)
	defer func() {
		user = nil
	}()
	flag,err := util.Engine.Where("wx_unionid=?",t.UnionId).Get(user)
	util.CheckError(err)
	if flag {
		//就返回数据库中的该条数据
		return ctx.JSON(200,"true",true)
	}else {
		//插入一条新的数据再返回

		return ctx.JSON(200,"该用户不存在",true)
	}

}
func (t *WxLogin) Doc() faygo.Doc {
	return faygo.Doc{
		Note:   "微信登录的api",
		Return: "返回个人信息的json",
	}
}




