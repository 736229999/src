package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	"github.com/caojunxyz/gotu"
	apiproto "github.com/caojunxyz/mimi-api/proto"
	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/lottery"
	"github.com/caojunxyz/mimi-server/lottery/site/zhongfu"
	"github.com/caojunxyz/mimi-server/lottery/validator"
	"github.com/caojunxyz/mimi-server/utils"
)

const (
	VENDOR_ZF = "中福民彩"
)

type BuycaiAgent struct {
	sync.RWMutex
	dbUc             dbproto.DbUsercenterAgentClient
	dbBuycai         dbproto.DbBuycaiAgentClient
	id               apiproto.LotteryId                  // 彩种id
	saleList         []*dbproto.BuycaiSaleIssue          // 在售期号列表
	saleIssues       map[string]*dbproto.BuycaiSaleIssue // 销售期号表
	curIssueIdx      int                                 // 当前期号索引
	endIssueIdx      int                                 // 可售期号结束索引
	chanUserOrders   chan *dbproto.BuycaiUserOrder       // 用户订单
	chanVendorOrders chan *dbproto.BuycaiVendorOrder     // 投注订单
	chaseUserOrders  []*dbproto.BuycaiUserOrder          // 追号订单
	chanOpen         chan *dbproto.BuycaiSaleIssue       // 开奖通道
}

func NewAgent(id apiproto.LotteryId, dbUc dbproto.DbUsercenterAgentClient, dbBuycai dbproto.DbBuycaiAgentClient /*, dbOpencai dbproto.DbOpencaiAgentClient */) *BuycaiAgent {
	return &BuycaiAgent{
		dbUc:             dbUc,
		dbBuycai:         dbBuycai,
		id:               id,
		chanUserOrders:   make(chan *dbproto.BuycaiUserOrder, 16),
		chanVendorOrders: make(chan *dbproto.BuycaiVendorOrder, 16),
		curIssueIdx:      -1,
		endIssueIdx:      -1,
		saleList:         make([]*dbproto.BuycaiSaleIssue, 0),
		saleIssues:       make(map[string]*dbproto.BuycaiSaleIssue),
		chaseUserOrders:  make([]*dbproto.BuycaiUserOrder, 0),
		chanOpen:         make(chan *dbproto.BuycaiSaleIssue, 1),
	}
}

func (agt *BuycaiAgent) AddToChaseList(userOrder *dbproto.BuycaiUserOrder) {
	agt.chaseUserOrders = append(agt.chaseUserOrders, userOrder)
}

func (agt *BuycaiAgent) RemoveFromChaseList(userOrderId int64) bool {
	for i, v := range agt.chaseUserOrders {
		if v.Id == userOrderId {
			agt.chaseUserOrders = append(agt.chaseUserOrders[:i], agt.chaseUserOrders[i+1:]...)
			return true
		}
	}
	return false
}

func (agt *BuycaiAgent) PickChaseOrders(issue string) []*dbproto.BuycaiUserOrder {
	result := []*dbproto.BuycaiUserOrder{}
	for _, v := range agt.chaseUserOrders {
		idx := v.GetChaseNo() + 1
		if idx < v.GetIssueNum() {
			im := v.GetIssues()[idx]
			if im.Issue == issue {
				result = append(result, v)
			}
		}
	}
	return result
}

func (agt *BuycaiAgent) IsOpenSale() bool {
	return (agt.curIssueIdx >= 0 && agt.curIssueIdx < len(agt.saleList))
}

func (agt *BuycaiAgent) GetCurIssue() *dbproto.BuycaiSaleIssue {
	if agt.curIssueIdx >= len(agt.saleList) || agt.curIssueIdx < 0 {
		return nil
	}
	return agt.saleList[agt.curIssueIdx]
}
func (agt *BuycaiAgent) GetNextIssue() *dbproto.BuycaiSaleIssue {
	idx := agt.curIssueIdx + 1
	if idx >= len(agt.saleList) || idx < 0 {
		return nil
	}
	return agt.saleList[idx]
}

func (agt *BuycaiAgent) GetLastIssue() *dbproto.BuycaiSaleIssue {
	idx := agt.curIssueIdx - 1
	if idx >= len(agt.saleList) || idx < 0 {
		return nil
	}
	return agt.saleList[idx]
}

