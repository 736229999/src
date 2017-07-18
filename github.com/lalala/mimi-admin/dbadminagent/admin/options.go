package admin

import (
	"database/sql"
	"fmt"
	"github.com/caojunxyz/mimi-admin/dbadminagent/proto"
	"golang.org/x/net/context"
	"log"
	"strings"
	"time"
)

// TABLE_FAQ 常见问题表名
const TABLE_FAQ string = "faq"

//更新客服信息.
func (agt *DbAdminAgent) UpdateContact(ctx context.Context, arg *dbproto.Contact) (*dbproto.Nil, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	contact, err := agt.QueryContact(context.Background(), &dbproto.Nil{})
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	var esql string
	if contact.GetId() < 1 {
		esql = fmt.Sprintf(`INSERT INTO contact(qq, wechat, email, telphone) VALUES ('%s', '%s', '%s', '%s')`, arg.GetQq(), arg.GetWechat(), arg.GetEmail(), arg.GetTelphone())
	} else {
		esql = fmt.Sprintf(`UPDATE contact SET qq = '%s', wechat = '%s', email = '%s', telphone = '%s' WHERE id = 1`, arg.GetQq(), arg.GetWechat(), arg.GetEmail(), arg.GetTelphone())
	}

	log.Println("sql:", esql)
	st, err := agt.optionsDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	_, err = st.Exec()
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	return &dbproto.Nil{}, nil
}

func (agt *DbAdminAgent) QueryContact(ctx context.Context, arg *dbproto.Nil) (*dbproto.Contact, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	esql := `SELECT id, qq, wechat, email, telphone FROM contact WHERE id = 1`
	st, err := agt.optionsDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	contact := &dbproto.Contact{}

	var qq, wechat, email, telphone sql.NullString
	if err = st.QueryRow().Scan(&contact.Id, &qq, &wechat, &email, &telphone); err != nil {
		if err == sql.ErrNoRows {
			return &dbproto.Contact{}, nil
		}
		log.Printf("%+v\n", err)
		return nil, err
	}

	if qq.Valid {
		contact.Qq = qq.String
	}

	if wechat.Valid {
		contact.Wechat = wechat.String
	}

	if email.Valid {
		contact.Email = email.String
	}

	if telphone.Valid {
		contact.Telphone = telphone.String
	}

	return contact, nil
}

