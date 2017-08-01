
package main

import (
	"flag"
	"fmt"
	"bat_messager/log"
	"bat_messager/libnet"
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
	fmt.Printf("router version %s Copyright (c) 2014-2015 Harold Miao (miaohong@miaohong.org)  \n", VERSION)
}

var InputConfFile = flag.String("conf_file", "router.json", "input conf file name")   

func main() {
	version()
	fmt.Printf("built on %s\n", BuildTime())
	flag.Parse()
	cfg := NewRouterConfig(*InputConfFile)
	err := cfg.LoadConfig()
	if err != nil {
		log.Error(err.Error())
		return
	}
	
	server, err := libnet.Listen(cfg.TransportProtocols, cfg.Listen)
	if err != nil {
		log.Error(err.Error())
		return
	}
	log.Info("server start: ", server.Listener().Addr().String())
	
	r := NewRouter(cfg)
	//TODO not use go
	go r.subscribeChannels()
	server.Serve(func(session *libnet.Session) {
	
	})
}
