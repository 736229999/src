package thirdapi

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"golang.org/x/net/context"
)

func (agt *DbThirdApiAgent) InsertSms(ctx context.Context, arg *dbproto.Sms) (*dbproto.Nil, error) {
	tx, err := agt.dbConn.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	query := fmt.Sprintf(`INSERT INTO sms_detail (sms_type, content, vendor, send_time, is_success, result, code, expire_time, sign) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`)
	stmt, err := tx.Prepare(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var detailId int64
	if err = stmt.QueryRow(arg.SmsType, arg.Content, arg.Vendor, arg.SendTime, arg.IsSuccess, arg.Result, arg.Code, arg.ExpireTime, arg.Sign).Scan(&detailId); err != nil {
		log.Println(err, query)
		tx.Rollback()
		return nil, err
	}

	query = fmt.Sprintf(`INSERT INTO sms_history (phone, detail_id) VALUES ($1, $2)`)
	stmt, err = tx.Prepare(query)
	if err != nil {
		log.Println(err, query)
		tx.Rollback()
		return nil, err
	}

	query = `INSERT INTO sms_stats (phone, count_time, latest_time, hourly_count, daily_count, total_count) 
		VALUES($1, $2, $3, $4, $5, $6)
		ON CONFLICT (phone) DO UPDATE SET 
		latest_time=$7, hourly_count=sms_stats.hourly_count+1, daily_count=sms_stats.daily_count+1, total_count=sms_stats.total_count+1`
	stmtUpsertStats, err := tx.Prepare(query)
	if err != nil {
		log.Println(err, query)
		tx.Rollback()
		return nil, err
	}

	now := time.Now()
	for _, phone := range arg.GetPhoneList() {
		_, err = stmt.Exec(phone, detailId)
		if err != nil {
			log.Println(err)
			tx.Rollback()
			return nil, err
		}

		_, err = stmtUpsertStats.Exec(phone, now.Unix(), now.Unix(), 1, 1, 1, now.Unix())
		if err != nil {
			log.Println(err)
			tx.Rollback()
			return nil, err
		}
	}

	return &dbproto.Nil{}, tx.Commit()
}

func (agt *DbThirdApiAgent) QuerySms(ctx context.Context, arg *dbproto.StringValue) (*dbproto.Sms, error) {
	cols := "id, sms_type, content, vendor, send_time, is_success, result, code, expire_time, sign"
	query := fmt.Sprintf("SELECT %s FROM sms_detail WHERE id=(SELECT detail_id FROM sms_history WHERE phone=$1 ORDER BY id DESC LIMIT 1)", cols)
	stmt, err := agt.dbConn.Prepare(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	result := &dbproto.Sms{}
	phone := arg.GetValue()
	err = stmt.QueryRow(phone).Scan(&result.Id, &result.SmsType, &result.Content, &result.Vendor, &result.SendTime, &result.IsSuccess, &result.Result, &result.Code, &result.ExpireTime, &result.Sign)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		return nil, err
	}
	return result, nil
}

func (agt *DbThirdApiAgent) UpdateSmsStats(ctx context.Context, arg *dbproto.SmsStats) (*dbproto.Nil, error) {
	query := "UPDATE sms_stats SET count_time = $1, latest_time=$2, hourly_count=$3, daily_count=$4, total_count=$5 WHERE id=$6"
	stmt, err := agt.dbConn.Prepare(query)
	if err != nil {
		log.Println(err, query)
		return nil, err
	}

	_, err = stmt.Exec(arg.CountTime, arg.LatestTime, arg.HourlyCount, arg.DailyCount, arg.TotalCount, arg.Id)
	if err != nil {
		log.Println(err, query)
		return nil, err
	}
	return &dbproto.Nil{}, nil
}

func (agt *DbThirdApiAgent) SetSmsExpired(ctx context.Context, arg *dbproto.IntValue) (*dbproto.Nil, error) {
	query := fmt.Sprintf("UPDATE sms_detail SET expire_time=%d WHERE id=%d", time.Now().Unix(), arg.GetValue())
	_, err := agt.dbConn.Exec(query)
	if err != nil {
		log.Println(err, query)
		return nil, err
	}
	return &dbproto.Nil{}, nil
}

func (agt *DbThirdApiAgent) QuerySmsStats(ctx context.Context, arg *dbproto.StringValue) (*dbproto.SmsStats, error) {
	query := "SELECT id, phone, latest_time, hourly_count, daily_count, total_count FROM sms_stats WHERE phone=$1"
	stmt, err := agt.dbConn.Prepare(query)
	if err != nil {
		log.Println(err, query)
		return nil, err
	}

	result := &dbproto.SmsStats{}
	if err = stmt.QueryRow(arg.GetValue()).Scan(&result.Id, &result.Phone, &result.LatestTime, &result.HourlyCount, &result.DailyCount, &result.TotalCount); err != nil {
		if err != sql.ErrNoRows {
			log.Println(err, query)
			return nil, err
		}
	}
	return result, nil

}

func (agt *DbThirdApiAgent) InsertIdcard(ctx context.Context, arg *dbproto.Idcard) (*dbproto.Nil, error) {
	query := `INSERT INTO idcard (cardno, realname, add_time) VALUES ($1, $2, $3)`
	stmt, err := agt.dbConn.Prepare(query)
	if err != nil {
		log.Println(err, query)
		return nil, err
	}
	_, err = stmt.Exec(arg.GetCardno(), arg.GetRealname(), time.Now().Unix())
	if err != nil {
		log.Println(err, query)
		return nil, err
	}
	return &dbproto.Nil{}, nil
}

func (agt *DbThirdApiAgent) QueryIdcard(ctx context.Context, arg *dbproto.StringValue) (*dbproto.Idcard, error) {
	query := "SELECT id, cardno, realname, add_time FROM idcard WHERE cardno=$1"
	stmt, err := agt.dbConn.Prepare(query)
	if err != nil {
		log.Println(err, query)
		return nil, err
	}

	result := &dbproto.Idcard{}
	if err = stmt.QueryRow(arg.GetValue()).Scan(&result.Id, &result.Cardno, &result.Realname, &result.AddTime); err != nil {
		if err != sql.ErrNoRows {
			log.Println(err, query)
			return nil, err
		}
	}
	return result, nil
}

func (agt *DbThirdApiAgent) InsertBankcard(ctx context.Context, arg *dbproto.Bankcard) (*dbproto.Nil, error) {
	cols := "idcard_no, bankcard_no, card_type, bankname, add_time, phone"
	query := fmt.Sprintf(`INSERT INTO bankcard (%s) VALUES ($1, $2, $3, $4, $5, $6)`, cols)
	stmt, err := agt.dbConn.Prepare(query)
	if err != nil {
		log.Println(err, query)
		return nil, err
	}

	idcardNo := arg.GetIdcardNo()
	bankcardNo := arg.GetBankcardNo()
	cardType := arg.GetCardType()
	bankname := arg.GetBankname()
	addTime := time.Now().Unix()
	phone := arg.GetPhone()
	_, err = stmt.Exec(idcardNo, bankcardNo, cardType, bankname, addTime, phone)
	if err != nil {
		log.Println(err, query)
		return nil, err
	}
	return &dbproto.Nil{}, nil
}

func (agt *DbThirdApiAgent) QueryBankcard(ctx context.Context, arg *dbproto.StringValue) (*dbproto.Bankcard, error) {
	cols := "bankcard.id, bankcard.idcard_no, bankcard.bankcard_no, bankcard.card_type, bankcard.bankname, bankcard.add_time, bankcard.phone, idcard.realname"
	query := fmt.Sprintf(`SELECT %s FROM (bankcard FULL JOIN idcard ON idcard_no=idcard.cardno) WHERE bankcard_no=$1`, cols)
	stmt, err := agt.dbConn.Prepare(query)
	if err != nil {
		log.Println(err, query)
		return nil, err
	}

	result := &dbproto.Bankcard{}
	var valPhone sql.NullString
	if err = stmt.QueryRow(arg.GetValue()).Scan(&result.Id, &result.IdcardNo, &result.BankcardNo,
		&result.CardType, &result.Bankname, &result.AddTime, &valPhone, &result.Realname); err != nil {
		if err != sql.ErrNoRows {
			log.Println(err, query)
			return nil, err
		}
	}
	if valPhone.Valid {
		result.Phone = valPhone.String
	}
	return result, nil
}
