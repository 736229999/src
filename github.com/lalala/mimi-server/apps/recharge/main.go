package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	ucproto "github.com/caojunxyz/mimi-server/apps/usercenter/proto"
	"github.com/caojunxyz/mimi-server/auth"
	"google.golang.org/grpc"
)

type RechargeServer struct {
	ucClient ucproto.UsercenterClient
	dbClient dbproto.DbRechargeAgentClient
}

func NewServer() *RechargeServer {
	return &RechargeServer{}
}

func (srv *RechargeServer) ConnectDb() {
	conn, err := grpc.Dial(*dbAgentAddress, grpc.WithInsecure())
	if err != nil {
		log.Panic(err)
	}
	srv.dbClient = dbproto.NewDbRechargeAgentClient(conn)
}

func (srv *RechargeServer) ConnectUc() {
	conn, err := grpc.Dial(*ucAgentAddress, grpc.WithInsecure())
	if err != nil {
		log.Panic(err)
	}
	srv.ucClient = ucproto.NewUsercenterClient(conn)
}

//用户中心代理.
func (srv *RechargeServer) UcAgentClient() ucproto.UsercenterClient {
	return srv.ucClient
}

//数据库代理.
func (srv *RechargeServer) DbAgentClient() dbproto.DbRechargeAgentClient {
	return srv.dbClient
}

//定义几个充值订单的状态.
const (
	WAIT_PAY    = 0 //待支付.
	PAY_SUCCESS = 1 //已支付.
	PAY_CLOSE   = 2 //支付取消.

	WECHAT_PAY = 0 //微信支付.
	ALIPAY_PAY = 1 //支付宝支付.
)

var httpPort = flag.Int("recharge", 7011, "http port")
var gw = flag.String("gw", "http://cp.kxkr.com:8088", "gateway address")
var dbAgentAddress = flag.String("dnAgent", "127.0.0.1:7000", "dbAgent address")
var ucAgentAddress = flag.String("ucAgent", "127.0.0.1:7004", "ucAgent address")
var debug = flag.Bool("debug", true, "debug")

func main() {

	flag.Parse()
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

	srv := NewServer()
	srv.ConnectDb()
	srv.ConnectUc()

	log.Println("http port:", *httpPort)
	log.Println("gw:", *gw)
	log.Println("dbAgentAddress:", *dbAgentAddress)
	log.Println("ucAgentAddress:", *ucAgentAddress)
	log.Println("debug:", *debug)

	http.HandleFunc("/recharge/wechat", auth.Validate(srv.HandleWechatRechargeOrderCommit))
	http.HandleFunc("/recharge/wechat/notify", srv.HandleWechatRechargeNotify)
	http.HandleFunc("/recharge/alipay", auth.Validate(srv.HandleAlipayRechargeCommitOrder))
	http.HandleFunc("/recharge/alipay/notify", srv.HandleAlipayRechargeNotify)

	addr := fmt.Sprintf(":%d", *httpPort)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Panicf("监听端口失败", err)
	}
}
