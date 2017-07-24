package main

import (
	"fmt"
	"flag"
	"bat_messager/protocol"
	"bat_messager/libnet"
	"bat_messager/common"
	"github.com/oikomi/FishChatServer/log"
)
func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")
}

type Protoproc struct {
	gateway   *Gateway
}
func NewProtoProc(gateway *Gateway) *Protoproc {
	return &Protoproc {
		gateway : gateway,
	}
}

func (self *Protoproc)procReqMsgServer(cmd protocol.Cmd,session *libnet.Session) error  {
	fmt.Println("处理转发请求的msg_server")
	msgServer := common.SelectServer(self.gateway.cfg.MsgServerList,self.gateway.cfg.MsgServerNum)

	resp := protocol.NewCmdSimple(protocol.SELECT_MSG_SERVER_FOR_CLIENT_CMD)
	resp.AddArg(msgServer)
	log.Info("Resp | ", resp)

	if session != nil {
		err := session.Send(libnet.Json(resp))
		if err != nil {
			log.Error(err.Error())
		}
		session.Close()
		log.Info("客户端关闭了连接")
	}
	return nil
}