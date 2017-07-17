package main

import (
	"fmt"
	"qiniu-go/util"
)

func main()  {
	a:=util.CurlPost("http://www.jiditv.com/OpenAPI/V1/Ticket/jd_Anchor?uid=711","name=123")
	fmt.Println(a)
}
