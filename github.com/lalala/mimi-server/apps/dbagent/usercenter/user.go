package usercenter

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	. "github.com/caojunxyz/mimi-server/apps/dbagent/helper"
	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"golang.org/x/net/context"
)

func (agt *DbUsercenterAgent) QueryUserInfo(ctx context.Context, arg *dbproto.IntValue) (*dbproto.UserInfo, error) {
	accountId := arg.GetValue()
	if accountId <= 0 {
		return nil, ErrInvalidParam
	}
	columns := []string{
		fmt.Sprintf("%s.icon AS icon", TABLE_USERINFO),                                  // icon
		fmt.Sprintf("%s.sex AS sex", TABLE_USERINFO),                                    // sex
		fmt.Sprintf("%s.nickname AS nickname", TABLE_USERINFO),                          // nickname
		fmt.Sprintf("%s.exp AS exp", TABLE_USERINFO),                                    // exp
		fmt.Sprintf("%s.lvl AS lvl", TABLE_USERINFO),                                    // level
		fmt.Sprintf("%s.invitation_code AS invitation_code", TABLE_USERINFO),            // invitation_code
		fmt.Sprintf("char_length(%s.pay_password) > 0 AS pay_password", TABLE_USERINFO), // pay_password
		fmt.Sprintf("%s.open_pay_password AS open_pay_password", TABLE_USERINFO),        // open_pay_password
		fmt.Sprintf("%s.daily_check_time AS daily_check_time", TABLE_USERINFO),          // daily_check_time
		fmt.Sprintf("%s.cont_check_days AS cont_check_days", TABLE_USERINFO),            // cont_check_days
		fmt.Sprintf("%s.idno AS idno", TABLE_IDCARD),                                    // idno
		fmt.Sprintf("%s.realname AS realname", TABLE_IDCARD),                            // idno
		fmt.Sprintf("%s.phone AS phone", TABLE_PHONE_USER),                              // phone
		fmt.Sprintf("%s.password AS login_password", TABLE_PHONE_USER),                  // login_password
		fmt.Sprintf("%s.openid AS qq_openid", TABLE_QQ_USER),                            // qq_openid
		fmt.Sprintf("%s.openid AS wx_openid", TABLE_WEIXIN_USER),                        // wx_openid
	}

	esql := fmt.Sprintf(`
		SELECT %s FROM ((((%s
		FULL JOIN %s ON %s.account_id=%d)
		FULL JOIN %s ON %s.account_id=%d)
		FULL JOIN %s ON %s.account_id=%d)
		FULL JOIN %s ON %s.account_id=%d)
		WHERE %s.account_id=%d`,
		strings.Join(columns, ","), TABLE_USERINFO,
		TABLE_IDCARD, TABLE_IDCARD, accountId,
		TABLE_PHONE_USER, TABLE_PHONE_USER, accountId,
		TABLE_QQ_USER, TABLE_QQ_USER, accountId,
		TABLE_WEIXIN_USER, TABLE_WEIXIN_USER, accountId,
		TABLE_USERINFO, accountId,
	)

	userInfo := &dbproto.UserInfo{}
	var valIcon, valNickname, valIdno, valRealname sql.NullString
	var valPhone, valLoginPwd, valQqOpenid, valWxOpenid sql.NullString
	var valSex sql.NullInt64
	var valPayPwd, valPayOpenPwd sql.NullBool
	var valInvitationCode sql.NullString
	var valDailyCheckTime, valContCheckDays sql.NullInt64
	values := []interface{}{
		&valIcon, &valSex, &valNickname,
		&userInfo.Exp, &userInfo.Level, &valInvitationCode,
		&valPayPwd, &valPayOpenPwd,
		&valDailyCheckTime, &valContCheckDays,
		&valIdno, &valRealname,
		&valPhone, &valLoginPwd, &valQqOpenid, &valWxOpenid,
	}

	if err := agt.dbConn.QueryRow(esql).Scan(values...); err != nil {
		log.Println(err, esql)
		return nil, err
	}

	if valIcon.Valid {
		userInfo.Icon = valIcon.String
	}

	if valSex.Valid {
		userInfo.Sex = int32(valSex.Int64)
	}

	if valNickname.Valid {
		userInfo.Nickname = valNickname.String
	}

	if valInvitationCode.Valid {
		userInfo.InvitationCode = valInvitationCode.String
	}

	if valPayPwd.Valid {
		userInfo.PayPassword = valPayPwd.Bool
	}

	if valPayOpenPwd.Valid {
		userInfo.PayOpenPassword = valPayOpenPwd.Bool
	}

	if valDailyCheckTime.Valid {
		userInfo.DailyCheckTime = valDailyCheckTime.Int64
	}

	if valContCheckDays.Valid {
		userInfo.ContCheckDays = int32(valContCheckDays.Int64)
	}

	if valIdno.Valid {
		userInfo.Idno = valIdno.String
	}

	if valRealname.Valid {
		userInfo.Realname = valRealname.String
	}

	if valPhone.Valid {
		userInfo.Phone = &dbproto.PhoneUser{Phone: valPhone.String, AccountId: accountId}
		if valLoginPwd.Valid {
			userInfo.Phone.Password = valLoginPwd.String
		}
	}

	if valQqOpenid.Valid {
		userInfo.Qq = &dbproto.QQUser{Openid: valQqOpenid.String, AccountId: accountId}
	}

	if valWxOpenid.Valid {
		userInfo.Weixin = &dbproto.WeixinUser{Openid: valWxOpenid.String, AccountId: accountId}
	}

	return userInfo, nil
}

