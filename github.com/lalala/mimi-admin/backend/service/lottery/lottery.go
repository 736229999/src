package lottery

import (
	"github.com/caojunxyz/mimi-admin/backend/core"
	"github.com/caojunxyz/mimi-admin/dbadminagent/proto"
	"github.com/caojunxyz/mimi-server/lottery"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"strconv"

	"github.com/caojunxyz/mimi-server/lottery/saleissue"

	"errors"
	"github.com/caojunxyz/mimi-server/lottery/validator/bjpk10"
	"github.com/caojunxyz/mimi-server/lottery/validator/cqssc"
	"github.com/caojunxyz/mimi-server/lottery/validator/dlt"
	"github.com/caojunxyz/mimi-server/lottery/validator/fc3d"
	"github.com/caojunxyz/mimi-server/lottery/validator/gd11x5"
	"github.com/caojunxyz/mimi-server/lottery/validator/pl3"
	"github.com/caojunxyz/mimi-server/lottery/validator/pl5"
	"github.com/caojunxyz/mimi-server/lottery/validator/ssq"
	"time"
)

type LotteryService struct {
	core.Service
}

type Lottery struct {
	Name     string
	Id       int
	Code     string
	DayMaxNo int
	Type     int
}

//获取彩种列表.
func lotteryList() []Lottery {

	lottery := lottery.GetLottery()

	var lotteryList []Lottery
	for _, v := range lottery {

		lottery := Lottery{
			Name:     v.Name,
			Id:       int(v.Id),
			Code:     v.Code,
			DayMaxNo: v.DayMaxNo,
			Type:     int(v.Type),
		}

		lotteryList = append(lotteryList, lottery)
	}
	return lotteryList
}

//获取彩票种类.
func (srv *LotteryService) HandleLotteryList(c *gin.Context) {

	srv.Log(c, nil, "获取彩票种类列表", core.QUERY_OPERATION)
	srv.Json(lotteryList(), http.StatusOK, c)
}

