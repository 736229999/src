package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	apiproto "github.com/caojunxyz/mimi-api/proto"
	"github.com/caojunxyz/mimi-server/utils"
)

var port = flag.Int("http", 7005, "http port")
var db = flag.String("db", "localhost:6013", "thidapi dbagent address")

func HandleLatest(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleLatest:", r.URL.Path)
	result := &apiproto.LatestOpen{
		List: getLatestOpen(),
		Tabs: []*apiproto.LotteryCollection{
			&apiproto.LotteryCollection{Name: "低频彩", Type: apiproto.LotteryType_LowFreq},
			&apiproto.LotteryCollection{Name: "高频彩", Type: apiproto.LotteryType_HighFreq},
		},
	}
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", result)
}

func HandleHistory(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleHistory:", r.URL.Path)
	id, err := utils.ParseLotteryIdArg(r)
	if err != nil {
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "无效参数", nil)
		return
	}
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", getHistory(id))
}

func HandleDigestHistory(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleDigestHistory:", r.URL.Path)
	id, err := utils.ParseLotteryIdArg(r)
	if err != nil {
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "无效参数", nil)
		return
	}
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", getDigestHistory(id))
}

func HandleTest(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleTest")
	w.Write([]byte("I am opencai server!"))
}

func main() {
	flag.Parse()
	fmt.Println("opencai...")
	fmt.Println("db: ", *db)
	fmt.Println("port: ", *port)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

	connectDb(*db)

	mux := http.NewServeMux()
	mux.HandleFunc("/opencai/latest", HandleLatest)
	mux.HandleFunc("/opencai/history/", HandleHistory)
	mux.HandleFunc("/opencai/dighistory/", HandleDigestHistory)
	mux.HandleFunc("/opencai/test", HandleTest)

	httpServer := http.Server{
		Addr:         fmt.Sprintf(":%d", *port),
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	if err := httpServer.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}
