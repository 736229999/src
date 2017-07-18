package main

import (
	"log"

	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"google.golang.org/grpc"
)

type OptionsServer struct {
	optionsClient dbproto.DbOptionsAgentClient
}

func NewServer() *OptionsServer {
	return &OptionsServer{}
}

func (srv *OptionsServer) ConnectDb() {

	conn, err := grpc.Dial(*db, grpc.WithInsecure())
	if err != nil {
		log.Panic(err)
	}
	srv.optionsClient = dbproto.NewDbOptionsAgentClient(conn)
}
