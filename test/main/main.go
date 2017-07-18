package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	//"github.com/xormplus/core"
	"github.com/xormplus/xorm"
)

var PostEngine *xorm.Engine

func init() {

	if PostEngine == nil {
		var err error
		//PostEngine, err = xorm.NewEngine("postgres", "jerry:root@(127.0.0.1:5432);dbname=pgdb?charset=utf8")
		PostEngine, err = xorm.NewPostgreSQL("postgres://jerry:root@localhost:5432/testpg?sslmode=disable")

		if err != nil {
			log.Println(err)
		}
	}
	//tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "sh_")
	//PostEngine.SetTableMapper(tbMapper)
	PostEngine.ShowSQL(true)
}
func main() {
	add()
	add()

}

func add(){
	t := new(Qwe)
	t.Name = "你好"
	_,err := PostEngine.Insert(t)
	log.Println(err)

	//PostEngine.Desc("id").Limit(1).Get(t)
	//sel(t.Id)
}

func delete()  {
	var t User
	//t.Id = 0
	_,err :=PostEngine.Id(0).Where("id=?",0).Unscoped().Delete(&t)
	log.Println(err)
}
func sel(id int)  {
	var user User
	PostEngine.Where("id=?",id).Get(&user)
	log.Println(user)
}
type User struct {
	Id   int
	Name string
}
type Qwe struct {
	Id int
	Name string
}