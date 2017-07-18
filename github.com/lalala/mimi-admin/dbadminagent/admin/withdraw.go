package admin

import (
	// "database/sql"
	"database/sql"
	"fmt"
	"github.com/caojunxyz/mimi-admin/dbadminagent/proto"
	"golang.org/x/net/context"
	"log"
	"strings"
	"time"
)

const (
	// TABLE_WITHDRAW_APPLY 提现申请表名
	TABLE_WITHDRAW_APPLY = "withdraw_apply"
	// TABLE_WITHDRAW_TRANSFER 提现转账表
	TABLE_WITHDRAW_TRANSFER = "withdraw_transfer"
	// TABLE_WITHDRAW_AUDIT_AUTH 提现申请审核权限表名
	TABLE_WITHDRAW_AUDIT_AUTH = "withdraw_audit_auth"
)

//获取用户的提现记录.
func (agt *DbAdminAgent) QueryUsercenterWithdrawById(ctx context.Context, arg *dbproto.IntValue) (*dbproto.WithdrawApplyList, error) {

	// 	defer func() {
	// 		if err := recover(); err != nil {
	// 			log.Println("recover from panic:", err)
	// 		}
	// 	}()

	// 	fieldList := []string{
	// 		"id", "account_id", "realname", "idcard_no", "phone",
	// 		"create_time", "amount", "in_no", "status", "auditor",
	// 		"audit_time", "audit_comment", "withdraw_type", "out_no", "out_sn", "trans_time",
	// 	}
	// 	esql := fmt.Sprintf(`SELECT %s FROM withdraw WHERE account_id = %d ORDER BY id DESC`, strings.Join(fieldList, ", "), arg.GetValue())
	// 	log.Println(esql)
	// 	st, err := agt.ucDbConn.Prepare(esql)
	// 	if err != nil {
	// 		log.Printf("%+v\n", err)
	// 		return nil, err
	// 	}

	// 	rows, err := st.Query()
	// 	if err != nil {
	// 		log.Printf("%+v\n", err)
	// 		return nil, err
	// 	}

	// 	list := &dbproto.WithdrawList{}
	// 	for rows.Next() {

	// 		var auditor, audit_comment, out_no, out_sn sql.NullString
	// 		var audit_time, trans_time sql.NullInt64
	// 		withdraw := &dbproto.Withdraw{}
	// 		value := []interface{}{
	// 			&withdraw.Id, &withdraw.AccountId, &withdraw.Realname, &withdraw.IdcardNo, &withdraw.Phone,
	// 			&withdraw.CreateTime, &withdraw.Amount, &withdraw.InNo, &withdraw.Status, &auditor,
	// 			&audit_time, &audit_comment, &withdraw.WithdrawType, &out_no, &out_sn, &trans_time,
	// 		}
	// 		if err = rows.Scan(value...); err != nil {
	// 			if err == sql.ErrNoRows {
	// 				return &dbproto.WithdrawList{}, nil
	// 			}
	// 			log.Printf("%+v\n", err)
	// 			return nil, err
	// 		}

	// 		if auditor.Valid {
	// 			withdraw.Auditor = auditor.String
	// 		}

	// 		if audit_comment.Valid {
	// 			withdraw.AuditComment = audit_comment.String
	// 		}

	// if out_no.Valid {
	// 	withdraw.OutNo = out_no.String
	// }

	// if out_sn.Valid {
	// 	withdraw.OutSn = out_sn.String
	// }

	// if audit_time.Valid {
	// 	withdraw.AuditTime = int32(audit_time.Int64)
	// }

	// if trans_time.Valid {
	// 	withdraw.TransTime = int32(trans_time.Int64)
	// }

	// 		list.List = append(list.List, withdraw)
	// 	}

	return nil, nil
}

