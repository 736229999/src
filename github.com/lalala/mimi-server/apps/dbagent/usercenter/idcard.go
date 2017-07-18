package usercenter

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/utils"
	"golang.org/x/net/context"
)

func (agt *DbUsercenterAgent) InsertAccountIdcard(ctx context.Context, arg *dbproto.AccountIdcard) (*dbproto.Nil, error) {
	accountId := arg.GetAccountId()
	idno := arg.GetIdno()
	realname := arg.GetRealname()
	now := time.Now().Unix()
	code := utils.GenerateInvitationCode(accountId)

	tx, err := agt.dbConn.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	esql := fmt.Sprintf("INSERT INTO %s(account_id, idno, realname, add_time) VALUES(%d, '%s', '%s', %d)",
		TABLE_IDCARD, accountId, idno, realname, now)
	_, err = tx.Exec(esql)
	if err != nil {
		log.Println(err, esql)
		tx.Rollback()
		return nil, err
	}

	esql = fmt.Sprintf("UPDATE %s SET invitation_code = '%s' WHERE account_id = %d", TABLE_USERINFO, code, accountId)
	_, err = tx.Exec(esql)
	if err != nil {
		log.Printf("%+v\n", err, esql)
		tx.Rollback()
		return nil, err
	}

	return &dbproto.Nil{}, tx.Commit()
}

func (agt *DbUsercenterAgent) QueryAccountIdcard(ctx context.Context, arg *dbproto.IntValue) (*dbproto.AccountIdcard, error) {
	accountId := arg.GetValue()
	esql := fmt.Sprintf("SELECT account_id, realname, idno, add_time FROM %s WHERE account_id=%d", TABLE_IDCARD, accountId)
	card := &dbproto.AccountIdcard{}
	if err := agt.dbConn.QueryRow(esql).Scan(&card.AccountId, &card.Realname, &card.Idno, &card.AddTime); err != nil {
		if err == sql.ErrNoRows {
			return card, nil
		}
		log.Println(err, esql)
		return nil, err
	}
	return card, nil
}
