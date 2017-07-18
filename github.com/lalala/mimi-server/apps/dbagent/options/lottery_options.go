package options

import (
	"database/sql"
	"log"

	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"golang.org/x/net/context"
)

//获取 LotteryOptions 列表.
func (agt *DbOptionsAgent) QueryLotteryOptionsList(ctx context.Context, arg *dbproto.Nil) (*dbproto.HomeParams, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	esql := `SELECT id, lottery_name, is_plus_award, info, stop_sale FROM lottery_options`
	st, err := agt.dbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	rows, err := st.Query()
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	homeParams := &dbproto.HomeParams{}
	for rows.Next() {
		option := &dbproto.LotteryOptions{}
		var boolValue sql.NullBool
		if err = rows.Scan(&option.Id, &option.LotteryName, &boolValue, &option.Info, &option.StopSale); err != nil {
			if err == sql.ErrNoRows {
				return &dbproto.HomeParams{}, nil
			}
			log.Printf("%+v\n", err)
			return nil, err
		}
		if boolValue.Valid {
			option.IsPlusAward = boolValue.Bool
		}
		homeParams.Lottery = append(homeParams.Lottery, option)
	}

	return homeParams, nil
}

//获取客服联系方式.
func (agt *DbOptionsAgent) QueryContact(ctx context.Context, arg *dbproto.Nil) (*dbproto.Contact, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	esql := `SELECT id, qq, wechat, email, telphone FROM contact WHERE id = 1`
	st, err := agt.dbConn.Prepare(esql)
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
