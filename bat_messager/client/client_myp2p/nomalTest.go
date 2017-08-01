
package main

import (
	"flag"
	"bat_messager/log"
	"fmt"
	"net"
	"bat_messager/protocol"
	"encoding/json"
	"bat_messager/base"
	"time"
)



func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")
}

type test struct {
	Name string
	Age int
	ProcType string
}
func main() {
	var err error
	gatewayClient, err := net.Dial("tcp","127.0.0.1:17000")
	//gatewayClient, err := net.Dial("tcp","192.168.2.29:17000")
	if err != nil {
		panic(err)
	}
	msg := &base.LoginData{
		ProcType:"login_msg_server",
		Time:time.Now().Unix(),
		ClientId:"111",
	}
	jsonData,_ := json.Marshal(msg)
	data := protocol.Packet(jsonData)
	log.Info(string(data))
	testData := protocol.Packet([]byte("123"))
	log.Info(testData)
	for i:=0;i<1;i++  {
		//_,err = gatewayClient.Write([]byte("nihao"))
		_,err = gatewayClient.Write(testData)
	}
	da := make([]byte,1024)
	_,err = gatewayClient.Read(da)
	if err != nil {
		log.Info(err)
		return
	}
	log.Info(string(da))
	if err != nil {
		fmt.Println(err)
		return
	}
	gatewayClient.Read(da)

	log.Flush()
}

