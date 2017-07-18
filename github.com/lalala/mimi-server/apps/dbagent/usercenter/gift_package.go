package usercenter

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"golang.org/x/net/context"
)

const (
	GIFT_PACKAGE_EXCHANGE = 0
	GIFT_PACKAGE_INVITE   = 1

	INVITER = 0
	INVITEE = 1
)

//type Gift struct {
//	Credits  int64 `json:"credits"`
//	Tickets []*dbproto.BuycaiTicket  `json:"tickets"`
//}

func (agt *DbUsercenterAgent) QueryCountActivity(ctx context.Context, arg *dbproto.GiftPackage) (*dbproto.IntValue, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	esql := fmt.Sprintf(`SELECT COUNT(*) AS num FROM %s WHERE gift_type = %d`, TABLE_GIFT_PACKAGE, GIFT_PACKAGE_EXCHANGE)
	var num int64 = 0

	if err := agt.dbConn.QueryRow(esql).Scan(&num); err != nil {
		log.Println(err, esql)
		return nil, err
	}
	return &dbproto.IntValue{Value: num}, nil
}

//创建礼包.
func (agt *DbUsercenterAgent) CreateGiftPackage(ctx context.Context, arg *dbproto.GiftCdkeyArg) (*dbproto.Nil, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	tx, err := agt.dbConn.Begin()
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	gift := arg.Gift
	title := gift.GetTitle()
	content_desc := gift.GetContentDesc()
	content := gift.GetContent()
	total_exchange_num := gift.GetTotalExchangeNum()
	gift_type := gift.GetGiftType()
	add_time := time.Now().Unix()

	esql := fmt.Sprintf(`INSERT INTO %s(title, content_desc, content, total_exchange_num, gift_type, add_time) VALUES ('%s', '%s','%s',%d,%d,%d) RETURNING id`, title, content_desc, content, total_exchange_num, gift_type, add_time)

	var id int64 = 0
	if err = tx.QueryRow(esql).Scan(&id); err != nil {
		log.Println(err, esql)
		tx.Rollback()
		return nil, err
	}

	arg.Cdkey.GiftPackageId = id
	if _, err = InsertCdKeyBatch(*tx, arg.Cdkey); err != nil {
		log.Printf("%+v\n", err)
		tx.Rollback()
		return nil, err
	}

	return &dbproto.Nil{}, tx.Commit()
}

//根据批次来查询礼包.
func (agt *DbUsercenterAgent) QueryGiftByBatch(ctx context.Context, arg *dbproto.IntValue) (*dbproto.StringValue, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	batch := arg.GetValue()

	log.Println("batch:", batch)

	esql := fmt.Sprintf(`SELECT content FROM gift_package WHERE id = (SELECT gift_package_id FROM cdkey_batch WHERE id = %d)`, batch)
	var content string
	if err := agt.dbConn.QueryRow(esql).Scan(&content); err != nil {
		if err == sql.ErrNoRows {
			return &dbproto.StringValue{}, nil
		}
		log.Println(err, esql)
		return nil, err
	}

	return &dbproto.StringValue{Value: content}, nil
}

func (agt *DbUsercenterAgent) InsertGiftUser(ctx context.Context, arg *dbproto.GiftArg) (*dbproto.Nil, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	log.Printf("进到了db:%+v\n", arg)
	//调用.
	tx, err := agt.dbConn.Begin()
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	tickList := arg.GetTicketList()
	if _, err = InsertTickets(tickList, *tx); err != nil {
		log.Printf("%+v\n", err)
		tx.Rollback()
		return nil, err
	}

	credits := arg.GetCredits()
	if _, err = agt.changeVirtualFund(*tx, "credits", credits); err != nil {
		log.Printf("%+v\n", err)
		tx.Rollback()
		return nil, err
	}

	if _, err = InsertGiftExchangeHistory(*tx, arg.GetCdKey(), arg.GetAccountId(), arg.Batch); err != nil {
		log.Printf("%+v\n", err)
		tx.Rollback()
		return nil, err
	}

	return &dbproto.Nil{}, tx.Commit()
}

func InsertGiftExchangeHistory(tx sql.Tx, code string, accountId int64, batch int64) (*dbproto.Nil, error) {

	esql := fmt.Sprintf(`INSERT INTO %s(account_id, cdkey, exchange_time,batch) VALUES (%d, '%s', %d, %d)`, TABLE_EXCHANGE_HISTORY, accountId, code, time.Now().Unix(), batch)
	if _, err := tx.Exec(esql); err != nil {
		log.Printf("%+v\n", err, esql)
		return nil, err
	}

	return nil, nil
}