//获取彩票的配置.
func (srv *LotteryService) HandleLotteryBuycaiOptions(c *gin.Context) {

	msg := dbproto.BuycaiOptionsReply{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Printf("BindJson:%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
	}

	if len(msg.Lottery) < 1 {
		srv.Json("数据获取失败", http.StatusForbidden, c)
		return
	}

	buycai, err := srv.Db().QueryBuycaiOptions(context.Background(), &msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取数据失败", http.StatusForbidden, c)
		return
	}

	log.Printf("%+v\n", buycai)

	info := dbproto.BuycaiOptionsReply{
		Page:    msg.GetPage(),
		Size:    msg.GetSize(),
		Total:   buycai.GetTotal(),
		Lottery: msg.GetLottery(),
		Buycai:  buycai.GetBuycai(),
	}

	srv.Log(c, msg, "获取彩票的配置", core.QUERY_OPERATION)
	srv.Json(info, http.StatusOK, c)
}

//添加彩票配置.
func (srv *LotteryService) HandleLotteryBuycaiOptionsAdd(c *gin.Context) {

	msg := dbproto.BuycaiOptionsIssue{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("添加失败", http.StatusForbidden, c)
		return
	}

	log.Printf("%+v\n", msg)
	_, err = srv.Db().InsertLotteryOptions(context.Background(), &msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("添加失败", http.StatusForbidden, c)
		return
	}

	srv.Log(c, msg, "添加彩票配置", core.ADD_OPERATION)
	srv.Json("添加成功", http.StatusOK, c)
}

//获取期次信息.
func (srv *LotteryService) HandleLotteryIssue(c *gin.Context) {

	msg := dbproto.BuycaiOptionsIssue{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	if msg.GetIssue() == "" {
		log.Printf("%+v\n", msg.GetIssue())
		srv.Json("期号不能为空", http.StatusForbidden, c)
		return
	}

	log.Printf("%+v\n", msg)
	var issueData interface{}
	num := int(msg.GetNum())

	switch msg.GetLottery() {

	case lottery.Bjpk10:
		issueData = bjpk10.MakeSaleIssueList(
			saleissue.SaleIssue{
				Issue:     msg.GetIssue(),
				StartTime: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local).Unix(),
			},
			num,
		)
		break
	case lottery.Cqssc:
		issueData = cqssc.MakeSaleIssueList(
			saleissue.SaleIssue{
				Issue:     msg.GetIssue(),
				StartTime: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local).Unix(),
			},
			num,
		)
		break
	case lottery.Gd11x5:
		issueData = gd11x5.MakeSaleIssueList(
			saleissue.SaleIssue{
				Issue:     msg.GetIssue(),
				StartTime: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local).Unix(),
			},
			num,
		)
		break
	case lottery.Dlt:
		st, et, ot, err := srv.getIssueByLottery(msg.GetLottery())
		if err != nil {
			log.Printf("%+v\n", err)
			srv.Json("获取失败", http.StatusForbidden, c)
			return
		}
		issueData = dlt.MakeSaleIssueList(
			saleissue.SaleIssue{
				Issue:     msg.GetIssue(),
				StartTime: time.Date(st.Year(), st.Month(), st.Day(), st.Hour(), st.Minute(), 0, 0, time.Local).Unix(),
				EndTime:   time.Date(et.Year(), et.Month(), et.Day(), et.Hour(), 0, 0, 0, time.Local).Unix(),
				OpenTime:  time.Date(ot.Year(), ot.Month(), ot.Day(), ot.Hour(), ot.Minute(), 0, 0, time.Local).Unix(),
			},
			num)
		break
	case lottery.Ssq:
		st, et, ot, err := srv.getIssueByLottery(msg.GetLottery())
		if err != nil {
			log.Printf("%+v\n", err)
			srv.Json("获取失败", http.StatusForbidden, c)
			return
		}
		issueData = ssq.MakeSaleIssueList(
			saleissue.SaleIssue{
				Issue:     msg.GetIssue(),
				StartTime: time.Date(st.Year(), st.Month(), st.Day(), st.Hour(), st.Minute(), 0, 0, time.Local).Unix(),
				EndTime:   time.Date(et.Year(), et.Month(), et.Day(), et.Hour(), 0, 0, 0, time.Local).Unix(),
				OpenTime:  time.Date(ot.Year(), ot.Month(), ot.Day(), ot.Hour(), ot.Minute(), 0, 0, time.Local).Unix(),
			},
			num)
		break
	case lottery.Fc3d:
		st, et, ot, err := srv.getIssueByLottery(msg.GetLottery())
		if err != nil {
			log.Printf("%+v\n", err)
			srv.Json("获取失败", http.StatusForbidden, c)
			return
		}
		issueData = fc3d.MakeSaleIssueList(
			saleissue.SaleIssue{
				Issue:     msg.GetIssue(),
				StartTime: time.Date(st.Year(), st.Month(), st.Day(), st.Hour(), st.Minute(), 0, 0, time.Local).Unix(),
				EndTime:   time.Date(et.Year(), et.Month(), et.Day(), et.Hour(), 0, 0, 0, time.Local).Unix(),
				OpenTime:  time.Date(ot.Year(), ot.Month(), ot.Day(), ot.Hour(), ot.Minute(), 0, 0, time.Local).Unix(),
			},
			num)
		break
	case lottery.Pl3:
		st, et, ot, err := srv.getIssueByLottery(msg.GetLottery())
		if err != nil {
			log.Printf("%+v\n", err)
			srv.Json("获取失败", http.StatusForbidden, c)
			return
		}
		issueData = pl3.MakeSaleIssueList(
			saleissue.SaleIssue{
				Issue:     msg.GetIssue(),
				StartTime: time.Date(st.Year(), st.Month(), st.Day(), st.Hour(), st.Minute(), 0, 0, time.Local).Unix(),
				EndTime:   time.Date(et.Year(), et.Month(), et.Day(), et.Hour(), 0, 0, 0, time.Local).Unix(),
				OpenTime:  time.Date(ot.Year(), ot.Month(), ot.Day(), ot.Hour(), ot.Minute(), 0, 0, time.Local).Unix(),
			},
			num)
		break
	case lottery.Pl5:
		st, et, ot, err := srv.getIssueByLottery(msg.GetLottery())
		if err != nil {
			log.Printf("%+v\n", err)
			srv.Json("获取失败", http.StatusForbidden, c)
			return
		}
		issueData = pl5.MakeSaleIssueList(
			saleissue.SaleIssue{
				Issue:     msg.GetIssue(),
				StartTime: time.Date(st.Year(), st.Month(), st.Day(), st.Hour(), st.Minute(), 0, 0, time.Local).Unix(),
				EndTime:   time.Date(et.Year(), et.Month(), et.Day(), et.Hour(), 0, 0, 0, time.Local).Unix(),
				OpenTime:  time.Date(ot.Year(), ot.Month(), ot.Day(), ot.Hour(), ot.Minute(), 0, 0, time.Local).Unix(),
			},
			num)
		break
	}

	srv.Log(c, msg, "获取期次信息", core.QUERY_OPERATION)
	srv.Json(issueData, http.StatusOK, c)
}

