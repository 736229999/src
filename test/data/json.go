package main

import (
	"fmt"
	"bytes"
	"encoding/binary"
)

type data struct {
	Name string	`json:"name"`
	Age int		`json:"age"`
}
func main() {
}

func test(i int) []byte {
	fmt.Println(i)
	return []byte("123")
}
func IntToBytes(n int64) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}