package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/chr4/pwgen"
	"google.golang.org/grpc"

	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/proto"
)

var httpPort = flag.Int("http", 7003, "http port")
var grpcPort = flag.Int("grpc", 7004, "grpc port")
var db = flag.String("db", "localhost:6008", "db agent")
var thirdapi = flag.String("thirdapi", "localhost:6600", "thirdapi server")

func main() {
	flag.Parse()
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	fmt.Println("usercenter...")
	srv := NewUcServer()
	srv.ConnectThirdapi()
	srv.ConnectDb()
	go srv.ServeGRPC()
	go srv.ServeHTTP()
	select {}
}

type PayPwdVerifySession struct {
	AccountId  int64
	Ip         string
	DeviceId   string
	Operate    string
	VerifyTime int64
}

type SetPayPwdToken struct {
	Token   string
	GenTime time.Time
}

type UcServer struct {
	sync.RWMutex
	randGenerator        *rand.Rand
	dbClient             dbproto.DbUsercenterAgentClient
	thirdapiClient       proto.ThirdApiClient
	payPwdVerifySessions map[int64]PayPwdVerifySession
	setPayPwdTokens      map[int64]SetPayPwdToken
}

func (srv *UcServer) RecordPayPwdVerify(accountId int64, ip string, deviceId string, operate string, verifyTime int64) {
	srv.Lock()
	defer srv.Unlock()
	srv.payPwdVerifySessions[accountId] = PayPwdVerifySession{
		AccountId: accountId, Ip: ip, DeviceId: deviceId, Operate: operate, VerifyTime: verifyTime,
	}
}

func (srv *UcServer) CheckPayPwdVerify(accountId int64, ip string, deviceId string, operate string, verifyTime int64) bool {
	srv.RLock()
	defer srv.RUnlock()
	sess, ok := srv.payPwdVerifySessions[accountId]
	if ok {
		if sess.Ip == ip && sess.DeviceId == deviceId && operate == sess.Operate && (time.Now().Unix()-verifyTime < 60) {
			return true
		}
	}
	return false
}

func (srv *UcServer) GenSetPayPwdToken(accountId int64) string {
	token := pwgen.AlphaNumSymbols(30)
	srv.Lock()
	defer srv.Unlock()
	srv.setPayPwdTokens[accountId] = SetPayPwdToken{Token: token, GenTime: time.Now()}
	return token
}

func (srv *UcServer) VerifySetPayPwdToken(accountId int64, token string) bool {
	srv.RLock()
	defer srv.RLock()
	val, ok := srv.setPayPwdTokens[accountId]
	if ok {
		return (token == val.Token && time.Now().Sub(val.GenTime) < time.Minute)
	}
	return false
}

func (srv *UcServer) ConnectDb() {
	conn, err := grpc.Dial(*db, grpc.WithInsecure())
	if err != nil {
		log.Panic(err)
	}
	// defer conn.Close()
	srv.dbClient = dbproto.NewDbUsercenterAgentClient(conn)
}

func (srv *UcServer) ConnectThirdapi() {
	conn, err := grpc.Dial(*thirdapi, grpc.WithInsecure())
	if err != nil {
		log.Panic(err)
	}
	// defer conn.Close()
	srv.thirdapiClient = proto.NewThirdApiClient(conn)
}

func NewUcServer() *UcServer {
	srv := &UcServer{
		randGenerator:        rand.New(rand.NewSource(time.Now().UnixNano())),
		payPwdVerifySessions: make(map[int64]PayPwdVerifySession),
		setPayPwdTokens:      make(map[int64]SetPayPwdToken),
	}
	return srv
}
