package main

import (
	"bufio"
	"fmt"
	"net"
	//"time"
)
func main(){
	var tcpAddr *net.TCPAddr
	tcpAddr ,err := net.ResolveTCPAddr("tcp",":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	tcpListener,err := net.ListenTCP("tcp",tcpAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("准备监听")
	defer tcpListener.Close()
	for {
		fmt.Println("进入循环监听")
		tcpConn,err := tcpListener.AcceptTCP()
		if err!=nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("现在有客户端进来，地址是：",tcpConn.RemoteAddr().String())
		go tcpPipe(tcpConn)
	}

}

func tcpPipe(conn *net.TCPConn)  {
	ipStr := conn.RemoteAddr().String()
	defer func() {
		fmt.Println(ipStr,"远程客户端关闭")
		conn.Close()
	}()
	reader := bufio.NewReader(conn)
	for  {
		message ,err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("收到的消息",message)
		//现在往里面写
		conn.Write([]byte("你好呀"))
	}
}