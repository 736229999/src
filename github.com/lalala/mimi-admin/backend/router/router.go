package router

import (
	"fmt"
	"log"

	"github.com/caojunxyz/mimi-admin/backend/filter"
	"github.com/caojunxyz/mimi-admin/backend/service/activity"
	"github.com/caojunxyz/mimi-admin/backend/service/logs"
	"github.com/caojunxyz/mimi-admin/backend/service/lottery"
	// "github.com/caojunxyz/mimi-admin/backend/service/discover"
	"github.com/caojunxyz/mimi-admin/backend/service/options"
	"github.com/caojunxyz/mimi-admin/backend/service/user"
	"github.com/caojunxyz/mimi-admin/backend/service/usercenter"
	"github.com/gin-gonic/gin"

	"github.com/caojunxyz/mimi-admin/backend/service/football"
	"github.com/caojunxyz/mimi-admin/backend/service/order"
	"github.com/caojunxyz/mimi-admin/backend/service/statistics"
)

func NewUserServer() *user.UserService {
	return &user.UserService{}
}

func NewLotteryServer() *lottery.LotteryService {
	return &lottery.LotteryService{}
}

func NewActivityServer() *activity.ActivityService {
	return &activity.ActivityService{}
}

// func NewDiscoverServer() *discover.DiscoverService {
// 	return &discover.DiscoverService{}
// }
func NewLogsServer() *logs.LogsService {
	return &logs.LogsService{}
}

func NewOptionsServer() *options.OptionsService {
	return &options.OptionsService{}
}

func NewUsercenterServer() *usercenter.UsercenterService {
	return &usercenter.UsercenterService{}
}

func NewFootballServer() *football.FootballService {
	return &football.FootballService{}
}

func NewStatisticsServer() *statistics.StatisticsService {
	return &statistics.StatisticsService{}
}

func NewOrderServer() *order.OrderService {
	return &order.OrderService{}
}

