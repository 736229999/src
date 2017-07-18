package admin

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	"github.com/caojunxyz/mimi-admin/dbadminagent/helper"
	dbproto "github.com/caojunxyz/mimi-admin/dbadminagent/proto"
	"google.golang.org/grpc"
)

type DbAdminAgent struct {
	dbConn         *sql.DB
	ucDbConn       *sql.DB
	buycaiDbConn   *sql.DB
	optionsDbConn  *sql.DB
	opencaiDbConn  *sql.DB
	footballDbConn *sql.DB
	rechargeDbConn *sql.DB
}

func NewAgent() *DbAdminAgent {
	return &DbAdminAgent{}
}

func (agt *DbAdminAgent) ConnectDb() {
	db, err := sql.Open("postgres", helper.DataSourceName(helper.DB_ADMIN, "disable"))
	if err != nil {
		log.Panic(err)
	}
	agt.dbConn = db
}
func (agt *DbAdminAgent) ConnectUcDb() {
	db, err := sql.Open("postgres", helper.DataSourceName(helper.DB_UC, "disable"))
	if err != nil {
		log.Panic(err)
	}
	agt.ucDbConn = db
}
func (agt *DbAdminAgent) ConnectBuycaiDb() {
	db, err := sql.Open("postgres", helper.DataSourceName(helper.DB_BUYCAI, "disable"))
	if err != nil {
		log.Panic(err)
	}
	agt.buycaiDbConn = db
}

func (agt *DbAdminAgent) ConnectOptionsDb() {
	db, err := sql.Open("postgres", helper.DataSourceName(helper.DB_OPTIONS, "disable"))
	if err != nil {
		log.Panic(err)
	}
	agt.optionsDbConn = db
}

func (agt *DbAdminAgent) ConnectOpencaiDb() {
	db, err := sql.Open("postgres", helper.DataSourceName(helper.DB_OPENCAI, "disable"))
	if err != nil {
		log.Panic(err)
	}
	agt.opencaiDbConn = db
}

func (agt *DbAdminAgent) ConnectFootballDb() {
	db, err := sql.Open("postgres", helper.DataSourceName(helper.DB_FOOTBALL, "disable"))
	if err != nil {
		log.Panic(err)
	}
	agt.footballDbConn = db
}

func (agt *DbAdminAgent) ConnectRechargeDb() {
	db, err := sql.Open("postgres", helper.DataSourceName(helper.DB_RECHARGE, "disable"))
	if err != nil {
		log.Panic(err)
	}
	agt.rechargeDbConn = db
}

func (agt *DbAdminAgent) Run(port int) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	agt.ConnectDb()
	agt.ConnectUcDb()
	agt.ConnectBuycaiDb()
	agt.ConnectOptionsDb()
	agt.ConnectOpencaiDb()
	agt.ConnectFootballDb()
	agt.ConnectRechargeDb()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Panic(err)
	}
	grpcServer := grpc.NewServer()
	dbproto.RegisterDbAdminAgentServer(grpcServer, agt)
	log.Println("DbAdminAgent is running on:", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Panic(err)
	}
}
