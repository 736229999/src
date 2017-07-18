package lottery

import (
	"github.com/caojunxyz/mimi-api/proto"

)

//-------------------------------------------------------------------------------------------------------------------------------------------
var JczqConfig = Config{
	Name:            "竞彩足球",
	Code:            Jczq,
	Id:              apiproto.LotteryId_Jczq,
	Type:            apiproto.LotteryType_Comp,
}

//-------------------------------------------------------------------------------------------------------------------------------------------
var JclqConfig = Config{
	Name:            "竞彩篮球",
	Code:            Jclq,
	Id:              apiproto.LotteryId_Jclq,
	Type:            apiproto.LotteryType_Comp,
}