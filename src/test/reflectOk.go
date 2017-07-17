package main
import "fmt"
import (
	"reflect"
	"log"
)

type S struct {
	i int
}

func (s *S) MyMissionMethod(a int64) (bool){
	fmt.Println("传过来的值：",a)
	return false
}

func HasCompleteTask(method interface{},parm int64)(interface{}){
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	fv := reflect.ValueOf(method)
	args := []reflect.Value{reflect.ValueOf(parm)}
	a := fv.Call(args)
	return a[0].Interface().(bool)
}
func main() {
	s := &S{}
	mission := s.MyMissionMethod
	a := HasCompleteTask(mission,1)
	fmt.Println("返回回来的值：",a)

}