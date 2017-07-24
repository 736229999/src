package main

import (
	"flag"
	"bat_messager/libnet"
	"github.com/oikomi/FishChatServer/log"
)

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")
}
var InputConfFile = flag.String("conf_file", "gateway.json", "input conf file name")

func handleSession(gw *Gateway,session *libnet.Session)  {
	session.Process(func(msg *libnet.InBuffer) error {
		err := gw.parseProtocol(msg.Data, session)
		if err != nil {
			log.Error(err.Error())
		}

		return nil
	})
}
func main() {
	flag.Parse()
	cfg := NewGatewayConfig(*InputConfFile)
	err := cfg.LoadConfig()
	if err != nil {
		log.Error(err.Error())
		return
	}
	gw := NewGateway(cfg)
	gw.server, err = libnet.Listen(cfg.TransportProtocols, cfg.Listen)
	if err != nil {
		log.Error(err.Error())
		return
	}
	log.Info("gateway server running at ", gw.server.Listener().Addr().String())
	gw.server.Serve(func(session *libnet.Session) {
		log.Info("client ", session.Conn().RemoteAddr().String(), " | come in")

		go handleSession(gw, session)
	})

}