//获取用户的反馈列表.
func (agt *DbAdminAgent) QueryFeedbackList(ctx context.Context, arg *dbproto.FeedbackList) (*dbproto.FeedbackList, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	esql := fmt.Sprintf(`SELECT id, email, name, content, status, create_time FROM feedback WHERE 1 = 1 `)
	totalSql := strings.Replace(esql, "id, email, name, content, status, create_time", "COUNT(*) AS num", -1)

	offset := (arg.GetPage() - 1) * arg.GetSize()
	esql += fmt.Sprintf(` ORDER BY status ASC, id DESC OFFSET %d LIMIT %d`, offset, arg.GetSize())

	st, err := agt.optionsDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	rows, err := st.Query()
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	feedbackList := &dbproto.FeedbackList{}
	for rows.Next() {
		feedback := &dbproto.Feedback{}
		var name sql.NullString
		if err = rows.Scan(&feedback.Id, &feedback.Email, &name, &feedback.Content, &feedback.Status, &feedback.CreateTime); err != nil {
			if err == sql.ErrNoRows {
				return feedbackList, nil
			}
			log.Printf("%+v\n", err)
			return nil, err
		}
		if name.Valid {
			feedback.Name = name.String
		}

		feedbackList.List = append(feedbackList.List, feedback)
	}

	st, err = agt.optionsDbConn.Prepare(totalSql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	log.Println("totalSql:", totalSql)
	if err = st.QueryRow().Scan(&feedbackList.Total); err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	return feedbackList, nil
}

//删除用户的反馈信息.
func (agt *DbAdminAgent) DeleteFeedbackById(ctx context.Context, arg *dbproto.IntValue) (*dbproto.Nil, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	esql := `DELETE FROM feedback WHERE id = $1`
	st, err := agt.optionsDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	_, err = st.Exec(arg.GetValue())
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	return &dbproto.Nil{}, nil
}

//获取用户反馈详细.
func (agt *DbAdminAgent) QueryFeedbackById(ctx context.Context, arg *dbproto.IntValue) (*dbproto.Feedback, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	esql := `SELECT id, email, name, content, create_time, status FROM feedback WHERE id = $1`
	st, err := agt.optionsDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	feedback := &dbproto.Feedback{}
	if err := st.QueryRow(arg.GetValue()).Scan(&feedback.Id, &feedback.Email, &feedback.Name, &feedback.Content, &feedback.CreateTime, &feedback.Status); err != nil {
		if err == sql.ErrNoRows {
			return &dbproto.Feedback{}, nil
		}
		log.Printf("%+v\n", err)
		return nil, err
	}

	return feedback, nil
}

//处理反馈.
func (agt *DbAdminAgent) UpdateFeedbackById(ctx context.Context, arg *dbproto.Feedback) (*dbproto.Nil, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	esql := `UPDATE feedback SET status = $1 WHERE id = $2`
	st, err := agt.optionsDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	_, err = st.Exec(arg.GetStatus(), arg.GetId())
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	return &dbproto.Nil{}, nil
}

// CreateFaq 创建一条常见问题
func (agt *DbAdminAgent) CreateFaq(ctx context.Context, arg *dbproto.Faq) (*dbproto.IntValue, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	var rid int64
	esql := fmt.Sprintf("INSERT INTO %s(title, content_url, create_time, update_time, content, is_visible, html, sort) VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id", TABLE_FAQ)
	err := agt.optionsDbConn.QueryRow(esql, arg.GetTitle(), arg.GetContentUrl(), time.Now().Unix(), time.Now().Unix(), arg.GetContent(), arg.GetIsVisible(), arg.GetHtml(), arg.GetSort()).Scan(&rid)
	if err != nil {
		log.Printf("error %v, arg: %+v\n", err, arg)
		log.Printf("esql is %v", esql)
		return nil, err
	}
	return &dbproto.IntValue{Value: rid}, nil
}

// QueryFaqList 后台查询Faq列表
func (agt *DbAdminAgent) QueryFaqList(ctx context.Context, arg *dbproto.QueryFaqArg) (*dbproto.FaqList, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	var (
		start                     int32
		total                     int64
		esql, where, order, pager string
	)
	if arg.GetPage() > 1 {
		start = (arg.GetPage() - 1) * arg.PageSize
	} else {
		start = 0
	}
	pager = fmt.Sprintf(" LIMIT %d OFFSET %d", arg.PageSize, start)
	// where = " WHERE is_visible = true"
	order = " ORDER BY is_visible DESC, sort DESC, create_time DESC"
	esql = fmt.Sprintf("SELECT id, title, content_url, create_time, update_time, content, is_visible, html, sort FROM %s", TABLE_FAQ)
	if arg.GetTitle() != "" {
		where += fmt.Sprintf(" WHERE title LIKE '%%%s%%'", arg.GetTitle())
	}
	rows, err := agt.optionsDbConn.Query(esql + where + order + pager)
	log.Println("SQL", esql+where+order+pager)
	if err != nil {
		log.Println(err, "SQL: ", esql+where+order)
		return nil, err
	}

	faqList := make([]*dbproto.Faq, 0)

	for rows.Next() {
		faq := &dbproto.Faq{}
		err := rows.Scan(&faq.Id,
			&faq.Title,
			&faq.ContentUrl,
			&faq.CreateTime,
			&faq.UpdateTime,
			&faq.Content,
			&faq.IsVisible,
			&faq.Html,
			&faq.Sort,
		)
		if err != nil {
			log.Println(err, esql)
			return nil, err
		}
		faqList = append(faqList, faq)
	}
	//todo
	totalEsql := fmt.Sprintf("SELECT count(*) FROM %s", TABLE_FAQ)
	row := agt.optionsDbConn.QueryRow(totalEsql + where)
	row.Scan(&total)
	log.Println("esql: ", totalEsql+where)
	// log.Printf("bannerList is %+v\n", bannerList)
	// esql := fmt.Sprintf("")
	log.Printf("total is %v\n", total)

	return &dbproto.FaqList{
		List:  faqList,
		Total: total,
	}, nil
}

// QueryFaqById 查询一条Faq信息
func (agt *DbAdminAgent) QueryFaqById(ctx context.Context, arg *dbproto.FaqId) (*dbproto.Faq, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	var (
		esql string
	)
	esql = fmt.Sprintf("SELECT id, title, content_url, create_time, update_time, content, is_visible, html, sort FROM %s WHERE id = $1", TABLE_FAQ)

	log.Println("esql", esql, "id ", arg.GetId())
	st, err := agt.optionsDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	faq := &dbproto.Faq{}
	err = st.QueryRow(arg.GetId()).Scan(&faq.Id, &faq.Title, &faq.ContentUrl, &faq.CreateTime, &faq.UpdateTime, &faq.Content, &faq.IsVisible, &faq.Html, &faq.Sort)
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

// UpdateFaq 更新Faq信息
func (agt *DbAdminAgent) UpdateFaq(ctx context.Context, arg *dbproto.Faq) (*dbproto.IntValue, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	esql := fmt.Sprintf("UPDATE %s "+
		"SET title = $1, content_url = $2, update_time = $3, content = $4, is_visible = $5, html = $6, sort = $7 WHERE id = $8", TABLE_FAQ)
	st, err := agt.optionsDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	log.Println("SQL", esql)
	res, err := st.Exec(arg.GetTitle(), arg.GetContentUrl(), time.Now().Unix(), arg.GetContent(), arg.GetIsVisible(), arg.GetHtml(), arg.GetSort(), arg.GetId())
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