// QueryWithdrawApplyList 查询提现申请列表
func (agt *DbAdminAgent) QueryWithdrawApplyList(ctx context.Context, arg *dbproto.QueryWithdrawApplyArg) (*dbproto.WithdrawApplyList, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	var (
		start, total              int64
		pager, order, esql, where string
	)

	if arg.GetPage() > 1 {
		start = int64((arg.GetPage() - 1) * arg.GetPageSize())
	} else {
		start = 0
	}
	pager = fmt.Sprintf(" LIMIT %d OFFSET %d", arg.PageSize, start)
	order = " ORDER BY create_time DESC"
	esql = fmt.Sprintf("SELECT id, account_id, phone, realname, idcard_no, in_no, withdraw_type, amount, step, is_success, auditor, audit_time, audit_comment, create_time FROM %s", TABLE_WITHDRAW_APPLY)
	log.Println(start, pager, order)

	if arg.GetStep() != dbproto.WithdrawApply_STEP_ALL {
		where = fmt.Sprintf(" WHERE step = %d AND is_success = %v", arg.GetStep(), arg.GetIsSuccess())
	}

	if arg.GetRealname() != "" {
		condition := fmt.Sprintf(" realname LIKE '%%%s%%'", arg.GetRealname())
		if strings.Contains(where, "WHERE") {
			where = fmt.Sprintf("%s AND %s", where, condition)
		} else {
			where = fmt.Sprintf(" WHERE %s", condition)
		}
	}
	if arg.GetPhone() != "" {
		condition := fmt.Sprintf(" phone LIKE '%%%s%%'", arg.GetPhone())
		if strings.Contains(where, "WHERE") {
			where = fmt.Sprintf("%s AND %s", where, condition)
		} else {
			where = fmt.Sprintf(" WHERE %s", condition)
		}
	}
	if arg.GetInNo() != "" {
		condition := fmt.Sprintf(" in_no LIKE '%%%s%%'", arg.GetInNo())
		if strings.Contains(where, "WHERE") {
			where = fmt.Sprintf("%s AND %s", where, condition)
		} else {
			where = fmt.Sprintf(" WHERE %s", condition)
		}
	}
	if arg.GetAuditor() != "" {
		condition := fmt.Sprintf(" auditor LIKE '%%%s%%'", arg.GetAuditor())
		if strings.Contains(where, "WHERE") {
			where = fmt.Sprintf("%s AND %s", where, condition)
		} else {
			where = fmt.Sprintf(" WHERE %s", condition)
		}
	}

	if arg.GetMinAmount() != 0 {
		condition := fmt.Sprintf(" amount >= %d", arg.GetMinAmount())
		if strings.Contains(where, "WHERE") {
			where = fmt.Sprintf("%s AND %s", where, condition)
		} else {
			where = fmt.Sprintf(" WHERE %s", condition)
		}
	}

	if arg.GetMaxAmount() != 0 {
		condition := fmt.Sprintf(" amount < %d", arg.GetMaxAmount())
		if strings.Contains(where, "WHERE") {
			where = fmt.Sprintf("%s AND %s", where, condition)
		} else {
			where = fmt.Sprintf(" WHERE %s", condition)
		}
	}

	if arg.GetStart() != 0 {
		condition := fmt.Sprintf(" create_time >= %d", arg.GetStart())
		if strings.Contains(where, "WHERE") {
			where = fmt.Sprintf("%s AND %s", where, condition)
		} else {
			where = fmt.Sprintf(" WHERE %s", condition)
		}
	}
	if arg.GetEnd() != 0 {
		condition := fmt.Sprintf(" create_time <= %d", arg.GetEnd())
		if strings.Contains(where, "WHERE") {
			where = fmt.Sprintf("%s AND %s", where, condition)
		} else {
			where = fmt.Sprintf(" WHERE %s", condition)
		}
	}

	rows, err := agt.ucDbConn.Query(esql + where + order + pager)
	if err != nil {
		log.Println(err, esql)
		return nil, err
	}

	withdrawList := make([]*dbproto.WithdrawApply, 0)

	for rows.Next() {
		withdraw := &dbproto.WithdrawApply{}
		err := rows.Scan(&withdraw.Id,
			&withdraw.AccountId,
			&withdraw.Phone,
			&withdraw.Realname,
			&withdraw.IdcardNo,
			&withdraw.InNo,
			&withdraw.WithdrawType,
			&withdraw.Amount,
			&withdraw.Step,
			&withdraw.IsSuccess,
			&withdraw.Auditor,
			&withdraw.AuditTime,
			&withdraw.AuditComment,
			&withdraw.CreateTime,
		)
		if err != nil {
			log.Println(err, esql)
			return nil, err
		}
		withdrawList = append(withdrawList, withdraw)
	}

	totalEsql := fmt.Sprintf("SELECT COUNT(*) FROM %s", TABLE_WITHDRAW_APPLY)

	row := agt.ucDbConn.QueryRow(totalEsql + where)
	row.Scan(&total)
	// log.Println("SQL: ", esql+where+order+pager)
	// log.Println("totalSQL: ", totalEsql+where)
	// log.Printf("total is %v", total)
	return &dbproto.WithdrawApplyList{
		List:  withdrawList,
		Total: total,
	}, nil
}

