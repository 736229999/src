package shop

import (
	"github.com/henrylee2cn/faygo"
	"qiniu-go/util"
	"qiniu-go/model/apiModel"
)

/*************************商品属性********************/




//增加属性
type AddAttr struct {
	Name     string     `param:"<in:query><required> <desc:属性名称>"`
}

/**
	添加商品属性api
 */
func (t *AddAttr) Serve(ctx *faygo.Context) error {

	attr := new(apiModel.ShopAttr)
	defer func() {
		attr = nil
	}()
	hasAttr ,e := util.Engine.Where("name=?",t.Name).Get(attr)
	util.CheckError(e)
	if hasAttr {
		return ctx.JSON(200,"该属性已经存在",true)
	}

	attr.Name = t.Name
	_,err := util.Engine.Insert(attr)
	util.CheckError(err)
	return ctx.JSON(200,"添加属性成功",true)
}
func (t *AddAttr) Doc() faygo.Doc {
	return faygo.Doc{
		Note:   "添加商品属性",
		Return: "添加成功与否",
	}
}

//删除商品属性
type DelAttr struct {
	Id     int     `param:"<in:query><required> <desc:属性名称>"`
}

/**
	删除商品属性api
 */
func (t *DelAttr) Serve(ctx *faygo.Context) error {

	attr := new(apiModel.ShopAttr)
	defer func() {
		attr = nil
	}()
	hasAttr ,e := util.Engine.Where("id=?",t.Id).Get(attr)
	util.CheckError(e)
	if hasAttr {
		_,err := util.Engine.Id(t.Id).Delete(attr)
		util.CheckError(err)
		return ctx.JSON(200,"删除商品属性成功",true)
	}

	return ctx.JSON(200,"删除失败，该属性不存在",true)
}
func (t *DelAttr) Doc() faygo.Doc {
	return faygo.Doc{
		Note:   "删除商品属性",
		Return: "删除商品成功与否",
	}
}
