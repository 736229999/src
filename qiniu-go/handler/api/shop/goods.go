package shop
//
import (
	"qiniu-go/util"
	"github.com/henrylee2cn/faygo"
	//"log"
	"time"
	"qiniu-go/model/apiModel"
	"strconv"
	"log"
)


//增加商品
type AddGoods struct {
	Uid	 string	    `param:"<in:query><required> <desc:商品拥有者id>"`
	Name     string     `param:"<in:query><required> <desc:商品名称>"`
	Cat_id   string        `param:"<in:query><required> <desc:分类id>"`
	//Pre_price   float32        `param:"<in:query> <desc:以前的价格>"`
	Now_price   float32        `param:"<in:query><required> <desc:现在的价格>"`
	Small_logo   string        `param:"<in:query><required> <desc:商品小图>"`
	//Big_logo   string        `param:"<in:query> <desc:商品大图>"`
	Des   string        `param:"<in:query><required> <desc:商品描述>"`
	Weight   float32        `param:"<in:query> <desc:商品重量>"`
	Is_postfee   string        `param:"<in:query> <desc:是否包邮，1：包邮；0：不包邮>"`
	//Attr_id   string        `param:"<in:query> <desc:属性id，多个属性用英文逗号隔开>"`
	Goods_num   string        `param:"<in:query> <desc:库存量>"`

}
func (t *AddGoods) Serve(ctx *faygo.Context) error {
	goods := new(apiModel.ShopGoods)
	//goodsAttr := new(apiModel.Goodsattr)
	goodsCategory := new(apiModel.ShopGoodsCategory)
	defer func() {
		goods = nil
		//goodsAttr = nil
		goodsCategory = nil
	}()

	goods.Uid ,_= strconv.Atoi(t.Uid)
	goods.Name = t.Name
	goods.Cat_id ,_= strconv.Atoi(t.Cat_id)
	//goods.Pre_price = t.Pre_price
	goods.Now_price = t.Now_price
	goods.Small_logo = t.Small_logo
	//goods.Big_logo = t.Big_logo
	goods.Des = t.Des
	goods.Weight = t.Weight
	i ,_:=  strconv.Atoi(t.Is_postfee)
	goods.Is_postfee = int8(i)
	goods.Num,_ = strconv.Atoi(t.Goods_num)
	goods.Is_onsale = 1
	goods.Addtime = time.Now().Unix()
	//再加个商品编号
	goods.Serialnum = util.RandSeq(6)
	_,err :=util.ApiEngine.Insert(goods)
	util.CheckError(err)


	_,e1 := util.ApiEngine.Desc("addtime").Limit(1).Get(goods)
	util.CheckError(e1)


	////goodsAttr.Attr_id = t.Attr_id
	//goodsAttr.Goods_id = goods.Id
	//goodsAttr.Goods_num = t.Goods_num
	//_,e := util.ApiEngine.Insert(goodsAttr)
	//util.CheckError(e)


	//还有商品分类_商品
	goodsCategory.Goods_id = goods.Id
	goodsCategory.Category_id ,_= strconv.Atoi(t.Cat_id)
	util.ApiEngine.Insert(goodsCategory)

	return ctx.JSON(200,util.GetResponseMap(200,"添加商品成功"), true)
}
func (t *AddGoods) Doc() faygo.Doc {
	return faygo.Doc{
		Note:   "添加商品",
		Return: "添加成功与否",
	}
}




//显示当前销售中的商品
type GetOnsaleGoods struct {
	Uid	 string	    `param:"<in:query><required> <desc:商家id>"`
}
func (t *GetOnsaleGoods) Serve(ctx *faygo.Context) error {
	log.Println(t.Uid)
	onsaleGoods := make([]apiModel.ShopGoods, 0)
	defer func() {
		onsaleGoods = nil
	}()
	uid,_ := strconv.Atoi(t.Uid)
	err := util.ApiEngine.Cols("now_price", "name","small_logo","des","num").Where("uid=? and is_onsale=?", uid,1).Find(&onsaleGoods)
	util.CheckError(err)
	if len(onsaleGoods) == 0 {
		log.Println(t)
		return ctx.JSON(200,util.GetResponseMap(400,"没有在售的商品"), true)
	}
	return ctx.JSON(200,util.GetResponseMap(200,onsaleGoods),true)
}
func (t *GetOnsaleGoods) Doc() faygo.Doc {
	return faygo.Doc{
		Note:   "显示所有上架的商品",
		Return: "name：商品名称，now_price：商品现价，small_logo：商品图片，des：商品描述，num：商品库存",
	}
}

//显示已经下架的商品
type GetDownSaleGoods struct {
	Uid	 string	    `param:"<in:query><required> <desc:商家id>"`
}
func (t *GetDownSaleGoods) Serve(ctx *faygo.Context) error {
	downSaleGoods := make([]apiModel.ShopGoods, 0)
	defer func() {
		downSaleGoods = nil
	}()
	uid,_ := strconv.Atoi(t.Uid)
	err := util.ApiEngine.Cols("now_price", "name","small_logo","des","num").Where("uid=? and is_onsale=?", uid,0).Find(&downSaleGoods)
	util.CheckError(err)
	if len(downSaleGoods) == 0 {
		return ctx.JSON(200,util.GetResponseMap(400,"没有下架的商品"), true)
	}
	return ctx.JSON(200,util.GetResponseMap(200,downSaleGoods), true)
}
func (t *GetDownSaleGoods) Doc() faygo.Doc {
	return faygo.Doc{
		Note:   "显示所有下架的商品",
		Return: "name：商品名称，now_price：商品现价，small_logo：商品图片，des：商品描述，num：商品库存",
	}
}

