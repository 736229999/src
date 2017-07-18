package main

import (
	"flag"
	"fmt"
	"github.com/caojunxyz/mimi-admin/dbadminagent/admin"
	_ "github.com/lib/pq"
	"log"
)

var adminPort = flag.Int("admin", 10000, "db admin port")

func main() {
	flag.Parse()
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	fmt.Println("dbagent...")
	go admin.NewAgent().Run(*adminPort)
	select {}
}
