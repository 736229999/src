package opencai

import (
	"fmt"
	"io"
	"log"
	"time"

	apiproto "github.com/caojunxyz/mimi-api/proto"
	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/apps/thirdapi/opencai/apiplus"
	"github.com/caojunxyz/mimi-server/lottery"
	"github.com/caojunxyz/mimi-server/utils"
	"golang.org/x/net/context"
)

type OpencaiVendor interface {
	QueryDateIssues(code, date string) []*dbproto.OpenInfo
	QueryLatest(code string) *dbproto.OpenInfo
	QueryLatestNext(code string, latestRows int) []*dbproto.OpenInfo
}

type OpencaiServer struct {
	vendor OpencaiVendor
	dbc    dbproto.DbThirdApiAgentClient
}

func NewServer(c dbproto.DbThirdApiAgentClient) *OpencaiServer {
	srv := &OpencaiServer{
		vendor: &apiplus.Apiplus{},
		dbc:    c,
	}
	srv.run()
	return srv
}

func (srv *OpencaiServer) dbQueryIssuesByOpendate(code, date string) map[string]*dbproto.OpenInfo {
	result := make(map[string]*dbproto.OpenInfo)
	arg := &dbproto.OpencaiQueryArg{
		Code: code,
		Args: []string{date},
	}
	stream, err := srv.dbc.OpencaiQueryByOpendate(context.Background(), arg)
	if err != nil {
		log.Println(err)
		return nil
	}

	for {
		v, err := stream.Recv()
		if err == nil {
			result[v.GetIssue()] = v
			continue
		}
		if err != io.EOF {
			log.Println(err)
		}
		break
	}
	return result
}

func (srv *OpencaiServer) dbQueryIssuesByLatestNum(code string, num int) map[string]*dbproto.OpenInfo {
	result := make(map[string]*dbproto.OpenInfo)
	arg := &dbproto.OpencaiQueryArg{
		Code: code,
		Args: []string{fmt.Sprint(num)},
	}
	stream, err := srv.dbc.OpencaiQueryByLatestNum(context.Background(), arg)
	if err != nil {
		log.Println(err)
		return nil
	}

	for {
		v, err := stream.Recv()
		if err == nil {
			result[v.GetIssue()] = v
			continue
		}
		if err != io.EOF {
			log.Println(err)
		}
		break
	}
	return result
}

func (srv *OpencaiServer) dbInsertOpeninfo(code string, info *dbproto.OpenInfo) error {
	arg := &dbproto.OpencaiInsertArg{
		Code: code,
		Info: info,
	}
	_, err := srv.dbc.OpencaiInsert(context.Background(), arg)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (srv *OpencaiServer) runAgent(conf lottery.Config) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	switch conf.Type {
	case apiproto.LotteryType_HighFreq:
		srv.initHighfreq(&conf)
	case apiproto.LotteryType_LowFreq:
		srv.initLowfreq(&conf)
	}
}

func (srv *OpencaiServer) run() {
	log.Println("opencai server is running...")
	for _, conf := range lottery.GetLottery() {
		go srv.runAgent(conf)
	}
}

func (srv *OpencaiServer) initHighfreq(conf *lottery.Config) {
	code := conf.Code
	// 获取最近3天开奖信息
	for n := 2; n >= 0; n-- {
		dayTime := utils.TimeBeforeDays(n)
		date := dayTime.Format("2006-01-02")
		result := srv.dbQueryIssuesByOpendate(code, date)
		log.Println(code, date, len(result))
		if len(result) < conf.DayMaxNo {
			list := srv.vendor.QueryDateIssues(code, date)
			log.Printf("%s: 抓取到%s日%d条数据\n", code, date, len(list))
			for _, v := range list {
				issue := v.GetIssue()
				if result[issue] == nil {
					if err := srv.dbInsertOpeninfo(code, v); err != nil {
						log.Panicf("%s: 插入%s期抓取数据失败!", code, issue)
					}
				}
			}
		}
	}
	go srv.scheduleGrab(conf)
}

func (srv *OpencaiServer) initLowfreq(conf *lottery.Config) {
	code := conf.Code
	// 获取最近20条信息
	num := 20
	result := srv.dbQueryIssuesByLatestNum(code, num)
	list := srv.vendor.QueryLatestNext(code, num)
	if len(list) < 1 {
		log.Panicf("%s: 无效数据!", code)
	}
	log.Printf("%s: 抓取到%d条数据\n", code, len(list[1:]))
	for _, v := range list[1:] {
		issue := v.GetIssue()
		if result[issue] == nil {
			if err := srv.dbInsertOpeninfo(code, v); err != nil {
				log.Panicf("%s: 插入%s期抓取数据失败!", code, issue)
			}
		}
	}
	go srv.scheduleGrab(conf)
}

func (srv *OpencaiServer) scheduleGrab(conf *lottery.Config) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	code := conf.Code
	var latestIssue string

	for {
		log.Printf("开始抓取%s最新开奖信息...\n", code)
		list := srv.vendor.QueryLatestNext(code, 1)
		if len(list) != 2 {
			log.Printf("抓取数据错误:%d\n", len(list))
			continue
		}

		next := list[0]
		latest := list[1]
		latestOpenTime := time.Unix(latest.GetOpenTime(), 0)
		nextOpenTime := time.Unix(next.GetOpenTime(), 0)
		log.Printf("[%s]最新: %s (%v), 下一期: %s (%v)\n", code, latest.GetIssue(), latestOpenTime, next.GetIssue(), nextOpenTime)
		if latest.GetIssue() != latestIssue {
			log.Printf("抓取到%s最新一期%s开奖信息\n", code, latest.GetIssue())
			if err := srv.dbInsertOpeninfo(code, latest); err != nil {
				log.Panicf("%s: 插入%s期抓取数据失败!", code, latest.GetIssue())
			}
			latestIssue = latest.GetIssue()
		}

		dur := nextOpenTime.Sub(time.Now())
		if dur > 0 {
			log.Printf("[%s] %v后开始抓取\n", code, dur)
			<-time.NewTimer(dur).C
		} else {
			dur := time.Second * 10
			if conf.Type == apiproto.LotteryType_HighFreq {
				dur = time.Minute * 5
			}
			<-time.NewTimer(dur).C
		}
	}
}
