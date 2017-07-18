package zhongfu

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// 投注方案
type Scheme struct {
	Type   string `json:"type"`    // 玩法名称
	TypeId string `json:"type_id"` // 玩法id
	Balls  string `json:"balls"`   // 投注号码
	Money  string `json:"money"`   // 金额
	Num    string `json:"num"`     // 注数
}

type Order struct {
	Name        string   `json:"name"`         // 账户名
	Secret      string   `json:"secret"`       // 专用密码
	Issue       string   `json:"issue"`        // 期号
	Schemes     []Scheme `json:"schemes"`      // 投注方案列表
	SumMoney    string   `json:"sum_money"`    // 总金额
	SumNum      string   `json:"sum_num"`      // 总注数
	Multiple    string   `json:"multiple"`     // 倍数
	LotteryId   string   `json:"lottery_id"`   // 彩种id
	LotteryName string   `json:"lottery_name"` // 彩种名称
	OrderId     string   `json:"order_id"`     // 透传参数
	IdNo        string   `json:"idno"`         // 投注者身份证号码
}

func FormatBalls(balls []int32, width int) []string {
	ret := []string{}
	format := fmt.Sprintf("%%0%dd", width)
	for _, v := range balls {
		ret = append(ret, fmt.Sprintf(format, v))
	}
	return ret
}

func MakeBalls(format string, balls ...[]string) string {
	list := []interface{}{}
	for _, v := range balls {
		list = append(list, strings.Join(v, " "))
	}
	result := fmt.Sprintf(format, list...)
	log.Println(result)
	return result
}

func MakeBalls2(format string, sep string, balls ...[]string) string {
	list := []interface{}{}
	for _, v := range balls {
		list = append(list, strings.Join(v, sep))
	}
	result := fmt.Sprintf(format, list...)
	log.Println(result)
	return result
}

// 测试
const (
	SERVER         = "http://120.77.34.33"
	KEY            = "60D81B05"
	ACCOUNT_NAME   = "xiaomi"
	ACCOUNT_SECRET = "xiaomi"
)

// 生产
// const (
// 	SERVER         = "http://47.88.168.86"
// 	KEY            = "60D81B05"
// 	ACCOUNT_NAME   = "cs008"
// 	ACCOUNT_SECRET = "xiaomi"
// )

type CommitResult struct {
	Status    string `json:"status"`
	Desc      string `json:"desc"`
	OrderId   string `json:"order_id"`
	ZfOrderId string `json:"zf_order_id"`
	IdNo      string `json:"idno"`
	TradeTime string `json:"trade_time"`
	SumMoney  string `json:"sum_money"`
	Balance   string `json:"balance"`
}

func CommitOrder(order *Order) (*CommitResult, error) {
	log.Printf("%+v\n", order)
	order.Name = ACCOUNT_NAME
	order.Secret = ACCOUNT_SECRET
	crypted, err := ZhongfuEncrypt(order)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	url := fmt.Sprintf("%s/api/LotteryApiSZ.ashx", SERVER)
	resp, err := http.Post(url, "text/plain", bytes.NewBuffer(crypted))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var result CommitResult
	err = ZhongfuDecrypt(data, &result)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Printf("%+v\n", result)
	return &result, nil
}

const (
	NOT_SOLD = 1 // 未出票
	SOLD     = 2 // 已出票
	WIN      = 3 // 中奖
	NOT_WIN  = 4 // 未中奖
	CANCEL   = 5 // 已撤销
)

type VerifyResult struct {
	Status      string `json:"status"`
	Desc        string `json:"desc"`
	OrderId     string `json:"order_id"`     // 我方(开心客人)订单id
	ZfOrderId   string `json:"zf_order_id"`  // 中福订单id
	IdNo        string `json:"idno"`         // 身份证号码
	OrderStatus string `json:"order_status"` // 订单状态：1-未出票, 2-已出票, 3-中奖, 4-未中奖, 5-已撤销
	Bonus       string `json:"bonus"`        // 中奖金额
}

func VerifyOrder(order_id, zf_order_id, idno string) (int, float64, error) {
	args := map[string]string{
		"name":        ACCOUNT_NAME,
		"secret":      ACCOUNT_SECRET,
		"order_id":    order_id,
		"zf_order_id": zf_order_id,
		"idno":        idno,
	}
	crypted, err := ZhongfuEncrypt(args)
	if err != nil {
		log.Println(err)
		return 0, 0, err
	}
	url := fmt.Sprintf("%s/api/OrderVerify.ashx", SERVER)
	resp, err := http.Post(url, "text/plain", bytes.NewBuffer(crypted))
	if err != nil {
		log.Println(err)
		return 0, 0, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return 0, 0, err
	}
	var result VerifyResult
	err = ZhongfuDecrypt(data, &result)
	if err != nil {
		log.Println(err)
		return 0, 0, err
	}
	log.Printf("result: %+v\n", result)
	if result.Status != "0" {
		err = fmt.Errorf("%s-%s", result.Status, result.Desc)
		return 0, 0, err
	}

	var status int
	var bonus float64
	status, err = strconv.Atoi(result.OrderStatus)
	if err != nil {
		log.Println(err)
		return 0, 0, err
	}
	if status != NOT_SOLD && status != SOLD && status != WIN && status != NOT_WIN && status != CANCEL {
		err = fmt.Errorf("无效订单状态: %d", status)
		return 0, 0, err
	}
	if status == WIN {
		bonus, err = strconv.ParseFloat(result.Bonus, 64)
	}
	return status, bonus, nil
}
