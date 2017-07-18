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

//添加日志.
func (agt *DbAdminAgent) InsertLog(ctx context.Context, arg *dbproto.Log) (*dbproto.Nil, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	st, err := agt.dbConn.Prepare("INSERT INTO log(user_id, path, operating, params, message, create_time) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	_, err = st.Exec(arg.GetUserId(), arg.GetPath(), arg.GetOperating(), arg.GetParams(), arg.GetMessage(), time.Now().Unix())
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	return &dbproto.Nil{}, nil
}

//查询日志.
func (agt *DbAdminAgent) QueryLog(ctx context.Context, arg *dbproto.LogReply) (*dbproto.LogReply, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	esql := fmt.Sprintf(`SELECT l.id, u.username, l.message, l.path, l.create_time, l.operating FROM log AS l LEFT JOIN "user" AS u ON l.user_id = u.id WHERE 1 = 1`)

	//是否根据账号查询日志.
	if arg.GetAccount() != "" {
		esql += fmt.Sprintf(` AND u.email = '%s'`, arg.GetAccount())
	}

	//是否有权限查询所有的日志.
	if !arg.GetAuthorization() {
		esql += fmt.Sprintf(` AND l.user_id = %d`, arg.GetUserId())
	}

	totalSql := strings.Replace(esql, "l.id, u.username, l.message, l.path, l.create_time, l.operating", "COUNT(*) AS num", 1)
	log.Println("total sql:", totalSql)

	offset := (arg.GetPage() - 1) * arg.GetSize()
	esql += fmt.Sprintf(` ORDER BY l.id DESC OFFSET %d LIMIT %d`, offset, arg.GetSize())

	st, err := agt.dbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err, esql)
		return nil, err
	}

	rows, err := st.Query()
	if err != nil {
		if err != sql.ErrNoRows {
			log.Printf("%+v\n", err, esql)
			return nil, err
		}
		return &dbproto.LogReply{}, nil
	}

	logReply := &dbproto.LogReply{}

	for rows.Next() {

		logReplyArr := &dbproto.Log{}
		if err = rows.Scan(&logReplyArr.Id, &logReplyArr.Username, &logReplyArr.Message, &logReplyArr.Path, &logReplyArr.CreateTime, &logReplyArr.Operating); err != nil {
			log.Printf("%+v\n", err)
			return nil, err
		}

		logReply.Log = append(logReply.Log, logReplyArr)
	}

	st, err = agt.dbConn.Prepare(totalSql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	var num int64 = 0
	if err = st.QueryRow().Scan(&num); err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	logReply.Total = num

	return logReply, nil
}