//获取彩票的最后一期期号数据.
func (srv *LotteryService) getIssueByLottery(lottery string) (st time.Time, et time.Time, ot time.Time, err error) {

	buycai, err := srv.Db().QueryLotteryIssue(context.Background(), &dbproto.StringValue{Value: lottery})
	if err != nil {
		log.Printf("%+v\n", err)
		return time.Now(), time.Now(), time.Now(), err
	}
	st = time.Unix(buycai.GetStartTime(), 0)
	et = time.Unix(buycai.GetEndTime(), 0)
	ot = time.Unix(buycai.GetOpenTime(), 0)

	return st, et, ot, nil
}

//获取最新的期号.
func (srv *LotteryService) HandleLotteryNewIssue(c *gin.Context) {

	msg := dbproto.BuycaiOptionsIssue{}

	err := c.BindJSON(&msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	buycai, err := srv.Db().QueryLotteryIssue(context.Background(), &dbproto.StringValue{Value: msg.GetLottery()})
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	srv.Log(c, msg, "获取最新的期号", core.QUERY_OPERATION)
	srv.Json(buycai, http.StatusOK, c)
}

//获取玩法时间设置的列表.
func (srv *LotteryService) HandleLotteryPlaytimeList(c *gin.Context) {

	list, err := srv.Db().QueryPlayTimeSettingList(context.Background(), &dbproto.Nil{})
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	var name string
	for _, v := range list.GetList() {
		name = lottery.GetLotteryName(v.LotteryId)
		log.Println(v.Id, ":", name)
		v.Name = name
	}

	srv.Log(c, nil, "获取玩法时间设置的列表", core.QUERY_OPERATION)
	srv.Json(list.GetList(), http.StatusOK, c)
}

//添加玩法时间设置.
func (srv *LotteryService) HandleLotteryPlaytimeAdd(c *gin.Context) {

	msg := dbproto.PlayTimeSettings{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("添加失败", http.StatusForbidden, c)
		return
	}

	log.Printf("msg:%+v\n", &msg)
	_, err = srv.Db().InsertPlayTimeSettings(context.Background(), &msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("添加失败", http.StatusForbidden, c)
		return
	}

	srv.Log(c, msg, "添加玩法时间设置", core.ADD_OPERATION)
	srv.Json("添加成功", http.StatusOK, c)
}

//更新玩法时间间隔设置
func (srv *LotteryService) HandleLotteryPlaytimeUpdate(c *gin.Context) {

	msg := &dbproto.PlayTimeSettings{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("更新失败", http.StatusForbidden, c)
		return
	}

	_, err = srv.Db().UpdatePlayTimeSettings(context.Background(), msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("更新失败", http.StatusForbidden, c)
		return
	}

	srv.Log(c, msg, "更新玩法时间间隔设置", core.UPDATE_OPERATION)
	srv.Json("更新成功", http.StatusOK, c)
}

//获取彩票类型列表.
func (srv *LotteryService) HandleLotteryTypeList(c *gin.Context) {

	list := lottery.GetLotteryTypeList()

	srv.Log(c, nil, "获取彩票类型列表", core.QUERY_OPERATION)
	srv.Json(list, http.StatusOK, c)
}

//添加彩种配置.
func (srv *LotteryService) HandleLotteryHomeOptionsAdd(c *gin.Context) {

	msg := &dbproto.LotteryOptions{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("添加失败", http.StatusForbidden, c)
		return
	}

	_, err = srv.Db().InsertLotteryOption(context.Background(), msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("添加失败", http.StatusForbidden, c)
		return
	}

	srv.Log(c, msg, "添加彩种配置", core.ADD_OPERATION)
	srv.Json("添加成功", http.StatusOK, c)
}

//获取彩票种类配置的列表.
func (srv *LotteryService) HandleLotteryHomeOptionsList(c *gin.Context) {

	list, err := srv.Db().QueryLotteryOptionsList(context.Background(), &dbproto.Nil{})
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	srv.Log(c, nil, "获取彩票种类配置的列表", core.QUERY_OPERATION)
	srv.Json(list.GetList(), http.StatusOK, c)
}

//编辑彩票配置.
func (srv *LotteryService) HandleLotteryHomeOptionsEdit(c *gin.Context) {

	msg := &dbproto.LotteryOptions{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("更新失败", http.StatusForbidden, c)
		return
	}

	_, err = srv.Db().UpdateLotteryOptionsById(context.Background(), msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("更新失败", http.StatusForbidden, c)
		return
	}

	srv.Log(c, msg, "编辑彩票配置", core.UPDATE_OPERATION)
	srv.Json("更新成功", http.StatusOK, c)
}

//根据id获取彩票配置.
func (srv *LotteryService) HandleGetLotteryHomeOptionsById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	option, err := srv.Db().GetLotteryOptionsById(context.Background(), &dbproto.IntValue{Value: int64(id)})
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	srv.Log(c, id, "根据id获取彩票配置", core.QUERY_OPERATION)
	srv.Json(option, http.StatusOK, c)
}

