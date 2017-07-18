package admin

import (
	"golang.org/x/net/context"
	"github.com/caojunxyz/mimi-admin/dbadminagent/proto"
	"log"
	"fmt"
	"strings"
	"database/sql"
	"time"
	"encoding/json"
)

//获取礼包模板的列表.
func (agt *DbAdminAgent) QueryGiftTemplateList (ctx context.Context, arg *dbproto.GiftTemplateList) (*dbproto.GiftTemplateList, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	filedList := []string {
		"id", "title", "content_desc", "content", "add_time", "creator",
	}

	whereList := []string{}
	if len(arg.GetTitle()) > 0 {
		whereList = append(whereList, fmt.Sprintf(" WHERE title LIKE '%s'", "%"+arg.GetTitle()+"%"))
	}

	offset := (arg.GetPage() - 1) * arg.GetSize()
	whereList = append(whereList, fmt.Sprintf("ORDER BY id DESC OFFSET %d LIMIT %d", offset, arg.GetSize()))

	esql := fmt.Sprintf("SELECT %s FROM gift_template %s", strings.Join(filedList, ", "), strings.Join(whereList, " "))

	totalSql := strings.Replace(esql, strings.Join(filedList, ", "), "COUNT(id)",-1)
	totalSql = strings.Replace(totalSql, whereList[len(whereList) - 1], "", -1)

	st, err := agt.ucDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	defer st.Close()

	rows, err := st.Query()
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	giftTempList := &dbproto.GiftTemplateList{}
	for rows.Next() {
		giftTemplate := &dbproto.GiftTemplate{}
		var contentJson  string
		value := []interface{}{
			&giftTemplate.Id, &giftTemplate.Title, &giftTemplate.ContentDesc,
			&contentJson, &giftTemplate.AddTime, &giftTemplate.Creator,
		}

		if err = rows.Scan(value...); err != nil {
			if err == sql.ErrNoRows {
				return &dbproto.GiftTemplateList{}, nil
			}
			log.Printf("%+v\n", err)
			return nil, err
		}

		if err = json.Unmarshal([]byte(contentJson), &giftTemplate.Content); err != nil {
			log.Printf("%+v\n", err)
			return nil, err
		}

		giftTempList.List = append(giftTempList.List, giftTemplate)
	}

	stmt, err := agt.ucDbConn.Prepare(totalSql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	defer stmt.Close()

	if err = stmt.QueryRow().Scan(&giftTempList.Total); err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	return giftTempList, nil
}

//添加礼包模板.
func (agt *DbAdminAgent) InsertGiftTemplate (ctx context.Context, arg *dbproto.GiftTemplate) (*dbproto.Nil, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	filedList := []string {
		"title", "content_desc", "content", "add_time", "creator",
	}

	content_byte, err := json.Marshal(arg.GetContent())
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	esql := fmt.Sprintf("INSERT INTO gift_template (%s) VALUES ('%s', '%s', '%s', %d, '%s')",
		strings.Join(filedList, ", "), arg.GetTitle(), arg.GetContentDesc(), string(content_byte), time.Now().Unix(), arg.GetCreator())

	log.Printf("esql:", esql)
	st, err := agt.ucDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	defer st.Close()

	_, err = st.Exec()
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	return &dbproto.Nil{}, nil
}

//根据id查询礼包模板.
func (agt *DbAdminAgent) QueryGiftTemplateById (ctx context.Context, arg *dbproto.IntValue) (*dbproto.GiftTemplate, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	filedList := []string{
		"id", "title", "content_desc", "content", "add_time", "creator",
	}
	esql := fmt.Sprintf("SELECT %s FROM gift_template WHERE id = %d", strings.Join(filedList, ", "), arg.GetValue())
	st, err := agt.ucDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	defer st.Close()

	template := &dbproto.GiftTemplate{}
	value := []interface{}{
		&template.Id, &template.Title, &template.ContentDesc,
		&template.ContentJson, &template.AddTime, &template.Creator,
	}
	if err = st.QueryRow().Scan(value...); err != nil {
		if err == sql.ErrNoRows {
			return &dbproto.GiftTemplate{}, nil
		}
		log.Printf("%+v\n", err)
		return nil, err
	}

	if  err = json.Unmarshal([]byte(template.GetContentJson()), &template.Content); err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	return template, nil
}

//更新礼包模板.
func (agt *DbAdminAgent) UpdateGiftTemplateById (ctx context.Context, arg *dbproto.GiftTemplate) (*dbproto.Nil, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	content_byte, err := json.Marshal(arg.GetContent())
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	fieldList := []string{
		fmt.Sprintf("title = '%s'", arg.GetTitle()),
		fmt.Sprintf("content_desc = '%s'", arg.GetContentDesc()),
		fmt.Sprintf("content = '%s'", string(content_byte)),
	}

	esql := fmt.Sprintf("UPDATE gift_template SET %s WHERE id = %d", strings.Join(fieldList, ", "), arg.GetId())
	st, err := agt.ucDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err, esql)
		return nil, err
	}
	defer st.Close()

	_, err = st.Exec()
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	return &dbproto.Nil{}, nil
}

//删除.
func (agt *DbAdminAgent) DeleteGiftTemplateById (ctx context.Context, arg *dbproto.IntValue) (*dbproto.Nil, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	esql := fmt.Sprintf("DELETE FROM gift_template WHERE id = %d", arg.GetValue())
	st, err := agt.ucDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	defer st.Close()

	_, err = st.Exec()
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	return &dbproto.Nil{}, nil
}

//获取所有的礼包列表.
func (agt *DbAdminAgent) QueryGiftTemplateListAll (ctx context.Context, arg *dbproto.Nil) (*dbproto.GiftTemplateList, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	filedList := []string{
		"id", "title", "content_desc", "content", "add_time", "creator",
	}
	esql := fmt.Sprintf("SELECT %s FROM gift_template ORDER BY id DESC", strings.Join(filedList, ", "))
	st, err := agt.ucDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err, esql)
		return nil, err
	}
	defer st.Close()

	rows, err := st.Query()
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	list := &dbproto.GiftTemplateList{}
	for rows.Next() {
		template := &dbproto.GiftTemplate{}

		value := []interface{}{
			&template.Id, &template.Title, &template.ContentDesc,
			&template.ContentJson, &template.AddTime, &template.Creator,
		}

		if err = rows.Scan(value...); err != nil {
			if err == sql.ErrNoRows {
				return &dbproto.GiftTemplateList{}, nil
			}
			log.Printf("%+v\n", err)
			return nil, err
		}

		if err = json.Unmarshal([]byte(template.GetContentJson()), &template.Content); err != nil {
			log.Printf("%+v\n", err)
			return nil, err
		}

		list.List = append(list.List, template)
	}

	return list, nil
}

