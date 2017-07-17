package shop

import (
	"github.com/henrylee2cn/faygo"
	"qiniu-go/util"
	"log"
	"qiniu-go/model/apiModel"
)


/**************商品类别，包括增删改查，以及无限极分类************/



//增加类别
type AddCategory struct {
	Name     string     `param:"<in:query><required> <desc:类别名称>"`
	Pid	 int	    `param:"<in:query><required><range:0:100> <desc:上级分类的id，0：顶级>"`
}

/**
	添加类别api
 */
func (t *AddCategory) Serve(ctx *faygo.Context) error {
	gcategory := new(apiModel.ShopGcategory)
	defer func() {
		gcategory = nil
	}()
	hasCategory ,e := util.ApiEngine.Where("name=?",t.Name).Get(gcategory)
	util.CheckError(e)
	if hasCategory {
		return ctx.JSON(200,util.GetResponseMap(400,"该分类已存在"),true)
	}

	gcategory.Name = t.Name
	gcategory.Pid = t.Pid
	_,err := util.ApiEngine.Insert(gcategory)

	util.CheckError(err)

	return ctx.JSON(200,util.GetResponseMap(200,"添加分类成功"),true)
}

//删除类别
type DelCategory struct {
	Id	 int	    `param:"<in:query><required> <desc:要删除的类别id>"`
}
/**
	删除类别api
 */
func (t *DelCategory) Serve(ctx *faygo.Context) error {
	gcategory := new(apiModel.ShopGcategory)
	defer func() {
		gcategory = nil
	}()
	hasCategory ,_ := util.Engine.Where("id=?",t.Id).Get(gcategory)
	if hasCategory {
		_ ,err := util.Engine.Where("id=?",t.Id).Delete(gcategory)
		util.CheckError(err)
		return ctx.JSON(200,"删除类别成功", true)
	}
	log.Println(gcategory)

	return ctx.JSON(200,"删除类别失败，该类别不存在", true)
}

//修改类别
type ModifyCategory struct {
	Name     string     `param:"<in:query><required> <desc:类别名称>"`
}

func (t *ModifyCategory) Serve(ctx *faygo.Context) error {

	return ctx.JSON(200, t, true)
}

//查询所有分类（无限极分类）
type AllCategory struct {

}

func (t *AllCategory) Serve(ctx *faygo.Context) error {
	// 这儿的无限极分类是调用wechat那边php写的
	uri := util.JidiPhpApi+"Front/Util/goGetAllCategory"
	allCategory := util.CurlGetReturnString(uri)
	log.Println(allCategory)
	return ctx.String(200,allCategory)
}


func (t *AddCategory) Doc() faygo.Doc {
	return faygo.Doc{
		Note:   "增加类别",
		Return: "成功与否",
	}
}

func (t *DelCategory) Doc() faygo.Doc {
	return faygo.Doc{
		Note:   "删除类别",
		Return: "成功与否",
	}
}
func (t *AllCategory) Doc() faygo.Doc {
	return faygo.Doc{
		Note:   "获取所有商品分类",
		Return: "所有商品分类的树状json",
	}
}

