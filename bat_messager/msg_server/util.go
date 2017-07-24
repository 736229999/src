package main

import (
	"fmt"
	"net"
	"sync"
	"time"
	"github.com/oikomi/FishChatServer/log"
	"github.com/oikomi/FishChatServer/libnet"
	"github.com/oikomi/FishChatServer/protocol"
)

type MonitorBeat struct {
	name       string
	session    *libnet.Session
	mu         sync.Mutex
	timeout    time.Duration
	expire     time.Duration
	fails      uint64
	threshold  uint64
}

func NewMonitorBeat(name string, timeout time.Duration, expire time.Duration, limit uint64) *MonitorBeat {
	return &MonitorBeat {
		name      : name,
		timeout   : timeout,
		expire    : expire,
		threshold : limit,
	}
}
func(self *MonitorBeat) ResetFailures() {
	self.mu.Lock()
	defer self.mu.Unlock()
	self.fails = 0
}
func (self *MonitorBeat) ChangeThreshold(thres uint64) {
	self.mu.Lock()
	defer self.mu.Unlock()
	self.threshold = thres
}
func (self *MonitorBeat)Beat(c *libnet.Channel,data *protocol.CmdMonitor)  {
	timer := time.NewTicker(10* time.Second)
	select {
	case <-timer.C:
		go func() {
			log.Info("正在发送心跳包")
			_, err := c.Broadcast(libnet.Json(data))
			if err != nil {
				log.Error(err.Error())
				//return err
			}
		}()
	}
}

func (self *MonitorBeat) Receive() {
	timeout := time.After(self.timeout)
	for {
		select {
		case <-timeout:
			self.fails = self.fails + 1
			if self.fails > self.threshold {
				return
			}
		}
	}
}