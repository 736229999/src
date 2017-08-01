package main

import (
	"fmt"
	"net"
)
type Stu func(int) int
func main() {
	tcp()
}
func test(s Stu) int {
	return 1

}
func tcp() {

	ln, err := net.Listen("tcp", ":17000")
	if err != nil {
		panic(err)
	}

	defer ln.Close()

	for {
		tcpConn, err := ln.Accept()
		if err != nil {
			continue
		}

		fmt.Println("A client connected : " + tcpConn.RemoteAddr().String())
		go tcpPipe(tcpConn)

		go func() {
			fmt.Println("你好")
		}()
	}

}

func tcpPipe(conn net.Conn) {
	ipStr := conn.RemoteAddr().String()
	defer func() {
		fmt.Println("disconnected :" + ipStr)
		//conn.Close()
	}()
	data:=make([]byte,1024)
	_,err := conn.Read(data)
	if err != nil {
		fmt.Println("接收失败")
		return
	}
	fmt.Println(string(data))
	conn.Write([]byte("你好，我收到了消息"))
}