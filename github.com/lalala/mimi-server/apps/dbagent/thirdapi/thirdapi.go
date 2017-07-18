package thirdapi

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	"github.com/caojunxyz/mimi-server/apps/dbagent/helper"
	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"google.golang.org/grpc"
)

type DbThirdApiAgent struct {
	dbConn *sql.DB
}

func NewAgent() *DbThirdApiAgent {
	return &DbThirdApiAgent{}
}

func (agt *DbThirdApiAgent) ConnectDb() {
	db, err := sql.Open("postgres", helper.DataSourceName(helper.DB_THIRDAPI, "disable"))
	if err != nil {
		log.Panic(err)
	}
	agt.dbConn = db
}

func (agt *DbThirdApiAgent) Run(port int) {
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
	dbproto.RegisterDbThirdApiAgentServer(grpcServer, agt)
	log.Println("DbThirdApiAgent is running on:", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Panic(err)
	}
}
