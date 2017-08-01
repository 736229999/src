package base

import (
	"time"
)

//请求聊天服务器数据
type ReqMsgServer struct {

}
//登录数据
type LoginData struct {
	ProcType string		`json:"procType"`
	Time 	int64		`json:"time"`
	ClientId string		`json:"clientId"`
}

//内部通信
type InnerRelationMapping struct {
	ClientId  	  string	`json:"clientId"`
	ClientIpPort  string	`json:"clientIpPort"`
	ServerIpPort  string	`json:"serverIpPort"`
	Time 		  int64		`json:"time"`
}
//单聊数据
type P2pData struct {
	ClientId 	string  	`json:"clientId"`
	ToUid 		string		`json:"toUid"`
	Data 		string		`json:"data"`
	Time 		int64 		`json:"time"`
}

func NewInnerRetionMapping(clientId ,ClientIpPort,ServerIpPort string) *InnerRelationMapping {
	return &InnerRelationMapping{
		ClientId:clientId,
		ClientIpPort:ClientIpPort,
		ServerIpPort:ServerIpPort,
		Time: time.Now().Unix(),
	}
}