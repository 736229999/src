package notify

import (
	. "github.com/caojunxyz/mimi-server/apps/dbagent/helper"
	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"

	"fmt"
	"log"

	"golang.org/x/net/context"
)

func (agt *DbNotifyAgent) ReadUserNotify(ctx context.Context, arg *dbproto.ReadUserNotifyArg) (*dbproto.IntValue, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	accountId := arg.GetAccountId()
	notifyId := arg.GetNotifyId()
	if accountId == 0 || notifyId == 0 {
		return nil, ErrInvalidParam
	}
	var (
		esql string
		err  error
	)
	esql = fmt.Sprintf("UPDATE %s SET is_read = TRUE WHERE account = $1 AND notify = $2", TABLE_USER_NOTIFY)
	stmt, err := agt.dbConn.Prepare(esql)
	if err != nil {
		log.Println("Prepare", err)
		return nil, err
	}
	res, err := stmt.Exec(accountId, notifyId)
	if err != nil {
		log.Println("stmt.Exec", err)
		return nil, err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		log.Println("res.RowsAffected", err)
		return nil, err
	}
	return &dbproto.IntValue{Value: affect}, nil
}
