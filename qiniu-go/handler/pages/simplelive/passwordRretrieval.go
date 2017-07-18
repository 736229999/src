package simplelive

import (
	"github.com/henrylee2cn/faygo"
)


var PasswordRretrieval = faygo.HandlerFunc(func(ctx *faygo.Context) error {


	return ctx.Render(200, faygo.JoinStatic("simpleLive/passwordRretrieval.html"), faygo.Map{

	})

})
