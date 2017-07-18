package rpc

import (
	"github.com/henrylee2cn/faygo"

	"net/rpc"
	"shop/util"
	"log"
)

var RpcClient = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	client,err := rpc.DialHTTP("tcp","localhost:9999/httpRpc")
	util.CheckError(err)
	var reply int
	a := 1
	e := client.Call("Num.Add",a,&reply)
	util.CheckError(e)
	log.Println(a)
	return nil
})

