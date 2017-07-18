package main

import (
	"flag"
	"net/http"
	"fmt"
	"time"
	"html/template"
	"log"
)

var httpPort = flag.Int("http", 7088, "http port")

func main(){
	fmt.Println("Listening and serving HTTP on :", *httpPort)
	ServerStart()
}

func ServerStart() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", pailie3)

	httpServer := http.Server{
		Addr:         fmt.Sprintf(":%d", *httpPort),
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	if err := httpServer.ListenAndServe(); err != nil {
		log.Println(err)
	}

}


func pailie3(w http.ResponseWriter, r *http.Request) {
	// filename,_:=parseLotteryIdArg(r)
	// fmt.Println(filename)
	//t, _ := template.ParseFiles("playguid/"+filename.String()+".html")
	t, _ := template.ParseFiles("serviceAgreement.html")
	fmt.Println(t.Execute(w, nil))
}

// func parseLotteryIdArg(r *http.Request) (apiproto.LotteryId, error) {
// 	arg := path.Base(r.URL.RawQuery)
// 	n, err := strconv.Atoi(arg)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return apiproto.LotteryId(n), nil
// }