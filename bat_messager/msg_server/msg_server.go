
package main

import (
	"fmt"
	"time"
	"flag"
	"bat_messager/log"
	"bat_messager/base"
	"bat_messager/libnet"
	"bat_messager/storage/redis_store"
	"bat_messager/protocol"
	"io"
)

/*
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
const char* build_time(void) {
	static const char* psz_build_time = "["__DATE__ " " __TIME__ "]";
	return psz_build_time;
}
*/
import "C"

var (
	buildTime = C.GoString(C.build_time())
)

func BuildTime() string {
	return buildTime
}

const VERSION string = "0.10"

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")
}

func version() {
	fmt.Printf("msg_server version %s Copyright (c) 2014 Harold Miao (miaohong@miaohong.org)  \n", VERSION)
}

var InputConfFile = flag.String("conf_file", "msg_server.json", "input conf file name")   

func handleSession(ms *MsgServer, session *libnet.Session) {
	//session.Process(func(msg *libnet.InBuffer) error {
	//	err := ms.parseProtocol(msg.Data, session)
	//	if err != nil {
	//		log.Error(err.Error())
	//	}
	//
	//	return nil
	//})
	tmpBuffer := make([]byte, 0)
	readerChannel := make(chan []byte, 16)
	go reader(readerChannel,session,ms)

	buffer := make([]byte, 1024)
	for {
		n, err := session.Conn().Read(buffer)
		if err != nil {
			log.Info(session.Conn().RemoteAddr().String(), " connection error: ", err)
			if err == io.EOF {
				session.Conn().Close()
				session.State = false
			}
			return
		}

		tmpBuffer = protocol.Unpack(append(tmpBuffer, buffer[:n]...), readerChannel)
	}
}
func reader(readerChannel chan []byte,session *libnet.Session,ms *MsgServer) {
	for {
		select {
		case data := <-readerChannel:
			log.Info(string(data))
			procMsgServerJson(data,session,ms)
		}
	}
}
func main() {
	version()
	fmt.Printf("built on %s\n", BuildTime())
	flag.Parse()
	cfg := NewMsgServerConfig(*InputConfFile)
	err := cfg.LoadConfig()
	if err != nil {
		log.Error(err.Error())
		return
	}
	
	rs := redis_store.NewRedisStore(&redis_store.RedisStoreOptions {
			Network        : "tcp",
			Address        : cfg.Redis.Port,
			ConnectTimeout : time.Duration(cfg.Redis.ConnectTimeout)*time.Millisecond,
			ReadTimeout    : time.Duration(cfg.Redis.ReadTimeout)*time.Millisecond,
			WriteTimeout   : time.Duration(cfg.Redis.WriteTimeout)*time.Millisecond,
			Database       : 1,
			KeyPrefix      : base.COMM_PREFIX,
		})

	ms := NewMsgServer(cfg, rs)

	ms.server, err = libnet.Listen(cfg.TransportProtocols, cfg.Listen)
	if err != nil {
		panic(err)
	}
	log.Info("msg_server running at  ", ms.server.Listener().Addr().String())
	
	ms.createChannels()

	go ms.scanDeadSession()

	go ms.sendMonitorData()

	ms.server.Serve(func(session *libnet.Session) {
		log.Info("a new client ", session.Conn().RemoteAddr().String(), " | come in")
		go handleSession(ms, session)
	})
}
