package usercenter

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"golang.org/x/net/context"

	. "github.com/caojunxyz/mimi-server/apps/dbagent/helper"
	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
)

// 普通订单才能使用购彩券
func useTicket(tx *sql.Tx, ticketId int64, orderId int64) error {
	esql := fmt.Sprintf("UPDATE %s SET order_id=%d WHERE id=%d", TABLE_TICKET, orderId, ticketId)
	_, err := tx.Exec(esql)
	return err
}

func (agt *DbUsercenterAgent) queryUserOrderTotalWinMoney(id int64) (float64, error) {
	esql := fmt.Sprintf("SELECT SUM(win_money) FROM %s WHERE user_order_id=%d", TABLE_BUYCAI_VENDOR_ORDER, id)
	var result float64
	err := agt.dbConn.QueryRow(esql).Scan(&result)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err, esql)
		return 0, err
	}
	return result, nil
}

// TODO: 存ticketId
func (agt *DbUsercenterAgent) BuycaiInsertUserOrder(ctx context.Context, arg *dbproto.BuycaiUserOrder) (*dbproto.IntValue, error) {
	varCai := arg.GetCai()
	varBalance := arg.GetBalance()
	sumMoney := arg.GetSumMoney()
	ticketId := arg.GetTicketId()
	ticketSubMoney := arg.GetTicketSubMoney()
	if varCai < 0 || varBalance < 0 || sumMoney <= 0 || ticketSubMoney < 0 {
		log.Println(varCai, varBalance, sumMoney, ticketSubMoney)
		return nil, ErrInvalidArgs
	}

	jsonIssues, err := json.Marshal(arg.GetIssues())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	jsonSchemeList, err := json.Marshal(arg.GetSchemeList())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	tx, err := agt.dbConn.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	accountId := arg.GetAccountId()
	lotteryId := arg.GetLotteryId()
	orderTime := arg.GetOrderTime()
	issueNum := arg.GetIssueNum()
	chaseNo := arg.GetChaseNo()
	status := arg.GetStatus()
	isWinStop := arg.GetIsWinStop()

	fieldList := []string{
		"account_id", "lottery_id", "issue_num", "chase_no",
		"cai", "balance", "sum_money", "issues", "scheme_list",
		"order_time", "ticket_sub_money", "status", "is_win_stop", "cost_cai", "cost_balance",
	}
	fmtList := []string{
		"%d", "%d", "%d", "%d",
		"%f", "%f", "%f", "'%s'", "'%s'",
		"%d", "%f", "%d", "%v", "%f", "%f",
	}
	values := []interface{}{
		accountId, lotteryId, issueNum, chaseNo,
		varCai, varBalance, sumMoney, jsonIssues, jsonSchemeList,
		orderTime, ticketSubMoney, status, isWinStop, float64(0), float64(0),
	}

	esql := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) RETURNING id", TABLE_BUYCAI_USER_ORDER, strings.Join(fieldList, ", "), strings.Join(fmtList, ", "))
	esql = fmt.Sprintf(esql, values...)
	var id int64
	if err := tx.QueryRow(esql).Scan(&id); err != nil {
		log.Println(err, esql)
		tx.Rollback()
		return nil, err
	}

	record := &dbproto.FundChangeRecord{
		AccountId:  accountId,
		VarBalance: -varBalance, VarFreezeBalance: varBalance,
		VarCai: -varCai, VarFreezeCai: varCai,
		ChangeType:    dbproto.FundChangeType_FREEZE,
		UserOrderId:   id,
		ChangeTime:    orderTime,
		ChangeComment: "订单冻结",
	}
	if err = changeFund(tx, record); err != nil {
		tx.Rollback()
		return nil, err
	}

	if ticketId > 0 {
		if err := useTicket(tx, ticketId, id); err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	return &dbproto.IntValue{Value: id}, tx.Commit()
}

