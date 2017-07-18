package helper

import (
	"fmt"
)

const (
	USER     = "postgres"
	PASSWORD = "Cp0422$()"
)

const (
	DB_UC       = "mimi-usercenter"
	DB_OPENCAI  = "mimi-opencai"
	DB_RECHARGE = "mimi-recharge"
	DB_ADMIN    = "mimi-admin"
	DB_BUYCAI   = "mimi-buycai"
	DB_OPTIONS  = "mimi-options"
	DB_FOOTBALL = "mimi-football"
)

func DataSourceName(dbname string, sslmode string) string {
	return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", USER, PASSWORD, dbname, sslmode)
}
