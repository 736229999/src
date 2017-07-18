package util

//import "github.com/garyburd/redigo/redis"

func test()  {

}
//import "github.com/garyburd/redigo/redis"
//
//var RedisConn redis.Conn
//func init()  {
//	if RedisConn == nil {
//		RedisConn,_=redis.Dial("tcp", "127.0.0.1:6379")
//	}
//}


//思路：做一个连接池，里面放入3个redis的连接，根据uid%3取余来判断用户要用哪个连接

//var RedisPool  [5]redis.Conn
//func initPoll(){
//
//}