package simplelive

import (
	"github.com/henrylee2cn/faygo"
)

/********************收货地址列表******************/

var Login = faygo.HandlerFunc(func(ctx *faygo.Context) error {


	return ctx.Render(200, faygo.JoinStatic("simpleLive/login.html"), faygo.Map{

	})

})
