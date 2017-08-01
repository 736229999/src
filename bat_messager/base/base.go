package base

import (
	"bat_messager/libnet"
	"bat_messager/protocol"
)

type ChannelMap map[string]*ChannelState
type SessionMap map[string]*libnet.Session

type AckMap map[string]map[string]bool

const COMM_PREFIX = "IM"

var ChannleList []string

func init() {
	ChannleList = []string{protocol.SYSCTRL_CLIENT_STATUS, protocol.SYSCTRL_TOPIC_STATUS, protocol.SYSCTRL_TOPIC_SYNC, 
		protocol.SYSCTRL_SEND, protocol.SYSCTRL_MONITOR, protocol.STORE_CLIENT_INFO, protocol.STORE_TOPIC_INFO}
}

type ChannelState struct {
	ChannelName   string
	Channel       *libnet.Channel
	ClientIDlist  []string
}

func NewChannelState(channelName string, channel *libnet.Channel) *ChannelState {
	return &ChannelState {
		ChannelName  : channelName,
		Channel      : channel,
		ClientIDlist : make([]string, 0),
	}
}

type SessionState struct {
	ClientID string
	Alive    bool
}

func NewSessionState(alive bool, cid string) *SessionState {
	return &SessionState {
		ClientID : cid,
		Alive    : alive,
	}
}

type Config interface {
	LoadConfig(configfile string) (*Config, error)
}
