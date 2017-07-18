//此文件中有两个路由：1、绑定代理人关系  2、三级分销

//三级分销
package service

import (
	"github.com/henrylee2cn/faygo"
	"fmt"
	"qiniu-go/util"
	"time"
)
//代理人表
type Agent struct {
	Uid  int   `xorm:"  unique 'uid'"`
	Pid  int
	Time int64  `xorm:"   unique 'time'"`
}

type Three_distribute struct {
	Uid int  `param:"<in:query> <required> <desc:用户id> "`
	Money float32 `param:"<in:query> <required> <desc:需要分销的钱> "`
}

var level int = 1			//用于递归中计算层数
var everyLevelGetMoney = []*getMoney{}	//用于返回所有的分成信息
var levelMoneyMap = make(map[int]float32)//各个等级对应的分成百分比

type getMoney struct {
	Money float32
	Content string
	Uid   int
}

//初始化各个等级对应的分成百分比
func init(){
	levelMoneyMap[1] = 0.4
	levelMoneyMap[2] = 0.3
	levelMoneyMap[3] = 0.2
}

/**
   三级分销的路由
 */
func (t *Three_distribute) Serve(ctx *faygo.Context) error {
	a:=calculateDistribute(t.Uid,t.Money)
	everyLevelGetMoney = nil
	return ctx.JSON(200,a ,true)
}



/**
   计算三级分销
 */
func calculateDistribute(uid int,money float32) interface{}{
	//思路：递归取得每一级对应的分成，然后再返回所有人（最多3个）获取的钱的详情json
	agent := new(Agent)
	responseData := new(getMoney)
	hasUser,_ := util.Engine.Where("uid=?",uid).Get(agent)
	if !hasUser{
		fmt.Println("肯定是哪个删了数据库")
		agent = nil
		responseData = nil
		return everyLevelGetMoney
	}

	//递归的两个出口
	//1、对应的pid为0 ; 2、分销的层数超过3级
	if agent.Pid == 0 || level > 3{
		level = 1
		agent = nil
		responseData = nil
		return everyLevelGetMoney
	}
	fmt.Println(agent.Uid,"的上级:",agent.Pid)

	responseData.Money = money*levelMoneyMap[level]
	responseData.Uid = agent.Pid
	responseData.Content = "对应等级分的钱"
	everyLevelGetMoney = append(everyLevelGetMoney,responseData)
	level++
	return calculateDistribute(agent.Pid,money)

}

/**
   三级分销api的文档
 */
func (t *Three_distribute) Doc() faygo.Doc {
	return faygo.Doc{
		Note:   "三级分销",
		Return: "json格式数据",
	}
}


//代理人关系
type BindAgentRelation struct {
	Uid int  `param:"<in:query> <required> <desc:用户id> "`
	Pid int  `param:"<in:query> <required> <desc:上级id，如果是直接通过平台注册进来的，pid为0> "`

}

/**
   绑定代理人关系
 */
func (b *BindAgentRelation) Serve(ctx *faygo.Context) error {
	//思路：直接注册进代理人关系表，uid唯一
	agent := new(Agent)

	hasAgent,_ := util.Engine.Where("uid=?",b.Uid).Get(agent)
	if hasAgent == true{
		return ctx.JSON(400,"false", true)
	}else {
		agent.Uid = b.Uid
		agent.Pid = b.Pid
		agent.Time = time.Now().Unix()
		fmt.Println(time.Now().Unix())
		util.Engine.Insert(agent)
		return ctx.JSON(200,"success", true)
	}


}


/**
   绑定代理人关系api的文档
 */
func (t *BindAgentRelation) Doc() faygo.Doc {
	return faygo.Doc{
		Note:   "绑定代理人关系",
		Return: "绑定成功与否",
	}
}