// QueryWithdrawApplyById 获取某个提现申请
func (agt *DbAdminAgent) QueryWithdrawApplyById(ctx context.Context, arg *dbproto.WithdrawApplyId) (*dbproto.WithdrawApply, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	var (
		esql string
	)
	esql = fmt.Sprintf("SELECT id, account_id, amount, create_time, phone, in_no, idcard_no, step, is_success, realname, auditor, audit_time, audit_comment, withdraw_type FROM %s WHERE id = $1", TABLE_WITHDRAW_APPLY)

	log.Println("esql", esql, "id ", arg.GetId())
	st, err := agt.ucDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	withdrawApply := &dbproto.WithdrawApply{}
	err = st.QueryRow(arg.GetId()).Scan(
		&withdrawApply.Id,
		&withdrawApply.AccountId,
		&withdrawApply.Amount,
		&withdrawApply.CreateTime,
		&withdrawApply.Phone,
		&withdrawApply.InNo,
		&withdrawApply.IdcardNo,
		&withdrawApply.Step,
		&withdrawApply.IsSuccess,
		&withdrawApply.Realname,
		&withdrawApply.Auditor,
		&withdrawApply.AuditTime,
		&withdrawApply.AuditComment,
		&withdrawApply.WithdrawType)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("提现申请不存在:%+v, %v\n", err, arg.GetId())
			return nil, err
		}
		log.Printf("%+v\n", err)
		return nil, err
	}
	log.Println("withdrawApply", withdrawApply)
	return withdrawApply, nil
}

// UpdateWithdrawApplyStatus 提交审核结果
func (agt *DbAdminAgent) UpdateWithdrawApplyStatus(ctx context.Context, arg *dbproto.UpdateWAStatusArg) (*dbproto.IntValue, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	var (
		affect int64
	)

	esql := fmt.Sprintf("UPDATE %s "+
		"SET step = $1, is_success = $2, audit_comment = $3, audit_time = $4 WHERE id = $5", TABLE_WITHDRAW_APPLY)

	if arg.GetStep() == dbproto.WithdrawApply_STEP_AUDIT && arg.GetIsSuccess() == true { // 审核通过逻辑
		tx, err := agt.ucDbConn.Begin()
		if err != nil {
			tx.Rollback()
			log.Printf("%+v\n", err)
			return nil, err
		}
		st, err := tx.Prepare(esql)
		if err != nil {
			tx.Rollback()
			log.Println(err)
			return nil, err
		}
		res, err := st.Exec(arg.GetStep(), arg.GetIsSuccess(), arg.GetAuditComment(), time.Now().Unix(), arg.GetId())
		if err != nil {
			tx.Rollback()
			log.Printf("%+v\n", err)
			// tx.Rollback()
			return nil, err
		}
		affect, err = res.RowsAffected()
		if err != nil {
			tx.Rollback()
			log.Printf("RowsAffected %v\n", err)
			// tx.Rollback()
			return nil, err
		}
		var rid int64
		isql := fmt.Sprintf("INSERT INTO %s(apply_id, transfer_amount, certificate_url, step, is_success, operator, operate_time, pay_no, pay_sn) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id", TABLE_WITHDRAW_TRANSFER)
		err = tx.QueryRow(isql, arg.GetId(), 0, "", 1, true, "", 0, "", "").Scan(&rid)
		if err != nil {
			tx.Rollback()
			log.Printf("error %v, arg: %+v\n", err, arg)
			log.Printf("esql is %v", esql)
			return nil, err
		}
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
	} else {

		st, err := agt.ucDbConn.Prepare(esql)
		// st, err := tx.Prepare(esql)
		if err != nil {
			log.Printf("%+v\n", err)
			// tx.Rollback()
			return nil, err
		}
		log.Println("SQL", esql)
		res, err := st.Exec(arg.GetStep(), arg.GetIsSuccess(), arg.GetAuditComment(), time.Now().Unix(), arg.GetId())
		if err != nil {
			log.Printf("%+v\n", err)
			// tx.Rollback()
			return nil, err
		}
		affect, err = res.RowsAffected()
		if err != nil {
			log.Printf("RowsAffected %v\n", err)
			// tx.Rollback()
			return nil, err
		}
	}
	return &dbproto.IntValue{
		Value: affect,
	}, nil
}

