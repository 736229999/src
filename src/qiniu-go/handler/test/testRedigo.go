package test

//import (
//	"github.com/henrylee2cn/faygo"
//	"fmt"
//	"github.com/henrylee2cn/qiniu-go/util"
//	"github.com/garyburd/redigo/redis"
//)
//
///*
//Test test struct handler
//*/
//type TestRedigo struct {
//
//	Name   string `param:"<in:query>  <desc:微信unionid> "`
//}
//
//func (t *TestRedigo) Serve(ctx *faygo.Context) error {
//	v,err := redis.String(util.RedisConn.Do("GET", "name"))
//	fmt.Println(v)
//	fmt.Println(err)
//	return ctx.JSON(200,
//		util.GetResponse{100,v}, true)
//}


