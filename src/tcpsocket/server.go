package main

import(
	"log"
	"net"
	"fmt"
)


//处理连接
func HandleConn(c net.Conn){
	defer c.Close()
	for  {
		//read from server
		var buf = make([]byte,10)
		log.Println("开始接收消息...：")
		n ,err :=c.Read(buf)
		if err != nil {
			log.Print(err)
			break
		}
		log.Print(n)
		log.Printf("%s",n)
	}
}

func main() {
	l,err := net.Listen("tcp",":8888")
	if err != nil {
		log.Println(err)
		return
	}
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			break
		}
		// start a new goroutine to handle
		// the new connection.
		go HandleConn(c)
	}
}

