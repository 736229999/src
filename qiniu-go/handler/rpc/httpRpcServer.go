package rpc

import (
	"github.com/henrylee2cn/faygo"

	"net/rpc"
)

var RegisterRpc = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	num := new(Num)
	num.A = 1
	num.B = 1

	rpc.Register(num)
	rpc.HandleHTTP()
	return nil
})

type Num struct {
	A,B int
}

func (num *Num) Add(a int,reply *int) error {
	*reply = num.A+num.B+a
	return nil
}