// CheckWithdrawApply 检查提现申请
func (agt *DbAdminAgent) CheckWithdrawApply(ctx context.Context, arg *dbproto.CheckWAArg) (*dbproto.CheckWARes, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	var (
		auditor                        string
		step                           int32
		existAuditor, isOwn, isSuccess bool
	)
	esql := fmt.Sprintf("SELECT auditor, step, is_success FROM %s WHERE id = $1 ", TABLE_WITHDRAW_APPLY)
	st, err := agt.ucDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	err = st.QueryRow(arg.GetWithdrawApplyId()).Scan(&auditor, &step, &isSuccess)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("不存在:%+v, %v\n", err, arg.GetWithdrawApplyId())
			return nil, err
		}
		log.Printf("%+v\n", err)
		return nil, err
	}

	if auditor == arg.GetAuditor() {
		existAuditor = true
		isOwn = true
	} else if auditor != "" {
		existAuditor = true
	}
	return &dbproto.CheckWARes{
		ExistAuditor: existAuditor,
		IsOwn:        isOwn,
		Step:         dbproto.WithdrawApply_WAStepBackend(step),
		IsSuccess:    isSuccess,
	}, nil
}

// ClaimWithdrawApply 认领提现申请
func (agt *DbAdminAgent) ClaimWithdrawApply(ctx context.Context, arg *dbproto.ClaimWAArg) (*dbproto.WithdrawApply, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	tx, err := agt.ucDbConn.Begin()

	esql := fmt.Sprintf("UPDATE %s SET auditor = $1 WHERE id = $2", TABLE_WITHDRAW_APPLY)
	st, err := tx.Prepare(esql)
	if err != nil {
		log.Printf("%v\n", err)
		return nil, err
	}
	log.Println("SQL", esql)
	res, err := st.Exec(arg.GetAuditor(), arg.GetWithdrawApplyId())
	if err != nil {
		log.Printf("%+v\n", err)
		tx.Rollback()
		return nil, err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		log.Printf("RowsAffected %v\n", err)
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		log.Println("tx.Commit error", err)
		return nil, err
	}

	log.Println("affect", affect)
	return agt.QueryWithdrawApplyById(ctx, &dbproto.WithdrawApplyId{Id: arg.GetWithdrawApplyId()})
}

// CreateWithdrawAuditAuth 创建一条提现审核信息
func (agt *DbAdminAgent) CreateWithdrawAuditAuth(ctx context.Context, arg *dbproto.WithdrawAuditAuth) (*dbproto.IntValue, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	var rid int64
	esql := fmt.Sprintf("INSERT INTO %s(user_id, min_amount, max_amount, create_time, creator, valid, type) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id", TABLE_WITHDRAW_AUDIT_AUTH)
	err := agt.dbConn.QueryRow(esql, arg.GetUserId(), arg.GetMinAmount(), arg.GetMaxAmount(), time.Now().Unix(), arg.GetCreator(), arg.GetValid(), arg.GetType()).Scan(&rid)
	if err != nil {
		log.Printf("error %v, arg: %+v\n", err, arg)
		log.Printf("esql is %v", esql)
		return nil, err
	}
	return &dbproto.IntValue{Value: rid}, nil
}

