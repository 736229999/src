package main

import (
//	"bytes"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// index页面
func serveHome(w http.ResponseWriter, r *http.Request) {
	// URL解析
	if r.URL.Path != "/" {
		http.Error(w, "页面没有找到！", 404)
		return
	}
	// 请求方法确认
	if r.Method != "GET" {
		http.Error(w, "请求方法不正确！", 405)
		return
	}
	// 加载index.html页面
	http.ServeFile(w, r, "./static/index.html")
}
func serveRoom(w http.ResponseWriter, r *http.Request) {
	// URL解析
	if r.URL.Path != "/room" {
		http.Error(w, "页面没有找到！", 404)
		return
	}
	// 请求方法确认
	if r.Method != "GET" {
		http.Error(w, "请求方法不正确！", 405)
		return
	}
	// 加载index.html页面
	http.ServeFile(w, r, "./static/room.html")
}

// 解析命令行参数
var addr = flag.String("addr", ":8080", "http service address")

// 入口
func main() {
	// 解析命令行参数 -addr = ：8080
	flag.Parse()
	// 调度器
	M := &Manage{
		Broadcast:  make(chan []byte),
		ToLogin:    make(chan map[*Client]string),
		ToUser:     make(chan map[*Client]string),
		ToRoom:     make(chan map[string]string),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
	}
	// 运行调度器
	go M.Run()
	// 首页展示页面 home.html  注意 该文件必须要有权限
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/room", serveRoom)

	// websocket 主要程序
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		log.Println("一个新连接")
		StrartWebSocket(M, w, r)
	})
	// 开启监听
	log.Printf("服务监听中.......[%v]", *addr)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// WebSocket

// 调度器
type Manage struct {
	// 广播
	Broadcast chan []byte
	// 用户登录
	ToLogin chan map[*Client]string
	// 基于房间号
	ToRoom chan map[string]string
	// 私聊
	ToUser chan map[*Client]string
	// 注册
	Register chan *Client
	// 注销
	Unregister chan *Client
	// 调度器核心 用户表
	Clients map[*Client]bool
}

// 运行调度器
func (self *Manage) Run() {
	log.Println("调度器运行中。。。。。")
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	// 阻塞调用select
	for {
		select {

		// 新用户 进行注册
		case newClient := <-self.Register:
			self.Clients[newClient] = true

		// 用户退出 进行注销
		case unClient := <-self.Unregister:

			if _, ok := self.Clients[unClient]; ok {
				// 注销前 推送一条消息
				delete(self.Clients, unClient)
				close(unClient.send)
			}
		// 有消息到达
		case Message := <-self.Broadcast:
			// 遍历所有客户端 并向他们推送消息
			for client := range self.Clients {
				// 消息分发器
				select {
				case client.send <- Message:
					log.Println("有消息分发成功")
				default:
					close(client.send)
					delete(self.Clients, client)
				}
			}
		// 用户登录消息回复
		case userMessage := <-self.ToLogin:
			for k, v := range userMessage {
				select {
				case k.send <- []byte(v):
					log.Println("登录消息")
				default:
					close(k.send)
				}
			}
		// 向用户推送消息
		case userMessage := <-self.ToUser:
			for k, v := range userMessage {
				// 检查当前客户端是否退出
				if _, ok := self.Clients[k]; ok {
					select {
					case k.send <- []byte(v):
						log.Println("用户消息")
					default:
						close(k.send)
					}
				}
			}
		// 向房间发送消息
		case roomMessage := <-self.ToRoom:
			for roomid, msg := range roomMessage {
				atRoom := getRoomUser(self, roomid)
				for k, _ := range atRoom {
					select {
					case k.send <- []byte(msg):
						log.Println("房间消息")
					default:
						close(k.send)
					}
				}
			}

		}
	}
}

// 获取房间客户端
func getRoomUser(clients *Manage, roomid string) map[*Client]bool {
	atRoom := make(map[*Client]bool)
	cli := clients.Clients
	for k, v := range cli {
		if k.User != nil {
			if k.User.RoomId == roomid {
				atRoom[k] = v
			}
		}
	}
	return atRoom
}
func getRoomUserInfo(clients *Manage, roomid string) []User {
	atRoom := []User{}
	cli := clients.Clients
	for k, _ := range cli {
		if k.User != nil {
			if k.User.RoomId == roomid {
				u := *k.User
				atRoom = append(atRoom, u)
			}
		}
	}
	return atRoom
}

