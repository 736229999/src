package main

import (
	"github.com/henrylee2cn/faygo"
	"qiniu-go/router"
)

func main() {
	router.Route(faygo.New("qiniu-go"))
	faygo.Run()
}
