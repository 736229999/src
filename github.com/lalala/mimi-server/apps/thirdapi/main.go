package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/apps/thirdapi/infoverify"
	"github.com/caojunxyz/mimi-server/apps/thirdapi/opencai"
	"github.com/caojunxyz/mimi-server/apps/thirdapi/sms"
	"github.com/caojunxyz/mimi-server/proto"
)

var dbagent = flag.String("dbagent", "localhost:6013", "thirdapi dbagent address")
var port = flag.Int("port", 6600, "grpc port")

func main() {
	flag.Parse()
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	fmt.Println("thirdapi...")
	fmt.Println("db:", *dbagent)
	fmt.Println("port:", *port)
	NewServer(*port).Run()
}

type ThirdApiServer struct {
	*sms.SmsServer
	*infoverify.VerifyServer
	*opencai.OpencaiServer
	dbc      dbproto.DbThirdApiAgentClient
	grpcPort int
}

func NewServer(port int) *ThirdApiServer {
	srv := &ThirdApiServer{grpcPort: port}
	srv.connectDb(*dbagent)
	srv.SmsServer = sms.NewServer(srv.dbc)
	srv.VerifyServer = infoverify.NewServer(srv.dbc)
	srv.OpencaiServer = opencai.NewServer(srv.dbc)
	return srv
}

func (srv *ThirdApiServer) connectDb(addr string) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Panic(err)
	}
	srv.dbc = dbproto.NewDbThirdApiAgentClient(conn)
}

func (srv *ThirdApiServer) Run() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", srv.grpcPort))
	if err != nil {
		log.Panic(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterThirdApiServer(grpcServer, srv)
	if err := grpcServer.Serve(lis); err != nil {
		log.Panic(err)
	}
}