// 客户结构体
type Client struct {
	// 连接对象
	Connet *websocket.Conn
	// 需要发送的消息
	send chan []byte
	// 用户信息
	User *User
}

// 用户
type User struct {
	Type    string `json:"type"`
	Id      string `json:"uid"`
	Name    string `json:"nickname"`
	RoomId  string `json:"roomnum"`
	Avatar  string `json:"avatar"`
	Content string `json:"content"`
}

// 响应
type Responce struct {
	ResponceType string `json:"type"`
	Data         []User `json:"data"`
}

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 180 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

// 读
func (client *Client) Read(manage *Manage) {
	// 内存回收
	defer func() {
		// 注销用户
		manage.Unregister <- client
		//		if client.User != nil {
		//			msg := bytes.TrimSpace(bytes.Replace([]byte(client.User.Name+"退出了聊天室"), newline, space, -1))
		//			manage.Broadcast <- msg
		//		} else {
		//			msg := bytes.TrimSpace(bytes.Replace([]byte("游客退出了聊天室"), newline, space, -1))
		//			manage.Broadcast <- msg
		//		}
		// 关闭连接
		client.Connet.Close()
	}()
	// 超时设置
	client.Connet.SetReadLimit(maxMessageSize)
	client.Connet.SetReadDeadline(time.Now().Add(pongWait))
	// 心跳设置
	client.Connet.SetPongHandler(func(string) error { client.Connet.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	// 收到消息 分发消息
	for {
		_, message, err := client.Connet.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}
		// 解析
		newUser := &User{}
		json.Unmarshal([]byte(message), &newUser)

		switch newUser.Type {
		// 登录
		case "login":
			if client.User == nil {
				client.User = newUser
				m := make(map[*Client]string)
				m[client] = string(message)
				manage.ToLogin <- m
				//				msg := bytes.TrimSpace(bytes.Replace([]byte(client.User.Name+"进入聊天室"), newline, space, -1))
				//				manage.Broadcast <- msg
			} else {
				client.User = newUser
				m := make(map[*Client]string)
				m[client] = newUser.Name + "已经登录聊天室,修改信息成功"
				manage.ToLogin <- m
			}
		// 聊天室消息分发
		case "PushToRoom":
			m := make(map[string]string)
			m[newUser.RoomId] = newUser.Content
			manage.ToRoom <- m
		// 获取当前房间人数
		case "getRoomUser":
			//			exit := make(chan *Client, 1000)
			go func() {
				for {
					// 检查当前客户端是否有退出
					if _, ok := manage.Clients[client]; ok {
						// 每隔2秒主动推送
						time.Sleep(2 * time.Second)
						data := getRoomUserInfo(manage, newUser.RoomId)

						rep := Responce{}
						rep.ResponceType = "AtRoom"
						rep.Data = data
						m := make(map[*Client]string)
						d, _ := json.Marshal(rep)
						m[client] = string(d)
						manage.ToUser <- m
					} else {
						log.Println("当前客户端已经退出，停止推送")
						break
					}
				}
			}()
		default:
			//			message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
			//			// 广播消息
			//			manage.Broadcast <- message
		}
	}
}

// 写
func (client *Client) Write(manage *Manage) {

	// 定时器 心跳功能
	ticker := time.NewTicker(pingPeriod)

	// 内存回收
	defer func() {
		// 关闭定时器
		ticker.Stop()
		// 关闭连接
		client.Connet.Close()
	}()

	// 消息写入
	for {
		select {
		// 需要发送的消息
		case message, ok := <-client.send:
			// 设置超时
			client.Connet.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				client.Connet.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			// 下一条消息写入入口
			w, err := client.Connet.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// 未理解
			n := len(client.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-client.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		// 心跳
		case <-ticker.C:
			client.Connet.SetWriteDeadline(time.Now().Add(writeWait))
			if err := client.Connet.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

// WebSocket连接对象
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WebSocket 实现
func StrartWebSocket(manage *Manage, w http.ResponseWriter, r *http.Request) {
	// websocket 连接对象
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// 客户连接信息
	client := &Client{Connet: ws, send: make(chan []byte, 256)}
	// 注册用户
	manage.Register <- client
	// 等待读写
	go client.Write(manage)
	go client.Read(manage)
}

