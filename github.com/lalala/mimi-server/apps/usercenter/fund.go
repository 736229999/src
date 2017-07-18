package main

import (
	"io"
	"log"
	"math"
	"net/http"
	"time"

	"github.com/caojunxyz/gotu"
	apiproto "github.com/caojunxyz/mimi-api/proto"
	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/utils"
	"golang.org/x/net/context"
)

func (srv *UcServer) queryFundInfo(accountId int64) *apiproto.FundInfo {
	if accountId <= 0 {
		return nil
	}
	ret, err := srv.dbClient.QueryFund(context.Background(), &dbproto.IntValue{Value: accountId})
	if err != nil {
		log.Println(accountId, err)
		return nil
	}
	info := &apiproto.FundInfo{
		Total:   ret.GetBalance() + ret.GetCai(),
		Balance: ret.GetBalance(),
		Cai:     ret.GetCai(),
		Freeze:  ret.GetFreezeBalance() + ret.GetFreezeCai(),
	}

	return info
}

func (srv *UcServer) queryBankcard(accountId int64) *apiproto.Bankcard {
	log.Println("queryBankcard")
	if accountId <= 0 {
		return nil
	}
	card, err := srv.dbClient.QueryAccountBankcard(context.Background(), &dbproto.IntValue{Value: accountId})
	if err != nil {
		log.Println(err)
		return nil
	}

	return &apiproto.Bankcard{
		Id:       card.Id,
		Cardno:   maskPrivacyBankcardNo(card.Cardno),
		Bankname: apiproto.Bankname(card.Bankname),
		Cardtype: card.Cardtype,
	}
}

func (srv *UcServer) queryBuycaiTicketInfo(accountId int64) *apiproto.BuycaiTicketInfo {
	log.Println("queryBuycaiTicketInfo:", accountId)
	if accountId <= 0 {
		return nil
	}
	stream, err := srv.dbClient.QueryBuycaiTickets(context.Background(), &dbproto.IntValue{Value: accountId})
	if err != nil {
		log.Println(err)
		return nil
	}
	info := &apiproto.BuycaiTicketInfo{}
	for {
		v, err := stream.Recv()
		if err == nil {
			status := utils.GetBuycaiTicketStatus(v)
			desc := utils.GetBuycaiTicketRestrictDesc(v)
			ticket := &apiproto.BuycaiTicket{
				Id:           v.GetId(),
				UseBase:      v.GetUseBase(),
				UseSub:       v.GetUseSub(),
				MaxStack:     v.GetMaxStack(),
				ValidStart:   v.GetValidStart(),
				ValidEnd:     v.GetValidEnd(),
				Addtime:      v.GetAddTime(),
				Title:        v.GetTitle(),
				RestrictDesc: desc,
				RestrictType: v.GetRestrictType(),
				RestrictId:   v.GetRestrictId(),
				Status:       status,
			}
			info.Tickets = append(info.Tickets, ticket)
			// log.Printf("%d ---> ticket: %+v\n", accountId, ticket)
			continue
		}
		if err != io.EOF {
			log.Println(err)
		}
		break
	}
	return info
}

func (srv *UcServer) HandleGetFundInfo(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleGetFundInfo")
	accountId, _, err := utils.ParseHttpRequest(w, r, nil)
	if err != nil {
		return
	}
	info := srv.queryFundInfo(accountId)
	// log.Printf("%d ---> %+v\n", accountId, info)
	if info == nil {
		http.Error(w, "queryFundInfo!", http.StatusInternalServerError)
		return
	}
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", info)
}

func (srv *UcServer) HandleGetBankcard(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleGetBankcard")
	accountId, _, err := utils.ParseHttpRequest(w, r, nil)
	if err != nil {
		return
	}
	card := srv.queryBankcard(accountId)
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", card)
}

func (srv *UcServer) HandleDeleteBankcard(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleDeleteBankcard")
	accountId, _, err := utils.ParseHttpRequest(w, r, nil)
	if err != nil {
		return
	}
	_, err = srv.dbClient.DeleteAccountBankcard(context.Background(), &dbproto.IntValue{Value: accountId})
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", &apiproto.Bankcard{})
}

func (srv *UcServer) HandleGetBuycaiTicketInfo(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleGetBuycaiTicketInfo")
	accountId, _, err := utils.ParseHttpRequest(w, r, nil)
	if err != nil {
		return
	}
	info := srv.queryBuycaiTicketInfo(accountId)
	if info == nil {
		http.Error(w, "queryBuycaiTicketInfo!", http.StatusInternalServerError)
		return
	}
	log.Printf("%d --> %+v\n", accountId, info)
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", info)
}

