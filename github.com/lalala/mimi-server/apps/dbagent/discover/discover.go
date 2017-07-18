package discover

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	"github.com/caojunxyz/mimi-server/apps/dbagent/helper"
	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"google.golang.org/grpc"
)

// TABLE_NEWS 新闻表名
const TABLE_NEWS string = "news"

// DbDiscoverAgent 发现模块数据库代理
type DbDiscoverAgent struct {
	dbConn *sql.DB
}

// NewAgent create a new DbDiscoverAgent
func NewAgent() *DbDiscoverAgent {
	return &DbDiscoverAgent{}
}

// ConnectDb 连接到discover
func (agt *DbDiscoverAgent) ConnectDb() {
	db, err := sql.Open("postgres", helper.DataSourceName("mimi-discover", "disable"))
	if err != nil {
		log.Panic(err)
	}
	agt.dbConn = db
}

// Run DbDiscoverAgent
func (agt *DbDiscoverAgent) Run(port int) {
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
	dbproto.RegisterDbDiscoveragentServer(grpcServer, agt)
	log.Println("DbDiscoverAgent is running on:", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Panic(err)
	}
}
