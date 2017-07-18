package admin

import (
	"log"
	"fmt"
	"database/sql"
	"strings"
	"time"
	"golang.org/x/net/context"
	"github.com/caojunxyz/mimi-admin/dbadminagent/proto"
)

//礼包列表.
func (agt *DbAdminAgent) QueryGiftList(ctx context.Context, arg *dbproto.GiftList) (*dbproto.GiftList, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	esql := `SELECT id, title, content_desc, content, add_time, gift_type FROM gift_package WHERE 1 = 1 `
	if arg.GetTitle() != "" {
		esql += fmt.Sprintf(`AND title LIKE '%s'`, "%"+arg.GetTitle()+"%")
	}

	//获取分页总条数的sql.
	totalSql := strings.Replace(esql, "id, title, content_desc, content, add_time, gift_type", "COUNT(*) AS num", 1)

	esql += ` ORDER BY id DESC `

	offset := (arg.GetPage() - 1) * arg.GetSize()
	esql += fmt.Sprintf(` OFFSET %d LIMIT %d`, offset, arg.GetSize())

	log.Println("sql:", esql)
	rows, err := agt.ucDbConn.Query(esql)
	if err != nil {
		log.Printf("%+v\n", err, esql)
		return nil, err
	}

	giftList := &dbproto.GiftList{}
	for rows.Next() {

		gift := &dbproto.GiftPackageArg{}
		if err = rows.Scan(&gift.Id, &gift.Title, &gift.ContentDesc, &gift.Content, &gift.AddTime, &gift.GiftType); err != nil {
			if err == sql.ErrNoRows {
				return &dbproto.GiftList{}, nil
			}
			log.Printf("%+v\n", err)
			return nil, err
		}

		boolValue, err := agt.GetGiftPackageIsUseById(ctx, &dbproto.IntValue{Value: gift.Id})
		if err != nil {
			log.Printf("%+v\n", err)
			return nil, err
		}

		gift.Status = boolValue.GetValue()

		log.Printf("gift:%+v\n", gift)
		giftList.List = append(giftList.List, gift)
	}

	row := agt.ucDbConn.QueryRow(totalSql)
	if err = row.Scan(&giftList.Total); err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	return giftList, nil
}

//添加礼包.
func (agt *DbAdminAgent) InsertGift(ctx context.Context, arg *dbproto.GiftPackageRequest) (*dbproto.Nil, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	esql := "INSERT INTO gift_package (title, content_desc, content, add_time, gift_type) VALUES ($1, $2, $3, $4, $5)"
	st, err := agt.ucDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	_, err = st.Exec(arg.GetTitle(), arg.GetContentDesc(), arg.GetContent(), time.Now().Unix(), arg.GetGiftType())
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	return &dbproto.Nil{}, nil
}

//获取礼包的详细.
func (agt *DbAdminAgent) QueryGiftDetailById(ctx context.Context, arg *dbproto.IntValue) (*dbproto.GiftPackageRequest, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	esql := `SELECT title, content_desc, content, add_time, gift_type FROM gift_package WHERE id = $1`
	st, err := agt.ucDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	gift := &dbproto.GiftPackageRequest{}

	//var content string
	if err = st.QueryRow(arg.GetValue()).Scan(&gift.Title, &gift.ContentDesc, &gift.Content, &gift.AddTime, &gift.GiftType); err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	return gift, nil
}

