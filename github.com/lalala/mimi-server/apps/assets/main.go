package main

import (
	"crypto/md5"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/caojunxyz/mimi-server/auth"
	"github.com/caojunxyz/mimi-server/utils"
)

var httpPort = flag.Int("http", 7009, "http port")

const ROOT_PATH = "/data/cp.kxkr.com/assets"

var IMAGE_PATH = path.Join(ROOT_PATH, "image")
var PAGES_PATH = path.Join(ROOT_PATH, "pages")
var HEADICON_PATH = path.Join(IMAGE_PATH, "headicon")

var NEWSIMAGE_PATH = path.Join(IMAGE_PATH, "news")
var TEAMICON_PATH = path.Join(IMAGE_PATH, "teamicon")

var serverAddr string

func StripPrefix(prefix string, h http.Handler) http.Handler {
	if prefix == "" {
		return h
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if p := strings.TrimPrefix(r.URL.Path, prefix); len(p) < len(r.URL.Path) {
			r.URL.Path = strings.ToLower(p)
			h.ServeHTTP(w, r)
		} else {
			http.NotFound(w, r)
		}
	})
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	fmt.Println("assets server run port:", *httpPort)
	flag.Parse()
	InitServerAddr()
	InitStorage()

	http.HandleFunc("/assets/upload/", auth.Validate(HandleUpload))
	http.HandleFunc("/assets/backend/upload/", HandleBackendUpload)
	http.HandleFunc("/assets/test", HandleTest)

	http.Handle("/assets/download/headicon/", http.StripPrefix("/assets/download/headicon/", http.FileServer(http.Dir(HEADICON_PATH))))

	http.Handle("/assets/playguid/pages/", StripPrefix("/assets/playguid/pages/", http.FileServer(http.Dir(PAGES_PATH))))

	http.Handle("/assets/download/news/", StripPrefix("/assets/download/news/", http.FileServer(http.Dir(NEWSIMAGE_PATH))))

	http.Handle("/assets/download/teamicon/", StripPrefix("/assets/download/teamicon/", http.FileServer(http.Dir(TEAMICON_PATH))))

	log.Panic(http.ListenAndServe(fmt.Sprintf(":%d", *httpPort), nil))
}

func InitServerAddr() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()
	serverAddr = strings.Split(conn.LocalAddr().String(), ":")[0]
	log.Println("本机地址:", serverAddr)
}

// 初始化存储空间
func InitStorage() {
	if !utils.IsDirExists(HEADICON_PATH) {
		if err := os.MkdirAll(HEADICON_PATH, os.ModePerm); err != nil {
			log.Panicf("path: %s, error: %v", HEADICON_PATH, err)
		}
	}
	if !utils.IsDirExists(PAGES_PATH) {
		if err := os.MkdirAll(PAGES_PATH, os.ModePerm); err != nil {
			log.Panicf("path: %s, error: %v", PAGES_PATH, err)
		}
	}
	if !utils.IsDirExists(NEWSIMAGE_PATH) {
		if err := os.MkdirAll(NEWSIMAGE_PATH, os.ModePerm); err != nil {
			log.Panicf("path: %s, error: %v", NEWSIMAGE_PATH, err)
		}
	}
	if !utils.IsDirExists(TEAMICON_PATH) {
		if err := os.MkdirAll(TEAMICON_PATH, os.ModePerm); err != nil {
			log.Panicf("path: %s, error: %v", TEAMICON_PATH, err)
		}
	}
}

const (
	SUCCESS = 0
	FAIL    = -1
)

type Response struct {
	Code   int    `json:"code"`
	Desc   string `json:"desc"`
	Result string `json:"result"`
}

func HandleTest(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleTest")
	resp := Response{Code: SUCCESS, Desc: "成功", Result: "资源服务器"}
	data, _ := json.Marshal(resp)
	w.Write(data)
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	typ := path.Base(r.URL.Path)
	log.Println("type:", typ)
	switch typ {
	case "headicon":
		uploadHeadIcon(w, r)
	default:
		http.Error(w, "", http.StatusNotFound)
	}
}

func HandleBackendUpload(w http.ResponseWriter, r *http.Request) {
	typ := path.Base(r.URL.Path)
	log.Println("type:", typ)
	switch typ {
	case "news":
		HandleNewsUpload(w, r)
	case "teamicon":
		HandleTeamIconUpload(w, r)
	default:
		http.Error(w, "", http.StatusNotFound)
	}
}

// HandleBackendUpload 后台新闻图片上传
func HandleNewsUpload(w http.ResponseWriter, r *http.Request) {
	HandlePicture(w, r, NEWSIMAGE_PATH)
}

// 后台球队队徽图片上传
func HandleTeamIconUpload(w http.ResponseWriter, r *http.Request) {
	HandlePicture(w, r, TEAMICON_PATH)
}

func HandlePicture(w http.ResponseWriter, r *http.Request, fielPath string) {
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	log.Println("Method:", r.Method)
	log.Println(r.Header.Get("Content-Type"))

	r.ParseMultipartForm(32 << 20)
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Println("error", err)
		return
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		log.Println("error", err)
		return
	}
	hashByte := hash.Sum(nil)
	HashStr := fmt.Sprintf("%x", hashByte)
	log.Printf("HashStr: %v", HashStr)
	res, err := file.Seek(0, io.SeekStart)
	if err != nil {
		log.Println("error", err)
		return
	}
	log.Println("res", res)
	f, err := os.OpenFile(fielPath+"/"+HashStr, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	if _, err := io.Copy(f, file); err != nil {
		log.Println("error", err)
	}

	typ := path.Base(fielPath)

	uri := fmt.Sprintf("/assets/download/%s/%s", typ, HashStr)
	resp := Response{Code: SUCCESS, Desc: "成功", Result: uri}
	data, _ := json.Marshal(resp)
	w.Write(data)
}

// 处理上传头像图标
func uploadHeadIcon(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	accountId, _ := ctx.Value("accountId").(int64)
	log.Println("accountId:", accountId)
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, "read body", http.StatusInternalServerError)
		return
	}

	hash := fmt.Sprintf("%x", md5.Sum(data))
	log.Println("size:", len(data), "md5: ", hash)
	name := fmt.Sprintf("%d-%s", accountId, hash)

	filePath := path.Join(HEADICON_PATH, name)
	file, err := os.Create(filePath)
	if err != nil {
		log.Println(err)
		http.Error(w, "read body", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	n, err := file.Write(data)
	if err != nil || len(data) != n {
		log.Println(err, len(data), n)
		http.Error(w, "write file", http.StatusInternalServerError)
		return
	}
	// TODO: 检查图片大小和分辨率
	// TODO: 头像多次修改做覆盖处理
	uri := fmt.Sprintf("/assets/download/headicon/%s", name)
	resp := Response{Code: SUCCESS, Desc: "成功", Result: uri}
	data, _ = json.Marshal(resp)
	w.Write(data)
}
