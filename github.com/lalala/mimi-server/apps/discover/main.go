package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net"
	"os/exec"
	"strings"

	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"google.golang.org/grpc"
)

// DiscoverServer 发现模块服务
type DiscoverServer struct {
	dbClient   dbproto.DbDiscoveragentClient
	ServerAddr string
}

var httpPort = flag.Int("http", 7013, "http port")
var grpcPort = flag.Int("grpc", 7014, "grpc port")
var db = flag.String("db", "localhost:6011", "db agent")
var isLocal = flag.Bool("local", false, "is local test")

var serverAddr string

func main() {
	flag.Parse()
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	fmt.Println("discover...")
	srv := NewDiscoverServer()
	srv.ConnectDb()
	InitServerAddr()
	// go srv.ServeGRPC()
	go srv.ServeHTTP()
	select {}
}

// ConnectDb 连接到mimi-discover 数据库
func (srv *DiscoverServer) ConnectDb() {
	conn, err := grpc.Dial(*db, grpc.WithInsecure())
	if err != nil {
		log.Panic(err)
	}
	// defer conn.Close()
	srv.dbClient = dbproto.NewDbDiscoveragentClient(conn)
}

// NewDiscoverServer 生成一个新的NewDiscoverServer
func NewDiscoverServer() *DiscoverServer {
	return &DiscoverServer{}
}

func InitServerAddr() {
	if *isLocal {
		conn, err := net.Dial("udp", "www.baidu.com:80")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer conn.Close()
		serverAddr = strings.Split(conn.LocalAddr().String(), ":")[0]
	} else {
		cmd := exec.Command("/bin/sh", "-c", "curl ifconfig.co") //调用Command函数
		var out bytes.Buffer                                     //缓冲字节

		cmd.Stdout = &out //标准输出
		err := cmd.Run()  //运行指令 ，做判断
		if err != nil {
			log.Fatal(err)
		}

		serverAddr = strings.Replace(out.String(), "\n", "", -1)
	}

	log.Println("本机地址:", serverAddr)
}
