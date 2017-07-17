package apiModel

//直播认证orm
type Liveauth struct {
	Uid int `xorm:"unique"`		//最好是和微信or支付宝支付之后的订单号一致
	Frontimg string
	Backimg string
	Holdimg string
	Status int8
	Addtime int64
	Lasttime int64
}