func (agt *BuycaiAgent) GetSaleIssue(issue string) *dbproto.BuycaiSaleIssue {
	agt.RLock()
	ret := agt.saleIssues[issue]
	agt.RUnlock()
	if ret == nil {
		for _, v := range agt.saleList {
			if v.Issue == issue {
				ret = v
				break
			}
		}

		if ret == nil {
			arg := &dbproto.BuycaiQueryIssueArg{Code: lottery.GetConfig(agt.id).Code, Issue: issue}
			saleIssue, err := agt.dbBuycai.BuycaiQueryIssue(context.Background(), arg)
			if err != nil {
				log.Println(err, issue)
				return nil
			}
			ret = saleIssue
		}

		if ret != nil {
			agt.Lock()
			agt.saleIssues[issue] = ret
			agt.Unlock()
		}
	}
	return ret
}

func (agt *BuycaiAgent) GetOnSaleList() []*dbproto.BuycaiSaleIssue {
	// 低频最多销售最近50期, 高频最多销售当天期号
	// log.Println(agt.id, agt.curIssueIdx, agt.endIssueIdx, len(agt.saleList))
	ret := agt.saleList[agt.curIssueIdx:agt.endIssueIdx]
	// log.Println(agt.id, agt.curIssueIdx, agt.endIssueIdx, len(agt.saleList), len(ret))
	return ret
}

func (agt *BuycaiAgent) updateEndIssueIdx() {
	log.Println("updateEndIssueIdx...")
	cfg := lottery.GetConfig(agt.id)
	log.Println(cfg.Type, apiproto.LotteryType_LowFreq, apiproto.LotteryType_HighFreq, len(agt.saleList))
	if cfg.Type == apiproto.LotteryType_LowFreq {
		agt.endIssueIdx = agt.curIssueIdx + 50
		if agt.endIssueIdx > len(agt.saleList) {
			agt.endIssueIdx = len(agt.saleList)
		}
	} else if cfg.Type == apiproto.LotteryType_HighFreq {
		now := time.Now()
		agt.endIssueIdx = agt.curIssueIdx + 1
		for i := 0; i < cfg.DayMaxNo; i++ {
			idx := agt.curIssueIdx + i
			if idx >= len(agt.saleList) || idx < 0 || len(agt.saleList) == 0 {
				break
			}
			si := agt.saleList[idx]
			startTime := time.Unix(si.StartTime, 0)
			if !gotu.IsSameDay(now, startTime) {
				break
			}
			agt.endIssueIdx = idx + 1
		}
	}
	if agt.curIssueIdx >= 0 && agt.curIssueIdx < len(agt.saleList) {
		agt.chanOpen <- agt.saleList[agt.curIssueIdx]
	}
	log.Printf("可购结束期号: %s %d %d %d\n", cfg.Name, agt.curIssueIdx, agt.endIssueIdx, len(agt.saleList))
}

// 派奖
func (agt *BuycaiAgent) distributeBonus(cfg *lottery.Config, issue string, list []*dbproto.BuycaiVendorOrder) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	log.Println("distributeBonus:", cfg.Code, len(list))
	for _, v := range list {
		log.Println(agt.id, v.GetId())
		orderNo := lottery.MakeOrderNo(cfg, issue, v.GetId())
		status, bonus, err := zhongfu.VerifyOrder(orderNo, v.GetVendorRespId(), fmt.Sprint(v.GetAccountId()))
		if err != nil {
			log.Println(err, v.GetId())
			continue
		}
		log.Println(agt.id, v.GetId(), status, bonus, err)
		arg := &dbproto.BuycaiUpdateVendorStatusArg{
			VendorOrderId: v.GetId(),
		}
		switch status {
		case zhongfu.WIN:
			log.Println("中奖:", v.GetId(), bonus)
			arg.Status = dbproto.VendorOrderStatus_VO_Win
			arg.WinMoney = bonus
			agt.dbUc.BuycaiUpdateVendorOrder(context.Background(), arg)
		case zhongfu.NOT_WIN:
			log.Println("未中奖:", v.GetId())
			arg.Status = dbproto.VendorOrderStatus_VO_NotWin
			agt.dbUc.BuycaiUpdateVendorOrder(context.Background(), arg)
		}
	}
}

