package model
//这里面的所有struct都不在数据库中，只是单独抽离出来
type distributeRule struct {
	level1 float32	//一级分销对应的分成百分比
	level2 float32
	level3 float32
}