func (agt *DbUsercenterAgent) QueryExchangeHistoryByAccountId(ctx context.Context, arg *dbproto.IntValue) (*dbproto.StringValue, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	accountId := arg.GetValue()

	esql := fmt.Sprintf(`SELECT cdkey FROM %s WHERE account_id = %d`, TABLE_EXCHANGE_HISTORY, accountId)
	var cdkey string
	if err := agt.dbConn.QueryRow(esql).Scan(&cdkey); err != nil {
		if err == sql.ErrNoRows {
			return &dbproto.StringValue{Value: ""}, nil
		}
		log.Printf("%+v\n", err, esql)
		return nil, err
	}

	return &dbproto.StringValue{Value: cdkey}, nil
}

func InsertCdKeyBatch(tx sql.Tx, arg *dbproto.CdKeyBatchArg) (*dbproto.Nil, error) {

	max_exchang := arg.GetMaxExchange()
	title := arg.GetTitle()
	gift_package_id := arg.GetGiftPackageId()
	valid_start := arg.GetValidStart()
	valid_end := arg.GetValidEnd()
	add_time := time.Now().Unix()

	esql := fmt.Sprintf(`INSERT INTO %s(max_exchange,title,gift_package_id,valid_start,valid_end,add_time) VALUES(%d,%d,'%s',%d,%d,%d,%d)`, TABLE_CDKEY_BATCH, max_exchang, title, gift_package_id, valid_start, valid_end, add_time)
	if _, err := tx.Exec(esql); err != nil {
		log.Printf("%+v\n", err, esql)
		return nil, err
	}
	return &dbproto.Nil{}, nil
}

func (agt *DbUsercenterAgent) QueryMaxExchangeByBatch(ctx context.Context, arg *dbproto.IntValue) (*dbproto.IntValue, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	num := arg.GetValue()

	esql := fmt.Sprintf(`SELECT max_exchange FROM %s WHERE id = %d`, TABLE_CDKEY_BATCH, num)

	var max int64 = 0
	if err := agt.dbConn.QueryRow(esql).Scan(&max); err != nil {
		log.Printf("%+v\n", err, esql)
		return nil, err
	}

	return &dbproto.IntValue{Value: max}, nil
}

func (agt *DbUsercenterAgent) QueryExchangeNumByBatch(ctx context.Context, arg *dbproto.IntValue) (*dbproto.IntValue, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	num := arg.GetValue()

	esql := fmt.Sprintf(`SELECT COUNT(*) AS num FROM %s WHERE batch = %d`, TABLE_EXCHANGE_HISTORY, num)
	var count int64
	if err := agt.dbConn.QueryRow(esql).Scan(&count); err != nil {
		log.Printf("%+v\n", err, esql)
		return nil, err
	}

	return &dbproto.IntValue{Value: count}, nil
}

func (agt *DbUsercenterAgent) SetUserInviteRelation(ctx context.Context, arg *dbproto.Gift) (*dbproto.Nil, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	tx, err := agt.dbConn.Begin()
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	//邀请者的购彩券.
	tickListInviter, creditInviter := Gift(arg, &dbproto.IntValue{Value: INVITER})
	if _, err = agt.giftPackage(*tx, tickListInviter, creditInviter); err != nil {
		log.Printf("%+v\n", err)
		tx.Rollback()
		return nil, err
	}

	//被邀请者的购彩券.
	tickListInvitee, creditInvitee := Gift(arg, &dbproto.IntValue{Value: INVITEE})
	if _, err = agt.giftPackage(*tx, tickListInvitee, creditInvitee); err != nil {
		log.Printf("%+v\n", err)
		tx.Rollback()
		return nil, err
	}

	//获取购彩圈张数.

	tickets_num := int64(len(arg.GetTicketList()))

	//获取兑换券的价值.
	var money int64 = 0
	if tickets_num > 0 {
		for _, v := range arg.GetTicketList() {
			log.Printf("%+v\n", v)
			money += int64(v.UseSub)
		}
	}

	//设置邀请关系.
	if _, err = InsertUserInviteRelation(*tx, arg.GetInviter(), arg.GetInvitee(), arg.GetCredits(), tickets_num, money); err != nil {
		log.Printf("设置邀请关系%+v\n", err)
		tx.Rollback()
		return nil, err
	}

	log.Println("邀请关系设置成功")
	return &dbproto.Nil{}, tx.Commit()
}

