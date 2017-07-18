package usercenter

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"golang.org/x/net/context"
)

func (agt *DbUsercenterAgent) QueryPhoneUser(ctx context.Context, arg *dbproto.StringValue) (*dbproto.PhoneUser, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	value := arg.GetValue()
	esql := fmt.Sprintf("SELECT password, account_id FROM %s WHERE phone='%s'", TABLE_PHONE_USER, value)
	var valPwd sql.NullString
	var valAccountId sql.NullInt64
	if err := agt.dbConn.QueryRow(esql).Scan(&valPwd, &valAccountId); err != nil {
		if err != sql.ErrNoRows {
			log.Println(err, esql)
			return nil, err
		}
	}
	ret := &dbproto.PhoneUser{Phone: value}
	if valPwd.Valid {
		ret.Password = valPwd.String
	}
	if valAccountId.Valid {
		ret.AccountId = valAccountId.Int64
	}
	return ret, nil
}

func (agt *DbUsercenterAgent) SetPhonePassword(ctx context.Context, arg *dbproto.PhonePassword) (*dbproto.Nil, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	tx, err := agt.dbConn.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var userId, accountId int64
	var esql string
	esql = fmt.Sprintf("UPDATE %s SET password='%s' WHERE phone='%s' RETURNING id, account_id", TABLE_PHONE_USER, arg.GetPassword(), arg.GetPhone())
	log.Println(esql)
	if err = tx.QueryRow(esql).Scan(&userId, &accountId); err != nil {
		log.Println(err, esql)
		tx.Rollback()
		return nil, err
	}

	// 账户变动历史表
	esql = fmt.Sprintf(`INSERT INTO %s(account_id, change_type, user_type, user_id, change_time, ip) VALUES(%d, %d, %d, %d, %d, '%s') RETURNING id`,
		TABLE_ACCOUNT_HISTORY, accountId, dbproto.AccountChangeType_Password, dbproto.UserType_Phone, userId, time.Now().Unix(), arg.GetIp(),
	)
	_, err = tx.Exec(esql)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}
	return &dbproto.Nil{}, tx.Commit()
}
