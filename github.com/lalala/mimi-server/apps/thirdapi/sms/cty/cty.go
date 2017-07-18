package cty

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/net/html/charset"

	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/utils"
)

/*
（2）测试内容必须添加签名，即在短信内容最后添加【短信签名】
注：短信内容最后必须包含签名，除签名外，内容中不能有中括号即【】或 [ ]，如内容中出现【】或[]，会导致提交运营商时被系统拦截，从而导致短信下发失败
	a、可以提前将签名报备给CTU，由CTU统一添加
	b、未提前报备，则必须在提交时，加到内容的最后
	c、签名格式为【短信签名】
	（3）单发标准模版样例：
	验证码类：您的验证码为：8888【畅天游】
   注：单个号码测试验证码短信，尽量1分钟1条，10分钟不超过5条
       若10分钟超过5条，会被运营商加入到临时黑名单，和运营商沟通后可以进行下一步解除
	通知类：恭喜您5元话费已经到账，如有疑问请联系客服电话010-62115926【畅天游】
4）群发内容之后，签名之前，一定要加“退订回T”，如不加会导致运营商下发信息失败
5）内容中有链接时，需要在链接后加上空格再加其他内容，避免用户收到短信时链接打不开
*/
const (
	// account = "ctyswse-24"
	// pwd     = "f766e3"
	account = "sckxkr-1"
	pwd     = "ce2cf6"
	address = "https://sms.800617.com:4410/sms/SendSMS.aspx"
)

type CtySms struct{}

func (cty *CtySms) Name() string {
	return "畅天游"
}

type Reply struct {
	XMLName   xml.Name `xml:"Root"`
	Result    string   `xml:"Result"`
	CtuId     string   `xml:"CtuId"`
	SendNum   string   `xml:"SendNum"`
	MobileNum string   `xml:"MobileNum"`
}

func (reply *Reply) isSuccess() bool {
	return reply.Result == "1"
}

func (cty *CtySms) send(sign string, phoneList []string, content string) (*Reply, error) {
	if sign == "" {
		return nil, fmt.Errorf("缺少签名")
	}

	if len(phoneList) > 300 {
		return nil, fmt.Errorf("群发号码个数超限")
	}

	content = fmt.Sprintf("【%s】%s", sign, content)
	data, err := utils.Utf8ToGbk([]byte(content))
	if err != nil {
		log.Println("ERROR:", err, content)
		return nil, err
	}

	mobile := strings.Join(phoneList, ",")
	msg := url.QueryEscape(string(data))
	url := fmt.Sprintf("%s?un=%s&pwd=%s&mobile=%s&msg=%s", address, account, pwd, mobile, msg)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("ERROR:", err, url)
		return nil, err
	}
	defer resp.Body.Close()
	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ERROR:", err)
		return nil, err
	}

	var reply Reply
	reader := bytes.NewReader(data)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&reply)
	if err != nil {
		log.Println("ERROR:", err)
		return nil, err
	}
	log.Printf("%+v, %v\n", reply, reply.Result)
	return &reply, nil
}

func (cty *CtySms) SendMessage(sign string, phoneList []string, content string) (*dbproto.Sms, error) {
	reply, err := cty.send(sign, phoneList, content)
	if err != nil {
		return nil, err
	}
	result := &dbproto.Sms{
		PhoneList: phoneList,
		Content:   content,
		Vendor:    cty.Name(),
		SendTime:  time.Now().Unix(),
		IsSuccess: reply.isSuccess(),
		Result:    reply.Result,
		Sign:      sign,
	}
	return result, nil
}

func (cty *CtySms) SendCode(sign string, phone string, code string, valid time.Duration) (*dbproto.Sms, error) {
	content := fmt.Sprintf("验证码是 %s ，请在%d分钟内完成操作。如非本人操作，请忽略此短信。", code, int(valid.Minutes()))
	reply, err := cty.send(sign, []string{phone}, content)
	if err != nil {
		return nil, err
	}
	result := &dbproto.Sms{
		PhoneList:  []string{phone},
		Content:    content,
		Vendor:     cty.Name(),
		SendTime:   time.Now().Unix(),
		IsSuccess:  reply.isSuccess(),
		Result:     reply.Result,
		Code:       code,
		ExpireTime: time.Now().Add(valid).Unix(),
		Sign:       sign,
	}
	return result, nil
}
