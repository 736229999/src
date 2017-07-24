package main

import (
	"time"
	"flag"
	"sync"
	"encoding/json"
	"github.com/oikomi/FishChatServer/log"
	"github.com/oikomi/FishChatServer/libnet"
	"github.com/oikomi/FishChatServer/base"
	"github.com/oikomi/FishChatServer/protocol"
	"github.com/oikomi/FishChatServer/storage/redis_store"
	"github.com/oikomi/FishChatServer/storage/mongo_store"
)

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")
}
type MsgServer struct {
	cfg               *MsgServerConfig
	sessions          base.SessionMap
	channels          base.ChannelMap
	topics            protocol.TopicMap
	server            *libnet.Server
	sessionCache      *redis_store.SessionCache
	topicCache        *redis_store.TopicCache
	offlineMsgCache   *redis_store.OfflineMsgCache
	p2pAckStatus      base.AckMap
	scanSessionMutex  sync.Mutex
	p2pAckMutex       sync.Mutex
	//还差个数据库存储
}

func NewMsgServer(cfg *MsgServerConfig, rs *redis_store.RedisStore) *MsgServer {
	return &MsgServer {
		cfg                : cfg,
		sessions           : make(base.SessionMap),
		channels           : make(base.ChannelMap),
		topics             : make(protocol.TopicMap),
		server             : new(libnet.Server),
		sessionCache       : redis_store.NewSessionCache(rs),
		topicCache         : redis_store.NewTopicCache(rs),
		offlineMsgCache    : redis_store.NewOfflineMsgCache(rs),
		//mongoStore         : mongo_store.NewMongoStore(cfg.Mongo.Addr, cfg.Mongo.Port, cfg.Mongo.User, cfg.Mongo.Password),
		p2pAckStatus       : make(base.AckMap),
	}
}

func (self *MsgServer)createChannels() {
	log.Info("createChannels")
	for _, c := range base.ChannleList {
		channel := libnet.NewChannel(self.server.Protocol())
		self.channels[c] = base.NewChannelState(c, channel)
	}
}

func (self *MsgServer)sendMonitorData() error {
	log.Info("sendMonitorData")
	resp := protocol.NewCmdMonitor()

	// resp.SessionNum = (uint64)(len(self.sessions))

	// log.Info(resp)

	mb := NewMonitorBeat("monitor", self.cfg.MonitorBeatTime, 40, 10)

	if self.channels[protocol.SYSCTRL_MONITOR] != nil {
		for{
			resp.SessionNum = (uint64)(len(self.sessions))

			//log.Info(resp)
			mb.Beat(self.channels[protocol.SYSCTRL_MONITOR].Channel, resp)
		}
		// _, err := self.channels[protocol.SYSCTRL_MONITOR].Channel.Broadcast(libnet.Json(resp))
		// if err != nil {
		// 	glog.Error(err.Error())
		// 	return err
		// }
	}

	return nil
}

func (self *MsgServer) scanDeadSession()  {
	log.Info("扫描过期会话")
	timer := time.NewTicker(self.cfg.ScanDeadSessionTimeout * time.Second)
	ttl := time.After(self.cfg.Expire * time.Second)
	for  {
		select {
		case <-timer.C:
			log.Info("已过期")
			go func() {
				for id, s := range self.sessions {
					self.scanSessionMutex.Lock()
					//defer self.scanSessionMutex.Unlock()
					if (s.State).(*base.SessionState).Alive == false {
						log.Info("delete" + id)
						delete(self.sessions, id)
					} else {
						s.State.(*base.SessionState).Alive = false
					}
					self.scanSessionMutex.Unlock()
				}
			}()
		case ttl:
			break
		}
	}
}

func (self *MsgServer)parseProtocol(cmd []byte,session *libnet.Session) error{
	//解析命令行传过来的值
	var c protocol.CmdSimple
	err := json.Unmarshal(cmd, &c)
	if err != nil {
		log.Error("error:", err)
		return err
	}
	pp := NewProtoProc(self)
	switch c.GetCmdName() {
	case protocol.SEND_PING_CMD:
		
	}
}