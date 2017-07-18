package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"google.golang.org/grpc"

	apiproto "github.com/caojunxyz/mimi-api/proto"
	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/auth"
)

var httpPort = flag.Int("http", 7007, "http port")
var grpcPort = flag.Int("grpc", 7008, "grpc port")
var dbBuycai = flag.String("dbbuycai", "localhost:6007", "db buycai agent")
var dbUc = flag.String("dbuc", "localhost:6008", "db uc agent")

func main() {
	fmt.Println("buycai...")
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	buycaiServer := NewServer()
	buycaiServer.connectDb()
	buycaiServer.initAgents()
	buycaiServer.ServeHttp()
}

type BuycaiServer struct {
	sync.RWMutex
	dbUc     dbproto.DbUsercenterAgentClient
	dbBuycai dbproto.DbBuycaiAgentClient
	agents   map[apiproto.LotteryId]*BuycaiAgent
}

func NewServer() *BuycaiServer {
	srv := &BuycaiServer{}
	srv.agents = make(map[apiproto.LotteryId]*BuycaiAgent)
	return srv
}

func (srv *BuycaiServer) connectDb() {
	conn, err := grpc.Dial(*dbUc, grpc.WithInsecure())
	if err != nil {
		log.Panic(err)
	}
	// defer conn.Close()
	srv.dbUc = dbproto.NewDbUsercenterAgentClient(conn)

	conn, err = grpc.Dial(*dbBuycai, grpc.WithInsecure())
	if err != nil {
		log.Panic(err)
	}
	// defer conn.Close()
	srv.dbBuycai = dbproto.NewDbBuycaiAgentClient(conn)
}

func (srv *BuycaiServer) initAgents() {
	if srv.dbUc == nil {
		log.Panic("nil dbUc")
	}
	lotteryIdList := []apiproto.LotteryId{
		apiproto.LotteryId_Ssq,
		apiproto.LotteryId_Dlt,
		apiproto.LotteryId_Fc3d,
		apiproto.LotteryId_Pl5,
		apiproto.LotteryId_Pl3,
		apiproto.LotteryId_Bjpk10,
		apiproto.LotteryId_Gd11x5,
		apiproto.LotteryId_Cqssc,
	}
	for _, id := range lotteryIdList {
		agt := NewAgent(id, srv.dbUc, srv.dbBuycai /* , srv.dbOpencai */)
		agt.Run()
		srv.agents[id] = agt
	}
}

func (srv *BuycaiServer) getAgent(id apiproto.LotteryId) *BuycaiAgent {
	srv.RLock()
	defer srv.RUnlock()
	return srv.agents[id]
}

func (srv *BuycaiServer) ServeHttp() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	log.Println("ServeHTTP on port ", *httpPort)
	mux := http.NewServeMux()
	mux.HandleFunc("/buycai/info/", srv.HandleBuycaiInfo)
	mux.HandleFunc("/buycai/order/commit", auth.Validate(srv.HandleCommitOrder))
	mux.HandleFunc("/buycai/history/vo", auth.Validate(srv.HandleVendorOrderHistory))
	mux.HandleFunc("/buycai/history/uo", auth.Validate(srv.HandleUserOrderHistory))
	mux.HandleFunc("/buycai/uo/detail", auth.Validate(srv.HandleUserOrderDetail))
	mux.HandleFunc("/buycai/vo/detail", auth.Validate(srv.HandleVendorOrderDetail))
	mux.HandleFunc("/buycai/stopchase", auth.Validate(srv.HandleStopChase))

	httpServer := http.Server{
		Addr:         fmt.Sprintf(":%d", *httpPort),
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	if err := httpServer.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}
