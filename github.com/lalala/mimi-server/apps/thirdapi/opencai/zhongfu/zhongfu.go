package zhongfu

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/lottery"
)

const NAME = "中福民彩"
const URL = "http://120.77.34.33/api/"
const (
	KEY     = "60D81B05"
	ACCOUNT = "xiaomi"
	SECRET  = "xiaomi"
)

//------------------------------------------------------------------------------------------------------------------------------
func formatBalls(code string, balls string) string {
	switch code {
	case lottery.Cqssc:
		list := []string{}
		for _, v := range balls {
			list = append(list, string(v))
		}
		return strings.Join(list, ",")
	case lottery.Bjpk10:
		return strings.Replace(balls, " ", ",", -1)
	case lottery.Gd11x5:
		return strings.Replace(balls, " ", ",", -1)
	}
	return balls
}

type QueryItem struct {
	Issue    string `json:"issue"`
	Opencode string `json:"opencode"`
	Opentime string `json:"opentime"`
	TryCode  string `json:"trycode"`
}

func (item *QueryItem) ToOpenInfo(t int64, source string) *dbproto.OpenInfo {
	ret := &dbproto.OpenInfo{
		Issue:      item.Issue,
		Balls:      item.Opencode,
		GrabTime:   t,
		GrabSource: source,
	}
	ot, err := time.ParseInLocation("2006-01-02 15:04", item.Opentime, time.Local)
	if err != nil {
		log.Panic(err)
		return nil
	}
	ret.OpenTime = ot.Unix()
	return ret
}

type QueryDateResult struct {
	Status string      `json:"status"`
	Desc   string      `json:"desc"`
	Result []QueryItem `json:"result"`
}

func QueryDateIssues(code string, zfId string, date string) []*dbproto.OpenInfo {
	arg := map[string]string{
		"name":       ACCOUNT,
		"secret":     SECRET,
		"lottery_id": zfId,
		"date":       date,
	}

	url := fmt.Sprintf("%sDateOpen.ashx", URL)
	data, err := ZhongfuEncrypt(arg)
	if err != nil {
		log.Println(err)
		return nil
	}
	resp, err := http.Post(url, "", bytes.NewBuffer(data))
	if err != nil {
		log.Println(err)
		return nil
	}
	defer resp.Body.Close()
	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil
	}

	var result QueryDateResult
	if err := ZhongfuDecrypt(data, &result); err != nil {
		log.Println(err)
		return nil
	}

	if result.Status != "0" {
		log.Println(result.Status, result.Desc, zfId, date)
		return nil
	}

	list := []*dbproto.OpenInfo{}
	t := time.Now().Unix()
	for _, v := range result.Result {
		info := v.ToOpenInfo(t, NAME)
		info.Balls = formatBalls(code, info.Balls)
		list = append(list, info)
	}
	return list
}

//------------------------------------------------------------------------------------------------------------------------------
