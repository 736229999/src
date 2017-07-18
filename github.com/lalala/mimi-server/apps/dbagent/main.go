package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/caojunxyz/mimi-server/apps/dbagent/activity"
	"github.com/caojunxyz/mimi-server/apps/dbagent/buycai"
	"github.com/caojunxyz/mimi-server/apps/dbagent/discover"
	"github.com/caojunxyz/mimi-server/apps/dbagent/notify"
	"github.com/caojunxyz/mimi-server/apps/dbagent/options"
	"github.com/caojunxyz/mimi-server/apps/dbagent/recharge"
	"github.com/caojunxyz/mimi-server/apps/dbagent/thirdapi"
	"github.com/caojunxyz/mimi-server/apps/dbagent/usercenter"
	_ "github.com/lib/pq"
)

var buycaiPort = flag.Int("buycai", 6007, "db buycai grpc port")
var ucPort = flag.Int("uc", 6008, "db usercenter grpc port")
var rechargePort = flag.Int("recharge", 7000, "db recharge grpc port")
var discoverPort = flag.Int("discover", 6011, "db discover grpc port")
var notifyPort = flag.Int("notify", 6010, "db notify grpc port")
var optionsPort = flag.Int("options", 6012, "db options grpc port")
var thirdapiPort = flag.Int("thirdapi", 6013, "db thirdapi grpc port")
var activityPort = flag.Int("activity", 6006, "db activity grpc port")

func main() {
	flag.Parse()
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	fmt.Println("dbagent...")
	go buycai.NewAgent().Run(*buycaiPort)
	go usercenter.NewAgent().Run(*ucPort)
	go recharge.NewAgent().Run(*rechargePort)
	go discover.NewAgent().Run(*discoverPort)
	go notify.NewAgent().Run(*notifyPort)
	go options.NewAgent().Run(*optionsPort)
	go thirdapi.NewAgent().Run(*thirdapiPort)
	go activity.NewAgent().Run(*activityPort)
	select {}
}
