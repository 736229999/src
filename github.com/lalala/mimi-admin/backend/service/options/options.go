package options

import (
	// dbproto "github.com/caojunxyz/mimi-server/dbagent/proto"
	// "google.golang.org/grpc"
	// "log"
	"github.com/caojunxyz/mimi-admin/backend/core"
)

// OPTIONS_DBAGENT_ADDR 平台配置Dbagent地址
const OPTIONS_DBAGENT_ADDR = "127.0.0.1:6012"

// OptionsService 平台配置相关服务
type OptionsService struct {
	core.Service
	// DbClient dbproto.DbOptionsAgentClient
}

// NewOptionsService ...
//func NewOptionsService() *OptionsService {
//	return &OptionsService{}
//}

// ConnectDb 连接到平台配置数据库
// func (srv *OptionsService) ConnectDb() dbproto.DbOptionsAgentClient {
// 	// conn, err := grpc.Dial(OPTIONS_DBAGENT_ADDR, grpc.WithInsecure())
// 	// if err != nil {
// 	// 	log.Panic(err)
// 	// }
// 	// srv.DbClient = dbproto.NewDbOptionsAgentClient(conn)
// 	// return srv.DbClient
// }

// Db 单例模式使用DB
// func (srv *OptionsService) Db() dbproto.DbOptionsAgentClient {
// 	if srv.DbClient == nil {
// 		srv.DbClient = NewOptionsService().ConnectDb()
// 	}
// 	return srv.DbClient
// }
