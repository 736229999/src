package usercenter

import (
	"golang.org/x/net/context"
	"github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"log"
	"fmt"
	"strings"
	"encoding/json"
	"time"
	"database/sql"
	"github.com/caojunxyz/mimi-server/proto"
)

//根据活动id获取礼包模板.
func (agt *DbUsercenterAgent) QueryGiftTemplateById (ctx context.Context, arg *dbproto.IntValue) (*dbproto.GiftTemplate, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	filed := "id, title, content_desc, content, add_time, creator"
	esql := fmt.Sprintf("SELECT %s FROM gift_template WHERE id = (SELECT package_id FROM act_activity WHERE id = %d)", filed, arg.GetValue())
	st, err := agt.dbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	template := &dbproto.GiftTemplate{}
	value := []interface{}{
		&template.Id, &template.Title, &template.ContentDesc,
		&template.ContentJson, &template.AddTime, &template.Creator,
	}
	if err = st.QueryRow().Scan(value...); err != nil {
		if err == sql.ErrNoRows {
			return template, nil
		}
		log.Printf("%+v\n", err)
		return nil, err
	}

	if err = json.Unmarshal([]byte(template.GetContentJson()), &template.Content); err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	return template, nil
}

func (agt *DbUsercenterAgent) InsertUserGift (tx sql.Tx, arg *dbproto.UserGiftPackage) (*dbproto.Nil, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	filedList := []string{
		"act_activity_id", "account_id",
		"content", "status", "create_time",
	}

	esql := fmt.Sprintf("INSERT INTO user_gift (%s) VALUES ($1, $2, $3, $4, $5)", strings.Join(filedList, ","))
	st, err := tx.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err, esql)
		return nil, err
	}

	content, err  := json.Marshal(arg.GetGift())
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	_, err = st.Exec(arg.GetActActivityId(), arg.GetAccountId(), content, proto.GiftStatus_WAIT_RECEIVE, time.Now().Unix())
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	return &dbproto.Nil{}, nil
}

//领取礼包.
func (agt *DbUsercenterAgent) ReceiveGift (ctx context.Context, arg *dbproto.UserGiftPackage) (*dbproto.Nil, error) {

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

	//添加到user_gift.
	_, err = agt.InsertUserGift(*tx, arg)
	if err != nil {
		log.Printf("%+v\n", err)
		tx.Rollback()
		return nil, err
	}

	//添加积分.
	credits := &dbproto.ChangeVirtualFundArg{
		Reason:    0,
		Var:       arg.GetGift().GetCredits(),
		AccountId: arg.GetAccountId(),
		Detail:    "赠送积分",
	}
	if _, err := agt.changeVirtualFund(*tx, "credits", credits); err != nil {
		log.Printf("赠送积分%+v\n", err)
		tx.Rollback()
		return nil, err
	}

	//添加购彩券.
	var tickList []*dbproto.BuycaiTicket
	for _, v := range arg.GetGift().GetTickets() {

		tick := &dbproto.BuycaiTicket{}
		tick.AccountId = arg.GetAccountId()
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

	if _, err := InsertTickets(tickList, *tx); err != nil {
		log.Printf("赠送购彩券%+v\n", err)
		tx.Rollback()
		return nil, err
	}

	return &dbproto.Nil{}, nil
}


