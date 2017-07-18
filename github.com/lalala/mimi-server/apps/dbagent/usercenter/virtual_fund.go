package usercenter

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"golang.org/x/net/context"
)

func (agt *DbUsercenterAgent) QueryVirtualFund(ctx context.Context, arg *dbproto.IntValue) (*dbproto.VirtualFund, error) {
	accountId := arg.GetValue()
	esql := fmt.Sprintf("SELECT credits, kxd FROM %s WHERE account_id=%d", TABLE_VIRTUAL_FUND, accountId)
	virtualFund := &dbproto.VirtualFund{}
	if err := agt.dbConn.QueryRow(esql).Scan(&virtualFund.Credits, &virtualFund.Kxd); err != nil {
		log.Println(err, esql)
		return nil, err
	}
	return virtualFund, nil
}

func (agt *DbUsercenterAgent) queryVirtualFundHistory(tbl string, arg *dbproto.QueryHistoryArg, stream dbproto.DbUsercenterAgent_QueryCreditsHistoryServer) error {
	accountId := arg.GetAccountId()
	limit := arg.GetLimit()
	offset := arg.GetOffset()
	esql := fmt.Sprintf(`SELECT var, remain, reason, change_time, detail FROM %s WHERE account_id=%d LIMIT %d OFFSET %d`, tbl, accountId, limit, offset)

	rows, err := agt.dbConn.Query(esql)
	if err != nil {
		log.Println(err, esql)
		return err
	}

	for rows.Next() {
		record := &dbproto.VirtualFundChangeRecord{}
		var detail sql.NullString
		if err := rows.Scan(&record.Var, &record.Remain, &record.Reason, &record.ChangeTime, &detail); err != nil {
			log.Println(err, esql)
			break
		}
		if detail.Valid {
			record.Detail = detail.String
		}
		stream.Send(record)
	}

	if err = rows.Err(); err != nil {
		log.Println(err, esql)
	}
	return err
}

func (agt *DbUsercenterAgent) QueryCreditsHistory(arg *dbproto.QueryHistoryArg, stream dbproto.DbUsercenterAgent_QueryCreditsHistoryServer) error {
	return agt.queryVirtualFundHistory(TABLE_CREDITS_HISTORY, arg, stream)
}

func (agt *DbUsercenterAgent) QueryKxdHistory(arg *dbproto.QueryHistoryArg, stream dbproto.DbUsercenterAgent_QueryKxdHistoryServer) error {
	return agt.queryVirtualFundHistory(TABLE_KXD_HISTORY, arg, stream)
}

func (agt *DbUsercenterAgent) changeVirtualFund(tx sql.Tx, name string, arg *dbproto.ChangeVirtualFundArg) (*dbproto.VirtualFund, error) {
	accountId := arg.GetAccountId()
	reason := arg.GetReason()
	changeVar := arg.GetVar()
	detail := arg.GetDetail()

	var esql string
	var tbHistory string
	if name == "credits" {
		esql = fmt.Sprintf("UPDATE %s SET credits=%s.credits + %d WHERE account_id=%d RETURNING credits, kxd", TABLE_VIRTUAL_FUND, TABLE_VIRTUAL_FUND, changeVar, accountId)
		tbHistory = TABLE_CREDITS_HISTORY
	} else if name == "kxd" {
		esql = fmt.Sprintf("UPDATE %s SET kxd=%s.kxd + %d WHERE account_id=%d RETURNING credits, kxd", TABLE_VIRTUAL_FUND, TABLE_VIRTUAL_FUND, changeVar, accountId)
		tbHistory = TABLE_KXD_HISTORY
	} else {
		log.Panicf("无效名称: %s", name)
	}

	var credits, kxd int32
	if err := tx.QueryRow(esql).Scan(&credits, &kxd); err != nil {
		log.Println(err, esql)
		//tx.Rollback()
		return nil, err
	}

	remain := int32(0)
	if name == "credits" {
		remain = credits
	} else if name == "kxd" {
		remain = kxd
	}

	now := time.Now().Unix()
	esql = fmt.Sprintf("INSERT INTO %s (account_id, var, remain, reason, change_time, detail) VALUES($1, $2, $3, $4, $5, $6)", tbHistory)
	_, err := tx.Exec(esql, accountId, changeVar, remain, reason, now, detail)
	if err != nil {
		log.Println(err, esql)
		//tx.Rollback()
		return nil, err
	}
	return &dbproto.VirtualFund{Credits: credits, Kxd: kxd}, nil
}

func (agt *DbUsercenterAgent) ChangeCredits(ctx context.Context, arg *dbproto.ChangeVirtualFundArg) (*dbproto.VirtualFund, error) {
	tx, err := agt.dbConn.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	VirtualFund, err := agt.changeVirtualFund(*tx, "credits", arg)
	if err != nil {
		log.Printf("%+v\n", err)
		tx.Rollback()
		return nil, err
	}

	return VirtualFund, tx.Commit()
}

func (agt *DbUsercenterAgent) ChangeKxd(ctx context.Context, arg *dbproto.ChangeVirtualFundArg) (*dbproto.VirtualFund, error) {
	tx, err := agt.dbConn.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	VirtualFund, err := agt.changeVirtualFund(*tx, "credits", arg)
	if err != nil {
		log.Printf("%+v\n", err)
		tx.Rollback()
		return nil, err
	}
	return VirtualFund, tx.Commit()
}

func (agt *DbUsercenterAgent) QueryAccountCreditsTaskInfo(ctx context.Context, arg *dbproto.IntValue) (*dbproto.CreditsTaskInfo, error) {
	accountId := arg.GetValue()
	esql := fmt.Sprintf(`SELECT sum(var) as sum_var, reason FROM %s WHERE account_id=%d GROUP BY reason`, TABLE_CREDITS_HISTORY, accountId)
	rows, err := agt.dbConn.Query(esql)
	if err != nil {
		log.Println(err, esql)
		return nil, err
	}

	result := &dbproto.CreditsTaskInfo{}
	for rows.Next() {
		task := &dbproto.CreditsTask{}
		if err := rows.Scan(&task.SumVar, &task.Reason); err != nil {
			log.Println(err, esql)
			break
		}
		result.TaskList = append(result.TaskList, task)
	}

	if err = rows.Err(); err != nil {
		log.Println(err, esql)
		return nil, err
	}
	vf, err := agt.QueryVirtualFund(ctx, arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	result.Credits = vf.GetCredits()
	return result, nil
}
