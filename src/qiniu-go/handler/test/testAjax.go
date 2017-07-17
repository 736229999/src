package test

import (
	"github.com/henrylee2cn/faygo"
)


var TestAjax = faygo.HandlerFunc(func(ctx *faygo.Context) error {
//testMap["age"] = "123"
//testMap["name"] = "jerry"
	return ctx.JSON(200,123)

	//a:=util.CurlGet("http://localhost:9090/add?one=1&other=2")
	//return ctx.String(200,"啦啦啦啦"+a)
})
