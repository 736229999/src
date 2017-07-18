package shop

import (
	"github.com/henrylee2cn/faygo"
	"qiniu-go/model/apiModel"
	"qiniu-go/util"
	"time"
)

/***********************购物车********************/

//增加购物车
type AddIntoCart struct {
	Uid     int     `param:"<in:query><required> <desc:用户id>"`
	Goods_id     int     `param:"<in:query><required> <desc:商品id>"`
	Goods_num     int     `param:"<in:query><required> <desc:商品数量>"`
	Attr_id     string     `param:"<in:query><required> <desc:属性id，多个属性用逗号隔开，eg:1,3,4>"`
}

/**
	添加商品到购物车api
 */
func (t *AddIntoCart) Serve(ctx *faygo.Context) error {
	shopCart := new(apiModel.ShopShopCart)
	defer func() {
		shopCart = nil
	}()

	shopCart.Uid = t.Uid
	shopCart.Goods_id = t.Goods_id
	shopCart.Goods_num = t.Goods_num
	shopCart.Attr_id = t.Attr_id
	shopCart.Time = time.Now().Unix()
	_,err := util.ApiEngine.Insert(shopCart)
	if err != nil{
		return ctx.JSON(200,"添加属性失败，检查下是不是数据格式没对，或者联系后台",true)
	}


	return ctx.JSON(200,"添加属性成功",true)
}
func (t *AddIntoCart) Doc() faygo.Doc {
	return faygo.Doc{
		Note:   "添加商品到购物车",
		Return: "添加成功与否",
	}
}

//删除购物车中商品
type DeleteFromCart struct {
	Uid     int     `param:"<in:query><required> <desc:用户id>"`
	Goods_id     int     `param:"<in:query><required> <desc:商品id>"`
	Goods_num     int     `param:"<in:query><required> <desc:商品数量>"`
	Attr_id     string     `param:"<in:query><required> <desc:属性id，多个属性用逗号隔开，eg:1,3,4>"`
}

/**
	删除购物车商品api
 */
func (t *DeleteFromCart) Serve(ctx *faygo.Context) error {
	shopCart := new(apiModel.ShopShopCart)
	defer func() {
		shopCart = nil
	}()

	shopCart.Uid = t.Uid
	shopCart.Goods_id = t.Goods_id
	shopCart.Goods_num = t.Goods_num
	shopCart.Attr_id = t.Attr_id
	shopCart.Time = time.Now().Unix()
	_,err := util.ApiEngine.Insert(shopCart)
	if err != nil{
		return ctx.JSON(200,"添加属性失败，检查下是不是数据格式没对，或者联系后台",true)
	}


	return ctx.JSON(200,"添加属性成功",true)
}
func (t *DeleteFromCart) Doc() faygo.Doc {
	return faygo.Doc{
		Note:   "添加商品到购物车",
		Return: "添加成功与否",
	}
}