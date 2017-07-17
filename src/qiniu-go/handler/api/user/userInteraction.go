package user

import (
	"github.com/henrylee2cn/faygo"
	"qiniu-go/model"
	"qiniu-go/util"
	"time"
)

/**********************用户交互的一些api 包括关注等********************/


type GetUserAttention struct {
	Uid	 int	    `param:"<in:query><required> <desc:用户id>"`
}
/**
	获取关注了自己的所有 uid
 */
func (t *GetUserAttention) Serve(ctx *faygo.Context) error {

	//

	var userAttention []*model.Attention
	if userAttention == nil{
		userAttention = make([]*model.Attention,0)
	}
	err := util.Engine.Where("attention_id=?",t.Uid).Find(&userAttention)
	util.CheckError(err)
	if len(userAttention) == 0 {
		return ctx.JSON(200,"没有人关注",true)
	}

	return ctx.JSON(200,userAttention,true)

}
func (t *GetUserAttention) Doc() faygo.Doc {
	return faygo.Doc{
		Note:   "获取关注了自己的人",
		Return: "返回所有关注了自己的uid",
	}
}

type AttentionUser struct {
	Uid	 int	    `param:"<in:query><required> <desc:用户id>"`
	Attention_id int    `param:"<in:query><required> <desc:被关注人的id>"`
}
/**
	关注某人
 */
func (t *AttentionUser) Serve(ctx *faygo.Context) error {

	//
	var attentionUser *model.Attention
	if attentionUser == nil {
		attentionUser = new(model.Attention)
	}
	//判断是否已经关注
	hasAttentioned ,_ := util.Engine.Where("user_id=? and attention_id=?",t.Uid,t.Attention_id).Get(attentionUser)
	if hasAttentioned {
		return ctx.JSON(200,"你已经关注了",true)
	}

	attentionUser.User_id = t.Uid
	attentionUser.Attention_id = t.Attention_id
	attentionUser.Time = time.Now().Unix()
	flag,err := util.Engine.Insert(attentionUser)
	util.CheckError(err)
	if flag<0{
		return ctx.JSON(200,"关注失败，服务器的问题",true)
	}
	return ctx.JSON(200,"关注成功",true)

}
func (t *AttentionUser) Doc() faygo.Doc {
	return faygo.Doc{
		Note:   "关注某人",
		Return: "成功与否",
	}
}



//=========================
type HasAttentionOther struct {
	Uid	 int	    `param:"<in:query><required> <desc:用户id>"`
	Attention_id int    `param:"<in:query><required> <desc:被关注人的id>"`
}
/**
	判断是否关注
 */
func (t *HasAttentionOther) Serve(ctx *faygo.Context) error {
	hasAttention := new(model.Attention)
	defer func() {
		hasAttention = nil
	}()
	//判断是否已经关注
	flag ,_ := util.Engine.Where("user_id=? and attention_id=?",t.Uid,t.Attention_id).Get(hasAttention)
	if flag {
		return ctx.JSON(200,"true",true)
	}

	return ctx.JSON(200,"false",true)

}
func (t *HasAttentionOther) Doc() faygo.Doc {
	return faygo.Doc{
		Note:   "判断uid是否关注attention_id",
		Return: "true or false",
	}
}