// 查询开奖号码
// func (agt *BuycaiAgent) queryOpenBalls(cfg *lottery.Config, issue string) string {
// 	log.Println("queryOpenBalls...")
// 	bt := time.Now()
// 	arg := &dbproto.OpencaiQueryArg{Code: cfg.Code, Args: []string{issue}}
// 	openInfo, err := agt.dbOpencai.OpencaiQueryByIssue(context.Background(), arg)
// 	if err != nil {
// 		ticker := time.NewTicker(time.Second * 5)
// 		for _ = range ticker.C {
// 			openInfo, err = agt.dbOpencai.OpencaiQueryByIssue(context.Background(), arg)
// 			if err == nil {
// 				ticker.Stop()
// 				log.Println("queryOpenBalls:", cfg.Id, issue, openInfo.Balls, time.Now().Sub(bt))
// 				return openInfo.Balls
// 			} else {
// 				log.Println(err)
// 			}
// 		}
// 	}
// 	log.Println("queryOpenBalls:", cfg.Id, issue, openInfo.Balls, time.Now().Sub(bt))
// 	return openInfo.Balls
// }

// 更新销售期号开奖号码
func (agt *BuycaiAgent) setSaleIssueOpenBalls(cfg *lottery.Config, issue string, openBalls string) {
	log.Println("setSaleIssueOpenBalls...")
	if si := agt.GetSaleIssue(issue); si != nil {
		si.OpenBalls = openBalls
		arg := &dbproto.BuycaiUpsertIssueArg{
			Code:      cfg.Code,
			SaleIssue: &dbproto.BuycaiSaleIssue{Issue: issue, OpenBalls: openBalls},
		}
		_, err := agt.dbBuycai.BuycaiUpdateOpenBalls(context.Background(), arg)
		if err != nil {
			log.Panic(err)
		}
		log.Println("设置成功!", cfg.Code, issue)
	} else {
		log.Panicf("无效期号: %v, %v", cfg.Id, issue)
	}
}

// 处理开奖
func (agt *BuycaiAgent) handleOpen() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	cfg := lottery.GetConfig(agt.id)
	for {
		openIssue := <-agt.chanOpen
		if openIssue != nil {
			openTime := time.Unix(openIssue.OpenTime, 0)
			dur := openTime.Sub(time.Now())
			log.Printf("%s %s期 %v 后开奖...\n", agt.id, openIssue.Issue, dur)
			timer := time.NewTimer(dur)
			<-timer.C
			list := agt.loadWaitOpenVendorOrders()
			if cfg.Type == apiproto.LotteryType_LowFreq {
				dur = time.Hour * 2
			} else if cfg.Type == apiproto.LotteryType_HighFreq {
				dur = time.Minute * 3
			}
			log.Printf("%s %s期延迟%v开奖, 购彩订单%d个...\n", agt.id, openIssue.Issue, dur, len(list))
			timer = time.NewTimer(dur)
			<-timer.C
			// TODO: 查询开奖号码
			if len(list) > 0 {
				// 派奖
				log.Printf("%s %s开始派奖\n", cfg.Code, openIssue.Issue)
				go agt.distributeBonus(cfg, openIssue.Issue, list)
			}
		}
	}
}

func (agt *BuycaiAgent) handleCurIssue() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	for {
		nextIssue := agt.GetNextIssue()
		if nextIssue == nil {
			log.Printf("%s没有下一期了\n", lottery.GetConfig(agt.id).Name)
			return
		}
		now := time.Now()
		startTime := time.Unix(nextIssue.StartTime, 0)
		if startTime.Before(now) {
			agt.curIssueIdx += 1
		} else {
			timer := time.NewTimer(startTime.Sub(now))
			<-timer.C
			agt.curIssueIdx += 1
		}

		agt.updateEndIssueIdx()
		if curIssue := agt.GetCurIssue(); curIssue != nil {
			log.Printf("%s当前期号: %+v\n", lottery.GetConfig(agt.id).Name, curIssue)
			// list := agt.PickOrders(curIssue.Issue)

		}
	}
}

func (agt *BuycaiAgent) Run() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	agt.loadOnSaleIssueList()
	if agt.curIssueIdx == -1 {
		agt.initSaleIssues()
		agt.loadOnSaleIssueList()
	}
	go agt.handleVendorOrders()
	go agt.handleUserOrders()
	go agt.loadUnfinishedUserOrder()
	go agt.loadUnfinishedVendorOrder()
	go agt.handleCurIssue()
	go agt.handleOpen()
}

