
package main

import (
	"flag"
	"bat_messager/log"
	"fmt"
	"unsafe"
	"bytes"
	binary "encoding/binary"
	"bat_messager/protocol"
	"bat_messager/libnet"
)

var InputConfFile = flag.String("conf_file", "client.json", "input conf file name")
const (
	ConstHeader         = "www.01happy.com"
	ConstHeaderLength   = 15
	ConstSaveDataLength = 4
)

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")
}
func Packet(message []byte) []byte {
	return append(append([]byte(ConstHeader), IntToBytes(len(message))...), message...)
}

//整形转换成字节
func IntToBytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

func main() {
	var err error
	gatewayClient, err := libnet.Dial("tcp","127.0.0.1:19001")
	//gatewayClient, err := net.Dial("tcp","192.168.2.29:17000")
	if err != nil {
		panic(err)
	}

	//msg := "要发送的数据"
	cmd := protocol.NewCmdSimple(protocol.SEND_MESSAGE_P2P_CMD)
	cmd.AddArg("id是：1")
	cmd.AddArg("测试数据")
	cmd.AddArg("哪个发送的："+gatewayClient.Conn().LocalAddr().String())
	err = gatewayClient.Send(libnet.Json(cmd))
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Flush()
}
type InBuffer struct {
	Data    []byte // Buffer data.
	ReadPos int    // Read position.
	isFreed bool
	next    unsafe.Pointer
}
func Write()  {

}
