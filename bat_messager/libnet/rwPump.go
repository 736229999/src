package libnet

import (
	"github.com/prometheus/common/log"
)

func (session *Session)ReadOnece(tmpBuffer,buffer []byte,readerChannel chan []byte)  {
	//声明一个临时缓冲区，用来存储被截断的数据


	//声明一个管道用于接收解包的数据

	n, err := session.Conn().Read(buffer)
	if err != nil {
		log.Info(session.Conn().RemoteAddr().String(), " connection error: ", err)
		return
	}

	tmpBuffer = Unpack(append(tmpBuffer, buffer[:n]...), readerChannel)

}
func (session *Session)ReadPump()  {
	tmpBuffer := make([]byte, 0)
	readerChannel := make(chan []byte, 16)
	go reader(readerChannel)
	buffer := make([]byte, 1024)
	for  {
		session.ReadOnece(tmpBuffer,buffer,readerChannel)
	}
}
func (session *Session)Write([]byte)  {

}

func reader(readerChannel chan []byte) {
	for {
		select {
		case data := <-readerChannel:
			log.Info("新的")
			log.Info(string(data))
		}
	}
}