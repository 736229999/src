
package main

import (
	"flag"
	"encoding/json"
	"bat_messager/log"
	"bat_messager/libnet"
	"bat_messager/protocol"
)

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")
}

type Gateway struct {
	cfg     *GatewayConfig
	server  *libnet.Server
}

func NewGateway(cfg *GatewayConfig) *Gateway {
	return &Gateway {
		cfg    : cfg,
		server : new(libnet.Server),
	}
}

func (self *Gateway)parseProtocol(cmd []byte, session *libnet.Session) error {
	var c protocol.CmdSimple
	err := json.Unmarshal(cmd, &c)
	if err != nil {
		log.Error("error:", err)
		return err
	}
	
	pp := NewProtoProc(self)

	switch c.GetCmdName() {
		case protocol.REQ_MSG_SERVER_CMD:
			err = pp.procReqMsgServer(&c, session)
			if err != nil {
				log.Error("error:", err)
				return err
			}
		}

	return err
}

