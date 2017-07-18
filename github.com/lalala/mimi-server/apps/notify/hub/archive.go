package hub

import (
	"github.com/gorilla/websocket"
)

// 保存的用户的连接结构
type Subscriber struct {
	AccountId int64          `json:"accout_id,omitempty"`
	Conn *websocket.Conn `json:"conn,omitempty"`
}

// 事件的类型
type EventType int

// 消息的类型
type ContentType int

// 返回状态
type ContentBackType int

const (
	// 加入
	EVENT_JOIN = iota
	// 离开
	EVENT_LEAVE
	// 发送的消息
	EVENT_MESSAGE
)

const (
	// 公告消息
	NOTIFY_ANNOUNCE ContentType = 0
	// 提醒消息
	NOTIFY_REMIND ContentType = 1
	// 私信消息
	NOTIFY_MESSAGE ContentType = 2
)

const (
	// 推送失败
	CONTENTBACK_FAILED ContentBackType = 0
	// 推送成功
	CONTENTBACK_SUCCESS ContentBackType = 1
)

type Event struct {
	Type      EventType
	User      string
	Timestamp int64
	Content   Notify
}

// 推送的消息
type Notify struct {
	Sender  int64
	// 消息类型，广播、单播、组播
	Type    ContentType
	ToUser int64
	Data interface{}
}

type ContentBack struct {
	Ret  ContentBackType
	Msg  string
	Data interface{}
}

//func (s *Subscriber)GetLocalHost() string {
//	addrs, err := net.InterfaceAddrs()
//
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//	for _, address := range addrs {
//		// 检查ip地址判断是否回环地址
//		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
//			if ipnet.IP.To4() != nil {
//				fmt.Println("ip is ", ipnet.IP.String())
//				return ipnet.IP.String() + config.RpcPort
//			}
//
//		}
//	}
//	return ""
//}
