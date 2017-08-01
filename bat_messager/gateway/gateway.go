
package main

import (
	"fmt"
	"flag"
	"bat_messager/log"
	"bat_messager/libnet"
	"bat_messager/protocol"
)

///*
//#include <stdlib.h>
//#include <stdio.h>
//#include <string.h>
//const char* build_time(void) {
//	static const char* psz_build_time = "["__DATE__ " " __TIME__ "]";
//	return psz_build_time;
//}
//*/
//import "C"

//var (
//	buildTime = C.GoString(C.build_time())
//)

//func BuildTime() string {
//	return buildTime
//}

const VERSION string = "0.24"

func version() {
	fmt.Printf("gateway version %s Copyright (c) 2014-2015 Harold Miao (miaohong@miaohong.org)  \n", VERSION)
}

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")
}

var InputConfFile = flag.String("conf_file", "gateway.json", "input conf file name") 

func handleSession(gw *Gateway, session *libnet.Session) {
	//声明一个临时缓冲区，用来存储被截断的数据
	tmpBuffer := make([]byte, 0)
	//声明一个管道用于接收解包的数据
	readerChannel := make(chan []byte, 16)
	go reader(readerChannel,session,gw)

	buffer := make([]byte, 1024)
	for {
		n, err := session.Conn().Read(buffer)
		if err != nil {
			log.Info(session.Conn().RemoteAddr().String(), " connection error: ", err)
			return
		}

		tmpBuffer,err  = protocol.Unpack(append(tmpBuffer, buffer[:n]...), readerChannel)
		if err != nil {
			session.Conn().Write([]byte("解包失败，封包没对"))
		}
	}
}
func reader(readerChannel chan []byte,session *libnet.Session,gw *Gateway) {
	for {
		select {
		case data := <-readerChannel:
			procLoginJson(data,session,gw)
		}
	}
}
func test()  {
	log.Info("执行了test")
}
func main() {
	version()
	//fmt.Printf("built on %s\n", BuildTime())
	flag.Parse()
	cfg := NewGatewayConfig(*InputConfFile)
	err := cfg.LoadConfig()
	if err != nil {
		log.Error(err.Error())
		return
	}
	
	gw := NewGateway(cfg)

	gw.server, err = libnet.Listen(cfg.TransportProtocols, cfg.Listen)
	if err != nil {
		log.Error(err.Error())
		return
	}
	log.Info("gateway server running at ", gw.server.Listener().Addr().String())

	gw.server.Serve(func(session *libnet.Session) {
		log.Info("client ", session.Conn().RemoteAddr().String(), " | come in")
		
		go handleSession(gw, session)
	})
}
