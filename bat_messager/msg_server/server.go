
package main

import (
	"time"
	"flag"
	"sync"
	"encoding/json"
	"bat_messager/log"
	"bat_messager/libnet"
	"bat_messager/base"
	"bat_messager/protocol"
	"bat_messager/storage/redis_store"
	"bat_messager/storage/mongo_store"
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
	mongoStore        *mongo_store.MongoStore
	p2pAckStatus      base.AckMap
	scanSessionMutex  sync.Mutex
	p2pAckMutex       sync.Mutex
	innerProtocol	  map[string]base.InnerRelationMapping
	innerSessions	  map[string]libnet.Session
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
		mongoStore         : mongo_store.NewMongoStore(cfg.Mongo.Addr, cfg.Mongo.Port, cfg.Mongo.User, cfg.Mongo.Password),
		p2pAckStatus       : make(base.AckMap),
		innerProtocol	   : make(map[string]base.InnerRelationMapping),
		innerSessions	   : make(map[string]libnet.Session),
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

func (self *MsgServer)scanDeadSession() {
	log.Info("scanDeadSession")
	timer := time.NewTicker(self.cfg.ScanDeadSessionTimeout * time.Second)
	ttl := time.After(self.cfg.Expire * time.Second)
	for {
		select {
		case <-timer.C:
			log.Info("scanDeadSession timeout")
			go func() {
				for id, s := range self.sessions {
					self.scanSessionMutex.Lock()
					//defer self.scanSessionMutex.Unlock()
					if s.State== false {
						log.Info("delete" + id)
						delete(self.sessions, id)
					}
					self.scanSessionMutex.Unlock()
				}
			}()
		case <-ttl:
			break
		}
	}
}

func (self *MsgServer)parseProtocol(cmd []byte, session *libnet.Session) error {
	var c protocol.CmdSimple
	err := json.Unmarshal(cmd, &c)
	if err != nil {
		log.Error("error:", err)
		return err
	}
	
	pp := NewProtoProc(self)

	switch c.GetCmdName() {
		case protocol.SEND_PING_CMD:
			err = pp.procPing(&c, session)
			if err != nil {
				log.Error("error:", err)
				return err
			}
		case protocol.SUBSCRIBE_CHANNEL_CMD:
			pp.procSubscribeChannel(&c, session)
		case protocol.SEND_CLIENT_ID_CMD:
			err = pp.procClientID(&c, session)
			if err != nil {
				log.Error("error:", err)
				return err
			}
		case protocol.SEND_MESSAGE_P2P_CMD:
			err = pp.procSendMessageP2P(&c, session)
			if err != nil {
				log.Error("error:", err)
				return err
			}
		case protocol.ROUTE_MESSAGE_P2P_CMD:
			err = pp.procRouteMessageP2P(&c, session)
			if err != nil {
				log.Error("error:", err)
				return err
			}
		case protocol.CREATE_TOPIC_CMD:
			err = pp.procCreateTopic(&c, session)
			if err != nil {
				log.Error("error:", err)
				return err
			}
		case protocol.JOIN_TOPIC_CMD:
			err = pp.procJoinTopic(&c, session)
			if err != nil {
				log.Error("error:", err)
				return err
			}
		case protocol.SEND_MESSAGE_TOPIC_CMD:
			err = pp.procSendMessageTopic(&c, session)
			if err != nil {
				log.Error("error:", err)
				return err
			}

		// p2p ack
		case protocol.P2P_ACK_CMD:
			err = pp.procP2pAck(&c, session)
			if err != nil {
				log.Error("error:", err)
				return err
			}
		}

	return err
}
