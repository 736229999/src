package usercenter

import (
	"database/sql"
	"fmt"
	"log"
	"math"
	"strings"
	"time"

	. "github.com/caojunxyz/mimi-server/apps/dbagent/helper"
	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"golang.org/x/net/context"
)

func changeFund(tx *sql.Tx, record *dbproto.FundChangeRecord) error {
	query := fmt.Sprintf(`UPDATE %s AS tbl SET 
		balance=tbl.balance+$1, freeze_balance=tbl.freeze_balance+$2, 
		cai=tbl.cai+$3, freeze_cai=tbl.freeze_cai+$4,  
		total_recharge=tbl.total_recharge+$5, total_win=tbl.total_win+$6,
		total_withdraw=tbl.total_withdraw+$7, total_cai=tbl.total_cai+$8,
		total_buycai=tbl.total_buycai+$9
		WHERE account_id=$10 RETURNING balance, freeze_balance, cai, freeze_cai`,
		TABLE_FUND)

	accountId := record.AccountId
	varBalance := record.VarBalance
	varFreezeBalance := record.VarFreezeBalance
	varCai := record.VarCai
	varFreezeCai := record.VarFreezeCai
	changeType := record.ChangeType

	var accRecharge, accWin, accWithdraw, accCai, accBuycai float64
	switch changeType {
	case dbproto.FundChangeType_RECHARGE:
		accRecharge = math.Abs(varBalance)
	case dbproto.FundChangeType_WIN:
		accWin = math.Abs(varBalance)
	case dbproto.FundChangeType_WITHDRAW:
		accWithdraw = math.Abs(varBalance)
	case dbproto.FundChangeType_BUYCAI:
		accBuycai = math.Abs(varBalance) + math.Abs(varCai)
	}
	if varCai > 0 {
		accCai = varCai
	}

	stmt0, err := tx.Prepare(query)
	if err != nil {
		log.Println(err, query)
		return err
	}
	defer stmt0.Close()
	var balance, freezeBalance, cai, freezeCai float64
	row := stmt0.QueryRow(varBalance, varFreezeBalance, varCai, varFreezeCai, accRecharge, accWin, accWithdraw, accCai, accBuycai, accountId)
	if err = row.Scan(&balance, &freezeBalance, &cai, &freezeCai); err != nil {
		log.Println(err, query)
		return err
	}

	if balance < 0 || cai < 0 || freezeBalance < 0 || freezeCai < 0 {
		log.Printf("%d --> 资金不足(%f, %f, %f, %f) (%f, %f, %f, %f)\n", accountId, balance, freezeBalance, cai, freezeCai, varBalance, varFreezeBalance, varCai, varFreezeCai)
		return ErrFundNotEnough
	}

	// 资金变动历史记录
	query = fmt.Sprintf(`INSERT INTO  %s
		(account_id, change_type, balance, freeze_balance, cai, freeze_cai, var_balance, var_freeze_balance, var_cai, var_freeze_cai, 
		vendor_order_id, user_order_id, withdraw_apply_id, recharge_order_no, change_time, change_comment)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)`, TABLE_FUND_HISTORY)

	stmt1, err := tx.Prepare(query)
	if err != nil {
		log.Println(err, query)
		return err
	}
	defer stmt1.Close()
	_, err = stmt1.Exec(accountId, changeType, balance, freezeBalance, cai, freezeCai, varBalance, varFreezeBalance, varCai, varFreezeCai,
		record.VendorOrderId, record.UserOrderId, record.WithdrawApplyId, record.RechargeOrderNo, record.ChangeTime, record.ChangeComment)
	if err != nil {
		log.Println(err, query)
		return err
	}
	return nil
}

func (agt *DbUsercenterAgent) QueryFund(ctx context.Context, arg *dbproto.IntValue) (*dbproto.Fund, error) {
	accountId := arg.GetValue()
	esql := fmt.Sprintf("SELECT balance, cai, freeze_balance, freeze_cai, total_recharge, total_win, total_withdraw, total_cai, total_buycai FROM %s WHERE account_id=%d", TABLE_FUND, accountId)
	fund := &dbproto.Fund{}
	if err := agt.dbConn.QueryRow(esql).Scan(&fund.Balance, &fund.Cai, &fund.FreezeBalance, &fund.FreezeCai, &fund.TotalRecharge, &fund.TotalWin, &fund.TotalWithdraw, &fund.TotalCai, &fund.TotalBuycai); err != nil {
		log.Println(err, esql)
		return nil, err
	}
	return fund, nil
}

