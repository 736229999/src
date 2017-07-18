package simplelive

import (
	"github.com/henrylee2cn/faygo"
)


var Index = faygo.HandlerFunc(func(ctx *faygo.Context) error {



	return ctx.Render(200, faygo.JoinStatic("simpleLive/index.html"), faygo.Map{

	})

})
