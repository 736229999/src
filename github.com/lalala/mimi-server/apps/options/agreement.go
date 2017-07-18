package main

import (
	"log"
	"net/http"
	"html/template"
)


// HandleGetAgreement 获取用户协议
func (srv *OptionsServer) HandleGetAgreement(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./views/serviceAgreement.html")
	if err != nil {
		log.Printf("ParseFiles error %v\n", err)
		return
	}
	t.Execute(w, nil)
}