// 加载在售期号列表
func (agt *BuycaiAgent) loadOnSaleIssueList() {
	log.Println("loadOnSaleIssueList:", agt.id)
	now := time.Now()
	agt.curIssueIdx = -1
	agt.saleList = make([]*dbproto.BuycaiSaleIssue, 0)
	arg := &dbproto.StringValue{Value: lottery.GetConfig(agt.id).Code}
	ret, err := agt.dbBuycai.BuycaiQuerySaleList(context.Background(), arg)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("可购期号:", agt.id, len(ret.GetList()))
		for i, v := range ret.GetList() {
			startTime := time.Unix(v.GetStartTime(), 0)
			endTime := time.Unix(v.GetEndTime(), 0)
			agt.saleList = append(agt.saleList, v)
			if now.After(startTime) && now.Before(endTime) {
				agt.curIssueIdx = i
				log.Printf("%s当前期号: %+v, %d\n", lottery.GetConfig(agt.id).Name, v.GetIssue(), agt.curIssueIdx)
			}
		}
	}

	// 处理空档期
	if agt.curIssueIdx == -1 && len(agt.saleList) > 0 {
		v := agt.saleList[0]
		startTime := time.Unix(v.GetStartTime(), 0)
		if startTime.After(now) {
			agt.curIssueIdx = 0
		}
	}
	log.Println(agt.id, agt.curIssueIdx, agt.endIssueIdx, time.Now())
	agt.updateEndIssueIdx()
	log.Println("loadOnSaleIssueList:", agt.id, len(agt.saleList), time.Now().Sub(now))
}

func (agt *BuycaiAgent) isIssueOnSale(issue string) bool {
	now := time.Now().Unix()
	for _, v := range agt.saleList {
		if issue == v.Issue && v.EndTime > now {
			return true
		}
	}
	return false
}

// 检查订单期号
func (agt *BuycaiAgent) checkIssues(order *apiproto.BuycaiOrder) bool {
	imList := order.GetIssues()
	for _, v := range imList {
		issue := v.GetIssue()
		if !agt.isIssueOnSale(issue) {
			log.Println("期号不可售:", issue)
			return false
		}
	}
	return true
}

// 检查账户资金是否充足
// 返回: 订单使用彩金、余额、购彩券抵扣金额、资金是否充足
func (agt *BuycaiAgent) checkMoney(accountId int64, order *apiproto.BuycaiOrder) (cai, balance, ticketSub float64, ok bool) {
	sumMoney := order.GetSumMoney()
	ticketId := order.GetTicketId()
	if ticketId > 0 {
		ticket, err := agt.dbUc.QueryBuycaiTicket(context.Background(), &dbproto.IntValue{Value: ticketId})
		if err != nil {
			log.Println(err)
			return
		}
		if !utils.IsBuycaiTicketCanUse(accountId, int32(agt.id), sumMoney, ticket) {
			log.Println(accountId, agt.id, sumMoney, ticketId)
			return
		}
		ticketSub = float64(ticket.GetUseSub())
	}
	fund, err := agt.dbUc.QueryFund(context.Background(), &dbproto.IntValue{Value: accountId})
	if err != nil {
		log.Println(accountId, err)
		return
	}

	cai = order.GetCai()
	balance = order.GetBalance()
	// log.Println(cai, balance, sumMoney, ticketSub, !validator.IsEqualMoney(cai+balance+ticketSub, sumMoney), cai+balance+ticketSub, sumMoney)
	if cai > fund.GetCai() || balance > fund.GetBalance() || !validator.IsEqualMoney(cai+balance+ticketSub, sumMoney) {
		log.Printf("资金错误: %d (%+v) --> %f (%f, %f, %f)\n", accountId, fund, sumMoney, cai, balance, ticketSub)
		return
	}
	return cai, balance, ticketSub, true
}

