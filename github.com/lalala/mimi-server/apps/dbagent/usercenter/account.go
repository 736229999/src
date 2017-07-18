package usercenter

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	. "github.com/caojunxyz/mimi-server/apps/dbagent/helper"
	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"golang.org/x/net/context"
)

func (agt *DbUsercenterAgent) CreateAccount(ctx context.Context, arg *dbproto.CreateAccountArg) (*dbproto.IntValue, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	phoneUser := arg.GetPhoneUser()
	qqUser := arg.GetQqUser()
	wxUser := arg.GetWxUser()
	if phoneUser == nil && qqUser == nil && wxUser == nil {
		log.Println("无效参数")
		return nil, ErrInvalidParam
	}

	tx, err := agt.dbConn.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var esql string
	userId := int64(0)
	nickname := "nickname"
	var detail string
	var tbUser, fieldUserId string

	userType := arg.GetUserType()
	switch userType {
	case dbproto.UserType_Phone:
		phone := phoneUser.GetPhone()
		password := phoneUser.GetPassword()
		esql = fmt.Sprintf("INSERT INTO %s(phone, password) VALUES('%s', '%s') ON CONFLICT (phone) DO UPDATE SET phone='%s' RETURNING id", TABLE_PHONE_USER, phone, password, phone)
		if err = tx.QueryRow(esql).Scan(&userId); err != nil {
			log.Println(err, esql)
			tx.Rollback()
			return nil, err
		}
		tbUser = TABLE_PHONE_USER
		fieldUserId = "phone_id"
		nickname = phone
		detail = fmt.Sprintf("phone: %s", phone)
	case dbproto.UserType_QQ:
		openid := qqUser.GetOpenid()
		esql = fmt.Sprintf("INSERT INTO %s(openid) VALUES('%s') ON CONFLICT (openid) DO UPDATE SET openid='%s' RETURNING id", TABLE_QQ_USER, openid, openid)
		if err = tx.QueryRow(esql).Scan(&userId); err != nil {
			log.Println(err, esql)
			tx.Rollback()
			return nil, err
		}
		tbUser = TABLE_QQ_USER
		fieldUserId = "qq_id"
		nickname = fmt.Sprintf("qq%d", userId)
		detail = fmt.Sprintf("qq: %s", openid)
	case dbproto.UserType_Weixin:
		openid := wxUser.GetOpenid()
		esql = fmt.Sprintf("INSERT INTO %s(openid) VALUES('%s') ON CONFLICT (openid) DO UPDATE SET openid='%s' RETURNING id", TABLE_WEIXIN_USER, openid, openid)
		if err = tx.QueryRow(esql).Scan(&userId); err != nil {
			log.Println(err, esql)
			tx.Rollback()
			return nil, err
		}
		tbUser = TABLE_WEIXIN_USER
		fieldUserId = "weixin_id"
		nickname = fmt.Sprintf("wx%d", userId)
		detail = fmt.Sprintf("wx: %s", openid)
	default:
		err = ErrInvalidParam
		log.Println(err)
		tx.Rollback()
		return nil, err
	}

	now := time.Now().Unix()
	// 账户表
	accountId := int64(0)
	esql = fmt.Sprintf("INSERT INTO %s(%s) VALUES(%d) RETURNING id", TABLE_ACCOUNT, fieldUserId, userId)
	if err = tx.QueryRow(esql).Scan(&accountId); err != nil {
		log.Println(err, esql)
		tx.Rollback()
		return nil, err
	}

	// 用户表
	esql = fmt.Sprintf("UPDATE %s SET account_id=%d WHERE id=%d RETURNING id", tbUser, accountId, userId)
	if err = tx.QueryRow(esql).Scan(&userId); err != nil {
		log.Println(err, esql)
		tx.Rollback()
		return nil, err
	}

	// 账户变动历史表
	changeType := dbproto.AccountChangeType_Create
	esql = fmt.Sprintf("INSERT INTO %s(account_id, change_type, user_type, user_id, change_time, ip, detail) VALUES(%d, %d, %d, %d, %d, '%s', '%s') RETURNING id",
		TABLE_ACCOUNT_HISTORY, accountId, changeType, userType, userId, now, arg.GetIp(), detail,
	)
	_, err = tx.Exec(esql)
	if err != nil {
		log.Println(err, esql)
		tx.Rollback()
		return nil, err
	}

	exp := 0
	lvl := 1
	// 用户信息表
	esql = fmt.Sprintf("INSERT INTO %s(account_id, exp, lvl, nickname) VALUES(%d, %d, %d, '%s')", TABLE_USERINFO, accountId, exp, lvl, nickname)
	_, err = tx.Exec(esql)
	if err != nil {
		log.Println(err, esql)
		tx.Rollback()
		return nil, err
	}

	// 初始化资金表
	esql = fmt.Sprintf("INSERT INTO %s(account_id) VALUES(%d)", TABLE_FUND, accountId)
	_, err = tx.Exec(esql)
	if err != nil {
		log.Println(err, esql)
		tx.Rollback()
		return nil, err
	}

	// 初始化虚拟资金表
	esql = fmt.Sprintf("INSERT INTO %s(account_id, credits, kxd) VALUES(%d, %d, %d)", TABLE_VIRTUAL_FUND, accountId, 0, 0)
	_, err = tx.Exec(esql)
	if err != nil {
		log.Println(err, esql)
		tx.Rollback()
		return nil, err
	}
	return &dbproto.IntValue{Value: accountId}, tx.Commit()
}

