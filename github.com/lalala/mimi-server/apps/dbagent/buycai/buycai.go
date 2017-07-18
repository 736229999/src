package buycai

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	"github.com/caojunxyz/mimi-server/apps/dbagent/helper"
	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"google.golang.org/grpc"
)

type DbBuycaiAgent struct {
	dbConn *sql.DB
}

func NewAgent() *DbBuycaiAgent {
	return &DbBuycaiAgent{}
}

func (agt *DbBuycaiAgent) ConnectDb() {
	db, err := sql.Open("postgres", helper.DataSourceName(helper.DB_BUYCAI, "disable"))
	if err != nil {
		log.Panic(err)
	}
	agt.dbConn = db
}

func (agt *DbBuycaiAgent) Run(port int) {
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
	dbproto.RegisterDbBuycaiAgentServer(grpcServer, agt)
	log.Println("DbBuycaiAgent is running on:", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Panic(err)
	}
}
