package apiModel

/**************************商城里面包括类别，商品属性等orm**********************/

//商品类别的orm
type ShopGcategory struct {
	Id int
	Name string
	Pid int
	//Children *Goodscategory `xorm:"-"`
}
//商品_类别中间表orm
type ShopGoodsCategory struct {
	Goods_id int
	Category_id int
}

//商品orm
type ShopGoods struct {
	Id int		`json:"id"`
	Uid int		`json:"uid"`
	Name string	`json:"name"`
	Serialnum string	`json:"serialnum"`	//编号
	Cat_id int	`json:"code"`
	//Pre_price float32	`json:"code"`
	Now_price float32	`json:"now_price"`
	Small_logo string	`json:"small_logo"`
	//Big_logo string
	Des string	`json:"des"`
	Weight float32	`json:"weight"`
	Is_onsale int8	`json:"is_onsale"`
	Is_postfee int8	`json:"is_postfee"`
	Is_delete int8	`json:"is_delete"`
	Num int		`json:"num"`
	Addtime int64	`json:"addtime"`

}
//属性orm
type ShopAttr struct {
	Id int
	Name string
}

//商品_属性中间表orm
type ShopGoodsattr struct {
	Goods_id int
	Attr_id string	//存的是属性的冗余，多个属性用逗号隔开
	Goods_num int   //库存
}

//购物车orm
type ShopShopCart struct {
	Uid int
	Goods_id int
	Goods_num int
	Attr_id string 	//多个属性yoga逗号隔开
	Time int64
}


//收货地址orm
type ShopReceiptAddr struct {
	Id int
	Uid int
	Name string
	Phone int
	Province string
	City string
	District string
	Detail_addr string
	Is_default int
	Addtime int64
}

//V1订单表
type ShopOrderv1 struct {
	Id int
	Buyer_id int
	Saler_id int
	Goods_id int
	Goods_num int
	Order_sn string
	Order_status int8
	Pay_status int8
	Post_status int8
	Post_id int8
	Pay_id int8
	Goods_tprice float32
	Post_fee float32
	Total_fee float32
	Receipt_id int
	Add_time int64
	Update_time int64
}


//基础订单orm
type ShopBaseOrder struct {
	Order_sn string
	Order_status int8
	Pay_status int8
	Post_status int8
	Post_id int8
	Pay_id int
	Goods_tprice float32
	Post_fee float32
	Total_fee float32
	Receipt_id int		//收货地址id
	Add_time int64
	Update_time int64
}
//用户账单orm
type ShopUserBill struct {
	Uid int
	Order_id int
	Money float32
	Balance float32
	Addtime int64
	Sn string       	//账单流水号
}
//系统账单orm
type ShopSystemBill struct {
	Uid int			//默认为-1
	Order_id int
	Money float32
	Balance float32
	Addtime int64
	Sn string       	//账单流水号
}

//退款orm
type ShopRefund struct {
	Id int
	Uid int
	Order_id int
	Money float32
	Reason string
	Status int8
	Result string
	Detail string
	Time int64
	Lasttime int64
}

