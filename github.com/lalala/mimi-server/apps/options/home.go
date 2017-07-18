package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	apiproto "github.com/caojunxyz/mimi-api/proto"
	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/utils"
	"golang.org/x/net/context"
)

//获取app首页的banner，最新中奖， 彩票配置.
func (srv *OptionsServer) HandleHomeIndex(w http.ResponseWriter, r *http.Request) {

	homeParams := &apiproto.HomeParams{}

	//获取彩票配置.
	optionsList, err := srv.optionsClient.QueryLotteryOptionsList(context.Background(), &dbproto.Nil{})
	if err != nil {
		log.Printf("%+v\n", err)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "获取失败", nil)
		return
	}

	options := make(map[int32]*apiproto.LotteryOptions)
	for _, v := range optionsList.GetLottery() {
		option := &apiproto.LotteryOptions{
			Id:          v.GetId(),
			LotteryName: v.GetLotteryName(),
			IsPlusAward: v.GetIsPlusAward(),
			Info:        v.GetInfo(),
			StopSale:    v.GetStopSale(),
		}
		options[int32(v.GetId())] = option
	}
	homeParams.LotteryOptions = options

	data := []string{
		"热烈庆祝本平台合作投注站已达10家",
	}
	homeParams.WinList = data

	bannerList, err := srv.optionsClient.QueryClientBannerList(context.Background(), &dbproto.QueryClientBannerArg{Location: dbproto.QueryClientBannerArg_Location_Home})
	if err != nil {
		log.Printf("%+v\n", err)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "获取失败", nil)
		return
	}

	for _, v := range bannerList.GetList() {
		banner := &apiproto.Banner{
			Id:          v.GetId(),
			TargetLink:  v.GetTargetLink(),
			Description: v.GetDescription(),
			TargetType:  apiproto.Banner_TargetType(v.GetTargetType()),
			Url:         v.GetUrl(),
		}

		homeParams.Banner = append(homeParams.Banner, banner)
	}

	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", homeParams)
}

//获取客服的联系方式.
func (srv *OptionsServer) HandleContactInfo(w http.ResponseWriter, r *http.Request) {

	contact, err := srv.optionsClient.QueryContact(context.Background(), &dbproto.Nil{})
	if err != nil {
		log.Printf("%+v\n", err)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "获取失败", nil)
		return
	}

	param := &apiproto.Contact{
		Qq:       contact.GetQq(),
		Wechat:   contact.GetWechat(),
		Email:    contact.GetEmail(),
		Telphone: contact.GetTelphone(),
	}

	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", param)
}

//反馈.
func (srv *OptionsServer) HandleFeedbackAdd(w http.ResponseWriter, r *http.Request) {

	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization,DNT,User-Agent,Keep-Alive,Content-Type,accept,origin,X-Requested-With")
		w.WriteHeader(http.StatusOK)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	msg := &dbproto.Feedback{}
	if err = json.Unmarshal(b, msg); err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
	_, err = srv.optionsClient.InsertFeedback(context.Background(), msg)
	if err != nil {
		log.Printf("%+v\n", err)
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("添加ok"))
}
