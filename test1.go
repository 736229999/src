package main

import (
	"fmt"
	"bytes"
	"encoding/binary"
)
const (
	ConstHeader         = "www.bat_messager.com"
	ConstHeaderLength   = 15
	ConstSaveDataLength = 4
)

type T struct {
	name string
}
func main(){
	a := make(chan bool)
	b := make(chan bool)
	go test1(a)
	test2(b)
	fmt.Println(123)
}
func test1(data chan bool)  {
	for  {
		select {
		case a := <-data:
			fmt.Println("receive:",a)
		}
	}

}
func test2(data chan bool)  {
	data <- false
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