func (agt *DbUsercenterAgent) BuycaiInsertVendorOrder(ctx context.Context, arg *dbproto.BuycaiVendorOrder) (*dbproto.IntValue, error) {
	varCai := arg.GetCai()
	varBalance := arg.GetBalance()
	money := arg.GetMoney()
	if varCai < 0 || varBalance < 0 || money <= 0 {
		log.Println(varCai, varBalance, money)
		return nil, ErrInvalidArgs
	}

	chaseNo := arg.GetChaseNo()
	if chaseNo < 1 {
		log.Println(chaseNo)
		return nil, ErrInvalidArgs
	}

	jsonSchemeList, err := json.Marshal(arg.GetSchemeList())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	tx, err := agt.dbConn.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	accountId := arg.GetAccountId()
	userOrderId := arg.GetUserOrderId()
	lotteryId := arg.GetLotteryId()
	issue := arg.GetIssue()
	sumNum := arg.GetSumNum()
	multiple := arg.GetMultiple()
	vendor := arg.GetVendor()
	vendorReqTime := arg.GetVendorReqTime()
	vendorRespTime := arg.GetVendorRespTime()
	vendorRespId := arg.GetVendorRespId()
	winMoney := arg.GetWinMoney()
	status := arg.GetStatus()
	statusDesc := arg.GetStatusDesc()
	addTime := arg.GetAddTime()

	fieldList := []string{
		"user_order_id", "account_id", "lottery_id", "issue", "sum_num", "multiple",
		"money", "cai", "balance", "chase_no", "vendor", "scheme_list", "add_time",
		"status", "status_desc", "win_money",
		"vendor_req_time", "vendor_resp_time", "vendor_resp_id",
	}
	fmtList := []string{
		"%d", "%d", "%d", "'%s'", "%d", "%d",
		"%f", "%f", "%f", "%d", "'%s'", "'%s'", "%d",
		"%d", "'%s'", "%f",
		"%d", "%d", "'%s'",
	}
	values := []interface{}{
		userOrderId, accountId, lotteryId, issue, sumNum, multiple,
		money, varCai, varBalance, chaseNo, vendor, jsonSchemeList, addTime,
		status, statusDesc, winMoney,
		vendorReqTime, vendorRespTime, vendorRespId,
	}

	esql := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) RETURNING id", TABLE_BUYCAI_VENDOR_ORDER, strings.Join(fieldList, ", "), strings.Join(fmtList, ", "))
	esql = fmt.Sprintf(esql, values...)
	var id int64
	if err := tx.QueryRow(esql).Scan(&id); err != nil {
		log.Println(err, esql)
		tx.Rollback()
		return nil, err
	}

	// 更新用户订单当前追号序号
	esql = fmt.Sprintf("UPDATE %s SET chase_no=%d WHERE id=%d RETURNING issue_num", TABLE_BUYCAI_USER_ORDER, chaseNo, userOrderId)
	var issueNum int32
	if err := tx.QueryRow(esql).Scan(&issueNum); err != nil {
		log.Println(err, esql)
		tx.Rollback()
		return nil, err
	}
	esql = fmt.Sprintf("UPDATE %s SET status=%d WHERE id=%d AND status=%d AND chase_no=issue_num",
		TABLE_BUYCAI_USER_ORDER, dbproto.UserOrderStatus_UO_FinishStop, userOrderId, dbproto.UserOrderStatus_UO_Doing)
	_, err = tx.Exec(esql)
	if err != nil {
		log.Println(err, esql)
		tx.Rollback()
		return nil, err
	}
	return &dbproto.IntValue{Value: id}, tx.Commit()
}

func (agt *DbUsercenterAgent) BuycaiQueryUserOrder(arg *dbproto.BuycaiQueryUserOrderArg, stream dbproto.DbUsercenterAgent_BuycaiQueryUserOrderServer) error {
	lotteryId := arg.GetLotteryId()
	statusList := arg.GetStatusList()
	fieldList := []string{
		"id, account_id", "lottery_id", "issue_num", "chase_no",
		"cai", "balance", "sum_money", "issues", "scheme_list",
		"order_time", "ticket_sub_money", "status", "is_win_stop", "cost_cai", "cost_balance",
	}

	statusFilterList := []string{}
	for _, v := range statusList {
		statusFilterList = append(statusFilterList, fmt.Sprintf("status=%d", v))
	}
	esql := fmt.Sprintf("SELECT %s FROM %s WHERE lottery_id=%d AND (%s)",
		strings.Join(fieldList, ", "), TABLE_BUYCAI_USER_ORDER, lotteryId, strings.Join(statusFilterList, " OR "))

	// log.Println("esql:", esql)
	rows, err := agt.dbConn.Query(esql)
	if err != nil {
		log.Println(err, esql)
		return err
	}

	for rows.Next() {
		order := &dbproto.BuycaiUserOrder{}
		var issues, schemeList []byte
		values := []interface{}{
			&order.Id, &order.AccountId, &order.LotteryId, &order.IssueNum, &order.ChaseNo,
			&order.Cai, &order.Balance, &order.SumMoney, &issues, &schemeList,
			&order.OrderTime, &order.TicketSubMoney, &order.Status, &order.IsWinStop, &order.CostCai, &order.CostBalance,
		}
		if err = rows.Scan(values...); err != nil {
			log.Println(err, esql)
			break
		} else {
			if err := json.Unmarshal(issues, &order.Issues); err != nil {
				log.Println(err, esql)
				return err
			}
			if err := json.Unmarshal(schemeList, &order.SchemeList); err != nil {
				log.Println(err, esql)
				return err
			}
			totalWinMoney, err := agt.queryUserOrderTotalWinMoney(order.Id)
			if err != nil {
				return err
			}
			order.TicketSubMoney = totalWinMoney
			stream.Send(order)
		}
	}

	if err = rows.Err(); err != nil {
		log.Println(err, esql)
	}
	return err
}

