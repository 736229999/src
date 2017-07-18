package sms

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/caojunxyz/gotu"
	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/apps/thirdapi/sms/cty"
	"github.com/caojunxyz/mimi-server/proto"
	"golang.org/x/net/context"
)

const (
	SmsType_Message = 1
	SmsType_Code    = 2
)

type SmsVendor interface {
	SendMessage(string, []string, string) (*dbproto.Sms, error)           // 发送消息
	SendCode(string, string, string, time.Duration) (*dbproto.Sms, error) // 发送验证码
}

type SmsServer struct {
	sign   string
	vendor SmsVendor
	rnd    *rand.Rand
	dbc    dbproto.DbThirdApiAgentClient
}

func NewServer(c dbproto.DbThirdApiAgentClient) *SmsServer {
	return &SmsServer{
		sign:   "亨通彩",
		vendor: &cty.CtySms{},
		rnd:    rand.New(rand.NewSource(time.Now().UnixNano())),
		dbc:    c,
	}
}

func (srv *SmsServer) genCode() string {
	code := ""
	for i := 0; i < 6; i++ {
		code += fmt.Sprint(srv.rnd.Intn(10))
	}
	return code
}

// 不能发送太频繁, 至少60秒间隔
// 每日最多发送10条
func (srv *SmsServer) isSendLimit(phone string) bool {
	stats, err := srv.dbc.QuerySmsStats(context.Background(), &dbproto.StringValue{Value: phone})
	if err != nil {
		log.Println(err)
		return true
	}

	countTime := time.Unix(stats.CountTime, 0)
	now := time.Now()
	if !gotu.IsSameDay(now, countTime) || now.Sub(countTime) > time.Hour {
		stats.CountTime = stats.LatestTime
		if !gotu.IsSameDay(now, countTime) {
			stats.DailyCount = 0
		}
		if now.Sub(countTime) > time.Hour {
			stats.HourlyCount = 0
		}
		srv.dbc.UpdateSmsStats(context.Background(), stats)
	}

	log.Printf("%s --> %+v\n", phone, stats)
	latestTime := time.Unix(stats.LatestTime, 0)
	// 两条短信至少间隔60s
	if now.Sub(latestTime) < time.Minute {
		log.Println(phone, now, latestTime)
		return true
	}

	// 每个小时最多发送10条
	if stats.HourlyCount > 10 {
		log.Println(phone, stats.HourlyCount)
		return true
	}

	// 每日最多发送20条
	if stats.DailyCount > 20 {
		log.Println(phone, stats.DailyCount)
		return true
	}
	return false
}

func (srv *SmsServer) SendSmsMessage(ctx context.Context, arg *proto.SmsRequest) (*proto.Bool, error) {
	result, err := srv.vendor.SendMessage(srv.sign, arg.GetPhoneList(), arg.GetContent())
	if err != nil {
		return nil, err
	}

	result.SmsType = SmsType_Message
	_, err = srv.dbc.InsertSms(ctx, result)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &proto.Bool{Value: result.IsSuccess}, nil
}

func (srv *SmsServer) SendSmsCode(ctx context.Context, arg *proto.SmsRequest) (*proto.Bool, error) {
	phone := arg.GetPhoneList()[0]
	if srv.isSendLimit(phone) {
		log.Printf("发送限制: %s\n", phone)
		return &proto.Bool{Value: false}, nil
	}

	code := srv.genCode()
	valid := time.Duration(arg.GetCodeValidDur())
	if valid == 0 {
		valid = time.Minute * 5
	}
	result, err := srv.vendor.SendCode(srv.sign, phone, code, valid)
	if err != nil {
		return nil, err
	}

	result.SmsType = SmsType_Code
	_, err = srv.dbc.InsertSms(ctx, result)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &proto.Bool{Value: result.IsSuccess}, nil
}

func (srv *SmsServer) VerifySmsCode(ctx context.Context, arg *proto.SmsRequest) (*proto.Bool, error) {
	phone := arg.GetPhoneList()[0]
	code := arg.GetCode()
	result, err := srv.dbc.QuerySms(context.Background(), &dbproto.StringValue{Value: phone})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if result.Code != code {
		log.Println("验证码不一致:", phone, code, result.Code)
		return nil, fmt.Errorf("验证码错误!")
	}
	if time.Now().Unix() >= result.ExpireTime {
		log.Println("验证码已过期！")
		return nil, fmt.Errorf("验证码已过期!")
	}

	srv.dbc.SetSmsExpired(ctx, &dbproto.IntValue{Value: result.Id})
	log.Println("验证成功!", phone, code)
	return &proto.Bool{Value: result.IsSuccess}, nil
}