// 保存购彩用户订单
func (agt *BuycaiAgent) saveUserOrder(accountId int64, order *apiproto.BuycaiOrder) (*dbproto.BuycaiUserOrder, error) {
	cai, balance, ticketSub, ok := agt.checkMoney(accountId, order)
	if !ok {
		return nil, fmt.Errorf("资金验证失败")
	}
	sumMoney := cai + balance + ticketSub
	imList := order.GetIssues()
	schemeList := order.GetSchemeList()
	userOrder := &dbproto.BuycaiUserOrder{
		AccountId:      accountId,
		LotteryId:      int32(agt.id),
		OrderTime:      time.Now().Unix(),
		IssueNum:       int32(len(imList)),
		Cai:            cai,
		Balance:        balance,
		SumMoney:       sumMoney,
		TicketSubMoney: ticketSub,
		TicketId:       order.GetTicketId(),
		Status:         int32(dbproto.UserOrderStatus_UO_Doing),
		IsWinStop:      order.GetIsWinStop(),
	}

	for _, v := range imList {
		im := &dbproto.IssueMultiple{Issue: v.GetIssue(), Multiple: v.GetMultiple()}
		userOrder.Issues = append(userOrder.Issues, im)
	}

	for _, v := range schemeList {
		scheme := &dbproto.BuycaiScheme{Type: v.GetType(), Num: v.GetNum(), Money: v.GetMoney()}
		selectBalls := v.GetSelectBalls()
		scheme.SelectBalls = make(map[string]*dbproto.Balls)
		for k, ball := range selectBalls {
			ball0 := &dbproto.Balls{}
			for _, no := range ball.List {
				ball0.List = append(ball0.List, no)
			}
			scheme.SelectBalls[k] = ball0
		}
		userOrder.SchemeList = append(userOrder.SchemeList, scheme)
	}
	ret, err := agt.dbUc.BuycaiInsertUserOrder(context.Background(), userOrder)
	if err != nil {
		log.Println(accountId, err)
		return nil, fmt.Errorf("创建用户订单失败")
	}
	userOrder.Id = ret.GetValue()
	return userOrder, nil
}

// 计算购彩方案列表总注数
func getSchemeListSumNum(schemeList []*dbproto.BuycaiScheme) int32 {
	num := int32(0)
	for _, v := range schemeList {
		num += v.GetNum()
	}
	return num
}

// 计算购彩方案列表总金额
func getSchemeListSumMoney(schemeList []*dbproto.BuycaiScheme) float64 {
	money := float64(0)
	for _, v := range schemeList {
		money += v.GetMoney()
	}
	return money
}

// 计算投注订单金额构成
// 返回本次投注使用的彩金、余额、没有抵扣的总金额
func (agt *BuycaiAgent) calVendorOrderMoney(userOrder *dbproto.BuycaiUserOrder) (cai, balance, sumMoney float64, err error) {
	// 非追号订单才能使用购彩券(即只有一期)
	ticketSub := userOrder.GetTicketSubMoney()
	if ticketSub > 0 {
		cai = userOrder.GetCai()
		balance = userOrder.GetBalance()
		sumMoney = userOrder.GetSumMoney()
		return
	}

	accountId := userOrder.GetAccountId()
	var fund *dbproto.Fund
	fund, err = agt.dbUc.QueryFund(context.Background(), &dbproto.IntValue{Value: accountId})
	if err != nil {
		log.Println(err, accountId)
		return
	}

	chaseNo := userOrder.GetChaseNo()
	im := userOrder.GetIssues()[chaseNo]
	multiple := im.GetMultiple()
	sumMoney = getSchemeListSumMoney(userOrder.GetSchemeList()) * float64(multiple)
	freezeCai := fund.GetFreezeCai()
	freezeBalance := fund.GetFreezeBalance()
	if sumMoney > freezeCai+freezeBalance {
		err = fmt.Errorf("冻结资金不足: %d, %f, %f, %f, %f", accountId, sumMoney, freezeCai, freezeBalance)
		log.Println(err)
		return
	}

	cai = sumMoney
	if freezeCai < cai {
		cai = freezeCai
	}
	balance = sumMoney - cai
	return cai, balance, sumMoney, nil
}

// 选择投注站
func (agt *BuycaiAgent) selectVendor() string {
	return VENDOR_ZF
}

