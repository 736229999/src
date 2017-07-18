package main

import (
	"flag"
	"github.com/caojunxyz/mimi-admin/backend/router"
	"log"
)

var dbAgentAddress = flag.String("dbAgent", "127.0.0.1:10000", "dbAgent address")
var adminPort = flag.Int("adminPort", 11111, "admin port")

func main() {

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("grpc:", *dbAgentAddress)
	log.Println("admin port:", *adminPort)

	err := router.Init(*adminPort)
	if err != nil {
		log.Printf("%+v\n", err)
	}
}
