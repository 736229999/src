package main

import (
	"fmt"
	"reflect"
)

type roles struct {
	roleId   int
	roleName string
}
type User struct {
	Name     string
	Age      int
	Email    string
	NickName string
	Telphone int
	Roles    roles
}

func main() {
	u := User{Name: "Name", Age: 30, Email: "101@afanty3d.com", NickName: "omni360", Telphone: 123, Roles: roles{roleId: 1001, roleName: "administrator"}}
	fmt.Println(u)
	Info(u)

}

func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())

	v := reflect.ValueOf(o)
	fmt.Println("Fileds:")
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s : %v %v\n", f.Name, f.Type, val)

	}
}
