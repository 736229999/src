package apiModel
/**************************用户交互的model***********************/


//发红包orm
type SendRedpocket struct {
	Id int
	Uid int
	Title string
	Num int
	Count float32
	Time int64
}
