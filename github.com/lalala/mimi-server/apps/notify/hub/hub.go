package hub

import (
	"errors"
	"github.com/gorilla/websocket"
	"log"
	"sync"
	"time"
)

var (
	ErrUserNotExist = errors.New("user does not exist")
)

type hub struct {
	// 保存的所有连接
	connects map[int64]*Subscriber
	// 用户连接
	subscribe chan *Subscriber
	// 用户取消连接
	unsubscribe chan int64
	// 向用户推送消息
	push chan *Notify
	// 推送返回的信息
	pushback chan *ContentBack
}

var instance *hub
var once sync.Once

var connsLook sync.RWMutex

const writeWait = 2 * time.Second

//const (
//	user_ws_not_exist error = errors.New("user connect do not exist!")
//)

// hub 采用单例模式
func GetInstance() *hub {
	once.Do(func() {
		instance = &hub{
			connects:    map[int64]*Subscriber{},
			subscribe:   make(chan *Subscriber, 1024),
			unsubscribe: make(chan int64, 1024),
			push:        make(chan *Notify, 10240),
			pushback:    make(chan *ContentBack, 10240),
		}
	})
	return instance
}

// 用户加入hub
func (h *hub) Join(accountId int64, conn *websocket.Conn) {
	sub := &Subscriber{accountId, conn}
	h.subscribe <- sub
}

// 用户离开hub
func (h *hub) Leave(accountId int64) {
	h.unsubscribe <- accountId
}

// 向用户推送消息
func (h *hub) Push(content *Notify) (pb *ContentBack) {
	h.push <- content
	//pb = <-h.pushback
	return
}

// 向用户推送消息
func (h *hub) PushAnnounce(notify *Notify) (onlineUsers []int64) {
	h.push <- notify
	return h.GetOnlineUsers()
}

// 获取所有在线用户id
func (h *hub) GetOnlineUsers() []int64 {
	ousers := make([]int64, 0)
	for key, _ := range h.connects {
		ousers = append(ousers, key)
	}
	return ousers
}

// 检查某个用户是否在线
func (h *hub) CheckUserOnline(id int64) bool {
	_, ok := h.connects[id]
	return ok
}

// 返回推送的结果
func (h *hub) PushBack(contentback *ContentBack) {
	h.pushback <- contentback
}

// 运行hub
func (h *hub) Start() {
	for {
		select {
		case sub := <-h.subscribe:
			if suber, ok := h.connects[sub.AccountId]; ok {
				log.Println("\033[33;1m------>Old User:", suber.AccountId, "\033[0m")
			} else {
				log.Println("\033[33;1m------>New User:", sub.AccountId, "\033[0m")
			}
			connsLook.Lock()
			h.connects[sub.AccountId] = sub
			connsLook.Unlock()
			log.Printf("\033[0msubscribe h.connects %v\n", h.connects)
			log.Printf("The current number of online users is \033[32;1m------>%v\033[0m\n", len(h.connects))

		case push := <-h.push:
			log.Printf("push: %+v\n", *push)
			go h.handleNotify(push)

		case unsub := <-h.unsubscribe:
			sub, ok := h.connects[unsub]
			if ok {
				sub.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				sub.Conn.Close()
				log.Println("\033[31;1m------>Close Connect of", sub.AccountId, "\033[0m")
			}
			delete(h.connects, unsub)

			log.Printf("unsubscribe h.connects %v\n", h.connects)
			log.Printf("The current number of online users is \033[32;1m------>%v\033[0m\n", len(h.connects))

		}
	}
}

// 通知handle
func (h *hub) handleNotify(notify *Notify) {
	switch notify.Type {
	case NOTIFY_ANNOUNCE:
		h.NotifyAnnounce(notify)
	case NOTIFY_REMIND:
		h.NotifyRemind(notify)
	case NOTIFY_MESSAGE:
		//todo 私信消息
	}
}

// 通知到用户
func (h *hub) notifyUser(accountId int64, content string) error {
	connsLook.RLock()
	conn, ok := h.connects[accountId]
	connsLook.RUnlock()
	if !ok {
		return ErrUserNotExist
	}
	conn.Conn.WriteMessage(websocket.TextMessage, []byte(content))
	return nil
}

// 发布公告
func (h *hub) NotifyAnnounce(notify *Notify) {
	for _, sub := range h.connects {
		sub.Conn.SetWriteDeadline(time.Now().Add(writeWait))
		log.Println("Notify Announce", sub.AccountId)
		log.Println("Notify Data", string(notify.Data.([]byte)))
		err := sub.Conn.WriteMessage(websocket.BinaryMessage, notify.Data.([]byte))
		if err != nil {
			log.Println("WriteMessage", err)
			sub.Conn.WriteMessage(websocket.CloseMessage, []byte{})
			sub.Conn.Close()
		}
	}
}

// 推送提醒
func (h *hub) NotifyRemind(notify *Notify) {
	sub, ok := h.connects[notify.ToUser]
	if ok {
		sub.Conn.SetWriteDeadline(time.Now().Add(writeWait))
		sub.Conn.WriteMessage(websocket.BinaryMessage, notify.Data.([]byte))
	}
}

func init() {
	h := GetInstance()
	go h.Start()
}
