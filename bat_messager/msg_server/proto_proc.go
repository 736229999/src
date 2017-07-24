package main
import (
	"flag"
	"strconv"
	"github.com/oikomi/FishChatServer/log"
	"github.com/oikomi/FishChatServer/libnet"
	"github.com/oikomi/FishChatServer/base"
	"github.com/oikomi/FishChatServer/protocol"
	"github.com/oikomi/FishChatServer/common"
	"github.com/oikomi/FishChatServer/storage/redis_store"
	"github.com/oikomi/FishChatServer/storage/mongo_store"
)

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")
}

type ProtoProc struct {
	msgServer    *MsgServer
}

func NewProtoProc(msgServer *MsgServer) *ProtoProc {
	return &ProtoProc {
		msgServer : msgServer,
	}
}

//订阅消息
func (self *ProtoProc)procSubscribeChannel(cmd protocol.Cmd,session *libnet.Session) (error) {
	log.Info("订阅消息")
	channelName := cmd.GetArgs()[0]
	cUUID := cmd.GetArgs()[1]
	log.Info("channelName是:",channelName)
	if self.msgServer.channels[channelName] != nil{
		self.msgServer.channels[channelName].Channel.Join(session, nil)
		self.msgServer.channels[channelName].ClientIDlist = append(self.msgServer.channels[channelName].ClientIDlist, cUUID)
	}else {
		log.Warning(channelName + " is not exist")
	}
	log.Info("所有的channel：",self.msgServer.channels)
	return nil
}
//心跳包，看是否还在线
func (self *ProtoProc)procPing (cmd protocol.Cmd, session *libnet.Session) error {
	cid := session.State.(*base.SessionState).ClientID
	self.msgServer.scanSessionMutex.Lock()
	defer self.msgServer.scanSessionMutex.Unlock()
	self.msgServer.sessions[cid].State.(*base.SessionState).Alive = true
	return nil
}
//离线消息
func (self *ProtoProc)procOfflineMsg(session *libnet.Session,ID string) error {
	exist, err := self.msgServer.offlineMsgCache.IsKeyExist(ID)
	if exist.(int64) == 0 {
		return err
	}else{
		// omrd => offlineMsgRespData
		omrd ,err := common.GetOfflineMsgFromOwnerName(self.msgServer.offlineMsgCache,ID)
		if err != nil {
			log.Error(err.Error())
			return err
		}
		for _,v := range omrd.MsgList{
			resp := protocol.NewCmdSimple(protocol.RESP_MESSAGE_P2P_CMD)
			resp.AddArg(v.Msg)
			resp.AddArg(v.FromID)
			resp.AddArg(v.Uuid)
			if self.msgServer.sessions[ID] != nil {
				self.msgServer.sessions[ID].Send(libnet.Json(resp))
				if err != nil {
					log.Error(err.Error())
					return err
				}
			}
			omrd.ClearMsg()
			self.msgServer.offlineMsgCache.Set(omrd)
		}
	}
	return nil
}
//解析客户端ID
func (self *ProtoProc)procCleintID(cmd protocol.Cmd,session *libnet.Session) error {
	log.Info("解析客户端id")
	ID := cmd.GetArgs()[0]
	sessionCacheData := redis_store.NewSessionCacheData(cmd.GetArgs()[0], session.Conn().RemoteAddr().String(),
		self.msgServer.cfg.LocalIP, strconv.FormatUint(session.Id(), 10))
	log.Info(sessionCacheData)
	args := make([]string, 0)
	args = append(args, cmd.GetArgs()[0])
	CCmd := protocol.NewCmdInternal(protocol.CACHE_SESSION_CMD, args, sessionCacheData)
	log.Info(CCmd)
	if self.msgServer.channels[protocol.SYSCTRL_CLIENT_STATUS] != nil {
		_, err = self.msgServer.channels[protocol.SYSCTRL_CLIENT_STATUS].Channel.Broadcast(libnet.Json(CCmd))
		if err != nil {
			log.Error(err.Error())
			return err
		}
	}

}