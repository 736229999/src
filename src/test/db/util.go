package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"

	"github.com/xormplus/core"
)

var PostEngine *xorm.Engine

func init() {
	if PostEngine == nil {
		var err error
		PostEngine, err = xorm.NewEngine("mysql", "root:root@(127.0.0.1:3306)/db_bill?charset=utf8")
		CheckError(err)
	}
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "sh_")
	PostEngine.SetTableMapper(tbMapper)
	PostEngine.ShowSQL(false)
}
