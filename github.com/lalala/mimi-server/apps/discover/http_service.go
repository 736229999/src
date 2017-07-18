package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	apiproto "github.com/caojunxyz/mimi-api/proto"
	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/utils"
	"golang.org/x/net/context"
)

func (srv *DiscoverServer) ServeHTTP() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	log.Println("ServeHTTP on port ", *httpPort)
	mux := http.NewServeMux()
	mux.HandleFunc("/discover/news/test", srv.HandleTest)
	mux.HandleFunc("/discover/news/list", srv.HandleNewsList)
	mux.HandleFunc("/discover/newsdetails", srv.HandleGetANews)
	httpServer := http.Server{
		Addr:         fmt.Sprintf(":%d", *httpPort),
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	if err := httpServer.ListenAndServe(); err != nil {
		panic(err)
	}
}

// HandleTest 测试
func (srv *DiscoverServer) HandleTest(w http.ResponseWriter, r *http.Request) {
	log.Println("Handle Test")
	w.Write([]byte("I am discover server!"))
}

// HandleNewsList 获取新闻列表
func (srv *DiscoverServer) HandleNewsList(w http.ResponseWriter, r *http.Request) {
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
	res, err := srv.dbClient.QueryNewsListClient(context.Background(), dbArg)
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
			ContentUrl:  fmt.Sprintf("http://%s:%d%s?newsid=%d", serverAddr, 8088, "/discover/newsdetails", val.GetId()),
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

// HandleGetANews 获取一条新闻
func (srv *DiscoverServer) HandleGetANews(w http.ResponseWriter, r *http.Request) {
	log.Println("Handle get a news")
	// reqArg := &apiproto.NewsId{}
	// accountId, _, err := utils.ParseHttpRequest(w, r, reqArg)
	// if err != nil {
	// 	return
	// }
	query := r.URL.Query()
	if len(query) != 1 {
		http.Error(w, "参数错误", http.StatusBadRequest)
		return
	}

	nIDStr := query["newsid"][0]
	nID, err := strconv.ParseInt(nIDStr, 10, 0)
	if err != nil {
		http.Error(w, "参数错误", http.StatusBadRequest)
		return
	}
	// log.Println("accountId", accountId)
	// log.Println("reqArg", reqArg.Id)
	dbArg := &dbproto.NewsId{
		Id: nID,
	}

	// # dev mode
	// dbArg := &dbproto.NewsId{
	// 	Id: 1,
	// }

	res, err := srv.dbClient.GetANews(context.Background(), dbArg)
	if err != nil {
		log.Println("GetANews", err)
		http.Error(w, "GetANews!", http.StatusInternalServerError)
		return
	}
	if res.GetId() != 0 {
		_, err := srv.dbClient.ReadANews(context.Background(), &dbproto.NewsId{Id: res.GetId()})
		if err != nil {
			log.Println("ReadANews", err)
		}
	}
	// log.Printf("res %+v\n", res)
	log.Println("news Id is ", res.GetId())

	t := template.New("fieldname example")
	t, _ = t.Parse(res.GetHtml())
	// p := Person{UserName: "Astaxie"}
	t.Execute(w, res)

	// w.Write([]byte(res.GetHtml()))
	// utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", result)
}
