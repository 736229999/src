package shop

import (
	"github.com/henrylee2cn/faygo"
	"qiniu-go/model/apiModel"
	"time"
	"qiniu-go/util"
	"log"
)


//架构：在生成订单的时候，如果以后并发量高，需要对订单编号提前生成一个   订单号池，这样就可以直接在池里面去取订单号



//创建订单
type CreateOrder struct {
	BuyerId     int     `param:"<in:query><required> <desc:买家id>"`
	SalerId	    int	    `param:"<in:query><required> <desc:卖家id>"`
	GoodsId     int     `param:"<in:query><required> <desc:商品id>"`
	GoodsNum     int     `param:"<in:query><required><range:1:10000> <desc:购买的商品数量(必须大于1)>"`
	Post_id     int8     `param:"<in:query><required> <desc:0：顺丰，1：申通，2：圆通，>"`
	Pay_id     int8     `param:"<in:query><required> <desc:0：微信；1：支付宝>"`
	Goods_tprice     float32     `param:"<in:query><required> <desc:商品总价（不带邮费）>"`
	Post_fee     float32     `param:"<in:query><required> <desc:邮费>"`
	Receipt_id     int     `param:"<in:query><range:0:10000><required> <desc:收货地址id>"`

}

/**
	创建订单的api
 */
func (t *CreateOrder) Serve(ctx *faygo.Context) error {

	//1、创建一条订单数据	2、创建一条账单数据
	orderv1 := new(apiModel.ShopOrderv1)
	goods := new(apiModel.ShopGoods)
	defer func() {
		orderv1 = nil
		goods = nil
	}()
	flag ,_:=util.ApiEngine.Where("id=?",t.GoodsId).Get(goods)
	if !flag {
		return ctx.JSON(200,util.GetResponseMap(400,"创建订单失败，没有这个商品"),true)
	}

	orderv1.Buyer_id = t.BuyerId
	orderv1.Saler_id = t.SalerId
	orderv1.Goods_id = t.GoodsId
	orderv1.Goods_num = t.GoodsNum
	orderv1.Post_id = t.Post_id
	orderv1.Pay_id = t.Pay_id
	orderv1.Goods_tprice = t.Goods_tprice
	orderv1.Post_fee = t.Post_fee
	orderv1.Receipt_id = t.Receipt_id
	orderv1.Order_sn = util.RandSeq(6)
	orderv1.Add_time = time.Now().Unix()
	orderv1.Update_time = time.Now().Unix()

	_,err := util.ApiEngine.Insert(orderv1)
	if err != nil {
		return ctx.JSON(200,util.GetResponseMap(400,"创建订单失败，问后台"),true)
	}

	//同时还会再生成一张账单
	return ctx.JSON(200,util.GetResponseMap(200,"创建订单成功"),true)
}

func (t *CreateOrder) Doc() faygo.Doc {
	return faygo.Doc{
		Note:   "创建订单",
		Return: "成功与否",
	}
}

type returnData struct {
	OrderId int
	GoodsName string
	OrderCreateTime int64
	GoodsNum int
	TotalFee float32
	GoodsImg string
	OrderStatus int8
	PostStatus int8
}

//用户的所有订单
type UserAllOrder struct {
	Uid     int     `param:"<in:query><required> <desc:用户id>"`

}

/**
	用户的所有订单的api
 */
func (t *UserAllOrder) Serve(ctx *faygo.Context) error {
	userAllOrder := make([]apiModel.ShopOrderv1, 0)
	goods := new(apiModel.ShopGoods)
	returnDatas := []returnData{}
	res := new(returnData)
	returnMap := make(map[string]interface{})
	defer func() {
		userAllOrder = nil
		goods = nil
		returnMap = nil
		returnDatas = nil
		res = nil
	}()

	err := util.ApiEngine.Cols("id","goods_id","order_status","post_status","total_fee","add_time").Where("buyer_id=? ",t.Uid).Find(&userAllOrder)
	if err != nil {
		return ctx.JSON(200,util.GetResponseMap(400,"数据库错误，联系后台"), true)
	}
	if len(userAllOrder) == 0 {
		return ctx.JSON(200,util.GetResponseMap(400,"没有订单"), true)
	}
	if len(userAllOrder) > 30 {
		return ctx.JSON(200,util.GetResponseMap(400,"系统设置的订单最多为30，找一下后台"), true)
	}

	returnMap["code"] = 200

	for k,v := range userAllOrder{
		log.Println(k)
		//查出所有的商品
		util.ApiEngine.Where("id=?",v.Goods_id).Get(goods)
		res.OrderId = v.Id
		res.GoodsName = goods.Name
		res.OrderCreateTime = v.Add_time
		res.GoodsNum = v.Goods_num
		res.TotalFee = v.Total_fee
		res.GoodsImg = goods.Small_logo
		res.OrderStatus = v.Order_status
		res.PostStatus = v.Post_status
		returnDatas = append(returnDatas,*res)
	}
	returnMap["data"] = returnDatas
	return ctx.JSON(200,returnMap,true)
}
func (t *UserAllOrder) Doc() faygo.Doc {
	return faygo.Doc{
		Note:   "用户所有订单",
		Return: "GoodsName：商品名，OrderCreateTime：订单创建时间，GoodsNum：商品数量， TotalFee：总价（带邮费）GoodsImg：商品图片，OrderStatus：订单状态 0：交易中，1：交易完成，2：交易取消" +
			"Post_status：运送状态：0：未发货，1：已发货，2：已收货，3：退款中，4：已退款",
	}
}


//商家的所有订单
type SalerAllOrders struct {
	SalerId     int     `param:"<in:query><required> <desc:商家id>"`

}


