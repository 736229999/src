package handler

import (
	"github.com/henrylee2cn/faygo"
)

/*
	将微信证书绑定到服务器上
*/
var Micha = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	//return ctx.Render(200, faygo.JoinStatic("index.html"), faygo.Map{
	//	"TITLE":   "faygo",
	//	"VERSION": faygo.VERSION,
	//	"CONTENT": "Welcome To Faygo",
	//	"AUTHOR":  "HenryLee",
	//})
	return ctx.String(200,"FfXgomWspAkznBQA")
})

/**
	将另一个微信证书也绑定到服务器上，说明一个问题：比如在登录的时候让用户选取某个公众号进行登录这种方式是可行的！！！！！！！！！！！！！！！！
 */
var Jidi = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	//return ctx.Render(200, faygo.JoinStatic("index.html"), faygo.Map{
	//	"TITLE":   "faygo",
	//	"VERSION": faygo.VERSION,
	//	"CONTENT": "Welcome To Faygo",
	//	"AUTHOR":  "HenryLee",
	//})
	return ctx.String(200,"Z2Q4pyfE628bSDwp")
})