// QueryWithdrawAuditAuthList 获取提现审核列表
func (agt *DbAdminAgent) QueryWithdrawAuditAuthList(ctx context.Context, arg *dbproto.QueryWithdrawAuditAuthArg) (*dbproto.WithdrawAuditAuthList, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	var (
		start                     int32
		total                     int64
		esql, where, order, pager string
	)

	if arg.GetPage() > 1 {
		start = (arg.GetPage() - 1) * arg.PageSize
	} else {
		start = 0
	}
	pager = fmt.Sprintf(" LIMIT %d OFFSET %d", arg.PageSize, start)
	// where = " WHERE is_visible = true"
	order = " ORDER BY valid DESC, create_time DESC"
	esql = fmt.Sprintf("SELECT waa.id, waa.user_id, waa.min_amount, waa.max_amount, waa.create_time, waa.creator, waa.type, waa.valid, waa.unlimited, "+
		"u.username, u.email, u.mobile, cu.username, cu.email, cu.mobile FROM %s AS waa "+
		"LEFT JOIN \"user\" AS u ON waa.user_id = u.id "+
		"LEFT JOIN \"user\" AS cu ON waa.creator = cu.id"+
		"", TABLE_WITHDRAW_AUDIT_AUTH)
	if arg.GetUserId() != 0 {
		where += fmt.Sprintf(" WHERE waa.user_id = %d", arg.GetUserId())
	}

	if arg.GetType() != dbproto.WithdrawAuditAuth_TYPE_ALL {
		condition := fmt.Sprintf(" waa.type = %d", arg.GetType())
		if strings.Contains(where, "WHERE") {
			where = fmt.Sprintf("%s AND %s", where, condition)
		} else {
			where = fmt.Sprintf(" WHERE %s", condition)
		}
	}

	rows, err := agt.dbConn.Query(esql + where + order + pager)
	log.Println("SQL", esql+where+order+pager)
	if err != nil {
		log.Println(err, "SQL: ", esql+where+order)
		return nil, err
	}

	waaList := make([]*dbproto.WithdrawAuditAuth, 0)

	for rows.Next() {
		waa := &dbproto.WithdrawAuditAuth{
			UserInfo:        new(dbproto.AdminUserInfoArg),
			CreatorUserInfo: new(dbproto.AdminUserInfoArg),
		}
		err := rows.Scan(
			&waa.Id,
			&waa.UserId,
			&waa.MinAmount,
			&waa.MaxAmount,
			&waa.CreateTime,
			&waa.Creator,
			&waa.Type,
			&waa.Valid,
			&waa.Unlimited,
			&waa.UserInfo.Username,
			&waa.UserInfo.Email,
			&waa.UserInfo.Mobile,
			&waa.CreatorUserInfo.Username,
			&waa.CreatorUserInfo.Email,
			&waa.CreatorUserInfo.Mobile,
		)
		if err != nil {
			log.Println(err, esql)
			return nil, err
		}
		waaList = append(waaList, waa)
	}

	totalEsql := fmt.Sprintf("SELECT COUNT(*) FROM %s", TABLE_WITHDRAW_AUDIT_AUTH)

	row := agt.dbConn.QueryRow(totalEsql + where)
	row.Scan(&total)
	log.Println("SQL: ", esql+where+order+pager)
	log.Println("totalSQL: ", totalEsql+where)
	log.Printf("total is %v", total)

	return &dbproto.WithdrawAuditAuthList{
		List:  waaList,
		Total: total,
	}, nil
}

// QueryWithdrawAuditAuthById 获取提现申请审核信息
func (agt *DbAdminAgent) QueryWithdrawAuditAuthById(ctx context.Context, arg *dbproto.IntValue) (*dbproto.WithdrawAuditAuth, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	var (
		esql string
	)
	esql = fmt.Sprintf("SELECT waa.id, waa.user_id, waa.min_amount, waa.max_amount, waa.create_time, waa.creator, waa.type, waa.valid, "+
		"u.username, u.email, u.mobile, cu.username, cu.email, cu.mobile FROM %s AS waa "+
		"LEFT JOIN \"user\" AS u ON waa.user_id = u.id "+
		"LEFT JOIN \"user\" AS cu ON waa.creator = cu.id "+
		"WHERE waa.id = $1", TABLE_WITHDRAW_AUDIT_AUTH)

	log.Println("esql", esql, "id ", arg.GetValue())
	st, err := agt.dbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	waa := &dbproto.WithdrawAuditAuth{
		UserInfo:        new(dbproto.AdminUserInfoArg),
		CreatorUserInfo: new(dbproto.AdminUserInfoArg),
	}
	err = st.QueryRow(arg.GetValue()).Scan(&waa.Id,
		&waa.UserId,
		&waa.MinAmount,
		&waa.MaxAmount,
		&waa.CreateTime,
		&waa.Creator,
		&waa.Type,
		&waa.Valid,
		&waa.UserInfo.Username,
		&waa.UserInfo.Email,
		&waa.UserInfo.Mobile,
		&waa.CreatorUserInfo.Username,
		&waa.CreatorUserInfo.Email,
		&waa.CreatorUserInfo.Mobile,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("权限不存在:%+v, %v\n", err, arg.GetValue())
			return nil, err
		}
		log.Printf("%+v\n", err)
		return nil, err
	}
	log.Println("waa", waa)
	return waa, nil
}

