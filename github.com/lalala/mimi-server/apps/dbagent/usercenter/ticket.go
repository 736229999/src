package usercenter

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"strings"

	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"golang.org/x/net/context"
)

func (agt *DbUsercenterAgent) QueryBuycaiTicket(ctx context.Context, arg *dbproto.IntValue) (*dbproto.BuycaiTicket, error) {
	ticketId := arg.GetValue()
	var valOrderId sql.NullInt64
	ticket := &dbproto.BuycaiTicket{}
	fieldList := []string{
		"id", "account_id", "use_base", "use_sub", "title",
		"restrict_id", "restrict_type", "order_id",
		"valid_start", "valid_end", "add_time",
	}
	values := []interface{}{
		&ticket.Id, &ticket.AccountId, &ticket.UseBase, &ticket.UseSub, &ticket.Title,
		&ticket.RestrictId, &ticket.RestrictType, &valOrderId,
		&ticket.ValidStart, &ticket.ValidEnd, &ticket.AddTime,
	}
	esql := fmt.Sprintf("SELECT %s FROM %s WHERE id=%d", strings.Join(fieldList, ", "), TABLE_TICKET, ticketId)
	if err := agt.dbConn.QueryRow(esql).Scan(values...); err != nil {
		if err == sql.ErrNoRows {
			return ticket, nil
		}
		log.Println(err, esql)
		return nil, err
	}
	if valOrderId.Valid {
		ticket.OrderId = valOrderId.Int64
	}
	return ticket, nil
}

func (agt *DbUsercenterAgent) QueryBuycaiTickets(arg *dbproto.IntValue, stream dbproto.DbUsercenterAgent_QueryBuycaiTicketsServer) error {
	accountId := arg.GetValue()
	fieldList := []string{
		"id", "account_id", "use_base", "use_sub", "title",
		"restrict_id", "restrict_type", "order_id",
		"valid_start", "valid_end", "add_time",
	}
	esql := fmt.Sprintf("SELECT %s FROM %s WHERE account_id=%d ORDER BY id DESC", strings.Join(fieldList, ", "), TABLE_TICKET, accountId)
	rows, err := agt.dbConn.Query(esql)
	if err != nil {
		log.Println(err, esql)
		return err
	}

	for rows.Next() {
		var valOrderId sql.NullInt64
		ticket := &dbproto.BuycaiTicket{}
		values := []interface{}{
			&ticket.Id, &ticket.AccountId, &ticket.UseBase, &ticket.UseSub, &ticket.Title,
			&ticket.RestrictId, &ticket.RestrictType, &valOrderId,
			&ticket.ValidStart, &ticket.ValidEnd, &ticket.AddTime,
		}
		if err := rows.Scan(values...); err != nil {
			log.Println(err, esql)
			break
		}
		if valOrderId.Valid {
			ticket.OrderId = valOrderId.Int64
		}
		stream.Send(ticket)
	}

	if err = rows.Err(); err != nil {
		log.Println(err, esql)
	}
	return err
}

func (agt *DbUsercenterAgent) QueryValidBuycaiTickets(arg *dbproto.IntValue, stream dbproto.DbUsercenterAgent_QueryValidBuycaiTicketsServer) error {
	accountId := arg.GetValue()
	now := time.Now().Unix()
	fieldList := []string{
		"id", "account_id", "use_base", "use_sub", "title",
		"restrict_id", "restrict_type", "order_id",
		"valid_start", "valid_end", "add_time",
	}
	// AND (order_id=0 OR order_id=NULL)
	esql := fmt.Sprintf("SELECT %s FROM %s WHERE account_id=%d AND valid_start<=%d AND %d<=valid_end ORDER BY id DESC",
		strings.Join(fieldList, ", "), TABLE_TICKET, accountId, now, now)
	rows, err := agt.dbConn.Query(esql)
	if err != nil {
		log.Println(err, esql)
		return err
	}

	for rows.Next() {
		var valOrderId sql.NullInt64
		ticket := &dbproto.BuycaiTicket{}
		values := []interface{}{
			&ticket.Id, &ticket.AccountId, &ticket.UseBase, &ticket.UseSub, &ticket.Title,
			&ticket.RestrictId, &ticket.RestrictType, &valOrderId,
			&ticket.ValidStart, &ticket.ValidEnd, &ticket.AddTime,
		}
		if err := rows.Scan(values...); err != nil {
			log.Println(err, esql)
			break
		}
		if valOrderId.Valid {
			ticket.OrderId = valOrderId.Int64
		}
		if ticket.OrderId > 0 {
			continue
		}
		stream.Send(ticket)
	}

	if err = rows.Err(); err != nil {
		log.Println(err, esql)
	}
	return err
}