// 生成投注订单
func (agt *BuycaiAgent) makeVendorOrder(userOrder *dbproto.BuycaiUserOrder) *dbproto.BuycaiVendorOrder {
	if userOrder.GetStatus() != int32(dbproto.UserOrderStatus_UO_Doing) {
		return nil
	}

	accountId := userOrder.GetAccountId()
	userOrderId := userOrder.GetId()
	issueNum := userOrder.GetIssueNum()
	chaseNo := userOrder.GetChaseNo()
	imList := userOrder.GetIssues()
	if chaseNo >= issueNum {
		log.Println("error: 期号已购完:", userOrderId)
		// TODO: 更新数据库订单信息
		dbArg := &dbproto.BuycaiUpdateUserStatusArg{UserOrderId: userOrderId, Status: dbproto.UserOrderStatus_UO_FinishStop}
		agt.dbUc.BuycaiUpdateUserOrder(context.Background(), dbArg)
		return nil
	}

	now := time.Now()
	im := imList[chaseNo]
	issue := im.GetIssue()
	multiple := im.GetMultiple()
	curSale := agt.GetCurIssue()
	if issue != curSale.Issue || now.Unix() > curSale.EndTime {
		log.Printf("error: 期号不在销售，投注失败停止: %d, %s, %+v\n", userOrderId, issue, curSale)
		// TODO: 更新数据库订单信息
		dbArg := &dbproto.BuycaiUpdateUserStatusArg{UserOrderId: userOrderId, Status: dbproto.UserOrderStatus_UO_FailStop}
		agt.dbUc.BuycaiUpdateUserOrder(context.Background(), dbArg)
		return nil
	}

	cai, balance, sumMoney, err := agt.calVendorOrderMoney(userOrder)
	if err != nil {
		return nil
	}

	schemeList := userOrder.GetSchemeList()
	vendorOrder := &dbproto.BuycaiVendorOrder{
		AccountId:   accountId,
		UserOrderId: userOrderId,
		LotteryId:   int32(agt.id),
		Issue:       issue,
		SumNum:      getSchemeListSumNum(schemeList),
		Multiple:    multiple,
		Money:       sumMoney,
		Cai:         cai,
		Balance:     balance,
		ChaseNo:     chaseNo + 1,
		Vendor:      agt.selectVendor(),
		SchemeList:  schemeList,
		AddTime:     now.Unix(),
	}
	ret, err := agt.dbUc.BuycaiInsertVendorOrder(context.Background(), vendorOrder)
	if err != nil {
		log.Println(err)
		return nil
	}
	vendorOrder.Id = ret.GetValue()
	return vendorOrder
}

// 添加客户端购彩订单
func (agt *BuycaiAgent) AddOrder(accountId int64, order *apiproto.BuycaiOrder) (*dbproto.BuycaiUserOrder, error) {
	if !agt.IsOpenSale() {
		return nil, fmt.Errorf("停止销售")
	}
	if !agt.checkIssues(order) {
		return nil, fmt.Errorf("存在无效期号")
	}
	if !lottery.Validate(order, lottery.GetConfig(agt.id)) {
		return nil, fmt.Errorf("订单验证失败")
	}
	userOrder, err := agt.saveUserOrder(accountId, order)
	if err != nil {
		return nil, fmt.Errorf("下单失败")
	}
	agt.chanUserOrders <- userOrder
	return userOrder, nil
}

// 加载未完成的用户订单
func (agt *BuycaiAgent) loadUnfinishedUserOrder() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	log.Println("loadUnfinishedUserOrder...")
	arg := &dbproto.BuycaiQueryUserOrderArg{
		LotteryId:  int32(agt.id),
		StatusList: []dbproto.UserOrderStatus{dbproto.UserOrderStatus_UO_Doing},
	}
	stream, err := agt.dbUc.BuycaiQueryUserOrder(context.Background(), arg)
	if err != nil {
		log.Println(err)
		return
	}

	count := 0
	for {
		v, err := stream.Recv()
		if err == nil {
			count++
			agt.chanUserOrders <- v
			continue
		}
		if err != io.EOF {
			log.Println(err)
		}
		break
	}
	// log.Println(agt.id.String(), "count = ", count)
}

// 加载等待开奖的购彩订单
func (agt *BuycaiAgent) loadWaitOpenVendorOrders() []*dbproto.BuycaiVendorOrder {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	log.Println("loadWaitOpenVendorOrders...")
	arg := &dbproto.BuycaiQueryVendorOrderArg{
		LotteryId:  int32(agt.id),
		StatusList: []dbproto.VendorOrderStatus{dbproto.VendorOrderStatus_VO_BetSuccess},
	}
	stream, err := agt.dbUc.BuycaiQueryVendorOrder(context.Background(), arg)
	if err != nil {
		log.Println(err)
		return nil
	}

	result := []*dbproto.BuycaiVendorOrder{}
	for {
		v, err := stream.Recv()
		if err == nil {
			result = append(result, v)
			continue
		}
		if err != io.EOF {
			log.Println(err)
		}
		break
	}
	log.Println(agt.id, len(result))
	return result
}

