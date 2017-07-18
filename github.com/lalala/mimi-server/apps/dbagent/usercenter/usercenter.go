package usercenter

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/caojunxyz/mimi-server/apps/dbagent/helper"
	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
)

const (
	TABLE_ACCOUNT                   = "account"
	TABLE_PHONE_USER                = "phone_user"
	TABLE_QQ_USER                   = "qq_user"
	TABLE_WEIXIN_USER               = "weixin_user"
	TABLE_BANKCARD                  = "bankcard"
	TABLE_IDCARD                    = "idcard"
	TABLE_USERINFO                  = "userinfo"
	TABLE_CLIENT_DEVICE             = "client_device"
	TABLE_FUND                      = "fund"
	TABLE_VIRTUAL_FUND              = "virtual_fund"
	TABLE_BUYCAI_USER_ORDER         = "buycai_user_order"
	TABLE_BUYCAI_VENDOR_ORDER       = "buycai_vendor_order"
	TABLE_GIFT_PACKAGE              = "gift_package"
	TABLE_TICKET                    = "ticket"
	TABLE_ACCOUNT_HISTORY           = "account_history"
	TABLE_RECHARGE_HISTORY          = "recharge_history"
	TABLE_FUND_HISTORY              = "fund_history"
	TABLE_INVITE_HISTORY            = "invite_history"
	TABLE_KXD_HISTORY               = "kxd_history"
	TABLE_CREDITS_HISTORY           = "credits_history"
	TABLE_EXCHANGE_HISTORY          = "exchange_history"
	TABLE_CDKEY_BATCH               = "cdkey_batch"
	TABLE_PHONE_REGIST_GIFT_HISTORY = "phone_regist_gift_history"
)

type DbUsercenterAgent struct {
	dbConn *sql.DB
}

func NewAgent() *DbUsercenterAgent {
	return &DbUsercenterAgent{}
}

func (agt *DbUsercenterAgent) ConnectDb() {
	db, err := sql.Open("postgres", helper.DataSourceName("mimi-usercenter", "disable"))
	if err != nil {
		log.Panic(err)
	}
	agt.dbConn = db
}

func (agt *DbUsercenterAgent) Run(port int) {
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
	dbproto.RegisterDbUsercenterAgentServer(grpcServer, agt)
	log.Println("DbUsercenterAgent is running on:", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Panic(err)
	}
}
