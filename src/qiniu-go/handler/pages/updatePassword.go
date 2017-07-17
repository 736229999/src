package pages

import (
	"github.com/henrylee2cn/faygo"
	"qiniu-go/model/store"
)

/********************修改密码******************/

var UpdatePassword = faygo.HandlerFunc(func(ctx *faygo.Context) error {

	//获取微信js-sdk配置
	storeUser := new(store.StoreUser)
	storeUser.Id = 1
	storeUser.NickName = "jerry"
	storeUser.Avatar = "123"
	storeUser.Signature = "我的签名"
	return ctx.Render(200, faygo.JoinStatic("jidiv2/pages/passWord.html"), faygo.Map{
		"USER" :storeUser,
	})

})
