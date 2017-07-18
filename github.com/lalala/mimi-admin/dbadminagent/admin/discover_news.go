package admin

import (
	"fmt"
	"log"
	"time"

	"strings"

	dbproto "github.com/caojunxyz/mimi-admin/dbadminagent/proto"
	"golang.org/x/net/context"
)

// TABLE_NEWS 新闻表名
const TABLE_NEWS string = "news"

// CreateNews 创建一条新闻。
func (agt *DbAdminAgent) CreateNews(ctx context.Context, arg *dbproto.News) (*dbproto.IntValue, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	var rid int64
	esql := fmt.Sprintf("INSERT INTO %s(title, author, description, cover, content, html, created, updated, is_visible) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id", TABLE_NEWS)
	err := agt.optionsDbConn.QueryRow(esql, arg.GetTitle(), arg.GetAuthor(), arg.GetDescription(), arg.GetCover(), arg.GetContent(), arg.GetHtml(), time.Now().Unix(), time.Now().Unix(), arg.GetIsVisible()).Scan(&rid)
	if err != nil {
		log.Printf("error %v, arg: %+v\n", err, arg)
		log.Printf("esql is %v", esql)
		return nil, err
	}
	return &dbproto.IntValue{Value: rid}, nil
}

// QueryNewsList 后台查询新闻列表。
func (agt *DbAdminAgent) QueryNewsList(ctx context.Context, arg *dbproto.QueryNewsArg) (*dbproto.NewsList, error) {
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
	order = " ORDER BY is_visible DESC, created DESC"
	esql = fmt.Sprintf("SELECT id, content, title, description, pageviews, author, updated, html, cover, is_visible, news_class, created FROM %s", TABLE_NEWS)
	log.Println(start, pager, order)
	if arg.GetTitle() != "" {
		where = fmt.Sprintf(" WHERE title LIKE '%%%s%%'", arg.GetTitle())
	}
	if arg.GetAuthor() != "" {
		if strings.Contains(where, "WHERE") {
			where = fmt.Sprintf("%s AND author LIKE '%%%s%%'", where, arg.GetAuthor())
		} else {
			where = fmt.Sprintf(" WHERE author LIKE '%%%s%%'", arg.GetAuthor())
		}
	}
	if arg.GetStart() != 0 {
		if strings.Contains(where, "WHERE") {
			where = fmt.Sprintf("%s AND created >= %d", where, arg.GetStart())
		} else {
			where = fmt.Sprintf(" WHERE created >= %d", arg.GetStart())
		}
	}

	if arg.GetEnd() != 0 {
		if strings.Contains(where, "WHERE") {
			where = fmt.Sprintf("%s AND created <= %d", where, arg.GetEnd())
		} else {
			where = fmt.Sprintf(" WHERE created <= %d", arg.GetEnd())
		}
	}
	rows, err := agt.optionsDbConn.Query(esql + where + order + pager)
	if err != nil {
		log.Println(err, esql)
		return nil, err
	}
	newsList := make([]*dbproto.News, 0)
	for rows.Next() {
		news := &dbproto.News{}
		err := rows.Scan(&news.Id,
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
			&news.Created)
		if err != nil {
			log.Println(err, esql)
			return nil, err
		}
		newsList = append(newsList, news)
	}
	// totalEsql := strings.Replace(esql, "*", "COUNT(*)", 1)
	totalEsql := fmt.Sprintf("SELECT COUNT(*) FROM %s", TABLE_NEWS)
	row := agt.optionsDbConn.QueryRow(totalEsql + where)
	row.Scan(&total)
	log.Println("esql: ", esql+where+order+pager)
	log.Println("totalEsql: ", totalEsql+where)
	log.Printf("total is %v", total)
	return &dbproto.NewsList{List: newsList,
		Total: total}, nil
}

// QueryNewsById 获取一条新闻。
func (agt *DbAdminAgent) QueryNewsById(ctx context.Context, arg *dbproto.NewsId) (*dbproto.News, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	esql := fmt.Sprintf("SELECT id, content, title, description, pageviews, author, updated, html, cover, is_visible, news_class, created FROM %s WHERE id = $1", TABLE_NEWS)

	news := &dbproto.News{}
	err := agt.optionsDbConn.QueryRow(esql, arg.Id).Scan(
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

// QueryBakendSelectOfNews News的下拉框
func (agt *DbAdminAgent) QueryBakendSelectOfNews(ctx context.Context, arg *dbproto.QueryNewsOfSelect) (*dbproto.NewsList, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	var (
		esql, where, order string
		total              int64
	)

	where = fmt.Sprintf(" WHERE is_visible = true AND (title LIKE '%%%s%%"+
		"' OR author LIKE '%%%s%%')", arg.GetKeyWord(), arg.GetKeyWord())
	order = " ORDER BY created DESC"
	esql = fmt.Sprintf("SELECT id, title, description, author, created, cover FROM %s", TABLE_NEWS)
	rows, err := agt.optionsDbConn.Query(esql + where + order)
	if err != nil {
		log.Println(err, esql+where+order)
		return nil, err
	}
	newsList := make([]*dbproto.News, 0)
	for rows.Next() {
		news := &dbproto.News{}
		err := rows.Scan(&news.Id,
			&news.Title,
			&news.Description,
			&news.Author,
			&news.Created,
			&news.Cover,
		)
		if err != nil {
			log.Println(err, esql)
			return nil, err
		}
		newsList = append(newsList, news)
	}
	totalEsql := fmt.Sprintf("SELECT count(*) FROM %s", TABLE_NEWS)
	row := agt.optionsDbConn.QueryRow(totalEsql + where)
	row.Scan(&total)
	log.Println("esql: ", esql+where+order)
	// esql := fmt.Sprintf("")
	log.Printf("total is %v", total)
	return &dbproto.NewsList{List: newsList,
		Total: total}, nil
}

// UpdateNews 更新一条新闻信息
func (agt *DbAdminAgent) UpdateNews(ctx context.Context, arg *dbproto.News) (*dbproto.IntValue, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	esql := fmt.Sprintf("UPDATE %s "+
		"SET content = $1, title = $2, description = $3, author = $4, updated = $5, html = $6, cover = $7, is_visible = $8, news_class = $9 WHERE id = $10", TABLE_NEWS)
	st, err := agt.optionsDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	log.Println("SQL", esql)
	res, err := st.Exec(arg.GetContent(), arg.GetTitle(), arg.GetDescription(), arg.GetAuthor(), time.Now().Unix(), arg.GetHtml(), arg.GetCover(), arg.GetIsVisible(), arg.GetNewsClass(), arg.GetId())
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		log.Printf("RowsAffected %v\n", err)
		return nil, err
	}
	return &dbproto.IntValue{
		Value: affect,
	}, nil
}