func (srv *UcServer) HandleAddBankcard(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleAddBankcard")
	var msg apiproto.AddBankcardRequest
	accountId, _, err := utils.ParseHttpRequest(w, r, &msg)
	if err != nil {
		return
	}

	log.Printf("%+v\n", msg)
	cardno := msg.GetCardno()
	phone := msg.GetPhone()
	code := msg.GetSmsCode()
	bankname := msg.GetBankname()
	// TODO: 正则校验
	if !srv.verifySmsCode(phone, code) {
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "验证失败", nil)
		return
	}

	idcard, err := srv.dbClient.QueryAccountIdcard(context.Background(), &dbproto.IntValue{Value: accountId})
	if err != nil {
		log.Println(err)
		http.Error(w, "db：QueryIdcard!", http.StatusInternalServerError)
		return
	}

	if idcard.GetAccountId() == 0 {
		log.Println("需要先实名认证:", accountId)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "请先实名认证!", nil)
		return
	}

	bankcard, err := srv.dbClient.QueryAccountBankcard(context.Background(), &dbproto.IntValue{Value: accountId})
	if err != nil {
		log.Println(err)
		http.Error(w, "db：QueryBankcard!", http.StatusInternalServerError)
		return
	}
	if bankcard.Cardno == cardno {
		log.Println("重复绑定:", accountId, cardno)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "银行卡已绑定", nil)
		return
	}

	log.Printf("%d -> %+v\n", accountId, idcard)
	idno := idcard.GetIdno()
	realname := idcard.GetRealname()
	verifyResult := srv.verifyBankcard(idno, realname, cardno, phone)
	if verifyResult == nil {
		log.Println("银行卡核验失败:", accountId, idno, realname, cardno, phone)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "核验失败", nil)
		return
	}

	verifyBankname := convertBankname(verifyResult.Bankname)
	if verifyBankname != bankname {
		log.Println("开户行不一致:", accountId, verifyBankname, bankname)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "开户行不一致", nil)
		return
	}

	if verifyResult.CardType != "借记卡" && verifyResult.CardType != "储蓄卡" {
		log.Println("不是借记卡:", accountId, verifyResult.CardType)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "只能绑定借记卡", nil)
		return
	}

	bankcard = &dbproto.AccountBankcard{
		AccountId: accountId,
		Cardno:    verifyResult.BankcardNo,
		Bankname:  int32(bankname),
		Cardtype:  verifyResult.CardType,
		Phone:     verifyResult.Phone,
		Idno:      verifyResult.IdcardNo,
		Realname:  verifyResult.Realname,
	}
	_, err = srv.dbClient.InsertAccountBankcard(context.Background(), bankcard)
	if err != nil {
		log.Println(err)
		http.Error(w, "db：InsertAccountBankcard!", http.StatusInternalServerError)
		return
	}
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", srv.queryBankcard(accountId))
}

func getFundChangeTitle(typ dbproto.FundChangeType) string {
	switch typ {
	case dbproto.FundChangeType_RECHARGE:
		return "充值"
	case dbproto.FundChangeType_WIN:
		return "中奖"
	case dbproto.FundChangeType_WITHDRAW:
		return "提现"
	case dbproto.FundChangeType_FREEZE:
		return "冻结"
	case dbproto.FundChangeType_UNFREEZE:
		return "解冻"
	case dbproto.FundChangeType_BUYCAI:
		return "购彩"
	case dbproto.FundChangeType_ACTIVITY:
		return "活动"
	}
	return typ.String()
}

func getFundDirect(record *dbproto.FundChangeRecord) int32 {
	if record.VarFreezeBalance > 0 || record.VarFreezeCai > 0 {
		return 0
	}
	if record.VarBalance > 0 || record.VarCai > 0 {
		return 1
	}
	return 2
}

