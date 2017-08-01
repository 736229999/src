package main

import (
	"bat_messager/libnet"
	"github.com/prometheus/common/log"
	"bat_messager/base"
	"encoding/json"
	"bat_messager/protocol"
	"github.com/pkg/errors"
)

func procMsgServerJson(data []byte,session *libnet.Session,ms *MsgServer)  {
	var loginData base.LoginData
	err := json.Unmarshal(data,&loginData)
	if err != nil {
		log.Info(err)
		return
	}
	pp := NewProtoProc(ms)
	switch loginData.ProcType {
	case "login_msg_server":
		pp.procLoginMsgServer(&loginData,session,ms)
	case protocol.SUBSCRIBE_CHANNEL_CMD:
		pp.subscribeChannelMsg(loginData.ClientId,session)
	}
}
//登录聊天服务器
func (self *ProtoProc) procLoginMsgServer(loginData *base.LoginData,session *libnet.Session,ms *MsgServer) error {
	log.Info("procClientID")
	self.msgServer.sessions[loginData.ClientId] = session
	log.Info("所有的客户端",self.msgServer.sessions)
	session.State = true
	innerRelationMapping := base.NewInnerRetionMapping(loginData.ClientId,session.Conn().RemoteAddr().String(),session.Conn().LocalAddr().String())

	self.msgServer.innerProtocol[loginData.ClientId] = *innerRelationMapping	//好像可以不要这行
	//往内部服务器里面存一个映射（router，manager,monitor)（考虑了下，还是用tcp socket，效率快些）
	relationData ,err := json.Marshal(innerRelationMapping)
	if err != nil {
		log.Info(err)
		return err
	}
	err = self.innerBroadcast(relationData)
	if err != nil {
		log.Info(err)
		return err
	}
	return nil
}
//服务器内部订阅消息
func (self *ProtoProc)subscribeChannelMsg(clientId string,session *libnet.Session)  {
	self.msgServer.innerSessions[clientId] = *session
}
func (self *ProtoProc)innerBroadcast(relationMapping []byte) error {
	log.Info("所有的内部服务器有",self.msgServer.innerSessions)
	if self.msgServer.innerSessions == nil {
		return errors.New("router这些连接失败")
	}
	for _,v := range self.msgServer.innerSessions{
		//k=>clientId    v=>innerProtocol
		_,err := v.Conn().Write(relationMapping)
		if err != nil {
			log.Info(err)
			return err
		}
	}
	return nil
}

