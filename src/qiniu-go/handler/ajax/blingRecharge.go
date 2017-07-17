package ajax

import (
	"github.com/henrylee2cn/faygo"
	"log"
	"qiniu-go/util"
	"qiniu-go/model/store"
)

/**
	获取充值的金额
*/
var GetRechargeMoney = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	money := ctx.URL().Query().Get("money")
	log.Println("传过来的钱：",money)

	openid := store.StoreWxMap[10].OpenId
	log.Println(openid)
	wxPayConfig := util.WxPay(openid,money)



	return ctx.JSON(200,wxPayConfig,true)
})
