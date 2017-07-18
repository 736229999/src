package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	apiproto "github.com/caojunxyz/mimi-api/proto"
	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/apps/notify/hub"
	"github.com/caojunxyz/mimi-server/auth"
	"github.com/caojunxyz/mimi-server/utils"
	"github.com/gorilla/websocket"
	"golang.org/x/net/context"
)

const (
	writeWait = 2 * time.Second

	// 每隔多少时间发送一次Ping，必须小于pongWait;pingPeriod = 6 * time.Second
	pingPeriod = (pongWait * 9) / 10

	// 读取对方发送过来的pong的时间间隔，超过这个时间websocket状态就变为corrupt
	pongWait = 10 * time.Second
)

func (srv *NotifyServer) ServeHTTP() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	log.Println("ServeWebsocket on port ", *httpPort)
	mux := http.NewServeMux()
	// test mode
	mux.HandleFunc("/notify/test", srv.HandleTest)
	// auth.Validate()

	// prod mode
	mux.HandleFunc("/notify", auth.WsValidate(srv.HandleNotify))
	mux.HandleFunc("/notify/pull", auth.Validate(srv.HandlePullNotifyCenterDev))
	mux.HandleFunc("/notify/read", auth.Validate(srv.HandleReadNotifyDev))

	// dev mode
	mux.HandleFunc("/notify/dev", srv.HandleNotifyDev)
	mux.HandleFunc("/notify/dev/pull", auth.Validate(srv.HandlePullNotifyCenterDev))
	mux.HandleFunc("/notify/dev/read", auth.Validate(srv.HandleReadNotifyDev))

	httpServer := http.Server{
		Addr:         fmt.Sprintf(":%d", *httpPort),
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	if err := httpServer.ListenAndServe(); err != nil {
		panic(err)
	}
}