// UpdateWithdrawAuditAuth 修改提现申请权限
func (agt *DbAdminAgent) UpdateWithdrawAuditAuth(ctx context.Context, arg *dbproto.WithdrawAuditAuth) (*dbproto.IntValue, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	esql := fmt.Sprintf("UPDATE %s "+
		"SET min_amount = $1, max_amount = $2, valid = $3 WHERE id = $4", TABLE_WITHDRAW_AUDIT_AUTH)
	st, err := agt.dbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	log.Println("SQL", esql)
	res, err := st.Exec(arg.GetMinAmount(), arg.GetMaxAmount(), arg.GetValid(), arg.GetId())
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		log.Printf("RowsAffected %v\n", err)
		return nil, err
	}
	return &dbproto.IntValue{
		Value: affect,
	}, nil
}

// CreateWithdrawTransfer 创建提现转账
func (agt *DbAdminAgent) CreateWithdrawTransfer(ctx context.Context, arg *dbproto.WithdrawTransfer) (*dbproto.IntValue, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	var rid int64
	esql := fmt.Sprintf("INSERT INTO %s(apply_id, transfer_amount, certificate_url, step, is_success, operator, operate_time, pay_no, pay_sn) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id", TABLE_WITHDRAW_TRANSFER)
	err := agt.dbConn.QueryRow(esql, arg.GetAppyId(), arg.GetTransferAmount(), arg.GetCertificateUrl(), arg.GetStep(), arg.GetIsSuccess(), arg.GetOperator(), arg.GetOperateTime(), arg.GetPayNo()).Scan(&rid)
	if err != nil {
		log.Printf("error %v, arg: %+v\n", err, arg)
		log.Printf("esql is %v", esql)
		return nil, err
	}
	return &dbproto.IntValue{Value: rid}, nil
}

