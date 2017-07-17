package api

import "github.com/henrylee2cn/faygo"
import (
	"fmt"
	"io/ioutil"
)

var TestApi = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	ctx.R.ParseForm()
	result, _:= ioutil.ReadAll(ctx.R.Body)
	fmt.Println("id的值",ctx.R.PostFormValue("id"))
	defer ctx.R.Body.Close()
	fmt.Println(string(result))
	return nil
})