package user

import (
	"github.com/henrylee2cn/faygo"
	"qiniu-go/model/apiModel"
	//"time"
	"qiniu-go/util"
	"log"
)

/**********************发红包*********************/
var RedPocketMap = make(map[int]*apiModel.SendRedpocket)
//发红包
type SendRedPocket struct {
	Uid	 int	    `param:"<in:query><required> <desc:用户id>"`
	Num	 int	    `param:"<in:query><required> <desc:红包个数>"`
	Count	 float32    `param:"<in:query><required> <desc:红包总金额>"`
	Title	 string	    `param:"<in:query><required> <desc:红包标题>"`

}
/**
	发送红包的api
 */
func (t *SendRedPocket) Serve(ctx *faygo.Context) error {
	//1、加入到数据库中
	sendRedPocket := new(apiModel.SendRedpocket)
	//sendRedPocket.Uid = t.Uid
	//sendRedPocket.Count = t.Count
	//sendRedPocket.Num = t.Num
	//sendRedPocket.Title = t.Title
	//sendRedPocket.Time = time.Now().Unix()
	//util.ApiEngine.Insert(sendRedPocket)
	f,er := util.ApiEngine.Desc("id").Get(&sendRedPocket)
	log.Println(f)
	util.CheckError(er)
	log.Println(sendRedPocket)
	//2、
	return ctx.JSON(200,sendRedPocket,true)

}
func (t *SendRedPocket) Doc() faygo.Doc {
	return faygo.Doc{
		Note:   "发送红包api",
		Return: "返回红包数据的json",
	}
}
