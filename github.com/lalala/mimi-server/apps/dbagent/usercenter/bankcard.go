package usercenter

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"golang.org/x/net/context"
)

func (agt *DbUsercenterAgent) InsertAccountBankcard(ctx context.Context, arg *dbproto.AccountBankcard) (*dbproto.Nil, error) {
	accountId := arg.GetAccountId()
	cardNo := arg.GetCardno()
	bankname := arg.GetBankname()
	cardtype := arg.GetCardtype()
	phone := arg.GetPhone()
	addTime := time.Now().Unix()
	idno := arg.GetIdno()
	realname := arg.GetRealname()
	cols := "account_id, cardno, bankname, cardtype, phone, add_time, idno, realname"
	esql := fmt.Sprintf(`INSERT INTO %s(%s) VALUES($1, $2, $3, $4, $5, $6, $7, $8) 
		ON CONFLICT (account_id) DO UPDATE SET cardno=$9, bankname=$10, cardtype=$11, phone=$12, add_time=$13`, TABLE_BANKCARD, cols)
	_, err := agt.dbConn.Exec(esql, accountId, cardNo, bankname, cardtype, phone, addTime, idno, realname, cardNo, bankname, cardtype, phone, addTime)
	if err != nil {
		log.Println(err, esql)
		return nil, err
	}
	return &dbproto.Nil{}, nil
}

func (agt *DbUsercenterAgent) QueryAccountBankcard(ctx context.Context, arg *dbproto.IntValue) (*dbproto.AccountBankcard, error) {
	accountId := arg.GetValue()
	cols := "id, account_id, cardno, bankname, cardtype, phone, add_time, idno, realname"
	esql := fmt.Sprintf("SELECT %s FROM %s WHERE account_id=%d", cols, TABLE_BANKCARD, accountId)
	card := &dbproto.AccountBankcard{}
	if err := agt.dbConn.QueryRow(esql).Scan(&card.Id, &card.AccountId, &card.Cardno, &card.Bankname, &card.Cardtype, &card.Phone, &card.AddTime, &card.Idno, &card.Realname); err != nil {
		if err != sql.ErrNoRows {
			log.Println(err, esql)
			return nil, err
		}
	}
	return card, nil
}

func (agt *DbUsercenterAgent) DeleteAccountBankcard(ctx context.Context, arg *dbproto.IntValue) (*dbproto.Nil, error) {
	accountId := arg.GetValue()
	query := fmt.Sprintf("DELETE FROM %s WHERE account_id=$1", TABLE_BANKCARD)
	stmt, err := agt.dbConn.Prepare(query)
	if err != nil {
		log.Println(err, query)
		return nil, err
	}
	_, err = stmt.Exec(accountId)
	if err != nil {
		log.Println(err, query)
		return nil, err
	}
	return &dbproto.Nil{}, nil
}
