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
		getNewsList()
	case "D":
		getANews()
	case "B":
		getDiscoverBannerList()
	case "F":
		getFootBall()
	}

}

func getNewsList() {
	client := &http.Client{}

	msg := &apiproto.QueryNewsArg{
		Page:     1,
		PageSize: 10,
	}
	msgb, err := pb.Marshal(msg)
	if err != nil {
		log.Println("pb.Marshal error", err)
		return
	}
	reqBody := bytes.NewBuffer(msgb)
	req, err := http.NewRequest("POST", "http://127.0.0.1:8088/options/discover/news/list", reqBody)
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

	result := apiproto.NewsList{}
	err = pb.Unmarshal(response.Result, &result)
	if err != nil {
		log.Println(" pb.Unmarshal", err)
		return
	}

	log.Printf("result is %+v\n", result)
	log.Printf("result list is %+v\n", result.GetList())
	log.Printf("result total is %+v\n", result.GetTotal())
}

func getANews() {
	client := &http.Client{}

	// msg := &apiproto.NewsId{
	// 	Id: 1,
	// }
	// msgb, err := pb.Marshal(msg)
	// if err != nil {
	// 	log.Println("pb.Marshal error", err)
	// 	return
	// }
	// reqBody := bytes.NewBuffer(msgb)
	// fmt.Printf("reqBody", reqBody)
	req, err := http.NewRequest("GET", "http://192.168.10.184:8088/options/discover/news/detail?id=1", nil)
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
	// if err != nil {
	// 	log.Println("ioutil.ReadAll", err)
	// 	return
	// }
	// response := apiproto.Response{}
	// err = pb.Unmarshal(resBody, &response)
	// if err != nil {
	// 	log.Println(" pb.Unmarshal", err)
	// 	return
	// }

	// result := apiproto.News{}
	// err = pb.Unmarshal(response.Result, &result)
	// if err != nil {
	// 	log.Println(" pb.Unmarshal", err)
	// 	return
	// }

	log.Printf("\nresult is %+v\n", string(resBody))
	// log.Printf("result is Title %+v\n", result.GetTitle())
	// log.Printf("result Author is %+v\n", result.GetAuthor())
}

func getFootBall() {

	client := &http.Client{}

	reqBody := bytes.NewBuffer(nil)
	req, err := http.NewRequest("POST", "http://192.168.10.133:7020/football/games", reqBody)
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

	result := apiproto.QueryGameInfoRes{}
	err = pb.Unmarshal(response.Result, &result)
	if err != nil {
		log.Println(" pb.Unmarshal", err)
		return
	}

	log.Printf("result is %+v\n", result)
	for _, v := range result.GetList() {
		log.Println(v)
	}
}

func getDiscoverBannerList() {
	client := &http.Client{}

	// msg := &apiproto.QueryNewsArg{
	// 	Page:     1,
	// 	PageSize: 10,
	// }
	// msgb, err := pb.Marshal(msg)
	// if err != nil {
	// 	log.Println("pb.Marshal error", err)
	// 	return
	// }
	// reqBody := bytes.NewBuffer(msgb)
	req, err := http.NewRequest("GET", "http://cp.kxkr.com:8088/options/discover/banner", nil)
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

	result := apiproto.BannerList{}
	err = pb.Unmarshal(response.Result, &result)
	if err != nil {
		log.Println(" pb.Unmarshal", err)
		return
	}

	log.Printf("result is %+v\n", result)
	log.Printf("result list is %+v\n", result.GetList())
	log.Printf("result total is %+v\n", result.GetTotal())
}