func Gift(arg *dbproto.Gift, accountId *dbproto.IntValue) (tick []*dbproto.BuycaiTicket, credits *dbproto.ChangeVirtualFundArg) {

	var account_id int64
	if accountId.GetValue() == INVITER {
		account_id = arg.GetInviter()
	} else {
		account_id = arg.GetInvitee()
	}

	//邀请者的购彩券.
	var tickList []*dbproto.BuycaiTicket
	for _, v := range arg.GetTicketList() {

		tick := &dbproto.BuycaiTicket{}
		tick.AccountId = account_id
		tick.ValidEnd = v.ValidEnd
		tick.ValidStart = v.ValidStart
		tick.Title = v.Title
		tick.OrderId = v.OrderId
		tick.RestrictId = v.RestrictId
		tick.RestrictType = v.RestrictType
		tick.UseBase = v.UseBase
		tick.UseSub = v.UseSub
		tickList = append(tickList, tick)
	}

	//邀请的积分.
	credit := &dbproto.ChangeVirtualFundArg{
		Reason:    0,
		Var:       int32(arg.GetCredits()),
		AccountId: account_id,
		Detail:    "赠送积分",
	}

	return tickList, credit
}

//查询用户是否已经存在邀请关系.
func (agt *DbUsercenterAgent) QueryUserInviteRelation(ctx context.Context, arg *dbproto.IntValue) (*dbproto.BoolValue, error) {

	inviteeId := arg.GetValue()
	esql := fmt.Sprintf(`SELECT COUNT(*) AS num FROM %s WHERE invitee = %d`, TABLE_INVITE_HISTORY, inviteeId)

	var num int64
	if err := agt.dbConn.QueryRow(esql).Scan(&num); err != nil {
		log.Printf("查询用户是否已经存在邀请关系%+v\n", err, esql)
		return nil, err

	} else {
		if num > 0 {
			log.Println("该用户：", inviteeId, "存在邀请记录")
			return &dbproto.BoolValue{Value: true}, nil
		} else {
			//没有查询到记录.
			log.Println("没有查到这个:", inviteeId, "用户的邀请记录")
			return &dbproto.BoolValue{Value: false}, nil
		}
	}
}

//查询邀请礼包.
func (agt *DbUsercenterAgent) QueryInviteGift(ctx context.Context, age *dbproto.Nil) (*dbproto.StringValue, error) {

	esql := fmt.Sprintf(`SELECT content FROM %s WHERE gift_type = %d ORDER BY id DESC LIMIT 1`, TABLE_GIFT_PACKAGE, GIFT_PACKAGE_INVITE)

	var content string
	if err := agt.dbConn.QueryRow(esql).Scan(&content); err != nil {
		if err == sql.ErrNoRows {
			log.Println("没有赠送的东西")
			return nil, errors.New("没有赠送的东西")
		} else {
			log.Println(err, esql)
			return nil, err
		}
	}
	return &dbproto.StringValue{Value: content}, nil

	//gift := &dbproto.GiftPackageArg{}
	//err := json.Unmarshal([]byte(content), gift)
	//if err != nil {
	//	log.Printf("%+v\n", err)
	//	return nil, err
	//}
	//return gift, nil
}

//添加用户的邀请关系.
func InsertUserInviteRelation(tx sql.Tx, inviter int64, invitee int64, credits int64, tickets_num int64, tickets_money int64) (*dbproto.Nil, error) {

	now := time.Now().Unix()
	esql := fmt.Sprintf(`INSERT INTO %s(inviter,invitee,accept_time,credits,tickets_num,tickets_money) VALUES (%d,%d,%d,%d,%d,%d)`, TABLE_INVITE_HISTORY, inviter, invitee, now, credits, tickets_num, tickets_money)
	log.Println(inviter, invitee, now)
	if _, err := tx.Exec(esql); err != nil {
		log.Println("添加用户的邀请关系:", err, esql)
		return nil, err
	}

	return &dbproto.Nil{}, nil
}