// QueryWithdrawTransferList 查询提现转账列表
func (agt *DbAdminAgent) QueryWithdrawTransferList(ctx context.Context, arg *dbproto.QueryWithdrawTransferArg) (*dbproto.WithdrawTransferList, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	var (
		start, total              int64
		pager, order, esql, where string
	)

	if arg.GetPage() > 1 {
		start = int64((arg.GetPage() - 1) * arg.GetPageSize())
	} else {
		start = 0
	}
	pager = fmt.Sprintf(" LIMIT %d OFFSET %d", arg.PageSize, start)
	order = " ORDER BY create_time DESC"
	esql = fmt.Sprintf("SELECT wt.id, wt.apply_id, wt.step, wt.is_success, wt.transfer_amount, wt.certificate_url, wt.operator, wt.operate_time, wt.target_account, "+
		"wa.id as wa_id, wa.account_id, wa.realname, wa.create_time, wa.amount, wa.in_bankname, wa.in_no, wa.phone, wa.idcard_no, wa.step AS wa_step, wa.is_success, wa.auditor, wa.audit_time, wa.audit_comment, wa.withdraw_type "+
		"FROM %s AS wt "+
		"LEFT JOIN %s AS wa ON wt.apply_id = wa.id", TABLE_WITHDRAW_TRANSFER, TABLE_WITHDRAW_APPLY)
	log.Println(start, pager, order)

	if arg.GetStep() != dbproto.WithdrawTransfer_STEP_ALL {
		where = fmt.Sprintf(" WHERE wt.step = %d AND wt.is_success = %v", arg.GetStep(), arg.GetIsSuccess())
	}

	if arg.GetRealname() != "" {
		condition := fmt.Sprintf(" wa.realname LIKE '%%%s%%'", arg.GetRealname())
		if strings.Contains(where, "WHERE") {
			where = fmt.Sprintf("%s AND %s", where, condition)
		} else {
			where = fmt.Sprintf(" WHERE %s", condition)
		}
	}
	if arg.GetPhone() != "" {
		condition := fmt.Sprintf(" wa.phone LIKE '%%%s%%'", arg.GetPhone())
		if strings.Contains(where, "WHERE") {
			where = fmt.Sprintf("%s AND %s", where, condition)
		} else {
			where = fmt.Sprintf(" WHERE %s", condition)
		}
	}
	if arg.GetInNo() != "" {
		condition := fmt.Sprintf(" wa.in_no LIKE '%%%s%%'", arg.GetInNo())
		if strings.Contains(where, "WHERE") {
			where = fmt.Sprintf("%s AND %s", where, condition)
		} else {
			where = fmt.Sprintf(" WHERE %s", condition)
		}
	}
	if arg.GetOperator() != "" {
		condition := fmt.Sprintf(" wt.operator LIKE '%%%s%%'", arg.GetOperator())
		if strings.Contains(where, "WHERE") {
			where = fmt.Sprintf("%s AND %s", where, condition)
		} else {
			where = fmt.Sprintf(" WHERE %s", condition)
		}
	}

	if arg.GetMinAmount() != 0 {
		condition := fmt.Sprintf(" wa.amount >= %d", arg.GetMinAmount())
		if strings.Contains(where, "WHERE") {
			where = fmt.Sprintf("%s AND %s", where, condition)
		} else {
			where = fmt.Sprintf(" WHERE %s", condition)
		}
	}

	if arg.GetMaxAmount() != 0 {
		condition := fmt.Sprintf(" wa.amount < %d", arg.GetMaxAmount())
		if strings.Contains(where, "WHERE") {
			where = fmt.Sprintf("%s AND %s", where, condition)
		} else {
			where = fmt.Sprintf(" WHERE %s", condition)
		}
	}

	if arg.GetStart() != 0 {
		condition := fmt.Sprintf(" wt.operate_time >= %d", arg.GetStart())
		if strings.Contains(where, "WHERE") {
			where = fmt.Sprintf("%s AND %s", where, condition)
		} else {
			where = fmt.Sprintf(" WHERE %s", condition)
		}
	}
	if arg.GetEnd() != 0 {
		condition := fmt.Sprintf(" wt.operate_time <= %d", arg.GetEnd())
		if strings.Contains(where, "WHERE") {
			where = fmt.Sprintf("%s AND %s", where, condition)
		} else {
			where = fmt.Sprintf(" WHERE %s", condition)
		}
	}

	rows, err := agt.ucDbConn.Query(esql + where + order + pager)
	if err != nil {
		log.Println(err, esql)
		return nil, err
	}

	wtList := make([]*dbproto.WithdrawTransfer, 0)

	for rows.Next() {
		wt := &dbproto.WithdrawTransfer{WithdrawApply: new(dbproto.WithdrawApply)}
		err := rows.Scan(
			&wt.Id,
			&wt.AppyId,
			&wt.Step,
			&wt.IsSuccess,
			&wt.TransferAmount,
			&wt.CertificateUrl,
			&wt.Operator,
			&wt.OperateTime,
			&wt.TargetAccount,
			&wt.WithdrawApply.Id,
			&wt.WithdrawApply.AccountId,
			&wt.WithdrawApply.Realname,
			&wt.WithdrawApply.CreateTime,
			&wt.WithdrawApply.Amount,
			&wt.WithdrawApply.InBankname,
			&wt.WithdrawApply.InNo,
			&wt.WithdrawApply.Phone,
			&wt.WithdrawApply.IdcardNo,
			&wt.WithdrawApply.Step,
			&wt.WithdrawApply.IsSuccess,
			&wt.WithdrawApply.Auditor,
			&wt.WithdrawApply.AuditTime,
			&wt.WithdrawApply.AuditComment,
			&wt.WithdrawApply.WithdrawType,
		)
		if err != nil {
			log.Println(err, esql)
			return nil, err
		}
		wtList = append(wtList, wt)
	}

	totalEsql := fmt.Sprintf("SELECT COUNT(*) FROM %s", TABLE_WITHDRAW_TRANSFER)

	row := agt.ucDbConn.QueryRow(totalEsql + where)
	row.Scan(&total)
	// log.Println("SQL: ", esql+where+order+pager)
	// log.Println("totalSQL: ", totalEsql+where)
	// log.Printf("total is %v", total)
	return &dbproto.WithdrawTransferList{
		List:  wtList,
		Total: total,
	}, nil
}

