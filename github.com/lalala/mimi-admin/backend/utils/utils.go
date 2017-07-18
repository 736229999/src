package utils

import (
	"database/sql"
	"fmt"
	"log"
	"regexp"
)

// 密码格式
// 以字母开头，长度在6~18之间，只能包含字符、数字和下划线
func IsPassword(pwd string) bool {
	reg := regexp.MustCompile(`^[a-zA-Z]\w{5,17}$`)
	return reg.MatchString(pwd)
}

// 是否是手机号
func IsMobile(pn string) bool {
	reg := regexp.MustCompile(`^(\+86)?\d{11}$`)
	return reg.MatchString(pn) || pn == ""
}

func IsEmail(e string) bool {
	reg := regexp.MustCompile(`^(\w)+([\.\-]\w+)*@(\w)+((\.\w{2,})+)$`)
	return reg.MatchString(e)
}

// name 长度是否合适
func IsName(name string) bool {
	return len(name) < 22 && len(name) > 0
}

func CheckErr(err error) bool {
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

/**
获取上一条id
@conn   数据库连接
@table   数据表名
*/
func GetLastId(conn *sql.DB, table string) int {
	var id int
	sql := fmt.Sprintf("SELECT id FROM %s ORDER BY id DESC LIMIT 1 ", table)
	st, err := conn.Prepare(sql)
	CheckErr(err)
	row := st.QueryRow()
	row.Scan(&id)
	return id
}
