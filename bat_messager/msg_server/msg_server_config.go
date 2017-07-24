package main
import (
	"os"
	"encoding/json"
	"time"
	"github.com/oikomi/FishChatServer/log"
)

type MsgServerConfig struct {
	configfile               string
	LocalIP                  string
	TransportProtocols       string
	Listen                   string
	LogFile                  string
	ScanDeadSessionTimeout   time.Duration
	Expire                   time.Duration
	MonitorBeatTime          time.Duration
	SessionManagerServerList []string
	Redis struct {
		Addr string
		Port string
		ConnectTimeout time.Duration
		ReadTimeout time.Duration
		WriteTimeout time.Duration
	}
	//还有数据库配置
}

func NewMsgServerConfig(configfile string) *MsgServerConfig {
	return &MsgServerConfig{
		configfile : configfile,
	}
}

func (self *MsgServerConfig)LoadConfig() error {
	file, err := os.Open(self.configfile)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	defer file.Close()

	dec := json.NewDecoder(file)
	err = dec.Decode(&self)
	if err != nil {
		return err
	}
	return nil
}
