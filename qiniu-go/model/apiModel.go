package model
//api要用到的所有model

//买家订单的orm
type Buyerorder struct {
	Id int
	Order_sn string `xorm:"unique"`		//最好是和微信or支付宝支付之后的订单号一致
	Buyer_id int
	Saler_id int
	Goods_id int
	Money float32
	Status int8
	Content string
	Create_time int64
	Lastup_time int64
}

//卖家订单的orm
type Salerorder struct {
	Id int
	Order_sn string `xorm:"unique"`
	Buyer_id int
	Saler_id int
	Goods_id int
	Money float32
	Status int8
	Content string
	Create_time int64
	Lastup_time int64
}
//账单的orm
type Bill struct {
	Id int
	Bill_sn string `xorm:"unique"`
	Buyer_id int
	Saler_id int
	Goods_id int
	Money float32
	Buyermoney_change float32
	Buyermoney_remain float32
	Systemmoney_change float32
	Systemmoney_remain float32
	Salermoney_change float32
	Salermoney_remain float32
	Pay_way int8
	Content string
	Create_time int64
}

type User struct {
	Id int
	Remain float32
}
