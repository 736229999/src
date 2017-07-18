package store
/*****************这里面存储登录之后保存的一些信息***************/

//uid => user
var StoreUserMap map[int]*StoreUser
var StoreWxMap map[int]*StoreWx
var StoreUserMoneyMap  map[int]*StoreUserMoney
var StoreAnchorInfoMap map[int]*StoreAnchorInfo
//进入个人中心的时候new一个user
type StoreUser struct {
	Id int
	NickName string
	Avatar string
	Sex int
	Signature string
	Level int
	Exp	int	//经验值
	Phone int
	Age int8
	Constellation string	//星座
	Token string
	
}

//微信的相关信息
type StoreWx struct {
	UnionId string
	OpenId string

}

//个人财产
type StoreUserMoney struct {
	StarDiamond int
	StarTicket int
	Remain float32	//余额
	FrozenRemain float32
}

//主播基本信息
type StoreAnchorInfo struct {
	Id int
	NickName string
	Avatar string
	Signature string
	Phone int

}

func init(){
	StoreUserMap = make(map[int]*StoreUser)
	StoreWxMap = make(map[int]*StoreWx)
	StoreUserMoneyMap = make(map[int]*StoreUserMoney)
	//user应该是在进入个人中心的时候注册进的，这里先模拟存储一个
	user := new(StoreUser)
	user.Id = 1
	user.Sex = 0
	user.NickName = "小美"
	StoreUserMap[user.Id] = user
}