package options

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
	// TABLE_BANNER banner表名
	TABLE_BANNER = "banner"
	// TABLE_NEWS 新闻表名
	TABLE_NEWS = "news"
	// TABLE_FAQ faq表名
	TABLE_FAQ = "faq"
)

// DbOptionsAgent 平台配置数据库代理
type DbOptionsAgent struct {
	dbConn *sql.DB
}

// NewAgent ...
func NewAgent() *DbOptionsAgent {
	return &DbOptionsAgent{}
}

// ConnectDb ...
func (agt *DbOptionsAgent) ConnectDb() {
	db, err := sql.Open("postgres", helper.DataSourceName(helper.DB_OPTIONS, "disable"))
	if err != nil {
		log.Panic(err)
	}
	agt.dbConn = db
}

// Run DbOptionsAgent
func (agt *DbOptionsAgent) Run(port int) {
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
	dbproto.RegisterDbOptionsAgentServer(grpcServer, agt)
	log.Println("DbOptionsAgent is running on:", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Panic(err)
	}
}