func (agt *DbUsercenterAgent) BuycaiQueryVendorOrder(arg *dbproto.BuycaiQueryVendorOrderArg, stream dbproto.DbUsercenterAgent_BuycaiQueryVendorOrderServer) error {
	lotteryId := arg.GetLotteryId()
	statusList := arg.GetStatusList()
	fieldList := []string{
		"id", "user_order_id", "account_id", "lottery_id", "issue", "sum_num", "multiple",
		"money", "cai", "balance", "chase_no", "vendor", "scheme_list", "add_time",
		"status", "status_desc", "win_money",
		"vendor_req_time", "vendor_resp_time", "vendor_resp_id",
	}

	statusFilterList := []string{}
	for _, v := range statusList {
		statusFilterList = append(statusFilterList, fmt.Sprintf("status=%d", v))
	}
	esql := fmt.Sprintf("SELECT %s FROM %s WHERE lottery_id=%d AND (%s)",
		strings.Join(fieldList, ", "), TABLE_BUYCAI_VENDOR_ORDER, lotteryId, strings.Join(statusFilterList, " OR "))

	rows, err := agt.dbConn.Query(esql)
	if err != nil {
		log.Println(err, esql)
		return err
	}

	for rows.Next() {
		order := &dbproto.BuycaiVendorOrder{}
		var schemeList []byte
		values := []interface{}{
			&order.Id, &order.UserOrderId, &order.AccountId, &order.LotteryId, &order.Issue, &order.SumNum, &order.Multiple,
			&order.Money, &order.Cai, &order.Balance, &order.ChaseNo, &order.Vendor, &schemeList, &order.AddTime,
			&order.Status, &order.StatusDesc, &order.WinMoney,
			&order.VendorReqTime, &order.VendorRespTime, &order.VendorRespId,
		}
		if err = rows.Scan(values...); err != nil {
			log.Println(err, esql)
			break
		} else {
			if err := json.Unmarshal(schemeList, &order.SchemeList); err != nil {
				log.Println(err, esql)
				return err
			}
			stream.Send(order)
		}
	}

	if err = rows.Err(); err != nil {
		log.Println(err, esql)
	}
	return err
}

func (agt *DbUsercenterAgent) buycaiVendorFail(vendorOrderId int64, vendorReqTime int64, vendorRespTime int64) (*dbproto.Nil, error) {
	log.Println("buycaiVendorFail:", vendorOrderId, vendorReqTime, vendorRespTime)
	tx, err := agt.dbConn.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var userOrderId int64
	var accountId int64
	now := time.Now().Unix()

	esql := fmt.Sprintf("UPDATE %s SET vendor_req_time=%d, vendor_resp_time=%d, status=%d WHERE id=%d RETURNING user_order_id, account_id",
		TABLE_BUYCAI_VENDOR_ORDER, vendorReqTime, vendorRespTime, dbproto.VendorOrderStatus_VO_BetFail, vendorOrderId)
	if err = tx.QueryRow(esql).Scan(&userOrderId, &accountId); err != nil {
		log.Println(err, esql)
		return nil, err
	}

	var cai, balance, costCai, costBalance float64
	esql = fmt.Sprintf("SELECT cai, balance, cost_cai, cost_balance FROM %s WHERE id=%d", TABLE_BUYCAI_USER_ORDER, userOrderId)
	if err = tx.QueryRow(esql).Scan(&cai, &balance, &costCai, &costBalance); err != nil {
		log.Println(err, esql)
		return nil, err
	}

	// 取消订单
	esql = fmt.Sprintf("UPDATE %s SET status=%d WHERE id=%d", TABLE_BUYCAI_USER_ORDER, dbproto.UserOrderStatus_UO_FailStop, userOrderId)
	_, err = tx.Exec(esql)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	varCai := cai - costCai
	varBalance := balance - costBalance
	record := &dbproto.FundChangeRecord{
		AccountId:  accountId,
		VarBalance: varBalance, VarFreezeBalance: -varBalance,
		VarCai: varCai, VarFreezeCai: -varCai,
		VendorOrderId: vendorOrderId,
		ChangeType:    dbproto.FundChangeType_UNFREEZE,
		ChangeTime:    now,
		ChangeComment: "购彩失败解冻",
	}

	if err = changeFund(tx, record); err != nil {
		tx.Rollback()
		return nil, err
	}
	return &dbproto.Nil{}, tx.Commit()
}

