package main

import (
	"fmt"
	"time"
)

var c  = make(chan bool)
var i int = 0
func main()  {



	startTimer(say)
	fmt.Println(12)
}
func say()  {
	fmt.Println("say函数")
}
func startTimer(f func()) {
	go func() {
		for {
			f()
			now := time.Now()
			// 计算下一个零点
			next := now.Add(time.Second*10)
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			fmt.Println("正在执行")
			<-t.C
		}
	}()
}