func (agt *DbUsercenterAgent) QueryInviteRelation(ctx context.Context, arg *dbproto.InviteRelationArg) (*dbproto.BoolValue, error) {

	esql := fmt.Sprintf(`SELECT COUNT(*) AS num FROM %s WHERE invitee = %d and inviter = %d`, TABLE_INVITE_HISTORY, arg.GetInviter(), arg.GetInvitee())

	var num int64
	if err := agt.dbConn.QueryRow(esql).Scan(&num); err != nil {
		log.Println(err, esql)
		return nil, err
	}

	if num > 0 {
		log.Println("不能邀请自己的邀请人")
		return &dbproto.BoolValue{Value: false}, nil
	} else {
		return &dbproto.BoolValue{Value: true}, nil
	}
}

func (agt *DbUsercenterAgent) giftPackage(tx sql.Tx, Tickets []*dbproto.BuycaiTicket, credits *dbproto.ChangeVirtualFundArg) (*dbproto.Nil, error) {

	if _, err := InsertTickets(Tickets, tx); err != nil {
		log.Printf("赠送购彩券%+v\n", err)
		tx.Rollback()
		return nil, err
	}

	log.Println("收到购彩券")

	//赠送积分.
	log.Printf("ccc%+v\n", credits)
	if _, err := agt.changeVirtualFund(tx, "credits", credits); err != nil {
		log.Printf("赠送积分%+v\n", err)
		tx.Rollback()
		return nil, err
	}

	log.Println("收到积分")

	return &dbproto.Nil{}, nil
}

func (agt *DbUsercenterAgent) QueryUserInfoByInvitationCode(ctx context.Context, arg *dbproto.StringValue) (*dbproto.IntValue, error) {

	code := arg.GetValue()
	//查询邀请码是否存在.
	esql := fmt.Sprintf(`SELECT account_id FROM %s WHERE invitation_code = '%s'`, TABLE_USERINFO, code)

	//邀请者id.
	var inviterId int64
	if err := agt.dbConn.QueryRow(esql).Scan(&inviterId); err != nil {
		if err == sql.ErrNoRows {
			log.Println("没有查询到这个验证码")
			return &dbproto.IntValue{Value: 0}, nil
		} else {
			log.Println(err, esql)
			return nil, err
		}
	}

	log.Println("邀请者id：", inviterId)

	return &dbproto.IntValue{Value: inviterId}, nil
}

//查询用户是否已经实名认证.
func (agt *DbUsercenterAgent) QueryUserAuthenticateByAccountId(ctx context.Context, arg *dbproto.IntValue) (*dbproto.BoolValue, error) {

	inviteeId := arg.GetValue()
	esql := fmt.Sprintf(`SELECT COUNT(*) AS num FROM %s WHERE account_id = %d`, TABLE_IDCARD, inviteeId)

	var num int64
	if err := agt.dbConn.QueryRow(esql).Scan(&num); err != nil {
		log.Println(err, esql)
		return nil, err
	}

	log.Println("num:", num, "sql:", esql)

	if num > 0 {
		log.Println("已经实名认证")
		return &dbproto.BoolValue{Value: true}, nil
	} else {
		log.Println("没有实名认证")
		return &dbproto.BoolValue{Value: false}, nil
	}
}

//根据批次，查询批次的信息.
func (agt *DbUsercenterAgent) QueryCdkeyBatchByBatch(ctx context.Context, arg *dbproto.IntValue) (*dbproto.CdkeyBatch, error) {

	batch := arg.GetValue()

	esql := fmt.Sprintf(`SELECT id,max_exchange,title,gift_template_id,valid_start,valid_end,add_time FROM %s WHERE id = %d`, TABLE_CDKEY_BATCH, batch)

	info := dbproto.CdkeyBatch{}
	if err := agt.dbConn.QueryRow(esql).Scan(&info.Id, &info.MaxExchange, &info.Title, &info.GiftTemplateId, &info.ValidStart, &info.ValidEnd, &info.AddTime); err != nil {
		if err == sql.ErrNoRows {
			return &dbproto.CdkeyBatch{}, nil
		}
		log.Println(err, esql)
		return nil, err
	}

	return &info, nil
}

//验证当前兑换码的状态，判断是否已经被兑换.
func (agt *DbUsercenterAgent) QueryCdkeyStatus(ctx context.Context, arg *dbproto.StringValue) (*dbproto.BoolValue, error) {

	code := arg.GetValue()

	//开启事物，防止并发.
	tx, err := agt.dbConn.Begin()
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	esql := fmt.Sprintf(`SELECT COUNT(*) AS num FROM %s WHERE cdkey = '%s'`, TABLE_EXCHANGE_HISTORY, code)

	var num int64
	if err = tx.QueryRow(esql).Scan(&num); err != nil {
		log.Printf("%+v\n", err)
		tx.Rollback()
		return nil, err
	}

	if num > 0 {
		log.Println("该验证码：", code, "已经兑换过了")
		return &dbproto.BoolValue{Value: true}, tx.Commit()
	}

	return &dbproto.BoolValue{Value: false}, tx.Commit()
}

