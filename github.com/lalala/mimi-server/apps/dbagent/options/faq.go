package options

import (
	"database/sql"
	"fmt"
	"log"

	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"golang.org/x/net/context"
)

// QueryFaqList 获取新闻列表
func (agt *DbOptionsAgent) QueryFaqList(ctx context.Context, arg *dbproto.Nil) (*dbproto.FaqList, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	var (
		total              int64
		order, esql, where string
	)

	order = " ORDER BY sort DESC, create_time DESC"
	esql = fmt.Sprintf("SELECT id, title, content_url FROM %s", TABLE_FAQ)
	where = " WHERE is_visible = true"
	// log.Println(start, pager, order, total)
	rows, err := agt.dbConn.Query(esql + where + order)
	if err != nil {
		log.Println(err, esql)
		return nil, err
	}
	faqList := make([]*dbproto.Faq, 0)
	for rows.Next() {
		faq := &dbproto.Faq{}
		err := rows.Scan(&faq.Id,
			&faq.Title,
			&faq.ContentUrl,
		)
		if err != nil {
			log.Println(err, esql)
			return nil, err
		}
		faqList = append(faqList, faq)
	}
	totalEsql := fmt.Sprintf("SELECT count(*) FROM %s", TABLE_FAQ)
	row := agt.dbConn.QueryRow(totalEsql + where)
	row.Scan(&total)
	log.Println("esql: ", esql+where+order)
	// esql := fmt.Sprintf("")
	log.Printf("total is %v\n", total)

	return &dbproto.FaqList{
		List:  faqList,
		Total: total,
	}, nil
}

// QueryFaqById 查询一条Faq信息
func (agt *DbOptionsAgent) QueryFaqById(ctx context.Context, arg *dbproto.FaqId) (*dbproto.Faq, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	var (
		esql string
	)
	esql = fmt.Sprintf("SELECT id, title, html FROM %s WHERE id = $1", TABLE_FAQ)

	log.Println("esql", esql, "id ", arg.GetId())
	st, err := agt.dbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	faq := &dbproto.Faq{}
	err = st.QueryRow(arg.GetId()).Scan(&faq.Id, &faq.Title, &faq.Html)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Faq不存在:%+v, %v\n", err, arg.GetId())
			return nil, err
		}
		log.Printf("%+v\n", err)
		return nil, err
	}
	log.Println("faq", faq)
	return faq, nil
}
