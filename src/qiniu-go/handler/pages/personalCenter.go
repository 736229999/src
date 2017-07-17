package pages

import (
	"github.com/henrylee2cn/faygo"
	"fmt"
	"qiniu-go/model/store"
	"strconv"
)

/**
	个人中心页面
 */
var PersonalCenter = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	uid,_ := strconv.Atoi(ctx.URL().Query().Get("uid"))
	fmt.Println(uid)
	return ctx.Render(200, faygo.JoinStatic("jidiv2/pages/personalCenter.html"), faygo.Map{
		"USER":store.StoreUserMap[uid],
	})

})
