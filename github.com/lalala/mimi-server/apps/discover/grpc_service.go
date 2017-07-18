package main

// import (
// 	"fmt"
// 	dbproto "github.com/caojunxyz/mimi-server/dbagent/proto"
// 	discoverproto "github.com/caojunxyz/mimi-server/discover/proto"
// 	"golang.org/x/net/context"
// 	"google.golang.org/grpc"
// 	"log"
// 	"net"
// )

// // ServeGRPC Discover 的gprc调用接口
// func (srv *DiscoverServer) ServeGRPC() {
// 	log.Println("ServeGRPC")

// 	defer func() {
// 		if err := recover(); err != nil {
// 			log.Println("recover from panic:", err)
// 		}
// 	}()

// 	log.Println("ServeGRPC on port ", *grpcPort)
// 	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *grpcPort))
// 	if err != nil {
// 		panic(err)
// 	}
// 	grpcServer := grpc.NewServer()
// 	discoverproto.RegisterDiscoverServer(grpcServer, srv)
// 	if err := grpcServer.Serve(lis); err != nil {
// 		log.Println(err)
// 		return
// 	}
// }

// // CreateNews 新建一条新闻
// func (srv *DiscoverServer) CreateNews(ctx context.Context, arg *discoverproto.News) (*discoverproto.IntValue, error) {
// 	news := dbproto.News{
// 		Title:       arg.GetTitle(),
// 		Description: arg.GetDescription(),
// 		Content:     arg.GetContent(),
// 		Cover:       arg.GetCover(),
// 		Html:        arg.GetHtml(),
// 		Author:      arg.GetAuthor(),
// 	}
// 	res, err := srv.dbClient.CreateNews(context.Background(), &news)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &discoverproto.IntValue{Value: res.Value}, nil
// }

// // QueryNewsList 查询新闻列表
// func (srv *DiscoverServer) QueryNewsList(ctx context.Context, arg *discoverproto.QueryNewsArg) (*discoverproto.NewsList, error) {
// 	query := dbproto.QueryNewsArg{
// 		Title:    arg.GetTitle(),
// 		Author:   arg.GetAuthor(),
// 		Start:    arg.GetStart(),
// 		End:      arg.GetEnd(),
// 		Page:     arg.GetPage(),
// 		PageSize: arg.GetPageSize(),
// 	}
// 	res, err := srv.dbClient.QueryNewsList(context.Background(), &query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	newsList := make([]*discoverproto.News, 0)
// 	for _, v := range res.GetList() {
// 		news := &discoverproto.News{
// 			Id:          v.GetId(),
// 			Title:       v.GetTitle(),
// 			Author:      v.GetAuthor(),
// 			Description: v.GetDescription(),
// 			Cover:       v.GetCover(),
// 			Content:     v.GetContent(),
// 			PageViews:   v.GetPageViews(),
// 			Html:        v.GetHtml(),
// 			Created:     v.GetCreated(),
// 			Updated:     v.GetUpdated(),
// 			NewsClass:   v.GetNewsClass(),
// 			IsVisible:   v.GetIsVisible(),
// 		}
// 		newsList = append(newsList, news)
// 	}
// 	return &discoverproto.NewsList{List: newsList, Total: res.GetTotal()}, nil
// }

// // QueryBakendSelectOfNews banner管理下拉框搜索查询
// func (srv *DiscoverServer) QueryBakendSelectOfNews(ctx context.Context, arg *discoverproto.QueryNewsOfSelect) (*discoverproto.NewsList, error) {
// 	dbArg := dbproto.QueryNewsOfSelect{
// 		KeyWord: arg.GetKeyWord(),
// 	}
// 	res, err := srv.dbClient.QueryBakendSelectOfNews(context.Background(), &dbArg)
// 	if err != nil {
// 		log.Println("QueryBakendSelectOfNews", err)
// 		return nil, err
// 	}
// 	newsList := make([]*discoverproto.News, 0)
// 	for _, v := range res.GetList() {
// 		news := &discoverproto.News{
// 			Id:          v.GetId(),
// 			Title:       v.GetTitle(),
// 			Author:      v.GetAuthor(),
// 			Description: v.GetDescription(),
// 			Cover:       v.GetCover(),
// 			Created:     v.GetCreated(),
// 		}
// 		newsList = append(newsList, news)
// 	}
// 	return &discoverproto.NewsList{List: newsList, Total: res.GetTotal()}, nil
// }
