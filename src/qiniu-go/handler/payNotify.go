package handler

import (
	"github.com/henrylee2cn/faygo"
	"log"
	"io/ioutil"
)
var PayNotify = faygo.HandlerFunc(func(ctx *faygo.Context) error {


	data,err := ioutil.ReadAll(ctx.R.Body)
	log.Println(string(data))
	//这个时候就将这个xml交给php去处理

	log.Println(err)
	ctx.R.Body.Close()
	return err
})