/**
	商家的所有订单的api
 */
func (t *SalerAllOrders) Serve(ctx *faygo.Context) error {
	userAllOrder := make([]apiModel.ShopOrderv1, 0)
	goods := new(apiModel.ShopGoods)
	returnDatas := []returnData{}
	res := new(returnData)
	returnMap := make(map[string]interface{})
	defer func() {
		userAllOrder = nil
		goods = nil
		returnMap = nil
		returnDatas = nil
		res = nil
	}()
	
	err := util.ApiEngine.Cols("id","goods_id","order_status","post_status","total_fee","add_time").Where("saler_id=? ",t.SalerId).Find(&userAllOrder)
	if err != nil {
		return ctx.JSON(200,util.GetResponseMap(400,"数据库错误，联系后台"), true)
	}
	if len(userAllOrder) == 0 {
		return ctx.JSON(200,util.GetResponseMap(400,"没有订单"), true)
	}
	if len(userAllOrder) > 30 {
		return ctx.JSON(200,util.GetResponseMap(400,"系统设置的订单最多为30，找一下后台"), true)
	}

	returnMap["code"] = 200

	for k,v := range userAllOrder{
		log.Println(k)
		//查出所有的商品
		util.ApiEngine.Where("id=?",v.Goods_id).Get(goods)
		res.OrderId = v.Id
		res.GoodsName = goods.Name
		res.OrderCreateTime = v.Add_time
		res.GoodsNum = v.Goods_num
		res.TotalFee = v.Total_fee
		res.GoodsImg = goods.Small_logo
		res.OrderStatus = v.Order_status
		res.PostStatus = v.Post_status
		returnDatas = append(returnDatas,*res)
	}
	returnMap["data"] = returnDatas
	return ctx.JSON(200,returnMap,true)
}
func (t *SalerAllOrders) Doc() faygo.Doc {
	return faygo.Doc{
		Note:   "商家所有订单",
		Return: "GoodsName：商品名，OrderCreateTime：订单创建时间，GoodsNum：商品数量，TotalFee：总价（带邮费）GoodsImg：商品图片，OrderStatus：订单状态 0：交易中，1：交易完成，2：交易取消" +
			"Post_status：运送状态：0：未发货，1：已发货，2：已收货，3：退款中，4：已退款",
	}
}


//交易完成订单详情
type OrderDetail struct {
	OrderId     int     `param:"<in:query><required> <desc:商家id>"`

}

type orderDetail struct {
	SalerId int		//
	SalerName string	//
	GoodsName string	//
	GoodsDesribe string	//
	GoodsPrice float32	//
	GoodsNum int		//
	BuyerId int		//
	BuyerName string	//
	BuyerPhone int		//
	BuyerAddr string

	OrderId int		//
	OrderSn string		//
	CreateOrderTime int64	//
	PostStatus int8
	PayId int8
	TotalFee float32
	PostId int8
	PostSn string



}

/**
	订单详情的api
 */
func (t *OrderDetail) Serve(ctx *faygo.Context) error {
	order := new(apiModel.ShopOrderv1)
	goods := new(apiModel.ShopGoods)
	user := new(apiModel.User)
	receiptAddr := new(apiModel.ShopReceiptAddr)
	orderdetail := new(orderDetail)
	defer func() {
		order = nil
		goods = nil
		orderdetail = nil
		receiptAddr = nil
		user = nil
	}()

	//先查订单
	flag,err:=util.ApiEngine.Where("id=?",t.OrderId).Get(order)
	util.CheckError(err)
	if !flag {
		return ctx.JSON(200,util.GetResponseMap(400,"该订单不存在"),true)
	}
	orderdetail.OrderId = t.OrderId
	orderdetail.OrderSn = order.Order_sn
	orderdetail.CreateOrderTime = order.Add_time
	orderdetail.PostStatus = order.Post_status
	orderdetail.PayId = order.Pay_id
	orderdetail.TotalFee = order.Total_fee
	orderdetail.PostId = order.Post_id
	orderdetail.PostSn = "物流单号还没与物流公司对接"
	orderdetail.SalerId = order.Saler_id
	orderdetail.BuyerId = order.Buyer_id
	orderdetail.GoodsNum = order.Goods_num
	orderdetail.GoodsPrice = order.Goods_tprice

	//商家名称
	util.ApiEngine.Where("id=?",order.Saler_id).Cols("name").Get(user)
	orderdetail.SalerName = user.Name
	//商品信息
	util.ApiEngine.Where("id=?",order.Goods_id).Cols("name","des").Get(goods)
	orderdetail.GoodsName = goods.Name
	orderdetail.GoodsDesribe = goods.Des
	//收货地址信息
	util.ApiEngine.Where("id=?",order.Receipt_id).Get(receiptAddr)
	orderdetail.BuyerName = receiptAddr.Name
	orderdetail.BuyerPhone = receiptAddr.Phone
	orderdetail.BuyerAddr = receiptAddr.Province+"省"+receiptAddr.City+"市"+receiptAddr.District+"区"+receiptAddr.Detail_addr

	return ctx.JSON(200,util.GetResponseMap(200,orderdetail),true)
}
func (t *OrderDetail) Doc() faygo.Doc {
	return faygo.Doc{
		Note:   "订单详情",
		Return: "GoodsName：商品名，OrderCreateTime：订单创建时间，GoodsNum：商品数量，TotalFee：总价（带邮费）GoodsImg：商品图片，OrderStatus：订单状态 0：交易中，1：交易完成，2：交易取消" +
			"Post_status：运送状态：0：未发货，1：已发货，2：已收货，3：退款中，4：已退款",
	}
}