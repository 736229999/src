package shop

import (
	"github.com/henrylee2cn/faygo"
	"qiniu-go/model/apiModel"
	"time"
	"qiniu-go/util"
	"log"
)



//添加收货地址
type AddReceiptAddr struct {
	Uid     int     `param:"<in:query><required> <desc:用户id>"`
	Name	    string	    `param:"<in:query><required> <desc:收货人姓名>"`
	Phone     int     `param:"<in:query><required> <desc:收货人电话>"`
	Province       string     `param:"<in:query><required> <desc:省>"`
	City     string       `param:"<in:query><required> <desc:市>"`
	District     string       `param:"<in:query><required> <desc:区县>"`
	Detail_addr     string       `param:"<in:query><required> <desc:详细地址>"`

}

/**
	添加收货地址的api
 */
func (t *AddReceiptAddr) Serve(ctx *faygo.Context) error {
	receiptAddr := new(apiModel.ShopReceiptAddr)
	defer func() {
		receiptAddr = nil
	}()
	receiptAddr.Uid = t.Uid
	receiptAddr.Phone = t.Phone
	receiptAddr.Name = t.Name
	receiptAddr.Province = t.Province
	receiptAddr.City = t.City
	receiptAddr.District = t.District
	receiptAddr.Detail_addr = t.Detail_addr
	receiptAddr.Addtime = time.Now().Unix()


	_,err := util.ApiEngine.Insert(receiptAddr)
	util.CheckError(err)

	return ctx.JSON(200,util.GetResponseMap(200,"添加收货地址成功"),true)
}

/**
	订单账单的api文档说明
 */
func (t *AddReceiptAddr) Doc() faygo.Doc {
	return faygo.Doc{
		Note:   "添加收货地址",
		Return: "成功与否",
	}
}


//获取用户收货地址列表
type GetUserReceiptAddr struct {
	Uid     int     `param:"<in:query><required> <desc:用户id>"`
}

/**
	获取用户收货地址列表的api
 */
func (t *GetUserReceiptAddr) Serve(ctx *faygo.Context) error {
	receiptAddr := make([]apiModel.ShopReceiptAddr, 0)
	defer func() {
		receiptAddr = nil
	}()
	err := util.ApiEngine.Cols("id","uid","name","phone","province","city","district","detail_addr").Where("uid=?",t.Uid).Find(&receiptAddr)
	if len(receiptAddr) == 0{
		return ctx.JSON(200,util.GetResponseMap(400,"没有收货地址"),true)
	}
	util.CheckError(err)
	return ctx.JSON(200,util.GetResponseMap(200,receiptAddr),true)
}

/**
	获取用户收货地址列表的文档说明
 */
func (t *GetUserReceiptAddr) Doc() faygo.Doc {
	return faygo.Doc{
		Note:   "获取用户收货地址列表",
		Return: "name:用户名，province：省，district：区县，detail_addr：详细地址",
	}
}


//设置默认收货地址
type SetDefalutAddr struct {
	Uid     int     `param:"<in:query><required> <desc:用户id>"`
	AddrId     int     `param:"<in:query><required> <desc:默认地址id>"`

}

/**
	设置默认收货地址的api
 */
func (t *SetDefalutAddr) Serve(ctx *faygo.Context) error {

	//先把原先的默认地址设为0，
	receiptAddr := new(apiModel.ShopReceiptAddr)
	defer func() {
		receiptAddr = nil
	}()
	flag , err := util.ApiEngine.Where("uid=? and is_default=?",t.Uid,1).Desc("id").Get(receiptAddr)
	util.CheckError(err)

	if flag {
		sql :="update api_shop_receipt_addr set is_default = 0 where uid = ? and is_default=1"
		res, err := util.ApiEngine.Exec(sql, t.Uid)
		log.Println(res)
		util.CheckError(err)
	}
	//再把新的地址设为默认地址
	sql :="update api_shop_receipt_addr set is_default = 1 where id = ?"
	res, err := util.ApiEngine.Exec(sql,  t.AddrId)
	aff,_ := res.RowsAffected()
	if int(aff)<1 {
		return ctx.JSON(200,util.GetResponseMap(400,"设置失败，地址id或uid不存在"),true)
	}
	util.CheckError(err)
	return ctx.JSON(200,util.GetResponseMap(200,"设置成功"),true)
}

/**
	获取用户收货地址列表的文档说明
 */
func (t *SetDefalutAddr) Doc() faygo.Doc {
	return faygo.Doc{
		Note:   "设置默认收货地址",
		Return: "成功与否",
	}
}



//删除收货地址
type DelCeiptAddr struct {
	ReceiptAddrId   int	`param:"<in:query><required> <desc:收货地址id>"`
}

/**
	删除收货地址的api
 */
func (t *DelCeiptAddr) Serve(ctx *faygo.Context) error {

	//先把原先的默认地址设为0，
	receiptAddr := new(apiModel.ShopReceiptAddr)
	defer func() {
		receiptAddr = nil
	}()
	flag , err := util.ApiEngine.Where("id=?",t.ReceiptAddrId).Desc("id").Get(receiptAddr)
	util.CheckError(err)
	if !flag {
		return ctx.JSON(200,util.GetResponseMap(400,"地址不存在"),true)
	}
	aff ,e := util.ApiEngine.Id(t.ReceiptAddrId).Delete(receiptAddr)
	util.CheckError(e)
	if aff<0 {
		return ctx.JSON(200,util.GetResponseMap(400,"删除失败"),true)
	}
	return ctx.JSON(200,util.GetResponseMap(200,"删除成功"),true)
}

/**
	删除收货地址的文档说明
 */
func (t *DelCeiptAddr) Doc() faygo.Doc {
	return faygo.Doc{
		Note:   "删除收货地址",
		Return: "成功与否",
	}
}