package pages

import (
	"github.com/henrylee2cn/faygo"
	"qiniu-go/model/store"
	"strconv"
)

/********************个人设置******************/

var PersonalSet = faygo.HandlerFunc(func(ctx *faygo.Context) error {

	uid,_ := strconv.Atoi(ctx.URL().Query().Get("uid"))
	return ctx.Render(200, faygo.JoinStatic("jidiv2/pages/set.html"), faygo.Map{
		"USER":store.StoreUserMap[uid],
	})

})