func (agt *DbUsercenterAgent) InsertBuycaiTickets(ctx context.Context, arg *dbproto.InsertBuycaiTicketArg) (*dbproto.Nil, error) {
	tx, err := agt.dbConn.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	_, err = InsertTickets(arg.GetList(), *tx)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}
	return &dbproto.Nil{}, tx.Commit()
}

func InsertTickets(list []*dbproto.BuycaiTicket, tx sql.Tx) (*dbproto.Nil, error) {
	log.Printf("%+v\n", list)
	for _, ticket := range list {
		accountId := ticket.GetAccountId()
		validStart := ticket.GetValidStart()
		validEnd := ticket.GetValidEnd()
		addTime := time.Now().Unix()
		useBase := ticket.GetUseBase()
		useSub := ticket.GetUseSub()
		title := ticket.GetTitle()
		restrictId := ticket.GetRestrictId()
		restrictType := ticket.GetRestrictType()

		//判断是否已经过期.
		if validEnd >= time.Now().Unix() {
			esql := fmt.Sprintf("INSERT INTO %s (account_id, use_base, use_sub, valid_start, valid_end, add_time, title, restrict_id, restrict_type) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", TABLE_TICKET)
			_, err := tx.Exec(esql, accountId, useBase, useSub, validStart, validEnd, addTime, title, restrictId, restrictType)
			if err != nil {
				log.Println(err, esql)
				return nil, err
			}
		} else {
			log.Println("注意，有活动已经过期 开始时间:", validStart, "结束时间：", validEnd)
		}
	}
	return &dbproto.Nil{}, nil
}

func (agt *DbUsercenterAgent) DeleteBuycaiTickets(ctx context.Context, arg *dbproto.DeleteBuycaiTicketArg) (*dbproto.Nil, error) {
	tx, err := agt.dbConn.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for _, id := range arg.GetList() {
		esql := fmt.Sprintf("DELETE FROM %s WHERE id=%d", TABLE_TICKET, id)
		_, err = tx.Exec(esql)
		if err != nil {
			log.Println(err, esql)
			tx.Rollback()
			return nil, err
		}
	}
	return &dbproto.Nil{}, tx.Commit()
}

//新人领取礼包.
func (agt *DbUsercenterAgent) InsertPhoneRegistGiftReceived(ctx context.Context, arg *dbproto.PhoneRegistGift) (*dbproto.Nil, error) {

	tx, err := agt.dbConn.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	//添加购彩券.
	_, err = InsertTickets(arg.GetList(), *tx)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}

	//添加积分.
	credits := &dbproto.ChangeVirtualFundArg{
		AccountId: arg.GetAccountId(),
		Reason:    0,
		Var:       int32(arg.GetCredits()),
		Detail:    "新人领取礼包",
	}
	if _, err := agt.changeVirtualFund(*tx, "credits", credits); err != nil {
		log.Printf("赠送积分%+v\n", err)
		tx.Rollback()
		return nil, err
	}

	_, err = setPhoneRegistGiftReceived(&dbproto.IntValue{Value: arg.GetAccountId()}, *tx)
	if err != nil {
		log.Printf("%+v\n", err)
		tx.Rollback()
		return nil, err
	}

	return &dbproto.Nil{}, tx.Commit()
}

//设置领取新人礼包.
func setPhoneRegistGiftReceived(arg *dbproto.IntValue, tx sql.Tx) (*dbproto.Nil, error) {
	accountId := arg.GetValue()
	now := time.Now().Unix()
	esql := fmt.Sprintf("INSERT INTO %s (account_id, receive_time) VALUES (%d,%d)", TABLE_PHONE_REGIST_GIFT_HISTORY, accountId, now)
	_, err := tx.Exec(esql)
	if err != nil {
		log.Println(err, esql)
		return nil, err
	}
	return &dbproto.Nil{}, nil
}