func (agt *DbUsercenterAgent) QueryFundHistory(ctx context.Context, arg *dbproto.QueryFundHistoryArg) (*dbproto.FundChangeHistory, error) {
	accountId := arg.GetAccountId()
	limit := arg.GetLimit()
	offset := arg.GetOffset()
	changeType := arg.GetChangeType()
	timeStart := arg.GetTimeStart()
	timeEnd := arg.GetTimeEnd()
	filter := fmt.Sprintf("account_id=%d AND change_time>=%d AND change_time<=%d", accountId, timeStart, timeEnd)
	if changeType != dbproto.FundChangeType_UNKNOWN {
		filter += fmt.Sprintf(" AND change_type=%d", changeType)
	}

	cols := []string{
		"id", "account_id", "change_type",
		"cai", "freeze_cai", "balance", "freeze_balance",
		"var_balance", "var_freeze_balance", "var_cai", "var_freeze_cai",
		"vendor_order_id", "user_order_id", "withdraw_apply_id", "recharge_order_no",
		"change_time", "change_comment",
	}
	esql := fmt.Sprintf(`SELECT %s
		FROM %s WHERE %s 
		ORDER BY change_time DESC 
		LIMIT %d OFFSET %d`, strings.Join(cols, ", "), TABLE_FUND_HISTORY, filter, limit, offset)

	rows, err := agt.dbConn.Query(esql)
	if err != nil {
		log.Println(err, esql)
		return nil, err
	}

	result := &dbproto.FundChangeHistory{}
	for rows.Next() {
		record := &dbproto.FundChangeRecord{}
		err := rows.Scan(&record.Id, &record.AccountId, &record.ChangeType,
			&record.Cai, &record.FreezeCai, &record.Balance, &record.FreezeBalance,
			&record.VarBalance, &record.VarFreezeBalance, &record.VarCai, &record.VarFreezeCai,
			&record.VendorOrderId, &record.UserOrderId, &record.WithdrawApplyId, &record.RechargeOrderNo,
			&record.ChangeTime, &record.ChangeComment)
		if err != nil {
			log.Println(err, esql)
			break
		}
		result.List = append(result.List, record)
	}

	if err = rows.Err(); err != nil {
		log.Println(err, esql)
	}
	return result, err
}

func (agt *DbUsercenterAgent) QueryFundHistoryStats(ctx context.Context, arg *dbproto.QueryFundHistoryStatsArg) (*dbproto.FundHistoryStats, error) {
	accountId := arg.GetAccountId()
	timeStart := arg.GetTimeStart()
	timeEnd := arg.GetTimeEnd()
	filter := fmt.Sprintf("account_id=%d AND change_time>=%d AND change_time<=%d", accountId, timeStart, timeEnd)
	dbConn := agt.dbConn

	var val sql.NullFloat64
	stats := &dbproto.FundHistoryStats{}
	esql := fmt.Sprintf("SELECT sum(var_balance) AS val FROM %s WHERE %s AND change_type=%d", TABLE_FUND_HISTORY, filter, dbproto.FundChangeType_RECHARGE)
	if err := dbConn.QueryRow(esql).Scan(&val); err != nil {
		log.Println(err, esql)
		return nil, err
	}
	if val.Valid {
		stats.Recharge = val.Float64
	}

	esql = fmt.Sprintf("SELECT sum(var_balance) AS val FROM %s WHERE %s AND change_type=%d", TABLE_FUND_HISTORY, filter, dbproto.FundChangeType_WITHDRAW)
	if err := dbConn.QueryRow(esql).Scan(&val); err != nil {
		log.Println(err, esql)
		return nil, err
	}
	if val.Valid {
		stats.Withdraw = val.Float64
	}

	esql = fmt.Sprintf("SELECT sum(var_balance) AS val FROM %s WHERE %s AND change_type=%d", TABLE_FUND_HISTORY, filter, dbproto.FundChangeType_WIN)
	if err := dbConn.QueryRow(esql).Scan(&val); err != nil {
		log.Println(err, esql)
		return nil, err
	}
	if val.Valid {
		stats.Win = val.Float64
	}

	esql = fmt.Sprintf("SELECT sum(var_balance) + sum(var_cai) AS val FROM %s WHERE %s AND change_type=%d", TABLE_FUND_HISTORY, filter, dbproto.FundChangeType_BUYCAI)
	if err := dbConn.QueryRow(esql).Scan(&val); err != nil {
		log.Println(err, esql)
		return nil, err
	}
	if val.Valid {
		stats.Buycai = val.Float64
	}
	return stats, nil
}

func (agt *DbUsercenterAgent) UpdateRechargeResult(ctx context.Context, arg *dbproto.RechargeResult) (*dbproto.Nil, error) {
	tx, err := agt.dbConn.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	accountId := arg.GetAccountId()
	money := arg.GetMoney()
	orderNo := arg.GetOrderNo()
	rechargeTime := time.Now().Unix()
	method := arg.GetMethod()

	var id int64
	esql := fmt.Sprintf("SELECT id FROM %s WHERE recharge_order_no='%s'", TABLE_FUND_HISTORY, orderNo)
	if err := tx.QueryRow(esql).Scan(&id); err != nil {
		if err != sql.ErrNoRows {
			log.Println(err, esql)
			return nil, err
		}
	}

	if id > 0 {
		err = fmt.Errorf("充值订单%s重复回调！！！\n", orderNo)
		return nil, err
	}

	record := &dbproto.FundChangeRecord{
		AccountId:       accountId,
		ChangeType:      dbproto.FundChangeType_RECHARGE,
		VarBalance:      money,
		RechargeOrderNo: orderNo,
		ChangeTime:      rechargeTime,
		ChangeComment:   method,
	}
	if err = changeFund(tx, record); err != nil {
		log.Println(err)
		tx.Rollback()
	}
	return &dbproto.Nil{}, tx.Commit()
}