//根据id删除礼包.
func (agt *DbAdminAgent) DeleteGiftById(ctx context.Context, arg *dbproto.IntValue) (*dbproto.Nil, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	esql := `DELETE FROM gift_package WHERE id = $1`
	st, err := agt.ucDbConn.Prepare(esql)
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

//更新礼包.
func (agt *DbAdminAgent) UpdateGiftById(ctx context.Context, arg *dbproto.GiftPackageRequest) (*dbproto.Nil, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	esql := `UPDATE gift_package SET title = $1, content_desc = $2, content = $3, gift_type = $4 WHERE id = $5`
	st, err := agt.ucDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	log.Printf("arg:%+v\n", arg)
	_, err = st.Exec(arg.GetTitle(), arg.GetContentDesc(), arg.GetContent(), arg.GetGiftType(), arg.GetId())
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	return &dbproto.Nil{}, nil
}

//获取cdkey的列表.
func (agt *DbAdminAgent) QueryCdkeyList(ctx context.Context, arg *dbproto.CdkeyReply) (*dbproto.CdkeyReply, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	filedList := []string{
		"id", "max_exchange", "title", "gift_template_id", "valid_start",
		"valid_end", "add_time", "cdkey_desc",
	}
	esql := fmt.Sprintf(`SELECT %s FROM cdkey_batch WHERE 1 = 1 `, strings.Join(filedList, ", "))

	if arg.GetTitle() != "" {
		esql += fmt.Sprintf(` AND title LIKE '%s'`, "%"+arg.GetTitle()+"%")
	}

	totalSql := strings.Replace(esql, strings.Join(filedList, ", "), "COUNT(*) AS num", -1)

	offset := (arg.GetPage() - 1) * arg.GetSize()
	esql += fmt.Sprintf(` ORDER BY id DESC OFFSET %d LIMIT %d`, offset, arg.GetSize())

	st, err := agt.ucDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err, esql)
		return nil, err
	}
	rows, err := st.Query()
	if err != nil {
		log.Printf("%+v\n", err, esql)
		return nil, err
	}

	cdkeyReply := &dbproto.CdkeyReply{}

	for rows.Next() {
		cdkey := &dbproto.CdkeyBatch{}
		value := []interface{}{
			&cdkey.Id, &cdkey.MaxExchange, &cdkey.Title, &cdkey.GiftTemplateId,
			&cdkey.ValidStart, &cdkey.ValidEnd, &cdkey.AddTime, &cdkey.CdkeyDesc,
		}
		if err = rows.Scan(value...); err != nil {
			if err == sql.ErrNoRows {
				return &dbproto.CdkeyReply{}, nil
			}
			log.Printf("%+v\n", err)
			return nil, err
		}
		cdkeyReply.List = append(cdkeyReply.List, cdkey)
	}

	st, err = agt.ucDbConn.Prepare(totalSql)
	if err != nil {
		log.Printf("%+v\n", err, totalSql)
		return nil, err
	}

	if err = st.QueryRow().Scan(&cdkeyReply.Total); err != nil {
		log.Printf("%+v\n", err, totalSql)
		return nil, err
	}

	return cdkeyReply, nil
}

//根据礼包类型获取礼包列表.
//func (agt *DbAdminAgent) QueryGiftListByType(ctx context.Context, arg *dbproto.GiftListByType) (*dbproto.GiftListByType, error) {
//
//	defer func() {
//		if err := recover(); err != nil {
//			log.Println("recover from panic:", err)
//		}
//	}()
//
//	fieldList := []string{
//		"id", "title", "content_desc", "content", "add_time", "gift_type",
//	}
//	esql := fmt.Sprintf(`SELECT %s FROM gift_package WHERE gift_type = %d`, strings.Join(fieldList, ","), arg.GetType())
//	st, err := agt.ucDbConn.Prepare(esql)
//	if err != nil {
//		log.Printf("%+v\n", err)
//		return nil, err
//	}
//
//	rows, err := st.Query()
//	if err != nil {
//		log.Printf("%+v\n", err)
//		return nil, err
//	}
//
//	list := &dbproto.GiftListByType{}
//	for rows.Next() {
//		gift := &dbproto.GiftPackageArg{}
//		value := []interface{}{
//			&gift.Id, &gift.Title, &gift.ContentDesc,
//			&gift.Content, &gift.AddTime, &gift.GiftType,
//		}
//		if err = rows.Scan(value...); err != nil {
//			if err == sql.ErrNoRows {
//				return &dbproto.GiftListByType{}, nil
//			}
//			log.Printf("%+v\n", err)
//			return nil, err
//		}
//
//		list.GiftList = append(list.GiftList, gift)
//	}
//
//	return list, nil
//}

