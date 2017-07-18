package handler

import (

	//	"mime/multipart"
	"github.com/henrylee2cn/faygo"
)
var test = []string{"name","123"}
var Token int = 123
/*
	微信登录
*/
type WechatLogin struct {
	LoginType string `param:"<in:query> <required> <desc:登录类型，有wechat app h5>"`
	UnionId   string `param:"<in:query>  <desc:微信unionId> "`
	Phone   string `param:"<in:query>  <desc:电话号> "`
	VerifyCode string `param:"<in:query>  <desc:短信验证码> "`
}

//登录入口，通过LoginType来判断是wechat还是app登录
type Login struct {
	LoginType string `param:"<in:query> <required> <desc:登录类型，有wechat app>"`
	Unionid   string `param:"<in:query> <desc:微信unionid> "`

}

/**
	登录路由
 */
func (l *Login) Serve(ctx *faygo.Context) error {

	return ctx.JSON(200,"chenggong", true)
}


/**
   登录api的文档
 */
func (t *Login) Doc() faygo.Doc {
	return faygo.Doc{
		Note:   "按类型登录",
		Return: "// JSON\n{}",
	}
}
