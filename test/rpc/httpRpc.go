package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/rpc"
)

type Args struct {
	A,B int
}

type Quotient struct {
	Quo,Rem int
}

type Arith int

func (t *Arith) Multiply(arg *Args,reply *int)error{
	*reply = arg.A * arg.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("除数不能为0") }
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil

}
func main(){
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()

	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
