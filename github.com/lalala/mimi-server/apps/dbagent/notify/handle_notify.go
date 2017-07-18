package notify

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	"log"
	"time"

	. "github.com/caojunxyz/mimi-server/apps/dbagent/helper"
	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"golang.org/x/net/context"
)

// 创建一条通知
func (agt *DbNotifyAgent) CreateNotify(ctx context.Context, arg *dbproto.Notify) (*dbproto.IntValue, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	content := arg.GetContent()
	sender := arg.GetSender()
	nType := arg.GetType()
	target := arg.GetTarget()
	targetType := arg.GetTargetType()
	if content == "" {
		return nil, ErrInvalidParam
	}
	var esql string
	var notifyId int64
	now := time.Now().Unix()

	esql = fmt.Sprintf("INSERT INTO %s(content, type, target, target_type, sender, created) VALUES($1, $2, $3, $4, $5, $6) RETURNING id", TABLE_NOTIFY)
	err := agt.dbConn.QueryRow(esql, content, nType, target, targetType, sender, now).Scan(&notifyId)
	if err != nil {
		log.Println(err, esql)
		return nil, err
	}
	log.Printf("notifyId is %v", notifyId)

	return &dbproto.IntValue{Value: notifyId}, nil
}

// 创建一条用户通知
func (agt *DbNotifyAgent) CreateUserNotify(ctx context.Context, arg *dbproto.UserNotify) (*dbproto.IntValue, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	nId := arg.GetNotify()
	nType := arg.GetNotifyType()
	uId := arg.GetAccount()
	var esql string
	var rId int64
	now := time.Now().Unix()

	esql = fmt.Sprintf("INSERT INTO %s(account, notify, n_type, created) VALUES($1, $2, $3, $4) RETURNING id", TABLE_USER_NOTIFY)
	err := agt.dbConn.QueryRow(esql, uId, nId, nType, now).Scan(&rId)
	if err != nil {
		log.Println(err, esql)
		return nil, err
	}
	log.Printf("User notifyId is %v", rId)
	return &dbproto.IntValue{Value: rId}, nil
}

// 查询未加入用户消息队列的notify
func (agt *DbNotifyAgent) QueryNotifyUserMissed(ctx context.Context, arg *dbproto.QueryUserMissedArg) (*dbproto.QueryUserMissedRes, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	nType := arg.GetType()
	uId := arg.GetAccount()
	var esql string
	var err error
	var rows *sql.Rows
	agt.InitAccountUserNotify(uId)

	esql = fmt.Sprintf("SELECT id, type FROM %s WHERE  created > (SELECT MAX(created) AS last_time from %s WHERE account = $1)", TABLE_NOTIFY, TABLE_USER_NOTIFY)
	if nType != dbproto.NotifyType_All {
		esql += " AND type = $2"
		rows, err = agt.dbConn.Query(esql, uId, nType)
	} else {
		rows, err = agt.dbConn.Query(esql, uId)
	}
	if err != nil {
		log.Println(err, esql)
		return nil, err
	}
	ns := make([]*dbproto.Notify, 0)
	for rows.Next() {
		notice := &dbproto.Notify{}
		err := rows.Scan(&notice.Id, &notice.Type)
		if err != nil {
			log.Println(err, esql)
			return nil, err
		}
		ns = append(ns, notice)
	}

	log.Printf("User notifyIds is %v", ns)
	res := &dbproto.QueryUserMissedRes{Notices: ns}
	return res, nil
}

// 查询用户的消息队列
func (agt *DbNotifyAgent) QueryUserNotify(ctx context.Context, arg *dbproto.QueryUserNotifyArg) (*dbproto.QueryUserNotifyRes, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	var (
		err   error
		rows  *sql.Rows
		start int64
		pager string
		order string
	)
	if arg.Page > 1 {
		start = (arg.Page - 1) * arg.PageSize
	} else {
		start = 0
	}
	pager = fmt.Sprintf(" LIMIT %d OFFSET %d", arg.PageSize, start)
	order = " ORDER BY n.created DESC"
	esql := fmt.Sprintf("SELECT n.*, un.is_read, un.account FROM %s AS un JOIN %s AS n ON n.id = un.notify WHERE un.account = $1", TABLE_USER_NOTIFY, TABLE_NOTIFY)
	if arg.NotifyType != dbproto.NotifyType_All {
		esql += " AND n.type = $2"
		esql += order + pager
		rows, err = agt.dbConn.Query(esql, arg.Account, arg.NotifyType)
	} else {
		esql += order + pager
		rows, err = agt.dbConn.Query(esql, arg.Account)
	}
	if err != nil {
		log.Println(err, esql)
		return nil, err
	}
	log.Println("SQL:", esql)
	ns := make([]*dbproto.UserNotifyInfo, 0)
	for rows.Next() {
		noticeInfo := &dbproto.UserNotifyInfo{Notify: new(dbproto.Notify), UserNotify: new(dbproto.UserNotify)}
		log.Printf("noticeInfo is %+v\n", noticeInfo)
		err := rows.Scan(&noticeInfo.Notify.Id,
			&noticeInfo.Notify.Content, &noticeInfo.Notify.Target,
			&noticeInfo.Notify.TargetType, &noticeInfo.Notify.Action,
			&noticeInfo.Notify.Sender, &noticeInfo.Notify.Created,
			&noticeInfo.Notify.Updated, &noticeInfo.Notify.Type, &noticeInfo.UserNotify.IsRead,
			&noticeInfo.UserNotify.Account)
		if err != nil {
			log.Println(err, esql)
			return nil, err
		}
		ns = append(ns, noticeInfo)
	}
	log.Printf("User UserNotifyInfo is %+v\n", ns)
	res := &dbproto.QueryUserNotifyRes{UserNoticeInfos: ns}
	return res, nil
}

// 初始化用户user_notify
func (agt *DbNotifyAgent) InitAccountUserNotify(account int64) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	var count, rId int64
	esql := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE account = $1", TABLE_USER_NOTIFY)
	err := agt.dbConn.QueryRow(esql, account).Scan(&count)
	if err != nil {
		log.Println(err, esql)
		return
	}
	if count == 0 {
		esql = fmt.Sprintf("INSERT INTO %s(account, notify, n_type, created) VALUES($1, $2, $3, $4) RETURNING id", TABLE_USER_NOTIFY)
		err := agt.dbConn.QueryRow(esql, account, 0, 0, 0).Scan(&rId)
		if err != nil {
			log.Println(err, esql)
			return
		}
	}
	log.Printf("Init User notifyId is %v", rId)
}
