package usercenter

import (
	"fmt"
	"log"
	"time"

	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"golang.org/x/net/context"
)

func (agt *DbUsercenterAgent) InsertWithdrawApply(ctx context.Context, arg *dbproto.WithdrawApply) (*dbproto.WithdrawApply, error) {
	tx, err := agt.dbConn.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	query := "INSERT INTO withdraw_apply(account_id, realname, create_time, amount, in_bankname, in_no, step, is_success, withdraw_type) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id"
	stmt, err := tx.Prepare(query)
	if err != nil {
		log.Println(err, query)
		return nil, err
	}

	accountId := arg.GetAccountId()
	realname := arg.GetRealname()
	createTime := time.Now().Unix()
	amount := arg.GetAmount()
	inBankname := arg.GetInBankname()
	inNo := arg.GetInNo()
	step := arg.GetStep()
	isSuccess := arg.GetIsSuccess()
	withdrawType := arg.GetWithdrawType()
	var id int64
	if err = stmt.QueryRow(accountId, realname, createTime, amount, inBankname, inNo, step, isSuccess, withdrawType).Scan(&id); err != nil {
		log.Println(err, query)
		tx.Rollback()
		return nil, err
	}

	record := &dbproto.FundChangeRecord{
		AccountId:       accountId,
		ChangeType:      dbproto.FundChangeType_FREEZE,
		WithdrawApplyId: id,
		VarBalance:      -amount, VarFreezeBalance: amount,
		ChangeTime:    createTime,
		ChangeComment: "提现申请冻结",
	}
	if err = changeFund(tx, record); err != nil {
		log.Println(err, query)
		tx.Rollback()
		return nil, err
	}

	progress := arg
	progress.Id = id
	return progress, tx.Commit()
}

func (agt *DbUsercenterAgent) QueryWithdrawApply(ctx context.Context, arg *dbproto.IntValue) (*dbproto.WithdrawApply, error) {
	cols := "id, account_id, realname, create_time, amount, in_bankname, in_no, step, is_success, withdraw_type, auditor, audit_time, audit_comment"
	query := fmt.Sprintf("SELECT %s FROM withdraw_apply WHERE id=$1", cols)
	stmt, err := agt.dbConn.Prepare(query)
	if err != nil {
		log.Println(err, query)
		return nil, err
	}
	defer stmt.Close()
	id := arg.GetValue()
	result := &dbproto.WithdrawApply{}
	if err = stmt.QueryRow(id).Scan(&result.Id, &result.AccountId, &result.Realname, &result.CreateTime,
		&result.Amount, &result.InBankname, &result.InNo, &result.Step, &result.IsSuccess, &result.WithdrawType,
		&result.Auditor, &result.AuditTime, &result.AuditComment); err != nil {
		log.Println(err, query)
		return nil, err
	}
	return result, nil
}
