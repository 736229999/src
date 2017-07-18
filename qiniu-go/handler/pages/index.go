package pages


import (
	//"reflect"
	"github.com/henrylee2cn/faygo"
	"qiniu-go/util"
	"log"
	"qiniu-go/model/store"
)

/**
	直播首页
 */
var Index = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	//获取微信登录信息
	storeUser,storeWx := util.WxLogin(ctx)
	if storeUser != nil {
		storeUser.Signature = "我这儿设置了个性签名"
		storeUser.Id = 10
		store.StoreUserMap[10] = storeUser
		store.StoreWxMap[10] = storeWx

	}

	log.Println(storeUser)
	log.Println(storeWx)

		////获取微信js-sdk配置
		//wxJsConfig := util.WxJsSdkConfig(ctx)

		////获取微信支付的配置
		//openid = "oWb6w0fZHi_66S39pE40c3R93TR4"
		//wxPayConfig := util.WxPay(openid,"10")
	return ctx.Render(200, faygo.JoinStatic("jidiv2/index.html"), faygo.Map{
		"USER":storeUser,
		//"WXJSCONFIG":wxJsConfig,
		//"WXPAYCONFIG":wxPayConfig,
	})

})