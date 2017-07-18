package usercenter

import (
	"database/sql"
	"fmt"
	"log"

	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"golang.org/x/net/context"
)

func (agt *DbUsercenterAgent) QueryQQUser(ctx context.Context, arg *dbproto.StringValue) (*dbproto.QQUser, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	value := arg.GetValue()
	esql := fmt.Sprintf("SELECT account_id FROM %s WHERE openid='%s'", TABLE_QQ_USER, value)
	var valAccountId sql.NullInt64
	if err := agt.dbConn.QueryRow(esql).Scan(&valAccountId); err != nil {
		if err != sql.ErrNoRows {
			log.Println(err, esql)
			return nil, err
		}
	}
	ret := &dbproto.QQUser{Openid: value}
	if valAccountId.Valid {
		ret.AccountId = valAccountId.Int64
	}
	return ret, nil
}