func Init(addr int) error {

	router := gin.Default()
	user := NewUserServer()
	lottery := NewLotteryServer()
	activity := NewActivityServer()
	logs := NewLogsServer()
	options := NewOptionsServer()
	usercenter := NewUsercenterServer()
	football := NewFootballServer()
	statistics := NewStatisticsServer()
	order := NewOrderServer()
	router.POST("/login", user.HandleLogin)


	//authz := router.Use()
	authz := router.Use(filter.FilterLogin)
	{
		//用户.
		authz.GET("/user/list", user.HandleUserList)
		authz.POST("/user/update", user.HandleUserUpdate)

		//用户权限.
		authz.POST("/user/privilege/list", user.HandleUserPrivilegeList)
		authz.POST("/user/privilege/add", user.HandleUserPrivilegeAdd)
		authz.POST("/user/privileges/edit", user.HandleUserPrivilegesEdit)
		authz.POST("/user/privileges/delete", user.HandleUserPrivilegesDelete)

		//用户角色.
		authz.POST("/user/role/list", user.HandleUserRoleList)
		authz.GET("/user/roles/add", user.HandleUserRolesAdd)

		//彩票.
		authz.GET("/lottery/list", lottery.HandleLotteryList)
		authz.POST("/lottery/buycai/options", lottery.HandleLotteryBuycaiOptions)
		authz.POST("/lottery/buycai/options/add", lottery.HandleLotteryBuycaiOptionsAdd)
		authz.POST("/lottery/buycai/options/delete", lottery.HandleLotteryBuycaiOptionsDelete)
		authz.GET("/lottery/buycai/options/detail/:id/lottery/:lottery", lottery.HandleLotteryBuycaiOptionsDetail)
		authz.POST("/lottery/buycai/options/update", lottery.HandleLotteryBuycaiOptionsUpdate)
		authz.POST("/lottery/buycai/options/init", lottery.HandleLotteryBuycaiOptionsInit)
		authz.POST("/lottery/issue", lottery.HandleLotteryIssue)
		authz.POST("/lottery/newissue", lottery.HandleLotteryNewIssue)
		authz.POST("/lottery/playtime/list", lottery.HandleLotteryPlaytimeList)
		authz.POST("/lottery/playtimesetting/add", lottery.HandleLotteryPlaytimeAdd)
		authz.POST("/lottery/playtimesetting/update", lottery.HandleLotteryPlaytimeUpdate)
		authz.POST("/lottery/type/list", lottery.HandleLotteryTypeList)
		authz.POST("/lottery/home/options/add", lottery.HandleLotteryHomeOptionsAdd)
		authz.POST("/lottery/home/options/list", lottery.HandleLotteryHomeOptionsList)
		authz.GET("/lottery/home/options/:id", lottery.HandleGetLotteryHomeOptionsById)
		authz.POST("/lottery/home/options/edit", lottery.HandleLotteryHomeOptionsEdit)
		authz.POST("/lottery/home/options/notaddlist", lottery.HandleLotteryHomeOptionsNotAddList)

		//开奖信息.
		authz.POST("/lottery/open/list", lottery.HandleLotteryOpenList)
		authz.POST("/lottery/open/search", lottery.HandleLotteryOpenSearch)
		authz.POST("/lottery/open/history", lottery.HandleLotteryOpenHistory)
		//authz.POST("/lottery/home/winning/list", lottery.HandleLotteryHomeWinningList)

		//活动.
		authz.POST("/activity/task/addTask", activity.InserTask)
		authz.POST("/activity/task/taskList", activity.QueryTaskList)
		authz.POST("/activity/task/allTask", activity.QueryAllTask)
		authz.POST("/activity/task/delete", activity.DeleteTask)
		authz.POST("/activity/task/detail", activity.QueryTaskById)
		authz.POST("/activity/task/update", activity.UpdateTask)
		authz.POST("/activity/task/taskTypes", activity.QueryAllTaskType)
		authz.POST("/activity/cdkey/allGiftTemplate", activity.QueryAllGiftTemplate)
		authz.POST("/activity/activity/activityList", activity.QueryActivityList)
		authz.POST("/activity/activity/addActivity", activity.AddActivity)
		authz.POST("/activity/activity/delete", activity.DeleteActivity)
		authz.POST("/activity/activity/detail", activity.QueryActivityById)
		authz.POST("/activity/activity/update", activity.UpdateActivity)

		//平台配置
		//authz.POST("/options/banner", options.HandleBannerAdd)
		authz.POST("/options/update/contact", options.HandleOptionsUpdateContact)
		authz.POST("/options/query/contact", options.HandleOptionsQueryContact)
		authz.POST("/options/feedback/list", options.HandleOptionsFeedbackList)
		authz.POST("/options/feedback/del", options.HandleOptionsFeedbackDel)
		authz.GET("/options/feedback/detail/:id", options.HandleOptionsFeedbackDetail)
		authz.POST("/options/feedback/update", options.HandleOptionsFeedbackUpdate)
		authz.POST("/options/faq", options.HandleFaqAdd)
		authz.GET("/options/faq", options.HandleFaqList)
		authz.GET("/options/faq/detail/:id", options.HandleFaqDetail)
		authz.PUT("/options/faq/detail", options.HandleUpdateFaq)

		//平台配置，Banner轮播图，发现,新闻等
		authz.POST("/options/discover/news", options.HandleNewsAdd)
		authz.GET("/options/discover/news", options.HandleNewsListGet)
		authz.GET("/options/discover/news/detail/:id", options.GetNewsById)
		authz.PUT("/options/discover/news", options.HandleUpdateNews)
		authz.GET("/options/discover/news/select", options.QueryNewsOfSelect)
		authz.GET("/options/banner", options.HandleBannerList)
		authz.POST("/options/banner", options.HandleBannerAdd)
		authz.GET("/options/banner/detail/:id", options.HandleBannerDetail)
		authz.PUT("/options/banner/detail", options.HandleUpdateBanner)
		authz.GET("/options/banner/pre", options.HandlePreBanner)

		//日志.
		authz.POST("/log/list", logs.HandleLogList)

		//礼包.
		authz.POST("/gift/list", activity.HandleGiftLit)
		//authz.POST("/gift/type/list", activity.HandleGiftTypeList)
		authz.POST("/gift/add", activity.HandleGiftAdd)
		authz.GET("/gift/detail/:id", activity.HandleDetail)
		authz.POST("/gift/delete", activity.HandleDeleteGift)
		authz.POST("/gift/update", activity.HandleGiftUpdate)

		//usercenter.
		authz.POST("/usercenter/list", usercenter.HandleUserList)
		authz.GET("/usercenter/detail/:id", usercenter.HandleUserDetail)
		authz.GET("/usercenter/fund/:id", usercenter.HandleUserFund)
		// authz.GET("/usercenter/withdraw/:id", usercenter.HandleUserWithdraw)

		authz.GET("/usercenter/withdrawapply", usercenter.HandleWithdrawApplyList)
		authz.GET("/usercenter/withdrawapply/detail/:id", usercenter.HandleGetWithdrawApplyById)
		authz.PUT("/usercenter/withdrawapply/detail", usercenter.HandleUpdateWithdrawApplyStaus)
		authz.PUT("/usercenter/withdrawapply/claim", usercenter.HandleClaimWithdrawApply)
		authz.GET("/usercenter/withdrawapply/check/:id", usercenter.HandleCheckWithdrawApply)
		authz.POST("/usercenter/withdraw/auth", usercenter.HandleWithdrawAuditAuthAdd)
		authz.GET("/usercenter/withdraw/auth", usercenter.HanldeWithdrawAuditAuthList)
		authz.GET("/usercenter/withdraw/auth/detail/:id", usercenter.HandleWithdrawAuditAuth)
		authz.PUT("/usercenter/withdraw/auth/detail", usercenter.HandleUpdateWithdrawAuditAuth)
		authz.GET("/usercenter/withdrawtransfer", usercenter.HandleWithdrawTransferList)
		authz.GET("/usercenter/withdrawtransfer/detail/:id", usercenter.HandleGetWithdrawTransfer)
		authz.PUT("/usercenter/withdrawtransfer/detail", usercenter.HandleUpdateWithdrawTransfer)

		//足彩
		authz.GET("/football/games", football.HandleGamesListGet)
		authz.GET("/football/games/detail/:id", football.GetFbGameById)
		authz.POST("/football/games/add", football.HandleGameAdd)
		authz.PUT("/football/games/update", football.HandleGameUpdate)
		authz.GET("/football/league/select", football.HandleLeagueListGet)
		authz.GET("/football/team/select", football.HandleTeamListGet)
		authz.GET("/football/odds/detail/:id", football.GetOddsById)
		authz.PUT("/football/odds/update", football.HandleOddsUpdate)
		authz.POST("/football/odds/add", football.HandleOddsAdd)
		authz.GET("/football/opencai", football.HandleOpencaiListGet)
		authz.GET("/football/opencai/detail/:id", football.GetOpencaiById)
		authz.POST("/football/opencai/add", football.HandleOpencaiAdd)
		authz.PUT("/football/opencai/update", football.HandleOpencaiUpdate)
		authz.GET("/football/team", football.HandleTeamListGet)
		authz.POST("/football/team/add", football.HandleTeamAdd)
		authz.GET("/football/team/detail/:id", football.GetTeamById)
		authz.PUT("/football/team/update", football.HandleTeamUpdate)
		authz.GET("/football/league", football.HandleLeagueListGet)
		authz.POST("/football/league/add", football.HandleLeagueAdd)
		authz.PUT("/football/league/update", football.HandleLeagueUpdate)

		//数据统计.
		authz.POST("/data/orderandincome", statistics.HandleOrderAndIncome)
		authz.POST("/data/recharge/list", statistics.HandleGetRechargeList)
		authz.POST("/data/recharge/month", statistics.HandleGetRechargeByMonth)
		authz.POST("/data/recharge/year", statistics.HandleGetRechargeByYear)
		authz.POST("/data/user/statistics", statistics.HandleUserStatistics)
		//购彩.
		//authz.POST("/data/buycai/day")

		//礼包.
		authz.POST("/gift/template/add", activity.HandleGiftTemplateAdd)
		authz.POST("/gift/template/list", activity.HandleGiftTemplateList)
		authz.GET("/gift/template/:id", activity.HandleGiftTemplateById)
		authz.POST("/gift/template/update", activity.HandleGiftTemplateUpdate)
		authz.POST("/gift/template/delete", activity.HandleGiftTemplateDelete)

		//cdkey.
		authz.POST("/activity/cdkey/add", activity.HandleActivityCdKeyAdd)
		authz.POST("/activity/cdkey/list", activity.HandleActivityCdKeyList)
		authz.POST("/activity/gift/template/list", activity.HandleActivityGiftTemplateList)
		authz.GET("/activity/cdkey/detail/:id", activity.HandleActivityCdkeyDetail)
		authz.POST("/activity/cdkey/update", activity.HandleActivityCdKeyUpdate)
		authz.POST("/activity/cdkey/delete", activity.HandleActivityCdKeyDelete)
		authz.POST("/activity/cdkey/export", activity.HandleActivityCdKeyExport)

		//订单
		authz.POST("/order/user", order.QueryUserOrderList)					//用户订单列表
		authz.POST("/order/select", order.QueryUserOrderListWithCondition)	//查询用户订单
	}

	log.Println(addr)
	return router.Run(fmt.Sprintf(":%d", addr))
}