//添加cdkey .
func (agt *DbAdminAgent) InsertCdkeyBatch(ctx context.Context, arg *dbproto.CdkeyBatch) (*dbproto.Nil, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	filedList := []string{
		"max_exchange", "title", "gift_template_id", "valid_start", "valid_end", "add_time", "cdkey_desc",
	}
	esql := fmt.Sprintf(`INSERT INTO cdkey_batch (%s) VALUES ($1, $2, $3, $4, $5, $6, $7)`, strings.Join(filedList, ", "))
	st, err := agt.ucDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	_, err = st.Exec(arg.GetMaxExchange(), arg.GetTitle(), arg.GetGiftTemplateId(), arg.GetValidStart(), arg.GetValidEnd(), time.Now().Unix(), arg.GetCdkeyDesc())
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	return &dbproto.Nil{}, nil
}

//根据id获取cdkey的详细内容.
func (agt *DbAdminAgent) QueryCdkeyDetailById(ctx context.Context, arg *dbproto.IntValue) (*dbproto.CdkeyDetail, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	fieldList := []string{
		"c.id", "c.max_exchange", "c.title", "c.gift_template_id",
		"c.valid_start", "c.valid_end", "c.add_time", "g.content", "c.cdkey_desc",
	}
	esql := fmt.Sprintf(`SELECT %s FROM cdkey_batch AS c LEFT JOIN gift_template AS g ON c.gift_template_id = g.id WHERE c.id = %d`, strings.Join(fieldList, ", "), arg.GetValue())
	st, err := agt.ucDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	detail := &dbproto.CdkeyDetail{}
	value := []interface{}{
		&detail.Id, &detail.MaxExchange, &detail.Title, &detail.GiftTemplateId,
		&detail.ValidStart, &detail.ValidEnd, &detail.AddTime, &detail.Content, &detail.CdkeyDesc,
	}
	if err = st.QueryRow().Scan(value...); err != nil {
		if err == sql.ErrNoRows {
			return &dbproto.CdkeyDetail{}, nil
		}
		log.Printf("%+v\n", err)
		return nil, err
	}

	return detail, nil
}

//更新cdkey.
func (agt *DbAdminAgent) UpdateCdkeyById(ctx context.Context, arg *dbproto.CdkeyDetail) (*dbproto.Nil, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	filedList := []string{
		fmt.Sprintf(`max_exchange = %d`, arg.GetMaxExchange()),
		fmt.Sprintf(`title = '%s'`, arg.GetTitle()),
		fmt.Sprintf(`gift_template_id = %d`, arg.GetGiftTemplateId()),
		fmt.Sprintf(`valid_start = %d`, arg.GetValidStart()),
		fmt.Sprintf(`valid_end = %d`, arg.GetValidEnd()),
		fmt.Sprintf(`cdkey_desc = '%s'`, arg.GetCdkeyDesc()),
	}
	esql := fmt.Sprintf(`UPDATE cdkey_batch SET %s WHERE id = %d`, strings.Join(filedList, ", "), arg.GetId())
	st, err := agt.ucDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err, esql)
		return nil, err
	}

	_, err = st.Exec()
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	return &dbproto.Nil{}, nil
}

//删除cdkey.
func (agt *DbAdminAgent) DeleteCdkeyById(ctx context.Context, arg *dbproto.IntValue) (*dbproto.Nil, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	st, err := agt.ucDbConn.Prepare(`DELETE FROM cdkey_batch WHERE id = $1`)
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

//根据cdkey礼包id判断当前礼包是否正在使用当中.
func (agt *DbAdminAgent) GetGiftPackageIsUseById(ctx context.Context, arg *dbproto.IntValue) (*dbproto.BoolValue, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	sqlStr := fmt.Sprintf(`SELECT valid_end > extract(epoch from now())::int AS status FROM cdkey_batch WHERE gift_template_id  = %d`, arg.GetValue())
	st, err := agt.ucDbConn.Prepare(sqlStr)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	rws, err := st.Query()
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	var status bool
	for rws.Next() {
		if err = rws.Scan(&status); err != nil {
			if err == sql.ErrNoRows {
				return &dbproto.BoolValue{Value: false}, nil
			}
			log.Printf("%+v\n", err)
			return nil, err
		}

		//.status == false 表示已经过期，没有被使用了
		if status {
			return &dbproto.BoolValue{Value: true}, nil
		}
	}

	return &dbproto.BoolValue{Value: false}, nil
}