// QueryWithdrawTransferById 查询提现转账
func (agt *DbAdminAgent) QueryWithdrawTransferById(ctx context.Context, arg *dbproto.IntValue) (*dbproto.WithdrawTransfer, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	var (
		esql string
	)
	esql = fmt.Sprintf("SELECT wt.id, wt.apply_id, wt.step, wt.is_success, wt.transfer_amount, wt.certificate_url, wt.operator, wt.operate_time, wt.target_account, "+
		"wa.id as wa_id, wa.account_id, wa.realname, wa.create_time, wa.amount, wa.in_bankname, wa.in_no, wa.phone, wa.idcard_no, wa.step AS wa_step, wa.is_success, wa.auditor, wa.audit_time, wa.audit_comment, wa.withdraw_type "+
		"FROM %s AS wt "+
		"LEFT JOIN %s AS wa ON wt.apply_id = wa.id", TABLE_WITHDRAW_TRANSFER, TABLE_WITHDRAW_APPLY)

	log.Println("esql", esql, "id ", arg.GetValue())
	st, err := agt.dbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	wt := &dbproto.WithdrawTransfer{
		WithdrawApply: new(dbproto.WithdrawApply),
	}
	err = st.QueryRow(arg.GetValue()).Scan(
		&wt.Id,
		&wt.AppyId,
		&wt.Step,
		&wt.IsSuccess,
		&wt.TransferAmount,
		&wt.CertificateUrl,
		&wt.Operator,
		&wt.OperateTime,
		&wt.TargetAccount,
		&wt.WithdrawApply.Id,
		&wt.WithdrawApply.AccountId,
		&wt.WithdrawApply.Realname,
		&wt.WithdrawApply.CreateTime,
		&wt.WithdrawApply.Amount,
		&wt.WithdrawApply.InBankname,
		&wt.WithdrawApply.InNo,
		&wt.WithdrawApply.Phone,
		&wt.WithdrawApply.IdcardNo,
		&wt.WithdrawApply.Step,
		&wt.WithdrawApply.IsSuccess,
		&wt.WithdrawApply.Auditor,
		&wt.WithdrawApply.AuditTime,
		&wt.WithdrawApply.AuditComment,
		&wt.WithdrawApply.WithdrawType,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("提现申请不存在:%+v, %v\n", err, arg.GetValue())
			return nil, err
		}
		log.Printf("%+v\n", err)
		return nil, err
	}
	log.Println("waa", wt)
	return wt, nil
}

// UpdateWithdrawTransfer 修改提现转账
func (agt *DbAdminAgent) UpdateWithdrawTransfer(ctx context.Context, arg *dbproto.WithdrawTransfer) (*dbproto.IntValue, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	esql := fmt.Sprintf("UPDATE %s "+
		"SET step = $1, is_success = $2, certificate_url = $3, transfer_amount = $4, pay_no = $5, pay_sn = $6, target_account = $7, opetator = $8, operate_time = $9 WHERE id = $10", TABLE_WITHDRAW_TRANSFER)
	st, err := agt.dbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	log.Println("SQL", esql)
	res, err := st.Exec(arg.GetStep(), arg.GetIsSuccess(), arg.GetCertificateUrl(), arg.GetTransferAmount(), arg.GetPayNo(), arg.GetPaySn(), arg.GetTargetAccount(), arg.GetOperator(), arg.GetOperateTime, arg.GetId())
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		log.Printf("RowsAffected %v\n", err)
		return nil, err
	}
	return &dbproto.IntValue{
		Value: affect,
	}, nil
}
