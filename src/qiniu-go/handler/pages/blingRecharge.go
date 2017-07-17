package pages

import (
	"github.com/henrylee2cn/faygo"

	"qiniu-go/model/store"
	"strconv"
)

/**
	星钻余额和充值星钻
 */
var BlingRecharge = faygo.HandlerFunc(func(ctx *faygo.Context) error {

	uid,_ := strconv.Atoi(ctx.URL().Query().Get("uid"))
	return ctx.Render(200, faygo.JoinStatic("jidiv2/pages/blingRecharge.html"), faygo.Map{
		"USER":store.StoreUserMap[uid],
	})

})
