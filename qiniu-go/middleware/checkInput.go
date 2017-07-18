package middleware

import (
	"github.com/henrylee2cn/faygo"
)

var CheckInputData = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	//检查用户输入的数据，只能是字母+数字
	return nil
})
