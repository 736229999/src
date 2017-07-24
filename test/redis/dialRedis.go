package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)



func main() {
	var RedisConn redis.Conn
	if RedisConn == nil {
		RedisConn,_=redis.Dial("tcp", "127.0.0.1:6379")
	}
	name ,err := redis.String(RedisConn.Do("GET","topicName"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("获取到的name：",name)

}