func (agt *DbUsercenterAgent) buycaiVendorSuccess(vendorOrderId int64, vendorReqTime int64, vendorRespTime int64, vendorRespId string) (*dbproto.Nil, error) {
	log.Println("buycaiVendorSuccess:", vendorOrderId, vendorReqTime, vendorRespTime, vendorRespId)
	tx, err := agt.dbConn.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var userOrderId int64
	var accountId int64
	var varCai, varBalance float64

	esql := fmt.Sprintf("UPDATE %s SET vendor_req_time=%d, vendor_resp_time=%d, vendor_resp_id=%s, status=%d WHERE id=%d RETURNING user_order_id, account_id, cai, balance",
		TABLE_BUYCAI_VENDOR_ORDER, vendorReqTime, vendorRespTime, vendorRespId, dbproto.VendorOrderStatus_VO_BetSuccess, vendorOrderId)
	if err = tx.QueryRow(esql).Scan(&userOrderId, &accountId, &varCai, &varBalance); err != nil {
		log.Println(err, esql)
		return nil, err
	}

	var chaseNo, issueNum, userOrderStatus int32
	// 增加订单已消费金额
	esql = fmt.Sprintf("UPDATE %s SET cost_cai=%s.cost_cai + %f, cost_balance=%s.cost_balance + %f WHERE id=%d RETURNING chase_no, issue_num, status",
		TABLE_BUYCAI_USER_ORDER, TABLE_BUYCAI_USER_ORDER, varCai, TABLE_BUYCAI_USER_ORDER, varBalance, userOrderId)
	if err = tx.QueryRow(esql).Scan(&chaseNo, &issueNum, &userOrderStatus); err != nil {
		log.Println(err, esql)
		tx.Rollback()
		return nil, err
	}

	if userOrderStatus == int32(dbproto.UserOrderStatus_UO_Doing) && chaseNo == issueNum {
		// 正常结束停止
		esql = fmt.Sprintf("UPDATE %s SET status=%d WHERE id=%d", TABLE_BUYCAI_USER_ORDER, dbproto.UserOrderStatus_UO_FinishStop, userOrderId)
		_, err = tx.Exec(esql)
		if err != nil {
			log.Println(err, esql)
			tx.Rollback()
			return nil, err
		}
	}

	// 解冻
	record := &dbproto.FundChangeRecord{
		AccountId:  accountId,
		ChangeType: dbproto.FundChangeType_UNFREEZE,
		VarBalance: varBalance, VarFreezeBalance: -varBalance,
		VarCai: varCai, VarFreezeCai: -varCai,
		VendorOrderId: vendorOrderId,
		ChangeTime:    vendorReqTime,
		ChangeComment: "购彩解冻",
	}
	if err = changeFund(tx, record); err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}

	// 购彩
	record = &dbproto.FundChangeRecord{
		AccountId:     accountId,
		ChangeType:    dbproto.FundChangeType_BUYCAI,
		VarBalance:    -varBalance,
		VarCai:        -varCai,
		VendorOrderId: vendorOrderId,
		ChangeTime:    vendorReqTime,
		ChangeComment: "购彩",
	}
	if err = changeFund(tx, record); err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}
	return &dbproto.Nil{}, tx.Commit()
}

func (agt *DbUsercenterAgent) buycaiVendorWin(vendorOrderId int64, winMoney float64) (*dbproto.Nil, error) {
	log.Println("buycaiVendorWin:", vendorOrderId, winMoney)
	tx, err := agt.dbConn.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	now := time.Now().Unix()
	var userOrderId int64
	var accountId int64

	esql := fmt.Sprintf("UPDATE %s SET status=%d, win_money=%f WHERE id=%d RETURNING user_order_id, account_id",
		TABLE_BUYCAI_VENDOR_ORDER, dbproto.VendorOrderStatus_VO_Win, winMoney, vendorOrderId)
	if err = tx.QueryRow(esql).Scan(&userOrderId, &accountId); err != nil {
		log.Println(err, esql)
		tx.Rollback()
		return nil, err
	}

	var issueNum, status int32
	var isWinStop bool
	var cai, balance, costCai, costBalance float64
	esql = fmt.Sprintf("SELECT issue_num, status, is_win_stop, cai, balance, cost_cai, cost_balance FROM %s WHERE id=%d", TABLE_BUYCAI_USER_ORDER, userOrderId)
	if err := tx.QueryRow(esql).Scan(&issueNum, &status, &isWinStop, &cai, &balance, &costCai, &costBalance); err != nil {
		log.Println(err, esql)
		tx.Rollback()
		return nil, err
	}

	// 中奖追号停止
	if status == int32(dbproto.UserOrderStatus_UO_Doing) && issueNum > 1 && isWinStop {
		esql = fmt.Sprintf("UPDATE %s SET status=%d WHERE id=%d", TABLE_BUYCAI_USER_ORDER, dbproto.UserOrderStatus_UO_FinishStop, userOrderId)
		_, err = tx.Exec(esql)
		if err != nil {
			log.Println(err, esql)
			tx.Rollback()
			return nil, err
		}
		varCai := cai - costCai
		varBalance := balance - costBalance
		record := &dbproto.FundChangeRecord{
			AccountId:  accountId,
			ChangeType: dbproto.FundChangeType_UNFREEZE,
			VarBalance: varBalance, VarFreezeBalance: -varBalance,
			VarCai: varCai, VarFreezeCai: -varCai,
			UserOrderId:   userOrderId,
			ChangeTime:    now,
			ChangeComment: "中奖停止追号",
		}
		if err = changeFund(tx, record); err != nil {
			log.Println(err)
			tx.Rollback()
			return nil, err
		}
	}

	// 派奖
	record := &dbproto.FundChangeRecord{
		AccountId:     accountId,
		ChangeType:    dbproto.FundChangeType_WIN,
		VarBalance:    winMoney,
		VendorOrderId: vendorOrderId,
		ChangeTime:    now,
	}
	if err = changeFund(tx, record); err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}
	return &dbproto.Nil{}, tx.Commit()
}

