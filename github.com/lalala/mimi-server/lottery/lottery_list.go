package lottery

import (
	"log"
	"github.com/caojunxyz/mimi-api/proto"
)

var lottery = []Config {
	cqsscConfig,
	bjpk10Config,
	gd11x5Config,
	dltConfig,
	ssqConfig,
	pl3Config,
	pl5Config,
	fc3dConfig,
	JczqConfig,
	JclqConfig,
}

type LotteryType struct {
	Type apiproto.LotteryType `json:"type"`
	Name string `json:"name"`
}
func GetLottery() []Config {
	return lottery
}

func GetLotteryName(lottery_id int64) string {

	var name string
	for _,v := range lottery {
		if int64(v.Id) == lottery_id {
			name = v.Name
			log.Println(int32(v.Id), lottery_id, v.Name)
			break
		}
	}
	return name
}

//获取所有的彩票类型.
func GetLotteryTypeList() []LotteryType {

	list := []LotteryType{
		{apiproto.LotteryType_LowFreq, GetLotteryTypeName(apiproto.LotteryType_LowFreq)},
		{apiproto.LotteryType_HighFreq, GetLotteryTypeName(apiproto.LotteryType_HighFreq)},
		{apiproto.LotteryType_Comp, GetLotteryTypeName(apiproto.LotteryType_Comp)},
	}
	return list
}