// 加载未完成的购彩订单
func (agt *BuycaiAgent) loadUnfinishedVendorOrder() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	log.Println("loadUnfinishedVendorOrder...")
	arg := &dbproto.BuycaiQueryVendorOrderArg{
		LotteryId:  int32(agt.id),
		StatusList: []dbproto.VendorOrderStatus{dbproto.VendorOrderStatus_VO_NotBet},
	}
	stream, err := agt.dbUc.BuycaiQueryVendorOrder(context.Background(), arg)
	if err != nil {
		log.Println(err)
		return
	}

	count := 0
	for {
		v, err := stream.Recv()
		if err == nil {
			count++
			agt.chanVendorOrders <- v
			continue
		}
		if err != io.EOF {
			log.Println(err)
		}
		break
	}
	// log.Println(agt.id.String(), "count = ", count)
}

// 请求投注站投注
func (agt *BuycaiAgent) requestVendor(vendorOrder *dbproto.BuycaiVendorOrder) {
	log.Println("requestVendor...", vendorOrder.GetId(), vendorOrder.GetUserOrderId())
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	reqTime := time.Now()
	vendor := vendorOrder.GetVendor()
	if vendor == VENDOR_ZF {
		log.Println("中福投注:", vendorOrder.GetId())
		cfg := lottery.GetConfig(agt.id)
		if zfOrder := lottery.Convert(vendorOrder, cfg); zfOrder != nil {
			result, err := zhongfu.CommitOrder(zfOrder)
			if err != nil {
				log.Println("请求投注失败:", err)
				agt.chanVendorOrders <- vendorOrder
				return
			}
			respTime := time.Now()
			vendorOrderId := vendorOrder.GetId()
			arg := &dbproto.BuycaiUpdateVendorStatusArg{
				VendorOrderId:  vendorOrderId,
				VendorReqTime:  reqTime.Unix(),
				VendorRespTime: respTime.Unix(),
			}
			log.Printf("%+v, %v\n", result, err)
			if result.Status != "0" {
				if result.Status == "-809" {
					// 订单已存在
					log.Println("订单已存在:", vendorOrderId)
					// TODO:
				} else {
					// 投注失败
					log.Printf("投注失败: %+v\n%+v\n", zfOrder, result)
					arg.Status = dbproto.VendorOrderStatus_VO_BetFail
				}
			} else {
				log.Println("投注成功:", vendorOrder.GetId())
				arg.Status = dbproto.VendorOrderStatus_VO_BetSuccess
				arg.VendorRespId = result.ZfOrderId
				utils.AddCredits(agt.dbUc, vendorOrder.GetAccountId(), apiproto.CreditsTask_Buycai, fmt.Sprint(vendorOrderId), vendorOrder.GetMoney())
			}
			// TODO: 更新在中福账户余额
			// log.Printf("%d, %+v\n", vendorOrderId, arg)
			agt.dbUc.BuycaiUpdateVendorOrder(context.Background(), arg)
		} else {
			log.Println("转换订单失败:", vendorOrder.GetId())
		}
	}
}

func (agt *BuycaiAgent) handleUserOrders() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	for {
		select {
		case userOrder := <-agt.chanUserOrders:
			if userOrder.GetStatus() != int32(dbproto.UserOrderStatus_UO_Doing) {
				continue
			}
			if vendorOrder := agt.makeVendorOrder(userOrder); vendorOrder != nil {
				agt.chanVendorOrders <- vendorOrder
			}
		}
	}
}

func (agt *BuycaiAgent) handleVendorOrders() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	for {
		select {
		case vendorOrder := <-agt.chanVendorOrders:
			// log.Printf("%+v\n", vendorOrder)
			if vendorOrder.GetStatus() != int32(dbproto.VendorOrderStatus_VO_NotBet) {
				continue
			}
			go agt.requestVendor(vendorOrder)
		}
	}
}