func (srv *NotifyServer) HandleNotify(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handle Notify")
	accountId, _, err := utils.ParseHttpRequest(w, r, nil)
	if err != nil {
		return
	}
	wsConn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if err != nil {
		log.Printf("websocket %v\n", err)
		return
	}

	h := hub.GetInstance()
	h.Join(accountId, wsConn)
	defer h.Leave(accountId)

	wsConn.SetPongHandler(func(s string) error {
		log.Println("Handle Pong ", s)
		wsConn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	wsConn.SetPingHandler(HandlePing)
	//stdoutDone := make(chan struct{})
	if !h.CheckUserOnline(accountId) {
		go ping(wsConn)
	}
	//select {}
	for {
		_, message, err := wsConn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("Received: %s", message)
		//err = wsConn.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func (srv *NotifyServer) HandleNotifyDev(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandleNotifyCenter !Dev")
	r.ParseForm()
	uidStr := r.Form["uid"]
	uid, _ := strconv.ParseInt(uidStr[0], 10, 0)
	log.Printf("------>uid is %v\n", uid)
	wsConn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if err != nil {
		log.Printf("websocket %v\n", err)
		return
	}
	h := hub.GetInstance()
	h.Join(uid, wsConn)
	defer h.Leave(uid)

	wsConn.SetPongHandler(func(s string) error {
		log.Println("Handle Pong ", s)
		wsConn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	wsConn.SetPingHandler(HandlePing)
	//stdoutDone := make(chan struct{})
	if !h.CheckUserOnline(uid) {
		go ping(wsConn)
	}

	for {
		_, message, err := wsConn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("Received: %s", message)
		//err = wsConn.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}

}

//done chan struct{}
// ping用于心跳检测，检测连接是否可用
func ping(ws *websocket.Conn) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		ws.Close()
	}()
	for {
		select {
		case <-ticker.C:
			// 服务器在writeWait之后没有写入数据，则websocket状态变为corrupt
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				log.Println("Ping:", err)
				return
			}
		}
	}
}

func HandlePong(appData string) error {
	log.Println("Handle Pong ", appData)
	return nil
}

func HandlePing(appData string) error {
	log.Println("Handle Ping ", appData)
	return nil
}

func (srv *NotifyServer) HandlePullNotifyCenterDev(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handle Pull NotifyCenter !Dev")

	var msg apiproto.QueryUserNotifyArg
	accountid, ip, err := utils.ParseHttpRequest(w, r, &msg)
	if err != nil {
		log.Println("utils.ParseHttpRequest", err)
		return
	}
	//fmt.Printf("============uid is %v==========\n", uid)
	fmt.Printf("============uid is %v==========ip %v\n", accountid, ip)
	fmt.Printf("requset is %+v\n", msg)

	res, err := srv.dbClient.QueryNotifyUserMissed(context.Background(), &dbproto.QueryUserMissedArg{Account: accountid, Type: dbproto.NotifyType_All})
	if err != nil {
		log.Println("QueryNotifyUserMissed ", err)
		http.Error(w, "QueryNotifyUserMissed!", http.StatusInternalServerError)
		return
	}
	if len(res.Notices) > 0 {
		for _, value := range res.Notices {
			log.Println(value.Id, value.Type)
			_, err := srv.dbClient.CreateUserNotify(context.Background(), &dbproto.UserNotify{
				Notify:     value.Id,
				NotifyType: value.Type,
				Account:    accountid,
			})
			if err != nil {
				log.Println("CreateUserNotify ", err)
				http.Error(w, "QueryNotifyUserMissed!", http.StatusInternalServerError)
				return
			}
		}
	}
	nsInfo, err := srv.dbClient.QueryUserNotify(context.Background(), &dbproto.QueryUserNotifyArg{Account: accountid,
		NotifyType: dbproto.NotifyType(msg.NotifyType),
		Page:       msg.Page,
		PageSize:   msg.PageSize,
	})
	if err != nil {
		log.Println("QueryUserNotify ", err)
		http.Error(w, "QueryUserNotify!", http.StatusInternalServerError)
		return
	}
	result := &apiproto.QueryUserNotifyRes{}
	for _, value := range nsInfo.UserNoticeInfos {
		unInfo := &apiproto.UserNotifyInfo{
			Notify: &apiproto.Notify{
				Id:      value.Notify.Id,
				Content: value.Notify.Content,
				Type:    apiproto.NotifyType(value.Notify.Type),
				Sender:  value.Notify.Sender,
				Created: value.Notify.Created,
			},
			UserNotify: &apiproto.UserNotify{
				Account: value.UserNotify.Account,
				IsRead:  value.UserNotify.IsRead,
			},
		}
		result.List = append(result.List, unInfo)
	}
	log.Printf("nsInfo Is %v\n", nsInfo.UserNoticeInfos)

	log.Printf("result is %+v, %d\n", result, len(result.GetList()))

	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", result)
}

// 读取某个消息
func (srv *NotifyServer) HandleReadNotifyDev(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handle Read Notify! Dev")
	var msg = apiproto.ReadNotifyArg{}
	accountId, _, err := utils.ParseHttpRequest(w, r, &msg)
	if err != nil {
		log.Println("utils.ParseHttpRequest", err)
		return
	}
	log.Printf("accountId is %v", accountId)
	dbRes, err := srv.dbClient.ReadUserNotify(context.Background(), &dbproto.ReadUserNotifyArg{
		AccountId: msg.GetAccountId(),
		NotifyId:  msg.GetNotifyId(),
	})
	if err != nil {
		log.Println("dbClient.ReadUserNotify", err)
		http.Error(w, "ReadUserNotify!", http.StatusInternalServerError)
		return
	}
	result := &apiproto.IntValue{Value: dbRes.Value}
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "OK", result)

}

func (srv *NotifyServer) HandleTest(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleTest")
	wsConn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if err != nil {
		log.Printf("websocket %v\n", err)
		return
	}
	for {
		mt, message, err := wsConn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = wsConn.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
