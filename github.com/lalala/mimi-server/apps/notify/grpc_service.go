package main

import (
	"fmt"
	"log"
	"net"

	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	notifyproto "github.com/caojunxyz/mimi-server/apps/notify/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"time"

	apiproto "github.com/caojunxyz/mimi-api/proto"
	"github.com/caojunxyz/mimi-server/apps/notify/hub"
	pb "github.com/golang/protobuf/proto"
)

func (srv *NotifyServer) ServeGRPC() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	log.Println("ServeGRPC on port ", *grpcPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *grpcPort))
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	notifyproto.RegisterNotifyServer(grpcServer, srv)
	if err := grpcServer.Serve(lis); err != nil {
		log.Println(err)
		return
	}
}

// 推送公告
func (srv *NotifyServer) PushAnnounce(ctx context.Context, arg *notifyproto.PushAnnounceArg) (*notifyproto.NotifyId, error) {
	notifyId, err := srv.dbClient.CreateNotify(context.Background(), &dbproto.Notify{Type: dbproto.NotifyType_Announce,
		Content: arg.Content,
		Sender:  arg.Sender,
	})
	if err != nil {
		log.Println("srv.dbClient.CreateNotify", err)
		return nil, err
	}
	h := hub.GetInstance()
	users := h.GetOnlineUsers()
	if len(users) == 0 {
		log.Println("No user online!")
		return &notifyproto.NotifyId{Value: 0}, nil
	}
	msg := apiproto.Notify{
		Type:          apiproto.NotifyType_Announce,
		Content:       arg.GetContent(),
		ContentLength: int64(len(arg.GetContent())),
		Target:        arg.GetTarget(),
		Action:        arg.GetAction(),
		TargetType:    arg.GetTargetType(),
		Created:       time.Now().Unix(),
		Sender:        arg.GetSender(),
	}
	pbMsg, err := pb.Marshal(&msg)
	if err != nil {
		log.Println("pb.Marshal", err)
	}
	go h.Push(&hub.Notify{
		Type:   hub.NOTIFY_ANNOUNCE,
		Data:   pbMsg,
		Sender: arg.GetSender(),
	})
	//go h.Push(&hub.Notify{Type:hub.NOTIFY_ANNOUNCE,
	//	Data:arg.Content,
	//	Sender:arg.Sender,
	//})
	log.Printf("Online users is %v\n", users)
	for _, user := range users {
		srv.dbClient.CreateUserNotify(context.Background(), &dbproto.UserNotify{
			Account:    user,
			Notify:     notifyId.Value,
			NotifyType: dbproto.NotifyType_Announce,
		})
	}
	return &notifyproto.NotifyId{Value: notifyId.Value}, nil
}

// 推送提醒
func (srv *NotifyServer) PushRemind(ctx context.Context, arg *notifyproto.PushRemindArg) (*notifyproto.NotifyId, error) {
	notifyId, err := srv.dbClient.CreateNotify(context.Background(), &dbproto.Notify{Type: dbproto.NotifyType_Remind,
		Content:    arg.Content,
		Sender:     arg.Sender,
		Target:     arg.Target,
		TargetType: arg.TargetType,
		Action:     arg.Action,
	})
	if err != nil {
		log.Println("srv.dbClient.CreateNotify", err)
		return nil, err
	}
	h := hub.GetInstance()
	if h.CheckUserOnline(arg.ToAccount) {

		msg := apiproto.Notify{
			Type:          apiproto.NotifyType_Remind,
			Content:       arg.GetContent(),
			ContentLength: int64(len(arg.GetContent())),
			Target:        arg.GetTarget(),
			Action:        arg.GetAction(),
			TargetType:    arg.GetTargetType(),
			Created:       time.Now().Unix(),
			Sender:        arg.GetSender(),
		}
		pbMsg, err := pb.Marshal(&msg)
		if err != nil {
			log.Println("pb.Marshal", err)
		}

		h.Push(&hub.Notify{
			Type:   hub.NOTIFY_REMIND,
			ToUser: arg.ToAccount,
			Data:   pbMsg,
			Sender: arg.Sender,
		})
	}
	srv.dbClient.CreateUserNotify(context.Background(), &dbproto.UserNotify{
		Account:    arg.ToAccount,
		Notify:     notifyId.Value,
		NotifyType: dbproto.NotifyType_Remind,
	})
	return &notifyproto.NotifyId{Value: notifyId.Value}, nil
}