func (agt *DbUsercenterAgent) buycaiVendorNotWin(vendorOrderId int64) (*dbproto.Nil, error) {
	log.Println("buycaiVendorNotWin:", vendorOrderId)
	esql := fmt.Sprintf("UPDATE %s SET status=%d WHERE id=%d", TABLE_BUYCAI_VENDOR_ORDER, dbproto.VendorOrderStatus_VO_NotWin, vendorOrderId)
	_, err := agt.dbConn.Exec(esql)
	if err != nil {
		log.Println(err, esql)
		return nil, err
	}
	return &dbproto.Nil{}, nil
}

func (agt *DbUsercenterAgent) BuycaiUpdateVendorOrder(ctx context.Context, arg *dbproto.BuycaiUpdateVendorStatusArg) (*dbproto.Nil, error) {
	vendorOrderId := arg.GetVendorOrderId()
	status := arg.GetStatus()
	vendorReqTime := arg.GetVendorReqTime()
	vendorRespTime := arg.GetVendorRespTime()
	vendorRespId := arg.GetVendorRespId()
	switch status {
	case dbproto.VendorOrderStatus_VO_BetFail:
		return agt.buycaiVendorFail(vendorOrderId, vendorReqTime, vendorRespTime)
	case dbproto.VendorOrderStatus_VO_BetSuccess:
		return agt.buycaiVendorSuccess(vendorOrderId, vendorReqTime, vendorRespTime, vendorRespId)
	case dbproto.VendorOrderStatus_VO_NotWin:
		return agt.buycaiVendorNotWin(vendorOrderId)
	case dbproto.VendorOrderStatus_VO_Win:
		winMoney := arg.GetWinMoney()
		return agt.buycaiVendorWin(vendorOrderId, winMoney)

	}
	return &dbproto.Nil{}, fmt.Errorf("无效状态:%d", status)
}

func (agt *DbUsercenterAgent) BuycaiUpdateUserOrder(ctx context.Context, arg *dbproto.BuycaiUpdateUserStatusArg) (*dbproto.Nil, error) {
	userOrderId := arg.GetUserOrderId()
	status := arg.GetStatus()
	esql := fmt.Sprintf("UPDATE %s SET status=%d WHERE id=%d", TABLE_BUYCAI_USER_ORDER, status, userOrderId)
	_, err := agt.dbConn.Exec(esql)
	if err != nil {
		log.Println(err, esql)
		return nil, err
	}
	return &dbproto.Nil{}, fmt.Errorf("无效状态:%d", status)
}

