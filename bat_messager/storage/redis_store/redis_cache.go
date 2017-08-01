
package redis_store

import (
	"time"
	"sync"
	"errors"
	"github.com/garyburd/redigo/redis"
)

var (
	ErrNoKeyPrefix = errors.New("cannot get session keys without a key prefix")
)

type RedisStoreOptions struct {
	Network              string
	Address              string
	ConnectTimeout       time.Duration
	ReadTimeout          time.Duration
	WriteTimeout         time.Duration
	Database             int           // Redis database to use for session keys
	KeyPrefix            string        // If set, keys will be KeyPrefix:SessionID (semicolon added)
	BrowserSessServerTTL time.Duration // Defaults to 2 days
}

type RedisStore struct {
	opts        *RedisStoreOptions
	conn        redis.Conn
	rwMutex     sync.Mutex
}

// Create a redis session store with the specified options.
func NewRedisStore(opts *RedisStoreOptions) *RedisStore {
	var err error
	rs := &RedisStore{
		opts : opts, 
		conn : nil,
		}
	rs.conn, err = redis.DialTimeout(opts.Network, opts.Address, opts.ConnectTimeout,
		opts.ReadTimeout, opts.WriteTimeout)
	if err != nil {
		panic(err)
	}
	return rs
}
