package main

import (
	"log"
	"net"
	"time"
	"os"
)

func EstablishConn(i int) net.Conn  {
	conn,err := net.DialTimeout("tcp","106.75.16.180:8088",5*time.Second)
	if err !=nil {
		log.Println(err)
		return nil
	}
	log.Println("建立了第",i,"个连接")
	return conn
}

func main() {
	if len(os.Args) <=1 {
		return
	}
	log.Println("begin dial.....")
	conn ,err := net.DialTimeout("tcp",":8888",5*time.Second)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	//time.Sleep(2*time.Second)
	data := os.Args[1]
	conn.Write([]byte(data))
	time.Sleep(1000*time.Second)
}
