package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)
var quitSemaphore chan bool
func main() {
	var tcpAddr *net.TCPAddr
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("连接成功")
	conn.Write([]byte("客户端连接成功"))
	<-quitSemaphore
	defer conn.Close()

	
}
func onMessageRecieved(conn *net.TCPConn)  {
	reader := bufio.NewReader(conn)
	for  {
		msg, err := reader.ReadString('\n')
		fmt.Println(msg)
		if err != nil {
			quitSemaphore <- true
			break
		}
		time.Sleep(time.Second)
		b := []byte(msg)
		conn.Write(b)
	}

}