package apix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
)

const (
	key  = "49f4d0f03005452564b9fe24f591a444"
	addr = "http://v.apix.cn"
)

type apixResponse struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data json.RawMessage `json:"data"`
}

type Apix struct{}

func (_ *Apix) VerifyIdcard(idno, realname string) *dbproto.Idcard {
	typ := "idcard"
	url := fmt.Sprintf("%s/apixcredit/idcheck/idcard?type=%s&cardno=%s&name=%s", addr, typ, idno, realname)
	log.Println("url:", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err, url)
		return nil
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("apix-key", key)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err, url)
		return nil
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err, url)
		return nil
	}
	log.Println(string(data))
	var apixResp apixResponse
	if err := json.Unmarshal(data, &apixResp); err != nil {
		log.Println(err, string(data), url)
		return nil
	}

	if apixResp.Code != 0 {
		return nil
	}
	return &dbproto.Idcard{Cardno: idno, Realname: realname}
}

func (_ *Apix) VerifyBankcard4(idno, realname, bankcardNo, phone string) *dbproto.Bankcard {
	typ := "bankcard_four"
	args := fmt.Sprintf("type=%s&bankcardno=%s&name=%s&idcardno=%s&phone=%s", typ, bankcardNo, realname, idno, phone)
	url := fmt.Sprintf("%s/apixcredit/idcheck/bankcard?%s", addr, args)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err, url)
		return nil
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("apix-key", key)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err, url)
		return nil
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err, url)
		return nil
	}
	log.Println(string(data))
	var apixResp apixResponse
	if err := json.Unmarshal(data, &apixResp); err != nil {
		log.Println(err, string(data), url)
		return nil
	}
	if apixResp.Code != 0 {
		log.Println(url, string(data))
		return nil
	}

	result := make(map[string]string)
	if err := json.Unmarshal(apixResp.Data, &result); err != nil {
		log.Println(err, string(data), url)
		return nil
	}

	return &dbproto.Bankcard{
		IdcardNo: idno, Realname: realname, BankcardNo: bankcardNo,
		Phone: phone, CardType: result["cardtype"], Bankname: result["bankname"],
	}
}
