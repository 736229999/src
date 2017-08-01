
package main

import (
	"flag"
	"bat_messager/log"
	"bat_messager/libnet"
	"bat_messager/protocol"
	"bat_messager/storage/redis_store"
	"bat_messager/storage/mongo_store"
)

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")
}

type ProtoProc struct {
	Manager   *Manager
}

func NewProtoProc(m *Manager) *ProtoProc {
	return &ProtoProc {
		Manager : m,
	}
}

func (self *ProtoProc)procCacheSession(cmd protocol.Cmd, session *libnet.Session) error {
	log.Info("procCacheSession")
	var err error
	log.Info(cmd.GetAnyData())
	err = self.Manager.sessionCache.Set(cmd.GetAnyData().(*redis_store.SessionCacheData))
	if err != nil {
		return err
		log.Error("error:", err)
	}
	log.Info("set sesion id success")
	
	return nil
}

func (self *ProtoProc)procCacheTopic(cmd protocol.Cmd, session *libnet.Session) error {
	log.Info("procCacheTopic")
	var err error
	log.Info(cmd.GetAnyData())
	err = self.Manager.topicCache.Set(cmd.GetAnyData().(*redis_store.TopicCacheData))
	if err != nil {
		return err
		log.Error("error:", err)
	}
	log.Info("set sesion id success")
	
	return nil
}


func (self *ProtoProc)procStoreSession(data interface{}, session *libnet.Session) error {
	log.Info("procStoreSession")
	var err error
	log.Info(data)
	err = self.Manager.mongoStore.Update(mongo_store.DATA_BASE_NAME, mongo_store.CLIENT_INFO_COLLECTION, data)
	if err != nil {
		return err
		log.Error("error:", err)
	}
	
	return nil
}

func (self *ProtoProc)procStoreTopic(data interface{}, session *libnet.Session) error {
	log.Info("procStoreTopic")
	var err error
	log.Info(data)
	
	err = self.Manager.mongoStore.Update(mongo_store.DATA_BASE_NAME, mongo_store.TOPIC_INFO_COLLECTION, data)
	if err != nil {
		return err
		log.Error("error:", err)
	}
	
	return nil
}