func (agt *DbUsercenterAgent) ChangePhone(ctx context.Context, arg *dbproto.ChangePhoneArg) (*dbproto.Nil, error) {
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

	phone := arg.GetPhone()
	newPhone := arg.GetNewPhone()

	var esql string
	var valAccountId sql.NullInt64
	var valPwd sql.NullString

	esql = fmt.Sprintf("SELECT account_id FROM %s WHERE phone='%s'", TABLE_PHONE_USER, newPhone)
	if err = tx.QueryRow(esql).Scan(&valAccountId); err != nil {
		if err != sql.ErrNoRows {
			log.Println(err, esql)
			return nil, err
		}
	}
	if valAccountId.Valid && valAccountId.Int64 > 0 {
		return nil, fmt.Errorf("新手机号已绑定")
	}

	esql = fmt.Sprintf("SELECT account_id, password FROM %s WHERE phone='%s'", TABLE_PHONE_USER, phone)
	if err = tx.QueryRow(esql).Scan(&valAccountId, &valPwd); err != nil {
		log.Println(err, esql)
		return nil, err
	}
	if !valAccountId.Valid || valAccountId.Int64 == 0 {
		return nil, fmt.Errorf("旧手机号未绑定")
	}

	accountId := valAccountId.Int64
	// 插入或设置新手机号的account_id为旧手机号的account_id
	if valPwd.Valid {
		esql = fmt.Sprintf("INSERT INTO %s AS d(phone, password, account_id) VALUES('%s', '%s', %d) ON CONFLICT (phone) DO UPDATE SET account_id=%d WHERE d.account_id=0 RETURNING id",
			TABLE_PHONE_USER, newPhone, valPwd.String, accountId, accountId)
	} else {
		esql = fmt.Sprintf("INSERT INTO %s AS d(phone, account_id) VALUES('%s', %d) ON CONFLICT (phone) DO UPDATE SET account_id=%d WHERE d.account_id=0 RETURNING id",
			TABLE_PHONE_USER, newPhone, accountId, accountId)
	}

	userId := int64(0)
	if err = tx.QueryRow(esql).Scan(&userId); err != nil {
		log.Println(err, esql)
		tx.Rollback()
		return nil, err
	}

	// 设置旧手机账号的account_id为0
	esql = fmt.Sprintf("UPDATE %s SET account_id=0 WHERE phone='%s'", TABLE_PHONE_USER, phone)
	_, err = tx.Exec(esql)
	if err != nil {
		log.Println(err, esql)
		tx.Rollback()
		return nil, err
	}

	esql = fmt.Sprintf("UPDATE %s SET phone_id=%d WHERE id=%d", TABLE_ACCOUNT, userId, accountId)
	_, err = tx.Exec(esql)
	if err != nil {
		log.Println(err, esql)
		tx.Rollback()
		return nil, err
	}

	now := time.Now().Unix()
	// 修改默认手机号昵称
	nickname := fmt.Sprint(newPhone)
	esql = fmt.Sprintf("UPDATE %s SET nickname='%s' WHERE account_id=%d", TABLE_USERINFO, nickname, accountId)
	_, err = tx.Exec(esql)
	if err != nil {
		log.Println(err, esql)
		tx.Rollback()
		return nil, err
	}

	// 账户变动历史表
	esql = fmt.Sprintf(`INSERT INTO uc_account_history(account_id, change_type, user_type, user_id, change_time, ip, detail) 
		VALUES(%d, %d, %d, %d, '%s', '%s', '%s') RETURNING id`, TABLE_ACCOUNT_HISTORY,
		accountId, dbproto.AccountChangeType_ChangePhone, dbproto.UserType_Phone, userId,
		now, arg.GetIp(), fmt.Sprintf("%s -> %s", phone, newPhone),
	)
	_, err = tx.Exec(esql)
	if err != nil {
		log.Println(err, esql)
		tx.Rollback()
		return nil, err
	}

	return &dbproto.Nil{}, tx.Commit()
}

