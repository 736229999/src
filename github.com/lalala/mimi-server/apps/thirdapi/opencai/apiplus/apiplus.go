package apiplus

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/lottery"
)

// 开彩网付费接口
const APIPLUS_ACCOUNT = "70C9164176227F38"
const APIPLUS_PASSWORD = "5116CBADB564"
const NAME = "开彩网"

const SERVER = "http://b.apiplus.net/"

// const SERVER = "http://101.37.98.235/"

const MIN_NEWLY_DUR = time.Second * 3
const MIN_DAILY_DUR = time.Second * 11

type Apiplus struct {
	mtx           sync.Mutex
	lastNewlyTime time.Time
	lastDailyTime time.Time
}

//------------------------------------------------------------------------------------------------------------------------------
func formatIssue(code string, issue string) string {
	var ret string
	switch code {
	case lottery.Cqssc:
		ret = fmt.Sprintf("%s-%s", issue[:8], issue[8:])
	case lottery.Gd11x5:
		ret = fmt.Sprintf("%s%s", issue[2:8], issue[8:])
	default:
		ret = issue
	}
	return ret
}

type queryItem struct {
	Expect        string `json:"expect"`
	Opencode      string `json:"opencode"`
	Opentime      string `json:"opentime"`
	OpentimeStamp int64  `json:"opentimestamp"`
}

func (item *queryItem) toOpenInfo(t int64, source string) *dbproto.OpenInfo {
	ret := &dbproto.OpenInfo{
		Issue:      item.Expect,
		Balls:      item.Opencode,
		GrabTime:   t,
		GrabSource: source,
	}
	ot, err := time.ParseInLocation("2006-01-02 15:04:05", item.Opentime, time.Local)
	if err != nil {
		log.Panic(err)
	}
	ret.OpenTime = ot.Unix()
	return ret
}

type queryDateResult struct {
	Rows   int         `json:"rows"`
	Code   string      `json:"code"`
	Remain string      `json:"remain"`
	Data   []queryItem `json:"data"`
}

func (p *Apiplus) QueryDateIssues(code, date string) []*dbproto.OpenInfo {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	dur := time.Now().Sub(p.lastDailyTime)
	if dur < MIN_DAILY_DUR {
		<-time.NewTimer(MIN_DAILY_DUR - dur).C
	}
	p.lastDailyTime = time.Now()

	url := fmt.Sprintf("%sdaily.do?token=%s&code=%s&date=%s&format=json", SERVER, APIPLUS_ACCOUNT, code, date)
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil
	}
	var result queryDateResult
	if err := json.Unmarshal(data, &result); err != nil {
		log.Printf("code: %s, date: %s, url: %s\nerror: %v\ndata: %s\n", code, date, url, err, string(data))
		return nil
	}
	list := []*dbproto.OpenInfo{}
	t := time.Now().Unix()
	for _, v := range result.Data {
		info := v.toOpenInfo(t, NAME)
		info.Issue = formatIssue(code, info.Issue)
		list = append(list, info)
	}
	return list
}

func (p *Apiplus) QueryLatest(code string) *dbproto.OpenInfo {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	dur := time.Now().Sub(p.lastNewlyTime)
	if dur < MIN_NEWLY_DUR {
		<-time.NewTimer(MIN_NEWLY_DUR - dur).C
	}
	p.lastNewlyTime = time.Now()

	url := fmt.Sprintf("%sdaily.do?token=%s&code=%s&rows=1&format=json", SERVER, APIPLUS_ACCOUNT, code)
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil
	}
	var result queryDateResult
	if err := json.Unmarshal(data, &result); err != nil {
		log.Printf("code: %s, url: %s\nerror: %v\ndata: %s\n", code, url, err, string(data))
		return nil
	}
	if len(result.Data) == 1 {
		t := time.Now().Unix()
		item := result.Data[0]
		info := item.toOpenInfo(t, NAME)
		info.Issue = formatIssue(code, info.Issue)
		return info
	}
	return nil
}

type queryLatestNextResult struct {
	Rows   int         `json:"rows"`
	Code   string      `json:"code"`
	Remain string      `json:"remain"`
	Next   []queryItem `json:"next"`
	Open   []queryItem `json:"open"`
	Time   string      `json:"time"`
}

// 返回列表第一个元素为下一期信息
func (p *Apiplus) QueryLatestNext(code string, rows int) []*dbproto.OpenInfo {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	dur := time.Now().Sub(p.lastNewlyTime)
	if dur < MIN_NEWLY_DUR {
		<-time.NewTimer(MIN_NEWLY_DUR - dur).C
	}
	p.lastNewlyTime = time.Now()

	url := fmt.Sprintf("%snewly.do?token=%s&code=%s&rows=%d&format=json&extend=true", SERVER, APIPLUS_ACCOUNT, code, rows)
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil
	}
	var result queryLatestNextResult
	if err := json.Unmarshal(data, &result); err != nil {
		log.Println(err, string(data))
		return nil
	}

	if len(result.Next) != 1 || len(result.Open) == 0 {
		log.Panicf("数据错误:%s\n", string(data))
		return nil
	}

	// log.Printf("%+v\n", result)
	t := time.Now().Unix()
	next := result.Next[0].toOpenInfo(t, NAME)
	next.Issue = formatIssue(code, next.Issue)
	list := []*dbproto.OpenInfo{next}
	for _, v := range result.Open {
		info := v.toOpenInfo(t, NAME)
		info.Issue = formatIssue(code, info.Issue)
		list = append(list, info)
	}
	// log.Printf("%+v\n", list)
	return list
}
