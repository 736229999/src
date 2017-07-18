package lottery

import (
	"fmt"
	"log"
	"strings"

	apiproto "github.com/caojunxyz/mimi-api/proto"
	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/lottery/saleissue"
	"github.com/caojunxyz/mimi-server/lottery/site/zhongfu"
	"github.com/caojunxyz/mimi-server/lottery/validator"
)

const (
	Dlt    = "dlt"
	Fc3d   = "fc3d"
	Ssq    = "ssq"
	Cqssc  = "cqssc"
	Bjpk10 = "bjpk10"
	Gd11x5 = "gd11x5"
	Pl3    = "pl3"
	Pl5    = "pl5"
	Jczq   = "jczq"
	Jclq   = "jclq"
)

type Config struct {
	Name            string
	Code            string
	Id              apiproto.LotteryId
	Type            apiproto.LotteryType
	ZfId            string
	BlueNum         int32
	DayMaxNo        int   // 每日期数
	MaxMultiple     int32 // 最大倍数
	BonusNames      map[int]string
	PlayList        map[string]validator.PlayType
	IssuesGenerator func(saleissue.SaleIssue, int) []saleissue.SaleIssue
}

func GetConfig(id apiproto.LotteryId) *Config {
	switch id {
	case apiproto.LotteryId_Dlt:
		return &dltConfig
	case apiproto.LotteryId_Fc3d:
		return &fc3dConfig
	case apiproto.LotteryId_Ssq:
		return &ssqConfig
	case apiproto.LotteryId_Cqssc:
		return &cqsscConfig
	case apiproto.LotteryId_Bjpk10:
		return &bjpk10Config
	case apiproto.LotteryId_Gd11x5:
		return &gd11x5Config
	case apiproto.LotteryId_Pl3:
		return &pl3Config
	case apiproto.LotteryId_Pl5:
		return &pl5Config
	}
	return nil
}

func GetLotteryTypeName(typ apiproto.LotteryType) string {
	switch typ {
	case apiproto.LotteryType_LowFreq:
		return "低频彩"
	case apiproto.LotteryType_HighFreq:
		return "高频彩"
	case apiproto.LotteryType_Comp:
		return "竞彩"
	}
	return ""
}

func MakeOrderNo(cfg *Config, issue string, orderId int64) string {
	return fmt.Sprintf("%s%s-%d", strings.ToUpper(cfg.Code), issue, orderId)
}

//-------------------------------------------------------------------------------------------------------------------------------------------
func Convert(vendorOrder *dbproto.BuycaiVendorOrder, cfg *Config) *zhongfu.Order {
	ret := &zhongfu.Order{}
	schemeList := vendorOrder.GetSchemeList()
	for _, sch := range schemeList {
		typ := sch.GetType()
		playType, ok := cfg.PlayList[typ]
		if !ok {
			log.Println("无效玩法:", typ)
			return nil
		}
		zfScheme := playType.ZfConverter(sch, playType.ZfId, playType.BallsFormat)
		ret.Schemes = append(ret.Schemes, zfScheme)
	}
	issue := vendorOrder.GetIssue()
	ret.Issue = issue
	ret.OrderId = MakeOrderNo(cfg, issue, vendorOrder.GetId())
	ret.SumMoney = fmt.Sprint(vendorOrder.GetMoney())
	ret.SumNum = fmt.Sprint(vendorOrder.GetSumNum())
	ret.Multiple = fmt.Sprint(vendorOrder.GetMultiple())
	ret.LotteryId = fmt.Sprint(cfg.ZfId)
	ret.LotteryName = cfg.Name
	ret.IdNo = fmt.Sprint(vendorOrder.GetAccountId())

	return ret
}

func Validate(order *apiproto.BuycaiOrder, cfg *Config) bool {
	imList := order.GetIssues()
	if len(imList) == 0 {
		log.Println("没有期号")
		return false
	}
	if order.GetTicketId() > 0 && len(imList) > 1 {
		log.Println("追号订单不能使用购彩券")
		return false
	}

	schemeList := order.GetSchemeList()
	money := float64(0)
	for _, sch := range schemeList {
		typ := sch.GetType()
		playType, ok := cfg.PlayList[typ]
		if !ok {
			log.Println("无效玩法:", typ)
			return false
		}
		if !playType.Validator(sch) {
			return false
		}
		money += sch.GetMoney()
	}

	sumMoney := float64(0)
	uniqueIssues := map[string]bool{}
	for _, v := range imList {
		issue := v.GetIssue()
		if uniqueIssues[issue] {
			log.Println("重复期号:", issue)
			return false
		}
		uniqueIssues[issue] = true
		multiple := v.GetMultiple()
		if multiple < 1 || multiple > cfg.MaxMultiple {
			log.Println("无效倍数:", multiple)
			return false
		}
		sumMoney += (float64(multiple) * money)
	}

	if !validator.IsEqualMoney(sumMoney, order.GetSumMoney()) {
		log.Println("总金额不一致:", sumMoney, order.GetSumMoney())
		return false
	}
	return true
}
