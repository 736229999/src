
package main

import (
	"flag"
	"bat_messager/log"
	"bat_messager/libnet"
	"bat_messager/common"
	"bat_messager/protocol"
	"encoding/json"
)

func init() {
	flag.Set("alsologwostderr", "true")
	flag.Set("log_dir", "false")
}

type Data struct {
	Name string 		`json:"name"`
	Age int 			`json:"age"`
	ProcType string 	`json:"procType"`
}
type ProtoProc struct {
	gateway    *Gateway
}

func NewProtoProc(gateway *Gateway) *ProtoProc {
	return &ProtoProc {
		gateway : gateway,
	}
}

func (self *ProtoProc)procReqMsgServer(cmd protocol.Cmd, session *libnet.Session) error {
	//log.Info("procReqMsgServer")
	var err error
	msgServer := common.SelectServer(self.gateway.cfg.MsgServerList, self.gateway.cfg.MsgServerNum)

	resp := protocol.NewCmdSimple(protocol.SELECT_MSG_SERVER_FOR_CLIENT_CMD)
	resp.AddArg(msgServer)
	
	log.Info("Resp | ", resp)
	
	if session != nil {
		err = session.Send(libnet.Json(resp))
		if err != nil {
			log.Error(err.Error())
		}
		session.Close()
		log.Info("client ", session.Conn().RemoteAddr().String(), " | close")
	}
	return nil
}
func procLoginJson(data []byte,session *libnet.Session,gw *Gateway)  {
	log.Info(string(data))
	var login Data
	err := json.Unmarshal(data,&login)
	if err != nil {
		log.Info(err)
		return
	}
	log.Info(login.ProcType)
	switch login.ProcType {
	case "req_msg_server":
		procMsgServer(session,gw)
	}
}

//请求连接聊天服务器
func procMsgServer(session *libnet.Session,gw *Gateway)  {
	msgServer := common.SelectServer(gw.cfg.MsgServerList,gw.cfg.MsgServerNum)
	session.Conn().Write([]byte(msgServer))
	session.Close()
}