func (agt *DbUsercenterAgent) BuycaiUserCancelStopChase(ctx context.Context, arg *dbproto.IntValue) (*dbproto.Nil, error) {
	// WARNING: 重点是解冻金额
	tx, err := agt.dbConn.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	userOrderId := arg.GetValue()
	var accountId int64
	var cai, costCai, balance, costBalance float64
	// WARNING: 句尾一定要分号
	esql := fmt.Sprintf("UPDATE %s SET status=%d WHERE id=%d AND issue_num>1 AND chase_no<issue_num AND status=%d RETURNING account_id, cai, cost_cai, balance, cost_balance;",
		TABLE_BUYCAI_USER_ORDER, dbproto.UserOrderStatus_UO_CancelStop, userOrderId, dbproto.UserOrderStatus_UO_Doing)
	log.Println(esql)
	if err := tx.QueryRow(esql).Scan(&accountId, &cai, &costCai, &balance, &costBalance); err != nil {
		log.Println(err, esql)
		return nil, err
	}

	unfreezeCai := cai - costCai
	unfreezeBalance := balance - costBalance
	if unfreezeCai < 0 || unfreezeBalance < 0 {
		log.Println(userOrderId, cai, costCai, balance, costBalance)
		tx.Rollback()
		return nil, fmt.Errorf("金额错误")
	}

	record := &dbproto.FundChangeRecord{
		AccountId:  accountId,
		ChangeType: dbproto.FundChangeType_UNFREEZE,
		VarBalance: unfreezeBalance, VarFreezeBalance: -unfreezeBalance,
		VarCai: unfreezeCai, VarFreezeCai: -unfreezeCai,
		ChangeTime:    time.Now().Unix(),
		UserOrderId:   userOrderId,
		ChangeComment: "用户取消追号",
	}
	if err = changeFund(tx, record); err != nil {
		log.Println(err, userOrderId, accountId)
		tx.Rollback()
		return nil, err
	}
	return &dbproto.Nil{}, tx.Commit()
}

func (agt *DbUsercenterAgent) BuycaiQueryAccountUserOrder(arg *dbproto.BuycaiQueryAccountUserOrderArg, stream dbproto.DbUsercenterAgent_BuycaiQueryAccountUserOrderServer) error {
	accountId := arg.GetAccountId()
	lotteryId := arg.GetLotteryId()
	statusList := arg.GetStatusList()
	limit := arg.GetLimit()
	offset := arg.GetOffset()
	fieldList := []string{
		"id, account_id", "lottery_id", "issue_num", "chase_no",
		"cai", "balance", "sum_money", "issues", "scheme_list",
		"order_time", "ticket_sub_money", "status", "is_win_stop", "cost_cai", "cost_balance",
	}

	var esql string
	if len(statusList) > 0 {
		statusFilterList := []string{}
		for _, v := range statusList {
			statusFilterList = append(statusFilterList, fmt.Sprintf("status=%d", v))
		}
		esql = fmt.Sprintf("SELECT %s FROM %s WHERE account_id=%d AND lottery_id=%d AND (%s) ORDER BY id DESC LIMIT %d OFFSET %d",
			strings.Join(fieldList, ", "), TABLE_BUYCAI_USER_ORDER,
			accountId, lotteryId, strings.Join(statusFilterList, " OR "),
			limit, offset,
		)
	} else {
		esql = fmt.Sprintf("SELECT %s FROM %s WHERE account_id=%d AND lottery_id=%d ORDER BY id DESC LIMIT %d OFFSET %d",
			strings.Join(fieldList, ", "), TABLE_BUYCAI_USER_ORDER,
			accountId, lotteryId,
			limit, offset,
		)
	}

	// log.Println("esql:", esql)
	rows, err := agt.dbConn.Query(esql)
	if err != nil {
		log.Println(err, esql)
		return err
	}

	for rows.Next() {
		order := &dbproto.BuycaiUserOrder{}
		var issues, schemeList []byte
		values := []interface{}{
			&order.Id, &order.AccountId, &order.LotteryId, &order.IssueNum, &order.ChaseNo,
			&order.Cai, &order.Balance, &order.SumMoney, &issues, &schemeList,
			&order.OrderTime, &order.TicketSubMoney, &order.Status, &order.IsWinStop, &order.CostCai, &order.CostBalance,
		}
		if err = rows.Scan(values...); err != nil {
			log.Println(err, esql)
			break
		} else {
			if err := json.Unmarshal(issues, &order.Issues); err != nil {
				log.Println(err, esql)
				return err
			}
			if err := json.Unmarshal(schemeList, &order.SchemeList); err != nil {
				log.Println(err, esql)
				return err
			}
			totalWinMoney, err := agt.queryUserOrderTotalWinMoney(order.Id)
			if err != nil {
				return err
			}
			order.TotalWinMoney = totalWinMoney
			stream.Send(order)
		}
	}

	if err = rows.Err(); err != nil {
		log.Println(err, esql)
	}
	return err
}

