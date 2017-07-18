package options

import (
	"fmt"
	"log"

	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"golang.org/x/net/context"
)

// GetNewsById 获取一条新闻。
func (agt *DbOptionsAgent) QueryNewsById(ctx context.Context, arg *dbproto.NewsId) (*dbproto.News, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	esql := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", TABLE_NEWS)

	news := &dbproto.News{}
	err := agt.dbConn.QueryRow(esql, arg.Id).Scan(
		&news.Id,
		&news.Content,
		&news.Title,
		&news.Description,
		&news.PageViews,
		&news.Author,
		&news.Updated,
		&news.Html,
		&news.Cover,
		&news.IsVisible,
		&news.NewsClass,
		&news.Created,
	)
	if err != nil {
		log.Println("QueryRow", err)
		return nil, err
	}

	return news, nil
}

// QueryNewsList 获取新闻列表
func (agt *DbOptionsAgent) QueryNewsList(ctx context.Context, arg *dbproto.QueryNewsArg) (*dbproto.NewsList, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	var (
		start, total              int64
		pager, order, esql, where string
	)

	if arg.GetPage() > 1 {
		start = (arg.GetPage() - 1) * arg.PageSize
	} else {
		start = 0
	}
	pager = fmt.Sprintf(" LIMIT %d OFFSET %d", arg.PageSize, start)
	order = " ORDER BY created DESC"
	esql = fmt.Sprintf("SELECT id, title, description, pageviews, created, cover FROM %s", TABLE_NEWS)
	where = " WHERE is_visible = true"
	log.Println(start, pager, order, total)
	rows, err := agt.dbConn.Query(esql + where + order + pager)
	if err != nil {
		log.Println(err, esql)
		return nil, err
	}
	newsList := make([]*dbproto.News, 0)
	for rows.Next() {
		news := &dbproto.News{}
		err := rows.Scan(&news.Id,
			&news.Title,
			&news.Description,
			&news.PageViews,
			&news.Created,
			&news.Cover)
		if err != nil {
			log.Println(err, esql)
			return nil, err
		}
		newsList = append(newsList, news)
	}
	totalEsql := fmt.Sprintf("SELECT count(*) FROM %s", TABLE_NEWS)
	row := agt.dbConn.QueryRow(totalEsql + where)
	row.Scan(&total)
	log.Println("esql: ", esql+where+order+pager)
	// esql := fmt.Sprintf("")
	log.Printf("total is %v\n", total)

	return &dbproto.NewsList{
		List:  newsList,
		Total: total,
	}, nil
}

// ReadANews 阅读一条新闻，新闻的阅读量+1
func (agt *DbOptionsAgent) ReadANews(ctx context.Context, arg *dbproto.NewsId) (*dbproto.Nil, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	stmt, err := agt.dbConn.Prepare("update news set pageviews = pageviews + 1 where id=$1")
	if err != nil {
		log.Println("Prepare", err)
		return nil, err
	}
	res, err := stmt.Exec(arg.GetId())
	if err != nil {
		log.Println("stmt.Exec()", err)
		return nil, err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		log.Println("RowsAffected", err)
		return nil, err
	}
	log.Println("affect:", affect)
	return &dbproto.Nil{}, nil
}
