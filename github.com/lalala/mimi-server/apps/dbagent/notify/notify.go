package notify

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	"github.com/caojunxyz/mimi-server/apps/dbagent/helper"
	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"google.golang.org/grpc"
)

const (
	TABLE_NOTIFY              = "notify"
	TABLE_USER_NOTIFY         = "user_notify"
	TABLE_SUBSCRIPTION        = "subcription"
	TABLE_SUBSCRIPTION_CONFIG = "subcription_config"
)

type DbNotifyAgent struct {
	dbConn *sql.DB
}

func NewAgent() *DbNotifyAgent {
	return &DbNotifyAgent{}
}

func (agt *DbNotifyAgent) ConnectDb() {
	db, err := sql.Open("postgres", helper.DataSourceName("mimi-notify", "disable"))
	if err != nil {
		log.Panic(err)
	}
	agt.dbConn = db
}

func (agt *DbNotifyAgent) Run(port int) {
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
	dbproto.RegisterDbNotifyAgentServer(grpcServer, agt)
	log.Println("DbNotifyAgent is running on:", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Panic(err)
	}
}
