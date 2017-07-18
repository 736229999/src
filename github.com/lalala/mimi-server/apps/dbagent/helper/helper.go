package helper

import (
	"errors"
	"fmt"
)

const (
	USER     = "postgres"
	PASSWORD = "Cp0422$()" //Cp0422$()
)

const (
	DB_UC       = "mimi-usercenter"
	DB_OPENCAI  = "mimi-opencai"
	DB_BUYCAI   = "mimi-buycai"
	DB_RECHARGE = "mimi-recharge"
	DB_ADMIN    = "mimi-admin"
	DB_DISCOVER = "mimi-discover"
	DB_NOTIF    = "mimi-notify"
	DB_OPTIONS  = "mimi-options"
	DB_THIRDAPI = "mimi-thirdapi"
)

func DataSourceName(dbname string, sslmode string) string {
	return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", USER, PASSWORD, dbname, sslmode)
}

var ErrInvalidArgs = errors.New("无效参数")
var ErrFundNotEnough = errors.New("资金不足")
var ErrNotFound = errors.New("not_found")
var ErrInvalidParam = errors.New("invalid_param")
var ErrDatabase = errors.New("database_error")
var ErrPhoneNotRegist = errors.New("手机号还未注册")
var ErrPhoneNotBind = errors.New("手机号还未绑定账号")
var ErrPhoneBinded = errors.New("手机号已绑定其它账号")
