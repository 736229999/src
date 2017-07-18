package main

import (
	"errors"
	"fmt"
	"log"
	"net"

	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	usercenter "github.com/caojunxyz/mimi-server/apps/usercenter/proto"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

var ErrInvalidParam = errors.New("无效参数")

func (srv *UcServer) ServeGRPC() {
	log.Println("ServeGRPC")
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
	usercenter.RegisterUsercenterServer(grpcServer, srv)
	if err := grpcServer.Serve(lis); err != nil {
		log.Println(err)
		return
	}
}

func (srv *UcServer) NotifyRecharged(ctx context.Context, result *usercenter.RechargeResult) (*usercenter.Nil, error) {
	log.Printf("%+v\n", result)
	accountId := result.GetAccountId()
	if accountId <= 0 {
		log.Println("账户无效:", accountId)
		return nil, ErrInvalidParam
	}
	money := float64(result.GetMoney()) / float64(100)
	if money <= 0 {
		log.Println("金额无效:", money)
		return nil, ErrInvalidParam
	}

	orderNo := result.GetOrderNo()
	if orderNo == "" {
		log.Println("无效orderNo")
		return nil, ErrInvalidParam
	}
	method := result.GetMethod()
	if method == "" {
		log.Println("无效来源")
		return nil, ErrInvalidParam
	}
	arg := &dbproto.RechargeResult{AccountId: accountId, Money: money, Method: method, OrderNo: orderNo}
	_, err := srv.dbClient.UpdateRechargeResult(context.Background(), arg)
	if err != nil {
		return nil, err
	}
	return &usercenter.Nil{}, nil
}