//获取还没有添加到彩票配置中的彩种.
func (srv *LotteryService) HandleLotteryHomeOptionsNotAddList(c *gin.Context) {

	list, err := srv.Db().QueryLotteryOptionsList(context.Background(), &dbproto.Nil{})
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	lotteryList := lotteryList()
	options := list.GetList()

	for _, val := range options {
		for k, v := range lotteryList {
			if val.GetId() == int64(v.Id) {
				lotteryList = append(lotteryList[:k], lotteryList[k+1:]...)
			}
		}
	}

	srv.Log(c, nil, "获取没有添加的彩种", core.QUERY_OPERATION)
	srv.Json(lotteryList, http.StatusOK, c)
}

//根据id删除期号.
func (srv *LotteryService) HandleLotteryBuycaiOptionsDelete(c *gin.Context) {

	msg := &dbproto.BuycaiOptionsIssue{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("删除失败", http.StatusForbidden, c)
		return
	}

	//已经销售的无法删除.
	issue, err := srv.Db().QueryLotteryBuycaiOptionsById(context.Background(), msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取错误", http.StatusForbidden, c)
		return
	}

	if issue.StartTime <= time.Now().Unix() {
		srv.Json("当前这期，已经销售，不能够删除", http.StatusForbidden, c)
		return
	}

	_, err = srv.Db().DeleteIssueById(context.Background(), msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("删除失败", http.StatusForbidden, c)
		return
	}

	srv.Log(c, msg, "删除彩票配置", core.DELETE_OPERATION)
	srv.Json("删除成功", http.StatusOK, c)
}

//获取期号的详细数据.
func (srv *LotteryService) HandleLotteryBuycaiOptionsDetail(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("获取失败")
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	lottery := c.Param("lottery")
	if len(lottery) < 1 {
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	msg := &dbproto.BuycaiOptionsIssue{
		Id:      int64(id),
		Lottery: lottery,
	}
	issue, err := srv.Db().QueryLotteryBuycaiOptionsById(context.Background(), msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取错误", http.StatusForbidden, c)
		return
	}

	srv.Log(c, id, "获取期号的详细数据", core.QUERY_OPERATION)
	srv.Json(issue, http.StatusOK, c)
}

//更新彩票期号数据.
func (srv *LotteryService) HandleLotteryBuycaiOptionsUpdate(c *gin.Context) {

	msg := &dbproto.BuycaiOptionsUpdateIssue{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("更新失败", http.StatusForbidden, c)
		return
	}

	options := &dbproto.BuycaiOptionsIssue{
		Id:      msg.GetId(),
		Lottery: msg.GetLottery(),
	}
	//已经销售的无法删除.
	issue, err := srv.Db().QueryLotteryBuycaiOptionsById(context.Background(), options)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取错误", http.StatusForbidden, c)
		return
	}

	//已经开始销售的不能够修改.
	if issue.GetStartTime() <= time.Now().Unix() {
		srv.Json("当前这期，已经销售，不能够修改", http.StatusForbidden, c)
		return
	}

	_, err = srv.Db().UpdateLotteryBuycaiOptionsById(context.Background(), msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("更新失败", http.StatusForbidden, c)
		return
	}

	srv.Log(c, msg, "更新期号数据", core.UPDATE_OPERATION)
	srv.Json("更新成功", http.StatusOK, c)
}

