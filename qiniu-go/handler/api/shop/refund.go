package shop

import (
	"github.com/henrylee2cn/faygo"
	"qiniu-go/util"
	"qiniu-go/model/apiModel"
	"time"
)


//申请退款
type ApplyRefund struct {
	OrderId     int     `param:"<in:query><required> <desc:订单id>"`
	Uid	     int     `param:"<in:query><required> <desc:用户id>"`
	Money     float32     `param:"<in:query><required> <desc:退款金额>"`
	Reason     string     `param:"<in:query><required> <desc:退款原因>"`
	Detail     string     `param:"<in:query><required> <desc:退款说明>"`
}

/**
	申请退款api
 */
func (t *ApplyRefund) Serve(ctx *faygo.Context) error {
	order := new(apiModel.ShopOrderv1)
	refund := new(apiModel.ShopRefund)
	defer func() {
		refund = nil
		order = nil
	}()
	flag,e:= util.ApiEngine.Where("id=?",t.OrderId).Get(order)
	util.CheckError(e)
	if !flag {
		return ctx.JSON(200,util.GetResponseMap(400,"申请失败，该条订单不存在"),true)
	}
	f,er := util.ApiEngine.Where("order_id=?",t.OrderId).Get(refund)
	util.CheckError(er)
	if f {
		return ctx.JSON(200,util.GetResponseMap(400,"申请失败，已经申请了"),true)
	}

	refund.Order_id = t.OrderId
	refund.Money = t.Money
	refund.Reason = t.Reason
	refund.Detail = t.Detail
	refund.Uid  = t.Uid
	refund.Time = time.Now().Unix()
	refund.Lasttime = time.Now().Unix()

	_,err := util.ApiEngine.Insert(refund)
	util.CheckError(err)
	return ctx.JSON(200,util.GetResponseMap(200,"申请成功"),true)
}
func (t *ApplyRefund) Doc() faygo.Doc {
	return faygo.Doc{
		Note:   "申请退款",
		Return: "成功与否",
	}
}

//退款进度
type RefundProgress struct {
	RefundId     int     `param:"<in:query><required> <desc:退款单id>"`
}

/**
	退款进度api
 */
func (t *RefundProgress) Serve(ctx *faygo.Context) error {
	refund := new(apiModel.ShopRefund)
	defer func() {
		refund = nil
	}()
	flag,err := util.ApiEngine.Where("id=?",t.RefundId).Desc("id").Get(refund)
	util.CheckError(err)
	if !flag {
		return ctx.JSON(200,util.GetResponseMap(400,"不存在该退款单"),true)
	}

	return ctx.JSON(200,util.GetResponseMap(200,refund),true)
}
func (t *RefundProgress) Doc() faygo.Doc {
	return faygo.Doc{
		Note:   "退款进度",
		Return: "status: 0：退款中，1：退款成功，2：退款失败   resault:退款成功或失败都在这儿说明,   detail：退款详情",
	}
}

//获取某个用户的所有退款单
type GetAllRefund struct {
	Uid    int     `param:"<in:query><required> <desc:用户id>"`
}

/**
	获取某个用户的所有退款单api
 */
func (t *GetAllRefund) Serve(ctx *faygo.Context) error {
	allRefund := make([]apiModel.ShopRefund, 0)
	defer func() {
		allRefund = nil
	}()
	err := util.ApiEngine.Where("uid=?",t.Uid).Find(&allRefund)
	util.CheckError(err)
	if len(allRefund) == 0 {
		return ctx.JSON(200,util.GetResponseMap(400,"没有退款的订单"),true)
	}

	return ctx.JSON(200,util.GetResponseMap(200,allRefund),true)
}
func (t *GetAllRefund) Doc() faygo.Doc {
	return faygo.Doc{
		Note:   "退款进度",
		Return: "status: 0：退款中，1：退款成功，2：退款失败   resault:退款成功或失败都在这儿说明,   detail：退款详情",
	}
}

