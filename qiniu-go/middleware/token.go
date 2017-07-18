package middleware

import (
	"github.com/henrylee2cn/faygo"
	"github.com/henrylee2cn/faygo/errors"
)

/*
Token
*/
var Token = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	if ctx.URL().Query().Get("token") == "" {
		return errors.New("miss token")
	}
	ctx.Log().Debugf("[ware] token:%q", ctx.QueryParam("token"))
	return nil
})