func (srv *UcServer) HandleGetFundHistory(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleGetFundHistory")
	var msg apiproto.FundHistoryPage
	accountId, _, err := utils.ParseHttpRequest(w, r, &msg)
	if err != nil {
		return
	}

	pageSize := int32(20)
	offset := msg.GetPage() * pageSize
	dbArg := &dbproto.QueryFundHistoryArg{AccountId: accountId, Limit: pageSize, Offset: offset}
	switch msg.GetType() {
	case apiproto.FundHistoryPage_Recharge:
		dbArg.ChangeType = dbproto.FundChangeType_RECHARGE
	case apiproto.FundHistoryPage_Withdraw:
		dbArg.ChangeType = dbproto.FundChangeType_WITHDRAW
	case apiproto.FundHistoryPage_Buycai:
		dbArg.ChangeType = dbproto.FundChangeType_BUYCAI
	case apiproto.FundHistoryPage_Win:
		dbArg.ChangeType = dbproto.FundChangeType_WIN
	}

	// TODO: 独立为一个函数
	now := time.Now()
	dbArg.TimeEnd = now.Unix()
	switch msg.GetRange() {
	case apiproto.FundHistoryPage_AllTime:
		dbArg.TimeStart = 0
	case apiproto.FundHistoryPage_Today:
		dbArg.TimeStart = gotu.BeginningOfDay(now).Unix()
	case apiproto.FundHistoryPage_Week:
		dbArg.TimeStart = gotu.BeginningOfDay(now).AddDate(0, 0, -7).Unix()
	case apiproto.FundHistoryPage_Month:
		dbArg.TimeStart = gotu.BeginningOfMonth(now).Unix()
	case apiproto.FundHistoryPage_ThreeMonth:
		dbArg.TimeStart = gotu.BeginningOfMonth(now).AddDate(0, -2, 0).Unix()
	}

	history, err := srv.dbClient.QueryFundHistory(context.Background(), dbArg)
	if err != nil {
		http.Error(w, "db：QueryFundHistory!", http.StatusInternalServerError)
		return
	}

	result := &apiproto.FundHistory{PageSize: pageSize}
	for _, v := range history.List {
		typ := v.GetChangeType()
		tm := v.GetChangeTime()
		balance := v.GetBalance()
		varBalance := v.GetVarBalance()
		varFreezeBalance := v.GetVarFreezeBalance()
		cai := v.GetCai()
		varCai := v.GetVarCai()
		varFreezeCai := v.GetVarFreezeCai()
		comment := v.GetChangeComment()
		if varBalance != 0 {
			record := &apiproto.FundChangeRecord{
				Title:           getFundChangeTitle(typ),
				Desc:            comment,
				Time:            tm,
				Val:             varBalance,
				Remain:          balance,
				Name:            "余额",
				UserOrderId:     v.UserOrderId,
				VendorOrderId:   v.VendorOrderId,
				WithdrawApplyId: v.WithdrawApplyId,
				Direct:          getFundDirect(v),
			}
			result.List = append(result.List, record)
		}
		if varFreezeBalance != 0 {
			record := &apiproto.FundChangeRecord{
				Title:           getFundChangeTitle(typ),
				Desc:            comment,
				Time:            tm,
				Val:             varFreezeBalance,
				Remain:          balance,
				Name:            "余额",
				UserOrderId:     v.UserOrderId,
				VendorOrderId:   v.VendorOrderId,
				WithdrawApplyId: v.WithdrawApplyId,
				Direct:          getFundDirect(v),
			}
			result.List = append(result.List, record)
		}
		if varCai != 0 {
			record := &apiproto.FundChangeRecord{
				Title:           getFundChangeTitle(typ),
				Desc:            comment,
				Time:            tm,
				Val:             varCai,
				Remain:          cai,
				Name:            "彩金",
				UserOrderId:     v.UserOrderId,
				VendorOrderId:   v.VendorOrderId,
				WithdrawApplyId: v.WithdrawApplyId,
				Direct:          getFundDirect(v),
			}
			result.List = append(result.List, record)
		}
		if varFreezeCai != 0 {
			record := &apiproto.FundChangeRecord{
				Title:           getFundChangeTitle(typ),
				Desc:            comment,
				Time:            tm,
				Val:             varFreezeCai,
				Remain:          cai,
				Name:            "彩金",
				UserOrderId:     v.UserOrderId,
				VendorOrderId:   v.VendorOrderId,
				WithdrawApplyId: v.WithdrawApplyId,
				Direct:          getFundDirect(v),
			}
			result.List = append(result.List, record)
		}
	}

	statsArg := &dbproto.QueryFundHistoryStatsArg{AccountId: accountId, TimeStart: dbArg.GetTimeStart(), TimeEnd: dbArg.GetTimeEnd()}
	stats, err := srv.dbClient.QueryFundHistoryStats(context.Background(), statsArg)
	if err != nil {
		http.Error(w, "db：QueryFundHistoryStats!", http.StatusInternalServerError)
		return
	}
	result.Recharge = math.Abs(stats.GetRecharge())
	result.Withdraw = math.Abs(stats.GetWithdraw())
	result.Buycai = math.Abs(stats.GetBuycai())
	result.Win = math.Abs(stats.GetWin())

	log.Printf("%d ---> %+v\n", accountId, result)
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", result)
}
