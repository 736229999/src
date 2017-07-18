package recharge

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	"github.com/caojunxyz/mimi-server/apps/dbagent/helper"
	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"google.golang.org/grpc"
)

type DbRechargeAgent struct {
	dbConn *sql.DB
}

func NewAgent() *DbRechargeAgent {
	return &DbRechargeAgent{}
}

func (agt *DbRechargeAgent) ConnectDb() {
	db, err := sql.Open("postgres", helper.DataSourceName("mimi-recharge", "disable"))
	if err != nil {
		log.Panic(err)
	}
	agt.dbConn = db
}

func (agt *DbRechargeAgent) Run(port int) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	agt.ConnectDb()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Panic(err)
	}
	grpcServer := grpc.NewServer()
	dbproto.RegisterDbRechargeAgentServer(grpcServer, agt)
	log.Println("DbRechargeAgent is running on:", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Panic(err)
	}
}
