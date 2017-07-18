package main

import (
	"flag"
	"fmt"
	"log"
)

var httpPort = flag.Int("http", 7017, "http port")
var grpcPort = flag.String("activity", "localhost:6006", "db activity agent")


func main() {
	fmt.Println("活动专区...")
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	activityServer := NewServer()
	activityServer.connectDb()
	activityServer.ServeHTTP()
}






