package main

import (
	"flag"

)

var httpPort = flag.Int("http", 7015, "http port")
var db = flag.String("db", "localhost:6012", "db agent")

func main() {
	srv := NewServer()
	srv.ConnectDb()
	go srv.ServeHTTP()
	select {}
}
