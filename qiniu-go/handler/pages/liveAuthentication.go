package pages

import (
	"github.com/henrylee2cn/faygo"
	"qiniu-go/util"
	"qiniu-go/model/store"
)

/********************直播认证页面******************/

var LiveAuthentication = faygo.HandlerFunc(func(ctx *faygo.Context) error {

	//获取微信js-sdk配置
	WXJS := util.WxJsSdkConfig(ctx)
	storeUser := new(store.StoreUser)
	storeUser.Id = 123
	return ctx.Render(200, faygo.JoinStatic("jidiv2/pages/liveAuthentication.html"), faygo.Map{
		"WXJSCONFIG":WXJS,
		"USER" :storeUser,
	})

})