func (agt *DbUsercenterAgent) QueryVendorOrderHistory(arg *dbproto.QueryVendorOrderHistoryArg, stream dbproto.DbUsercenterAgent_QueryVendorOrderHistoryServer) error {
	accountId := arg.GetAccountId()
	lotteryId := arg.GetLotteryId()
	startTime := arg.GetStartTime()
	endTime := arg.GetEndTime()
	statusList := arg.GetStatusList()
	limit := arg.GetLimit()
	offset := arg.GetOffset()

	fields := []string{
		"v.id",
		"v.lottery_id",
		"v.issue",
		"v.money",
		"v.status",
		"u.issue_num",
		"v.chase_no",
		"v.win_money",
		"v.add_time",
	}
	var lotteryIdFilter, statusFilter string
	statusFilterList := []string{}
	for _, v := range statusList {
		statusFilterList = append(statusFilterList, fmt.Sprintf("v.status=%d", v))
	}
	if len(statusFilterList) > 0 {
		statusFilter = fmt.Sprintf(" AND (%s)", strings.Join(statusFilterList, " OR "))
	}
	if lotteryId > 0 {
		lotteryIdFilter = fmt.Sprintf(" AND v.lottery_id=%d", lotteryId)
	}
	timeFilter := fmt.Sprintf(" AND v.add_time>=%d AND v.add_time<=%d", startTime, endTime)

	esql := fmt.Sprintf(`SELECT %s FROM (%s v LEFT JOIN %s u ON v.user_order_id=u.id) WHERE v.account_id=%d 
			%s %s %s ORDER BY id DESC LIMIT %d OFFSET %d`,
		strings.Join(fields, ", "), TABLE_BUYCAI_VENDOR_ORDER, TABLE_BUYCAI_USER_ORDER, accountId,
		lotteryIdFilter, statusFilter, timeFilter, limit, offset,
	)

	rows, err := agt.dbConn.Query(esql)
	if err != nil {
		log.Println(err, esql)
		return err
	}

	for rows.Next() {
		record := &dbproto.VendorOrderRecord{}
		err := rows.Scan(&record.Id, &record.LotteryId, &record.Issue, &record.Money, &record.Status,
			&record.IssueNum, &record.ChaseNo, &record.WinMoney, &record.AddTime)
		if err != nil {
			log.Println(err, esql)
			break
		}
		stream.Send(record)
	}

	if err = rows.Err(); err != nil {
		log.Println(err, esql)
	}
	return err
}

func (agt *DbUsercenterAgent) QueryUserOrderHistory(arg *dbproto.QueryUserOrderHistoryArg, stream dbproto.DbUsercenterAgent_QueryUserOrderHistoryServer) error {
	accountId := arg.GetAccountId()
	lotteryId := arg.GetLotteryId()
	startTime := arg.GetStartTime()
	endTime := arg.GetEndTime()
	statusList := arg.GetStatusList()
	limit := arg.GetLimit()
	offset := arg.GetOffset()
	isChase := arg.GetIsChase()

	fields := []string{
		"u.id",
		"u.lottery_id",
		"u.sum_money",
		"u.status",
		"u.issue_num",
		"u.chase_no",
		"SUM(v.win_money) AS win_money",
		"u.order_time",
	}
	var lotteryIdFilter, statusFilter, maxIssueNumFilter string
	statusFilterList := []string{}
	for _, v := range statusList {
		statusFilterList = append(statusFilterList, fmt.Sprintf("u.status=%d", v))
	}
	if len(statusFilterList) > 0 {
		statusFilter = fmt.Sprintf(" AND (%s)", strings.Join(statusFilterList, " OR "))
	}
	if lotteryId > 0 {
		lotteryIdFilter = fmt.Sprintf(" AND u.lottery_id=%d", lotteryId)
	}
	timeFilter := fmt.Sprintf(" AND u.order_time>=%d AND u.order_time<=%d", startTime, endTime)

	if isChase {
		maxIssueNumFilter = " AND u.issue_num>1"
	} else {
		maxIssueNumFilter = " AND u.issue_num=1"
	}

	esql := fmt.Sprintf(`SELECT %s FROM (%s u LEFT JOIN %s v ON u.id=v.user_order_id) WHERE u.account_id=%d
		%s %s %s %s GROUP BY u.id ORDER BY u.id DESC LIMIT %d OFFSET %d 
	`, strings.Join(fields, ", "), TABLE_BUYCAI_USER_ORDER, TABLE_BUYCAI_VENDOR_ORDER, accountId,
		lotteryIdFilter, statusFilter, timeFilter, maxIssueNumFilter, limit, offset,
	)

	rows, err := agt.dbConn.Query(esql)
	if err != nil {
		log.Println(err, esql)
		return err
	}

	for rows.Next() {
		record := &dbproto.UserOrderRecord{}
		err := rows.Scan(&record.Id, &record.LotteryId, &record.Money, &record.Status,
			&record.IssueNum, &record.ChaseNo, &record.WinMoney, &record.AddTime)
		if err != nil {
			log.Println(err, esql)
			break
		}

		subEsql := fmt.Sprintf("SELECT issue FROM %s WHERE user_order_id=%d AND chase_no=%d", TABLE_BUYCAI_VENDOR_ORDER, record.Id, record.ChaseNo)
		if err := agt.dbConn.QueryRow(subEsql).Scan(&record.CurIssue); err != nil {
			log.Println(err, subEsql)
			return err
		}
		stream.Send(record)
	}

	if err = rows.Err(); err != nil {
		log.Println(err, esql)
	}
	return err
}

