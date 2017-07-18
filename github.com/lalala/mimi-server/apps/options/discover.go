package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	apiproto "github.com/caojunxyz/mimi-api/proto"
	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/utils"
	"golang.org/x/net/context"
)

// HandleDiscoverBanner 获取发现模块的轮播图
func (srv *OptionsServer) HandleDiscoverBanner(w http.ResponseWriter, r *http.Request) {
	log.Println("Handle Discover Banner")

	discBnner, err := srv.optionsClient.QueryBannerList(context.Background(), &dbproto.QueryBannerArg{Location: dbproto.Banner_Location_Discover})
	if err != nil {
		log.Printf("%+v\n", err)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "获取失败", nil)
		return
	}
	log.Printf("discBanner is %v\n", discBnner)
	reList := make([]*apiproto.Banner, 0)
	for _, val := range discBnner.GetList() {
		reBanner := &apiproto.Banner{
			Id:          val.GetId(),
			Url:         val.GetUrl(),
			Description: val.GetDescription(),
			TargetLink:  val.GetTargetLink(),
			TargetId:    val.GetTargetId(),
			TargetType:  apiproto.Banner_TargetType(val.GetTargetType()),
			Location:    apiproto.Banner_Location(val.GetLocation()),
		}
		reList = append(reList, reBanner)
	}
	result := &apiproto.BannerList{
		List:  reList,
		Total: discBnner.GetTotal(),
	}
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", result)
	return
}

// HandleGetNewsList 获取新闻列表
func (srv *OptionsServer) HandleGetNewsList(w http.ResponseWriter, r *http.Request) {
	log.Println("Handle News List")
	reqArg := &apiproto.QueryNewsArg{}
	accountId, _, err := utils.ParseHttpRequest(w, r, reqArg)
	if err != nil {
		return
	}
	log.Println("reqArg page", reqArg.GetPage())
	log.Println("reqArg pagesize", reqArg.GetPageSize())
	log.Println("accountId", accountId)
	dbArg := &dbproto.QueryNewsArg{
		Page:     reqArg.GetPage(),
		PageSize: reqArg.GetPageSize(),
	}
	res, err := srv.optionsClient.QueryNewsList(context.Background(), dbArg)
	if err != nil {
		http.Error(w, "QueryNewsListClient!", http.StatusInternalServerError)
		return
	}
	// log.Printf("res %+v", res)
	reList := make([]*apiproto.News, 0)
	for _, val := range res.GetList() {
		reNews := &apiproto.News{
			Id:          val.GetId(),
			Title:       val.GetTitle(), // id, title, description, pageviews, created,
			Cover:       val.GetCover(),
			Description: val.GetDescription(),
			PageViews:   val.GetPageViews(),
			Created:     val.GetCreated(),
			ContentUrl:  fmt.Sprintf("/options/discover/news/detail?id=%d", val.GetId()),
		}
		reList = append(reList, reNews)
	}
	result := &apiproto.NewsList{
		List:  reList,
		Total: res.GetTotal(),
	}
	log.Printf("\nresult total is %v\n", len(result.List))
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", result)
	return
}

// HandleGetNews 获取一条新闻
func (srv *OptionsServer) HandleGetNews(w http.ResponseWriter, r *http.Request) {
	log.Println("Handle get a news")
	query := r.URL.Query()
	if len(query) != 1 {
		http.Error(w, "参数错误", http.StatusBadRequest)
		return
	}
	nIDStr := query["id"][0]
	nID, err := strconv.ParseInt(nIDStr, 10, 0)
	if err != nil {
		http.Error(w, "参数错误", http.StatusBadRequest)
		return
	}
	dbArg := &dbproto.NewsId{
		Id: nID,
	}

	res, err := srv.optionsClient.QueryNewsById(context.Background(), dbArg)
	if err != nil {
		log.Println("GetANews", err)
		http.Error(w, "Get news error", http.StatusInternalServerError)
		return
	}
	if res.GetId() != 0 {
		_, err := srv.optionsClient.ReadANews(context.Background(), &dbproto.NewsId{Id: res.GetId()})
		if err != nil {
			log.Println("ReadANews", err)
		}
	}
	log.Printf("res %+v\n", res)
	log.Println("news Id is ", res.GetId())

	t := template.New("fieldname example")
	t, _ = t.Parse(res.GetHtml())
	t.Execute(w, res)
}