//初始化彩票配置数据.
func (srv *LotteryService) HandleLotteryBuycaiOptionsInit(c *gin.Context) {

	msg := &dbproto.BuycaiOptionsUpdateIssue{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("添加失败", http.StatusForbidden, c)
		return
	}

	_, err = srv.Db().InsertLotteryBuycaiOptions(context.Background(), msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("添加失败", http.StatusForbidden, c)
		return
	}

	srv.Log(c, msg, "初始化彩票期号数据", core.ADD_OPERATION)
	srv.Json("添加成功", http.StatusOK, c)
}

//中奖彩票列表.
func (srv *LotteryService) HandleLotteryOpenList(c *gin.Context) {

	lottery := lotteryList()

	list := &dbproto.WinningLotteryList{}
	for _, v := range lottery {
		if v.Type != int(dbproto.LotteryType_Comp) {
			lottery := &dbproto.WinningLottery{
				Lottery: v.Code,
			}
			list.List = append(list.List, lottery)
		}
	}

	winningNoList, err := srv.Db().QueryWinningLotteryList(context.Background(), list)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	winningList := winningNoList.List
	for k, v := range winningList {
		lottery_name, err := getLotteryNameByCode(v.GetLottery())
		if err != nil {
			log.Printf("%+v\n", err)
			srv.Json("获取失败", http.StatusForbidden, c)
			return
		}
		winningList[k].LotteryName = lottery_name
	}
	winningNoList.List = winningList

	srv.Log(c, lottery, "获取没种彩票最新的中奖数据", core.QUERY_OPERATION)
	srv.Json(winningNoList, http.StatusOK, c)
}

//根据code获取彩票名称.
func getLotteryNameByCode(str string) (string, error) {

	lottery := lotteryList()

	for _, v := range lottery {
		if v.Code == str {
			return v.Name, nil
		}
	}
	return "", errors.New("未获取到彩票名称")
}

//根据期号查询开奖信息.
func (srv *LotteryService) HandleLotteryOpenSearch(c *gin.Context) {

	msg := &dbproto.LotteryWinningNo{}
	if err := c.BindJSON(&msg); err != nil {
		log.Printf("%+v\n", err)
		srv.Json("查询失败", http.StatusForbidden, c)
		return
	}

	winning, err := srv.Db().QueryWinningByIssue(context.Background(), msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("查询失败", http.StatusForbidden, c)
		return
	}

	list := winning.GetList()
	for k, v := range list {
		lottery_name, err := getLotteryNameByCode(v.GetLottery())
		if err != nil {
			log.Printf("%+v\n", err)
			srv.Json("获取失败", http.StatusForbidden, c)
			return
		}
		list[k].LotteryName = lottery_name
	}

	winning.List = list

	srv.Json(winning, http.StatusOK, c)
}

//获取开奖的历史数据.
func (srv *LotteryService) HandleLotteryOpenHistory(c *gin.Context) {

	msg := &dbproto.WinningListByLottery{}
	if err := c.BindJSON(&msg); err != nil {
		log.Print("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	winning, err := srv.Db().QueryWinningHistoryListByLottery(context.Background(), msg)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	list := winning.GetList()
	for k, v := range list {
		lottery_name, err := getLotteryNameByCode(v.GetLottery())
		if err != nil {
			log.Printf("%+v\n", err)
			srv.Json("获取失败", http.StatusForbidden, c)
			return
		}
		list[k].LotteryName = lottery_name
	}

	winning.Page = msg.GetPage()
	winning.Size = msg.GetSize()
	winning.Lottery = msg.GetLottery()
	winning.Issue = msg.GetIssue()
	winning.List = list
	winning.StartTime = msg.GetStartTime()
	winning.EndTime = msg.GetEndTime()

	srv.Json(winning, http.StatusOK, c)
}
