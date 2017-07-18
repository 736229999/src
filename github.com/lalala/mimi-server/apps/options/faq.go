package main

import (
	"log"
	"net/http"

	apiproto "github.com/caojunxyz/mimi-api/proto"
	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/utils"
	"golang.org/x/net/context"
	// "fmt"
	"fmt"
	"html/template"
	"strconv"
)

// HandleGetFaq 获取一条Faq
func (srv *OptionsServer) HandleGetFaq(w http.ResponseWriter, r *http.Request) {
	log.Println("Handle get a faq")
	query := r.URL.Query()
	if len(query) != 1 {
		http.Error(w, "参数错误", http.StatusBadRequest)
		return
	}
	fIDStr := query["id"][0]
	fID, err := strconv.ParseInt(fIDStr, 10, 0)
	if err != nil {
		http.Error(w, "参数错误", http.StatusBadRequest)
		return
	}
	dbArg := &dbproto.FaqId{
		Id: fID,
	}

	faq, err := srv.optionsClient.QueryFaqById(context.Background(), dbArg)
	if err != nil {
		log.Println("QueryFaqById", err)
		http.Error(w, "Get faq error", http.StatusInternalServerError)
		return
	}
	type Faq struct {
		Title string
		Html  template.HTML
	}
	t, err := template.ParseFiles("./views/faq.html")
	if err != nil {
		log.Printf("ParseFiles error %v\n", err)
		return
	}
	resFaq := Faq{
		Title: faq.GetTitle(),
		Html:  template.HTML(faq.GetHtml()),
	}
	t.Execute(w, resFaq)
}

// HandleGetFaqList 获取新闻列表
func (srv *OptionsServer) HandleGetFaqList(w http.ResponseWriter, r *http.Request) {
	log.Println("Handle Get Faq List")
	dbResList, err := srv.optionsClient.QueryFaqList(context.Background(), &dbproto.Nil{})
	if err != nil {
		log.Printf("%+v\n", err)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "获取失败", nil)
		return
	}
	resList := make([]*apiproto.Faq, 0)
	for _, val := range dbResList.GetList() {
		refaq := &apiproto.Faq{
			Id:         val.GetId(),
			Title:      val.GetTitle(),
			ContentUrl: fmt.Sprintf("/options/faq/detail?id=%d", val.GetId()),
		}
		resList = append(resList, refaq)
	}
	result := &apiproto.FaqList{
		List:  resList,
		Total: dbResList.GetTotal(),
	}
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", result)
	return
}
