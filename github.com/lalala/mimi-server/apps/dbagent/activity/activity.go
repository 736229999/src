package activity

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	"github.com/caojunxyz/mimi-server/apps/dbagent/helper"
	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"google.golang.org/grpc"
)

type DbActivityAgent struct {
	dbConn *sql.DB
}

func NewAgent() *DbActivityAgent {
	return &DbActivityAgent{}
}

func (agt *DbActivityAgent) ConnectDb() {
	db, err := sql.Open("postgres", helper.DataSourceName(helper.DB_UC, "disable"))
	if err != nil {
		log.Panic(err)
	}
	agt.dbConn = db
}

func (agt *DbActivityAgent) Run(port int) {
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
	dbproto.RegisterDbActivityAgentServer(grpcServer, agt)
	log.Println("DbActivityAgent is running on:", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Panic(err)
	}
}