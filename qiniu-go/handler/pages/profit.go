package pages

import (
	"github.com/henrylee2cn/faygo"
	"fmt"
	"strconv"
	"qiniu-go/model/store"
)

/**
	收益
 */
var Profit = faygo.HandlerFunc(func(ctx *faygo.Context) error {


	uid,_ := strconv.Atoi(ctx.URL().Query().Get("uid"))
	fmt.Println(store.StoreUserMap[uid])
	return ctx.Render(200, faygo.JoinStatic("jidiv2/pages/income_anchor.html"), faygo.Map{
		"USER":store.StoreUserMap[uid],
	})

})
