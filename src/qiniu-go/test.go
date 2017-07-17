package main

import (

	"qiniu-go/util"
)

func test()  {
	sql :="insert into user(id,name) values (?, ?)"
	_, err := util.Engine.Exec(sql,12123,"lalla你好111")
	util.CheckError(err)
}