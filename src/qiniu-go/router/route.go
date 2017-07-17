package router

import (
	"github.com/henrylee2cn/faygo"
	"qiniu-go/handler"
	"qiniu-go/middleware"
	"qiniu-go/handler/ajax"
	"qiniu-go/handler/api/shop"
	"qiniu-go/handler/pages"
	"qiniu-go/handler/api"
	"qiniu-go/handler/pages/simplelive"
	"qiniu-go/handler/ajax/simpleLiveAjax"
	"qiniu-go/handler/rpc"
)

// Route register router in a tree style.
func Route(frame *faygo.Framework) {

	frame.Route(
		//这两个将两个不同的公众号绑定到同一个服务器
		frame.NewNamedAPI("test struct handler", "GET", "/MP_verify_FfXgomWspAkznBQA.txt",handler.Micha),
		frame.NewNamedAPI("test struct handler", "GET", "/MP_verify_Z2Q4pyfE628bSDwp.txt",handler.Jidi),

		//====================
		frame.NewNamedAPI("test struct handler", "POST", "/test", &handler.Test{}).Use(middleware.Token),
		frame.NewNamedAPI("login", "Get", "/login", &handler.Login{}),


		//页面层
		frame.NewNamedAPI("pages", "Get", "/", pages.Index),
		frame.NewNamedAPI("pages", "Get", "/index.html", pages.Index),
		frame.NewNamedAPI("pages", "Get", "/personalCenter.html", pages.PersonalCenter),
		frame.NewNamedAPI("pages", "Get", "/personalHome.html", pages.PersonalHome),
		frame.NewNamedAPI("pages", "Get", "/dynamic.html", pages.Dynamic),
		frame.NewNamedAPI("pages", "Get", "/grade.html", pages.Grade),
		frame.NewNamedAPI("pages", "Get", "/profit.html", pages.Profit),
		frame.NewNamedAPI("pages", "Get", "/blingRecharge.html", pages.BlingRecharge),
		frame.NewNamedAPI("pages", "Get", "/myConsumption.html", pages.MyConsumption),
		frame.NewNamedAPI("pages", "Get", "/liveAuthentication.html", pages.LiveAuthentication),
		frame.NewNamedAPI("pages", "Get", "/invitation.html", pages.Invitation),
		frame.NewNamedAPI("pages", "Get", "/personalSet.html", pages.PersonalSet),
		frame.NewNamedAPI("pages", "Get", "/updatePassword.html", pages.UpdatePassword),
		frame.NewNamedAPI("pages", "Get", "/addressList.html", pages.AddressList),

		frame.NewGroup("simpleLive",

			frame.NewNamedAPI("首页", "get", "/index.html", simplelive.Index),
			frame.NewNamedAPI("登录", "get", "/login.html", simplelive.Login),
			frame.NewNamedAPI("修改密码", "get", "/passwordRretrieval.html", simplelive.PasswordRretrieval),
			frame.NewNamedAPI("微信登录", "get", "/wxLogin.html", simplelive.WxLogin),

		),


		////业务层
		//frame.NewNamedAPI("login", "Get", "/threeDistribute", &service.Three_distribute{}),
		//frame.NewNamedAPI("login", "Get", "/bindAgent", &service.BindAgentRelation{}),



		//测试
		frame.NewNamedAPI("test", "POST", "/payNotify", handler.PayNotify),
		frame.NewNamedAPI("test", "POST", "/testAjax", api.TestApi),

		frame.NewNamedAPI("rpc", "Get", "/httpRpc", rpc.RegisterRpc),
		frame.NewNamedAPI("rpc", "Get", "/rpcClient", rpc.RpcClient),

		//api
		frame.NewGroup("api",

			frame.NewGroup("shop",
				frame.NewNamedAPI("商品分类", "POST", "/addCategory", &shop.AddCategory{}),
				frame.NewNamedAPI("商品分类", "POST", "/allCategory", &shop.AllCategory{}),
				//frame.NewNamedAPI("商品分类", "POST", "/deleteCategory", &shop.DelCategory{}),

				//frame.NewNamedAPI("商品属性", "POST", "/addAttr", &shop.AddAttr{}),
				//frame.NewNamedAPI("商品属性", "POST", "/delAttr", &shop.DelAttr{}),

				frame.NewNamedAPI("添加商品", "POST", "/addGoods", &shop.AddGoods{}),
				frame.NewNamedAPI("在售商品", "Get", "/getOnsaleGoods", &shop.GetOnsaleGoods{}),
				frame.NewNamedAPI("下架商品", "POST", "/getDownsaleGoods", &shop.GetDownSaleGoods{}),

				frame.NewNamedAPI("添加收货地址", "POST", "/addReceiptAddr", &shop.AddReceiptAddr{}),
				frame.NewNamedAPI("获取用户收货地址", "POST", "/getUserReceiptAddr", &shop.GetUserReceiptAddr{}),
				frame.NewNamedAPI("设置默认收货地址", "POST", "/setDefaultAddr", &shop.SetDefalutAddr{}),
				frame.NewNamedAPI("删除收货地址", "POST", "/delReceiptAddr", &shop.DelCeiptAddr{}),

				frame.NewNamedAPI("创建订单", "POST", "/createOrder", &shop.CreateOrder{}),
				frame.NewNamedAPI("用户订单列表", "POST", "/use   rAllOrders", &shop.UserAllOrder{}),
				frame.NewNamedAPI("商家订单列表", "POST", "/salerAllOrders", &shop.SalerAllOrders{}),
				frame.NewNamedAPI("订单详情", "POST", "/orderDetail", &shop.OrderDetail{}),

				frame.NewNamedAPI("申请退款", "POST", "/applyRefund", &shop.ApplyRefund{}),
				frame.NewNamedAPI("退款进度", "POST", "/refundProgress", &shop.RefundProgress{}),
				frame.NewNamedAPI("用户所有退款订单", "POST", "/userAllRefund", &shop.GetAllRefund{}),
				//frame.NewNamedAPI("购物车", "POST", "/addGoodsIntoCart", &shop.AddIntoCart{}),

				//frame.NewNamedAPI("支付阶段", "POST", "/createOrder", &shop.CreateOrder{}),

			),
			//frame.NewGroup("user",
			//
			//	frame.NewNamedAPI("用户", "POST", "/getUserInfo", &user.GetUserInfo{}),
			//	frame.NewNamedAPI("用户", "POST", "/getUserLevel", &user.GetUserLevel{}),
			//	frame.NewNamedAPI("发红包", "POST", "/sendRedPocket", &user.SendRedPocket{}),
			//),
			//frame.NewGroup("interaction",
			//
			//	frame.NewNamedAPI("交互", "POST", "/getAttentionUser", &user.GetUserAttention{}),
			//	frame.NewNamedAPI("交互", "POST", "/hasAttentionOther", &user.HasAttentionOther{}),
			//	frame.NewNamedAPI("交互", "POST", "/attentionUser", &user.AttentionUser{}),
			//),
		),
		//ajax
		frame.NewGroup("ajax",
			frame.NewNamedAPI("直播间的ajax", "Get", "/isLogin", ajax.IsLogin),
			frame.NewGroup("index",
				frame.NewNamedAPI("直播间的ajax", "Get", "/uploadImg", ajax.UploadImg),
			),
			frame.NewGroup("blingRecharge",
				frame.NewNamedAPI("星钻余额页面的ajax", "Get", "/getRechargeMoney", ajax.GetRechargeMoney),
			),
			frame.NewGroup("liveAuthentication",
				frame.NewNamedAPI("提交照片的ajax", "Get", "/uploadImg", ajax.LiveAuthUploadImg),
				frame.NewNamedAPI("提交照片的ajax", "Get", "/StoreIdCard", ajax.StoreIdCard),
			),
		),
		frame.NewGroup("simpleLiveAjax",
			frame.NewNamedAPI("简播微信登录ajax", "Get", "/wxLogin", simpleLiveAjax.WxLogin),
		),
	)

}
