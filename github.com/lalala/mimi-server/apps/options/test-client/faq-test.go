package main

import (
	"bytes"
	"flag"
	apiproto "github.com/caojunxyz/mimi-api/proto"
	pb "github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	"net/http"
)

var reqType = flag.String("type", "L", "req Type")

func main() {
	flag.Parse()
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("reqTYPE", *reqType)
	switch *reqType {
	case "L":
		getFaqList()
	case "D":
		// getANews()
	case "B":
		// getDiscoverBannerList()
	}

}

func getFaqList() {
	client := &http.Client{}

	msg := &apiproto.Nil{}
	msgb, err := pb.Marshal(msg)
	if err != nil {
		log.Println("pb.Marshal error", err)
		return
	}
	reqBody := bytes.NewBuffer(msgb)
	req, err := http.NewRequest("POST", "http://cp.kxkr.com:8088/options/faq/list", reqBody)
	if err != nil {
		log.Println("NewRequest", err)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("client.Do", err)
		return
	}
	defer resp.Body.Close()
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ioutil.ReadAll", err)
		return
	}
	response := apiproto.Response{}
	err = pb.Unmarshal(resBody, &response)
	if err != nil {
		log.Println(" pb.Unmarshal", err)
		return
	}

	result := apiproto.FaqList{}
	err = pb.Unmarshal(response.Result, &result)
	if err != nil {
		log.Println(" pb.Unmarshal", err)
		return
	}

	// log.Printf("result is %+v\n", result)
	log.Printf("result list is \n%+v\n", result.GetList())
	log.Printf("\nresult total is %+v\n", result.GetTotal())
}