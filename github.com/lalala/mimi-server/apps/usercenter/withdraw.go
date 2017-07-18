package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"net/http"

	apiproto "github.com/caojunxyz/mimi-api/proto"
	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/utils"
)

func (srv *UcServer) queryWithdrawInfo(accountId int64) *apiproto.WithdrawInfo {
	// statsArg := &dbproto.QueryFundHistoryStatsArg{AccountId: accountId, TimeStart: 0, TimeEnd: time.Now().Unix()}
	// stats, err := srv.dbClient.QueryFundHistoryStats(context.Background(), statsArg)
	// if err != nil {
	// 	log.Println(err)
	// 	return nil
	// }
	fund, err := srv.dbClient.QueryFund(context.Background(), &dbproto.IntValue{Value: accountId})
	if err != nil {
		log.Println(err)
		return nil
	}

	result := &apiproto.WithdrawInfo{TotalWin: fund.TotalWin, TotalWithdraw: fund.TotalWithdraw}
	result.CurWithdraw = fund.Balance
	v := fund.TotalWin - math.Abs(fund.TotalWithdraw)
	if fund.Balance > v {
		result.CurWithdraw = v
	}
	return result
}

func (srv *UcServer) queryWithdrawProgress(id int64) *apiproto.WithdrawProgress {
	apply, err := srv.dbClient.QueryWithdrawApply(context.Background(), &dbproto.IntValue{Value: id})
	if err != nil {
		return nil
	}

	resultStr := "失败"
	if apply.IsSuccess {
		resultStr = "成功"
	}
	stepList := []string{"1.提交申请", "2.平台审核", "3.银行处理"}
	result := &apiproto.WithdrawProgress{
		StepList:  stepList,
		CurStep:   apply.Step,
		IsSuccess: apply.IsSuccess,
		Desc:      fmt.Sprintf("%s%s", stepList[apply.Step], resultStr),
	}
	if apply.Step == 0 {
		result.DetailList = append(result.DetailList, &apiproto.WithdrawProgress_Detail{Key: "提现金额", Value: fmt.Sprintf("%.2f元", apply.Amount)})
		result.DetailList = append(result.DetailList, &apiproto.WithdrawProgress_Detail{
			Key:   "银行账户",
			Value: fmt.Sprintf("%s(尾号%s)", apply.InBankname, apply.InNo[len(apply.InNo)-4:]),
		})
		result.DetailList = append(result.DetailList, &apiproto.WithdrawProgress_Detail{Key: "真实姓名", Value: maskPrivacyRealname(apply.Realname)})
	} else {
		if apply.AuditComment != "" {
			result.DetailList = append(result.DetailList, &apiproto.WithdrawProgress_Detail{Key: "备注", Value: apply.AuditComment})
		}
	}
	return result
}

func (srv *UcServer) HandleGetWithdrawInfo(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleGetWithdrawInfo")
	accountId, _, err := utils.ParseHttpRequest(w, r, nil)
	if err != nil {
		return
	}

	info := srv.queryWithdrawInfo(accountId)
	if info == nil {
		log.Println("nil info")
		http.Error(w, "queryWithdrawInfo", http.StatusInternalServerError)
		return
	}
	log.Printf("%d -> %+f\n", accountId, info)
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", info)
}

func (srv *UcServer) HandleGetWithdrawProgress(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleGetWithdrawProgress")
	var msg apiproto.IntValue
	_, _, err := utils.ParseHttpRequest(w, r, &msg)
	if err != nil {
		return
	}

	progress := srv.queryWithdrawProgress(msg.GetValue())
	if progress == nil {
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "无效id", nil)
		return
	}
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", progress)
}

func (srv *UcServer) HandleWithdraw(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleWithdraw")
	var msg apiproto.WithdrawRequest
	accountId, _, err := utils.ParseHttpRequest(w, r, &msg)
	if err != nil {
		return
	}

	log.Printf("%d --> %+v\n", accountId, msg)
	if msg.Type != apiproto.WithdrawType_ToBankcard {
		log.Println(accountId, msg.Type)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "无效提现方式", nil)
		return
	}

	bankcard, err := srv.dbClient.QueryAccountBankcard(context.Background(), &dbproto.IntValue{Value: accountId})
	if err != nil {
		http.Error(w, "QueryAccountBankcard", http.StatusInternalServerError)
		return
	}
	if bankcard.Id == 0 {
		log.Println("未绑定银行卡", accountId)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "还未绑定银行卡", nil)
		return
	}

	if msg.Amount < 5 {
		log.Println(accountId, msg.Amount)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "提现金额最少5元", nil)
		return
	}

	withdrawInfo := srv.queryWithdrawInfo(accountId)
	if withdrawInfo == nil {
		http.Error(w, "queryWithdrawInfo", http.StatusInternalServerError)
		return
	}

	if withdrawInfo.CurWithdraw < msg.Amount {
		log.Printf("%d, %+v, %d\n", accountId, withdrawInfo, msg.Amount)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "可提现金额不足", nil)
		return
	}

	withdrawApply := &dbproto.WithdrawApply{
		AccountId:    accountId,
		Realname:     bankcard.Realname,
		InBankname:   convertReverseBankname(apiproto.Bankname(bankcard.Bankname)),
		InNo:         bankcard.Cardno,
		Amount:       msg.Amount,
		Step:         0,
		IsSuccess:    true,
		WithdrawType: 1,
	}
	progress, err := srv.dbClient.InsertWithdrawApply(context.Background(), withdrawApply)
	if err != nil {
		http.Error(w, "QueryAccountBankcard", http.StatusInternalServerError)
		return
	}
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", progress)
}
