package discover

// import (
// 	"log"

// 	"net/http"

// 	"github.com/caojunxyz/mimi-admin/backend/core"
// 	discoverproto "github.com/caojunxyz/mimi-server/discover/proto"
// 	"github.com/gin-gonic/gin"
// 	"golang.org/x/net/context"
// 	"strconv"
// 	"google.golang.org/grpc"
// )

// const DiscoverAdd string = "localhost:7014"

// // DiscoverService 新闻相关服务
// type DiscoverService struct {
// 	core.Service
// 	discoverClient             discoverproto.DiscoverClient
// }

// func NewDiscoverService() *DiscoverService {
// 	return &DiscoverService{}
// }

// // ConnectDiscover 连接到发现服务
// func (srv *DiscoverService)ConnectDiscover()  discoverproto.DiscoverClient {
// 	conn, err := grpc.Dial(DiscoverAdd, grpc.WithInsecure())
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	srv.discoverClient = discoverproto.NewDiscoverClient(conn)
// 	return srv.discoverClient
// }

// func (srv *DiscoverService) Discover() discoverproto.DiscoverClient {

// 	if srv.discoverClient == nil {
// 		srv.discoverClient = srv.ConnectDiscover()
// 	}
// 	return srv.discoverClient
// }

// // HandleNewsAdd 用于添加一条新闻
// func (srv *DiscoverService) HandleNewsAdd(c *gin.Context) {
// 	log.Println("Handle News Add")
// 	msg := discoverproto.News{}
// 	err := c.BindJSON(&msg)
// 	if err != nil {
// 		log.Println("BindJSON", err)
// 		srv.Json("数据格式错误", http.StatusBadRequest, c)
// 		return
// 	}
// 	// defer srv.Log(c, msg, "添加新闻", core.ADD_OPERATION)
// 	if len(msg.Title) < 1 {
// 		srv.Json("标题不能为空", http.StatusForbidden, c)
// 		return
// 	}
// 	if len(msg.Content) < 1 {
// 		srv.Json("内容不能为空", http.StatusForbidden, c)
// 		return
// 	}
// 	if len(msg.Description) < 1 {
// 		srv.Json("封面描述不能为空", http.StatusForbidden, c)
// 		return
// 	}
// 	log.Printf("msg is %+v\n", msg)
// 	rid, err := srv.Discover().CreateNews(context.Background(), &msg)
// 	if err != nil {
// 		log.Println("InsertNews ", err)
// 		srv.Json(err, http.StatusInternalServerError, c)
// 		return
// 	}
// 	log.Printf("Id is %v\n", rid)
// 	srv.Json("", http.StatusOK, c)
// 	return
// }

// // HandleNewsListGet 获取新闻列表
// func (srv *DiscoverService) HandleNewsListGet(c *gin.Context) {
// 	log.Println("Handle NewsList Get")
// 	srv.ConnectDiscover()
// 	page := c.DefaultQuery("page", "1")
// 	pageSize := c.DefaultQuery("pagesize", "10")
// 	start := c.DefaultQuery("start", "0")
// 	end := c.DefaultQuery("end", "0")
// 	log.Println(page, pageSize, start, end)
// 	startDb, err := strconv.ParseInt(start, 10, 0)
// 	if err != nil {
// 		log.Println("error", err)
// 		srv.Json("开始时间参数错误", http.StatusForbidden, c)
// 		return
// 	}
// 	endDb, err := strconv.ParseInt(end, 10, 0)
// 	if err != nil {
// 		log.Println("error", err)
// 		srv.Json("结束时间参数错误", http.StatusForbidden, c)
// 		return
// 	}
// 	pageDb, err := strconv.ParseInt(page, 10, 0)
// 	if err != nil {
// 		log.Println("error", err)
// 		srv.Json("页码参数错误", http.StatusForbidden, c)
// 		return
// 	}
// 	pageSizeDb, err := strconv.ParseInt(pageSize, 10, 0)
// 	if err != nil {
// 		log.Println("error", err)
// 		srv.Json("每页尺寸参数错误", http.StatusForbidden, c)
// 		return
// 	}
// 	msg := discoverproto.QueryNewsArg{
// 		Title: c.Query("title"),
// 		Author: c.Query("author"),
// 		Start: startDb,
// 		End: endDb,
// 		Page: pageDb,
// 		PageSize: pageSizeDb,
// 	}
// 	log.Printf("msg: %+v", msg)
// 	log.Printf("msg title: %+v", c.Query("title"))
// 	res, err := srv.Discover().QueryNewsList(context.Background(), &msg)
// 	if err != nil {
// 		log.Println("QueryNewsList ", err)
// 		srv.Json("服务器异常", http.StatusInternalServerError, c)
// 		return
// 	}
// 	log.Printf("res %+v\n", res)
// 	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 	c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
// 	c.JSON(http.StatusOK, res)
// 	// srv.Json(res, http.StatusOK, c)
// }

// // QueryNewsOfSelect banner管理中下拉框搜索接口
// func (srv *DiscoverService)QueryNewsOfSelect(c *gin.Context)  {
// 	keyword := c.DefaultQuery("keyword", "")
// 	log.Println("keyword", keyword)
// 	msg := discoverproto.QueryNewsOfSelect{
// 		KeyWord: keyword,
// 	}
// 	res, err := srv.Discover().QueryBakendSelectOfNews(context.Background(), &msg)
// 	if err != nil {
// 		log.Println("QueryBakendSelectOfNews ", err)
// 		srv.Json("服务器异常", http.StatusInternalServerError, c)
// 		return
// 	}
// 	log.Printf("res %+v\n", res)
// 	srv.Json(res, http.StatusOK, c)
// }
