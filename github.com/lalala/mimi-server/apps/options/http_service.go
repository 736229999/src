package main

import (
	"log"
	"net/http"
	"fmt"
	"time"
)

func (srv *OptionsServer) ServeHTTP() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	log.Println("ServeHTTP on port ", *httpPort)
	mux := http.NewServeMux()
	mux.HandleFunc("/options/home/index", srv.HandleHomeIndex)
	mux.HandleFunc("/options/contact/info", srv.HandleContactInfo)
	mux.HandleFunc("/options/feedback/add", srv.HandleFeedbackAdd)
	mux.HandleFunc("/options/discover/banner", srv.HandleDiscoverBanner)
	mux.HandleFunc("/options/discover/news/list", srv.HandleGetNewsList)
	mux.HandleFunc("/options/discover/news/detail", srv.HandleGetNews)
	mux.HandleFunc("/options/faq/list", srv.HandleGetFaqList)
	mux.HandleFunc("/options/faq/detail", srv.HandleGetFaq)
	mux.HandleFunc("/options/agreement", srv.HandleGetAgreement)

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
