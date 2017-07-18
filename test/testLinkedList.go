package main
import "fmt"
import "util"
func main() {
	//初始化一个头结点
	var h util.Node
	//往链表插入10个元素
	for i := 1; i <= 10; i++ {
		var d util.Node
		d.Data = i
		util.Insert(&h, &d, i)
		fmt.Println(util.GetLoc(&h, i))
	}

	fmt.Println(h.Next.Next.Data)
}