func (agt *DbUsercenterAgent) SetUserIcon(ctx context.Context, arg *dbproto.UserInfoArg) (*dbproto.Nil, error) {
	icon := arg.GetIcon()
	accountId := arg.GetAccountId()
	esql := fmt.Sprintf("UPDATE %s SET icon='%s' WHERE account_id=%d", TABLE_USERINFO, icon, accountId)
	_, err := agt.dbConn.Exec(esql)
	if err != nil {
		log.Println(err, esql)
		return nil, err
	}
	return &dbproto.Nil{}, nil
}

func (agt *DbUsercenterAgent) SetUserNickname(ctx context.Context, arg *dbproto.UserInfoArg) (*dbproto.Nil, error) {
	nickname := arg.GetNickname()
	accountId := arg.GetAccountId()
	esql := fmt.Sprintf("UPDATE %s SET nickname='%s' WHERE account_id=%d", TABLE_USERINFO, nickname, accountId)
	_, err := agt.dbConn.Exec(esql)
	if err != nil {
		log.Println(err, esql)
		return nil, err
	}
	return &dbproto.Nil{}, nil
}

func (agt *DbUsercenterAgent) SetUserSex(ctx context.Context, arg *dbproto.UserInfoArg) (*dbproto.Nil, error) {
	sex := arg.GetSex()
	accountId := arg.GetAccountId()
	esql := fmt.Sprintf("UPDATE %s SET sex=%d WHERE account_id=%d", TABLE_USERINFO, sex, accountId)
	_, err := agt.dbConn.Exec(esql)
	if err != nil {
		log.Println(err, esql)
		return nil, err
	}
	return &dbproto.Nil{}, nil
}

func (agt *DbUsercenterAgent) SetAccountPayPassword(ctx context.Context, arg *dbproto.PayPasswordArg) (*dbproto.Nil, error) {
	pwd := arg.GetPassword()
	accountId := arg.GetAccountId()
	esql := fmt.Sprintf("UPDATE %s SET pay_password='%s', open_pay_password=%v WHERE account_id=%d ", TABLE_USERINFO, pwd, true, accountId)
	_, err := agt.dbConn.Exec(esql)
	if err != nil {
		log.Println(err, esql)
		return nil, err
	}
	return &dbproto.Nil{}, nil
}

func (agt *DbUsercenterAgent) OpenAccountPayPassword(ctx context.Context, arg *dbproto.PayPasswordArg) (*dbproto.Nil, error) {
	open := arg.GetOpen()
	accountId := arg.GetAccountId()
	esql := fmt.Sprintf("UPDATE %s SET open_pay_password=%v WHERE account_id=%d", TABLE_USERINFO, open, accountId)
	_, err := agt.dbConn.Exec(esql)
	if err != nil {
		log.Println(err, esql)
		return nil, err
	}
	return &dbproto.Nil{}, nil
}

func (agt *DbUsercenterAgent) QueryAccountPaySettings(ctx context.Context, arg *dbproto.IntValue) (*dbproto.PaySettings, error) {
	accountId := arg.GetValue()
	var pwd sql.NullString
	var isOpen sql.NullBool
	esql := fmt.Sprintf("SELECT pay_password, open_pay_password FROM %s WHERE account_id=%d", TABLE_USERINFO, accountId)
	ret := &dbproto.PaySettings{}
	if err := agt.dbConn.QueryRow(esql).Scan(&pwd, &isOpen); err != nil {
		if err != sql.ErrNoRows {
			log.Println(err, esql)
			return nil, err
		}
	}
	if pwd.Valid {
		ret.Password = pwd.String
	}
	if isOpen.Valid {
		ret.Open = isOpen.Bool
	}
	return ret, nil
}

func (agt *DbUsercenterAgent) SetDailyCheck(ctx context.Context, arg *dbproto.DailyCheckArg) (*dbproto.Nil, error) {
	accountId := arg.GetAccountId()
	checkTime := time.Now().Unix()
	exp := arg.GetExp()
	level := arg.GetLevel()
	contCheckDays := arg.GetContCheckDays()
	esql := fmt.Sprintf("UPDATE %s SET exp=%d, lvl=%d, daily_check_time=%d, cont_check_days=%d WHERE account_id=%d", TABLE_USERINFO, exp, level, checkTime, contCheckDays, accountId)
	_, err := agt.dbConn.Exec(esql)
	if err != nil {
		log.Println(err, esql)
		return nil, err
	}
	return &dbproto.Nil{}, nil
}