//查询用户的邀请状态.
func (agt *DbUsercenterAgent) QueryUserInviteStatus(ctx context.Context, arg *dbproto.IntValue) (*dbproto.BoolValue, error) {

	esql := fmt.Sprintf(`SELECT COUNT(*) AS num FROM %s WHERE invitee = %d`, TABLE_INVITE_HISTORY, arg.GetValue())

	var num int64 = 0
	if err := agt.dbConn.QueryRow(esql).Scan(&num); err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	var status bool = false
	if num > 0 {
		status = true
	}
	return &dbproto.BoolValue{Value: status}, nil
}

//查询用户的邀请数据.
func (agt *DbUsercenterAgent) QueryUserInviteInfo(ctx context.Context, arg *dbproto.IntValue) (*dbproto.UserInviteInfo, error) {

	esql := fmt.Sprintf(`SELECT COUNT(invitee) AS invite_num, SUM(credits) AS credits, SUM(tickets_num) AS tickets_num, SUM(tickets_money) AS tickets_money, (SELECT COUNT(*) FROM %s where invitee = %d) as num from %s where inviter = %d`, TABLE_INVITE_HISTORY, arg.GetValue(), TABLE_INVITE_HISTORY, arg.GetValue())

	inviteInfo := &dbproto.UserInviteInfo{}
	var num int64 = 0
	var credits, tickets_num, tickets_money sql.NullInt64
	err := agt.dbConn.QueryRow(esql).Scan(&inviteInfo.InviteNum, &credits, &tickets_num, &tickets_money, &num)
	if err != nil {
		log.Println(err, esql)
		return nil, err
	}

	var status bool = false
	if num > 0 {
		status = true
	}
	inviteInfo.UserInviteStatus = status

	if credits.Valid {
		inviteInfo.Credits = credits.Int64
	}

	if tickets_num.Valid {
		inviteInfo.TicketsNum = tickets_num.Int64
	}

	if tickets_money.Valid {
		inviteInfo.TicketsMoney = tickets_money.Int64
	}

	return inviteInfo, nil
}

//func (agt *DbUsercenterAgent) SetPhoneRegistGiftReceived(ctx context.Context, arg *dbproto.IntValue) (*dbproto.Nil, error) {
//	accountId := arg.GetValue()
//	now := time.Now().Unix()
//	esql := fmt.Sprintf("INSERT INTO %s (account_id, receive_time) VALUES (%d,%d)", TABLE_PHONE_REGIST_GIFT_HISTORY, accountId, now)
//	_, err := agt.dbConn.Exec(esql)
//	if err != nil {
//		log.Println(err, esql)
//		return nil, err
//	}
//	return &dbproto.Nil{}, nil
//}

func (agt *DbUsercenterAgent) QueryPhoneRegistGiftReceived(ctx context.Context, arg *dbproto.IntValue) (*dbproto.BoolValue, error) {
	accountId := arg.GetValue()
	var count int32
	esql := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE account_id=%d", TABLE_PHONE_REGIST_GIFT_HISTORY, accountId)
	if err := agt.dbConn.QueryRow(esql).Scan(&count); err != nil {
		log.Println(err, esql)
		return nil, err
	}

	return &dbproto.BoolValue{Value: count > 0}, nil
}

//获取用户的注册礼包.
func (agt *DbUsercenterAgent) QueryPhoneUserRegisterGift(ctx context.Context, arg *dbproto.Nil) (*dbproto.StringValue, error) {

	esql := fmt.Sprintf(`SELECT content FROM gift_package WHERE gift_type = %d`, dbproto.GiftPackageType_register)
	st, err := agt.dbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err, esql)
		return nil, err
	}

	var giftJson string
	if err = st.QueryRow().Scan(&giftJson); err != nil {
		if err == sql.ErrNoRows {
			return &dbproto.StringValue{}, nil
		}
		log.Printf("%+v\n", err)
		return nil, err
	}

	return &dbproto.StringValue{Value: giftJson}, nil
}
