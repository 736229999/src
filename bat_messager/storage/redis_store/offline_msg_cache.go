
package redis_store

import (
	"sync"
	"time"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
)

type OfflineMsgCache struct {
	RS       *RedisStore
	rwMutex  sync.Mutex
}

func NewOfflineMsgCache(RS *RedisStore) *OfflineMsgCache {
	return &OfflineMsgCache {
		RS    : RS,
	}
}

type OfflineMsgData struct {
	Msg        string
	FromID     string
	Uuid       string
}

func NewOfflineMsgData(msg string, fromID string, uuid string) *OfflineMsgData {
	return &OfflineMsgData {
		Msg : msg,
		FromID : fromID,
		Uuid   : uuid,
	}
}

type OfflineMsgCacheData struct {
	OwnerName     string
	MsgList       []*OfflineMsgData
	MaxAge        time.Duration
}

func (self *OfflineMsgCacheData) AddMsg(d *OfflineMsgData) {
	self.MsgList = append(self.MsgList, d)
}

func (self *OfflineMsgCacheData) ClearMsg() {
	self.MsgList = self.MsgList[:0]
}

func NewOfflineMsgCacheData(ownerName string) *OfflineMsgCacheData {
	return &OfflineMsgCacheData {
		OwnerName : ownerName,
	}
}

// Get the session from the store.
func (self *OfflineMsgCache) Get(k string) (*OfflineMsgCacheData, error) {
	self.rwMutex.Lock()
	defer self.rwMutex.Unlock()
	key := k + OFFLINE_MSG_UNIQ_PREFIX
	if self.RS.opts.KeyPrefix != "" {
		key = self.RS.opts.KeyPrefix + ":" + k + OFFLINE_MSG_UNIQ_PREFIX
	}
	b, err := redis.Bytes(self.RS.conn.Do("GET", key))
	if err != nil {
		return nil, err
	}
	var sess OfflineMsgCacheData
	err = json.Unmarshal(b, &sess)
	if err != nil {
		return nil, err
	}
	return &sess, nil
}

// Save the session into the store.
func (self *OfflineMsgCache) Set(sess *OfflineMsgCacheData) error {
	self.rwMutex.Lock()
	defer self.rwMutex.Unlock()
	b, err := json.Marshal(sess)
	if err != nil {
		return err
	}
	key := sess.OwnerName + OFFLINE_MSG_UNIQ_PREFIX
	if self.RS.opts.KeyPrefix != "" {
		key = self.RS.opts.KeyPrefix + ":" + sess.OwnerName + OFFLINE_MSG_UNIQ_PREFIX
	}
	ttl := sess.MaxAge
	if ttl == 0 {
		// Browser session, set to specified TTL
		ttl = self.RS.opts.BrowserSessServerTTL
		if ttl == 0 {
			ttl = 2 * 24 * time.Hour // Default to 2 days
		}
	}
	_, err = self.RS.conn.Do("SETEX", key, int(ttl.Seconds()), b)
	if err != nil {
		return err
	}
	return nil
}

// Delete the session from the store.
func (self *OfflineMsgCache) Delete(id string) error {
	self.rwMutex.Lock()
	defer self.rwMutex.Unlock()
	key := id + OFFLINE_MSG_UNIQ_PREFIX
	if self.RS.opts.KeyPrefix != "" {
		key = self.RS.opts.KeyPrefix + ":" + id + OFFLINE_MSG_UNIQ_PREFIX
	}
	_, err := self.RS.conn.Do("DEL", key)
	if err != nil {
		return err
	}
	return nil
}

// Clear all sessions from the store. Requires the use of a key
// prefix in the store options, otherwise the method refuses to delete all keys.
func (self *OfflineMsgCache) Clear() error {
	self.rwMutex.Lock()
	defer self.rwMutex.Unlock()
	vals, err := self.getSessionKeys()
	if err != nil {
		return err
	}
	if len(vals) > 0 {
		self.RS.conn.Send("MULTI")
		for _, v := range vals {
			self.RS.conn.Send("DEL", v)
		}
		_, err = self.RS.conn.Do("EXEC")
		if err != nil {
			return err
		}
	}
	return nil
}

// Get the number of session keys in the store. Requires the use of a
// key prefix in the store options, otherwise returns -1 (cannot tell
// session keys from other keys).
func (self *OfflineMsgCache) Len() int {
	self.rwMutex.Lock()
	defer self.rwMutex.Unlock()
	vals, err := self.getSessionKeys()
	if err != nil {
		return -1
	}
	return len(vals)
}

func (self *OfflineMsgCache) getSessionKeys() ([]interface{}, error) {
	self.rwMutex.Lock()
	defer self.rwMutex.Unlock()
	if self.RS.opts.KeyPrefix != "" {
		return redis.Values(self.RS.conn.Do("KEYS", self.RS.opts.KeyPrefix+":*"))
	}
	return nil, ErrNoKeyPrefix
}

func (self *OfflineMsgCache) IsKeyExist(k string) (interface{}, error) {
	self.rwMutex.Lock()
	defer self.rwMutex.Unlock()
	
	key := k + OFFLINE_MSG_UNIQ_PREFIX
	if self.RS.opts.KeyPrefix != "" {
		key = self.RS.opts.KeyPrefix + ":" + k + OFFLINE_MSG_UNIQ_PREFIX
	}
	
	v, err := self.RS.conn.Do("EXISTS", key)
	if err != nil {
		return v, err
	}

	return v, err
}