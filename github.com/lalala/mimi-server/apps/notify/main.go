package main

import (
	"flag"
	"fmt"
	"log"
	"sync"

	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"google.golang.org/grpc"
)

var httpPort = flag.Int("http", 7010, "http port")
var grpcPort = flag.Int("grpc", 7012, "grpc port")
var db = flag.String("db", "localhost:6010", "db agent")

func main() {
	flag.Parse()
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	fmt.Println("notify server start...")
	srv := NewNotifyServer()
	srv.ConnectDb()
	go srv.ServeHTTP()
	go srv.ServeGRPC()
	select {}
}

type NotifyServer struct {
	sync.RWMutex
	dbClient dbproto.DbNotifyAgentClient
}

func (srv *NotifyServer) ConnectDb() {
	conn, err := grpc.Dial(*db, grpc.WithInsecure())
	if err != nil {
		log.Panic(err)
	}
	// defer conn.Close()
	srv.dbClient = dbproto.NewDbNotifyAgentClient(conn)
}

func NewNotifyServer() *NotifyServer {
	srv := &NotifyServer{}
	return srv
}
