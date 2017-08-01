
package main

import (
	"sync"
	"time"
	"encoding/json"
	"bat_messager/log"
	"bat_messager/base"
	"bat_messager/libnet"
	"bat_messager/protocol"
	"bat_messager/storage/redis_store"
	"bat_messager/storage/mongo_store"
)

type Router struct {
	cfg                 *RouterConfig
	msgServerClientMap  map[string]*libnet.Session
	sessionCache        *redis_store.SessionCache
	mongoStore          *mongo_store.MongoStore
	topicServerMap      map[string]string
	readMutex           sync.Mutex
}   

func NewRouter(cfg *RouterConfig) *Router {
	return &Router {
		cfg                : cfg,
		msgServerClientMap : make(map[string]*libnet.Session),
		sessionCache       : redis_store.NewSessionCache(redis_store.NewRedisStore(&redis_store.RedisStoreOptions {
					Network :   "tcp",
					Address :   cfg.Redis.Addr + cfg.Redis.Port,
					ConnectTimeout : time.Duration(cfg.Redis.ConnectTimeout)*time.Millisecond,
					ReadTimeout : time.Duration(cfg.Redis.ReadTimeout)*time.Millisecond,
					WriteTimeout : time.Duration(cfg.Redis.WriteTimeout)*time.Millisecond,
					Database :  1,
					KeyPrefix : base.COMM_PREFIX,
		})),
		mongoStore         : mongo_store.NewMongoStore(cfg.Mongo.Addr, cfg.Mongo.Port, cfg.Mongo.User, cfg.Mongo.Password),
		topicServerMap     : make(map[string]string),
	}
}

func (self *Router)connectMsgServer(ms string) (*libnet.Session, error) {
	client, err := libnet.Dial("tcp", ms)
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}

	return client, err
}

func (self *Router)handleMsgServerClient(msc *libnet.Session) {
	msc.Process(func(msg *libnet.InBuffer) error {
		log.Info("msg_server", msc.Conn().RemoteAddr().String()," say: ", string(msg.Data))
		var c protocol.CmdInternal
		pp := NewProtoProc(self)
		err := json.Unmarshal(msg.Data, &c)
		if err != nil {
			log.Error("error:", err)
			return err
		}
		switch c.GetCmdName() {
			case protocol.SEND_MESSAGE_P2P_CMD:
				err := pp.procSendMsgP2P(c, msc)
				if err != nil {
					log.Warning(err.Error())
				}
			case protocol.CREATE_TOPIC_CMD:
				err := pp.procCreateTopic(c, msc)
				if err != nil {
					log.Warning(err.Error())
				}
			case protocol.JOIN_TOPIC_CMD:
				err := pp.procJoinTopic(c, msc)
				if err != nil {
					log.Warning(err.Error())
				}
			case protocol.SEND_MESSAGE_TOPIC_CMD:
				err := pp.procSendMsgTopic(c, msc)
				if err != nil {
					log.Warning(err.Error())
				}
				
			}
		return nil
	})
}

func (self *Router)subscribeChannels() error {
	log.Info("route start to subscribeChannels")
	for _, ms := range self.cfg.MsgServerList {
		msgServerClient, err := self.connectMsgServer(ms)
		if err != nil {
			log.Error(err.Error())
			return err
		}
		msg := &base.LoginData{
			ProcType:protocol.SUBSCRIBE_CHANNEL_CMD,
			Time:time.Now().Unix(),
		}
		jsonData,_ := json.Marshal(msg)
		data := protocol.Packet(jsonData)
		_,err = msgServerClient.Conn().Write(data)
		if err != nil {
			log.Error(err.Error())
			return err
		}

		data1 := make([]byte,1024)
		n,err := msgServerClient.Conn().Read(data1)
		if err != nil {
			log.Info(err)
			return err
		}
		log.Info(string(data1))
		j := data1[:n]
		var orm base.InnerRelationMapping
		err = json.Unmarshal(j,&orm)
		if err != nil {
			log.Info(err)
			return err
		}
		log.Info(orm.ServerIpPort)
		//解析出来之后给对应的msg_server发送同样的relationMapping
		//msg_server收到之后给对应的客户端发送消息
		self.msgServerClientMap[ms] = msgServerClient
	}

	for _, msc := range self.msgServerClientMap {
		go self.handleMsgServerClient(msc)
	}
	return nil
}