func (agt *DbUsercenterAgent) SetAccountUser(ctx context.Context, arg *dbproto.SetUserArg) (*dbproto.Nil, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	log.Printf("%+v\n", arg)
	phoneUser := arg.GetPhoneUser()
	qqUser := arg.GetQqUser()
	wxUser := arg.GetWxUser()
	if phoneUser == nil && qqUser == nil && wxUser == nil {
		log.Println("无效参数")
		return nil, ErrInvalidParam
	}

	tx, err := agt.dbConn.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var esql string
	var fieldUserId string
	accountId := int64(0)
	userType := arg.GetUserType()
	setType := arg.GetSetType()
	var detail string

	switch userType {
	case dbproto.UserType_Phone:
		fieldUserId = "phone_id"
		accountId = phoneUser.GetAccountId()
		phone := phoneUser.GetPhone()
		if setType == dbproto.AccountChangeType_Bind {
			esql = fmt.Sprintf(`INSERT INTO %s AS d (phone, account_id) VALUES('%s', %d) 
				ON CONFLICT(phone) DO UPDATE SET account_id=%d WHERE d.account_id=0 RETURNING id`, TABLE_PHONE_USER, phone, accountId, accountId)
		} else if setType == dbproto.AccountChangeType_Unbind {
			esql = fmt.Sprintf("UPDATE %s SET account_id=0 WHERE phone='%s'", TABLE_PHONE_USER, phone)
		}
		detail = fmt.Sprintf("phone: %s", phone)
	case dbproto.UserType_QQ:
		fieldUserId = "qq_id"
		accountId = qqUser.GetAccountId()
		openId := qqUser.GetOpenid()
		if setType == dbproto.AccountChangeType_Bind {
			esql = fmt.Sprintf(`INSERT INTO %s AS d (openid, account_id) VALUES('%s',  %d) 
				ON CONFLICT(openid) DO UPDATE SET account_id=%d WHERE d.account_id=0 RETURNING id`, TABLE_QQ_USER, openId, accountId, accountId)
		} else if setType == dbproto.AccountChangeType_Unbind {
			esql = fmt.Sprintf("UPDATE %s SET account_id=0 WHERE openid='%s'", TABLE_QQ_USER, openId)
		}
		detail = fmt.Sprintf("qq: %s", openId)
	case dbproto.UserType_Weixin:
		fieldUserId = "weixin_id"
		accountId = wxUser.GetAccountId()
		openId := wxUser.GetOpenid()
		if setType == dbproto.AccountChangeType_Bind {
			esql = fmt.Sprintf(`INSERT INTO %s AS d (openid, account_id) VALUES('%s',  %d) 
				ON CONFLICT(openid) DO UPDATE SET account_id=%d WHERE d.account_id=0 RETURNING id`, TABLE_WEIXIN_USER, openId, accountId, accountId)
		} else if setType == dbproto.AccountChangeType_Unbind {
			esql = fmt.Sprintf("UPDATE %s SET account_id=0 WHERE openid='%s'", TABLE_WEIXIN_USER, openId)
		}
		detail = fmt.Sprintf("wx: %s", openId)
	}

	userId := int64(0)
	if setType == dbproto.AccountChangeType_Bind {
		var valUserId sql.NullInt64
		if err = tx.QueryRow(esql).Scan(&valUserId); err != nil {
			log.Println(err, esql)
			tx.Rollback()
			return nil, err
		}

		if valUserId.Valid {
			userId = valUserId.Int64
		} else {
			tx.Rollback()
			return nil, err
		}
		esql = fmt.Sprintf("UPDATE %s SET %s=%d WHERE id=%d", TABLE_ACCOUNT, fieldUserId, userId, accountId)
		_, err = tx.Exec(esql)
		if err != nil {
			log.Println(err, esql)
			tx.Rollback()
			return nil, err
		}
	} else if setType == dbproto.AccountChangeType_Unbind {
		log.Println(esql)
		_, err = tx.Exec(esql)
		if err != nil {
			log.Println(err, esql)
			tx.Rollback()
			return nil, err
		}

		esql = fmt.Sprintf("UPDATE %s SET %s=0 WHERE id=%d", TABLE_ACCOUNT, fieldUserId, accountId)
		log.Println(esql)
		_, err = tx.Exec(esql)
		if err != nil {
			log.Println(err, esql)
			tx.Rollback()
			return nil, err
		}
	}

	// 账户变动历史表
	esql = fmt.Sprintf(`INSERT INTO %s(account_id, change_type, user_type, user_id, change_time, ip, detail) 
		VALUES(%d, %d, %d, %d, %d, '%s', '%s') RETURNING id`, TABLE_ACCOUNT_HISTORY,
		accountId, setType, dbproto.UserType_Phone, userId, time.Now().Unix(), arg.GetIp(), detail,
	)
	_, err = tx.Exec(esql)
	if err != nil {
		log.Println(err, esql)
		tx.Rollback()
		return nil, err
	}
	return &dbproto.Nil{}, tx.Commit()
}