func (agt *DbUsercenterAgent) QueryBuycaiUserOrderById(ctx context.Context, arg *dbproto.IntValue) (*dbproto.BuycaiUserOrder, error) {
	fieldList := []string{
		"id, account_id", "lottery_id", "issue_num", "chase_no",
		"cai", "balance", "sum_money", "issues", "scheme_list",
		"order_time", "ticket_sub_money", "status", "is_win_stop", "cost_cai", "cost_balance",
	}

	order := &dbproto.BuycaiUserOrder{}
	var issues, schemeList []byte
	values := []interface{}{
		&order.Id, &order.AccountId, &order.LotteryId, &order.IssueNum, &order.ChaseNo,
		&order.Cai, &order.Balance, &order.SumMoney, &issues, &schemeList,
		&order.OrderTime, &order.TicketSubMoney, &order.Status, &order.IsWinStop, &order.CostCai, &order.CostBalance,
	}

	esql := fmt.Sprintf("SELECT %s FROM %s WHERE id=%d", strings.Join(fieldList, ", "), TABLE_BUYCAI_USER_ORDER, arg.GetValue())
	if err := agt.dbConn.QueryRow(esql).Scan(values...); err != nil {
		log.Println(err, esql)
		return nil, err
	}
	if err := json.Unmarshal(issues, &order.Issues); err != nil {
		log.Println(err, esql)
		return nil, err
	}
	if err := json.Unmarshal(schemeList, &order.SchemeList); err != nil {
		log.Println(err, esql)
		return nil, err
	}
	totalWinMoney, err := agt.queryUserOrderTotalWinMoney(order.Id)
	if err != nil {
		return nil, err
	}
	order.TotalWinMoney = totalWinMoney
	return order, nil
}

func (agt *DbUsercenterAgent) QueryBuycaiVendorOrderById(ctx context.Context, arg *dbproto.IntValue) (*dbproto.BuycaiVendorOrder, error) {
	fieldList := []string{
		"id", "user_order_id", "account_id", "lottery_id", "issue", "sum_num", "multiple",
		"money", "cai", "balance", "chase_no", "vendor", "scheme_list", "add_time",
		"status", "status_desc", "win_money",
		"vendor_req_time", "vendor_resp_time", "vendor_resp_id",
	}

	order := &dbproto.BuycaiVendorOrder{}
	var schemeList []byte
	values := []interface{}{
		&order.Id, &order.UserOrderId, &order.AccountId, &order.LotteryId, &order.Issue, &order.SumNum, &order.Multiple,
		&order.Money, &order.Cai, &order.Balance, &order.ChaseNo, &order.Vendor, &schemeList, &order.AddTime,
		&order.Status, &order.StatusDesc, &order.WinMoney,
		&order.VendorReqTime, &order.VendorRespTime, &order.VendorRespId,
	}
	esql := fmt.Sprintf("SELECT %s FROM %s WHERE id=%d", strings.Join(fieldList, ", "), TABLE_BUYCAI_VENDOR_ORDER, arg.GetValue())
	if err := agt.dbConn.QueryRow(esql).Scan(values...); err != nil {
		log.Println(err, esql)
		return nil, err
	}
	if err := json.Unmarshal(schemeList, &order.SchemeList); err != nil {
		log.Println(err, esql)
		return nil, err
	}
	return order, nil
}
func (agt *DbUsercenterAgent) BuycaiQueryUserOrderAllIssues(ctx context.Context, arg *dbproto.IntValue) (*dbproto.UserOrderAllIssues, error) {
	fieldList := []string{
		"issue", "multiple", "money", "win_money", "status", "chase_no", "id",
	}

	result := &dbproto.UserOrderAllIssues{}
	esql := fmt.Sprintf("SELECT %s FROM %s WHERE user_order_id=%d ORDER BY id DESC", strings.Join(fieldList, ", "), TABLE_BUYCAI_VENDOR_ORDER, arg.GetValue())

	rows, err := agt.dbConn.Query(esql)
	if err != nil {
		log.Println(err, esql)
		return nil, err
	}

	for rows.Next() {
		info := &dbproto.BuycaiIssueInfo{}
		values := []interface{}{
			&info.Issue, &info.Multiple, &info.Money, &info.WinMoney, &info.Status, &info.ChaseNo, &info.VendorOrderId,
		}
		if err = rows.Scan(values...); err != nil {
			log.Println(err, esql)
			return nil, err
		}
		result.List = append(result.List, info)
	}
	if err = rows.Err(); err != nil {
		log.Println(err, esql)
	}
	return result, err
}
