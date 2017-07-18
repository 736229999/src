package main

import (
	"io"
	"log"
	"net/http"
	"strings"
	"time"
	"github.com/caojunxyz/gotu"
	apiproto "github.com/caojunxyz/mimi-api/proto"
	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/utils"
	"golang.org/x/net/context"
	"math/rand"
)

const DefaultIcon = "/assets/download/headicon/default-7fa78623d1149128800a820c8ab33091"

func (srv *UcServer) queryUserInfo(accountId int64) *apiproto.UserInfo {
	if accountId <= 0 {
		return nil
	}
	ret, err := srv.dbClient.QueryUserInfo(context.Background(), &dbproto.IntValue{Value: accountId})
	if err != nil {
		return nil
	}

	boolValue, err := srv.dbClient.QueryUserInviteStatus(context.Background(), &dbproto.IntValue{Value: accountId})
	if err != nil {
		log.Printf("%+v\n", err)
		return nil
	}

	if err != nil {
		log.Printf("%+v\n", err)
		return nil
	}
	userInfo := &apiproto.UserInfo{
		AccountId:        ret.GetAccountId(),
		Nickname:         ret.GetNickname(),
		Icon:             ret.GetIcon(),
		Idno:             maskPrivacyIdno(ret.GetIdno()),
		Realname:         maskPrivacyRealname(ret.GetRealname()),
		Exp:              ret.GetExp(),
		Level:            ret.GetLevel(),
		Sex:              apiproto.Sex(ret.GetSex()),
		PayPassword:      ret.GetPayPassword(),
		PayOpenPassword:  ret.GetPayOpenPassword(),
		InvitationCode:   ret.InvitationCode,
		UserInviteStatus: boolValue.GetValue(),
		IsDailyCheck:     gotu.IsSameDay(time.Now(), time.Unix(ret.GetDailyCheckTime(), 0)),
	}

	if userInfo.GetIcon() == "" {
		userInfo.Icon = DefaultIcon
	}

	if v := ret.GetPhone(); v != nil {
		userInfo.Bindphone = v.GetPhone()
		pwd := v.GetPassword()
		if len(pwd) > 0 {
			userInfo.LoginPassword = true
		}
	}
	if v := ret.GetQq(); v != nil {
		userInfo.Bindqq = v.GetOpenid()
	}
	if v := ret.GetWeixin(); v != nil {
		userInfo.Bindwx = v.GetOpenid()
	}

	log.Println("userInfo:", userInfo)
	return userInfo
}

func (srv *UcServer) HandleGetUserInfo(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleGetUserInfo")
	accountId, _, err := utils.ParseHttpRequest(w, r, nil)
	if err != nil {
		return
	}

	result := srv.queryUserInfo(accountId)
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", result)
}

func (srv *UcServer) HandleSetIcon(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleSetIcon")
	var msg apiproto.SetIconRequest
	accountId, _, err := utils.ParseHttpRequest(w, r, &msg)
	if err != nil {
		return
	}

	dbArg := &dbproto.UserInfoArg{AccountId: accountId, Icon: msg.GetValue()}
	_, err = srv.dbClient.SetUserIcon(context.Background(), dbArg)
	if err != nil {
		http.Error(w, "db：SetUserIcon!", http.StatusInternalServerError)
		return
	}
	result := srv.queryUserInfo(accountId)
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", result)
}

func (srv *UcServer) HandleSetNickname(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleSetNickname")
	var msg apiproto.SetNicknameRequest
	accountId, _, err := utils.ParseHttpRequest(w, r, &msg)
	if err != nil {
		return
	}

	dbArg := &dbproto.UserInfoArg{AccountId: accountId, Nickname: msg.GetValue()}
	_, err = srv.dbClient.SetUserNickname(context.Background(), dbArg)
	if err != nil {
		http.Error(w, "db：SetUserNickname!", http.StatusInternalServerError)
		return
	}
	result := srv.queryUserInfo(accountId)
	log.Printf("result: %+v\n", result)
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", result)
}

func (srv *UcServer) HandleSetSex(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleSetSex")
	var msg apiproto.SetSexRequest
	accountId, _, err := utils.ParseHttpRequest(w, r, &msg)
	if err != nil {
		return
	}

	dbArg := &dbproto.UserInfoArg{AccountId: accountId, Sex: int32(msg.GetValue())}
	_, err = srv.dbClient.SetUserSex(context.Background(), dbArg)
	if err != nil {
		log.Println(err)
		http.Error(w, "db：SetUserSex!", http.StatusInternalServerError)
		return
	}
	result := srv.queryUserInfo(accountId)
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", result)
}

func (srv *UcServer) HandleSetPayPassword(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleSetPayPassword")
	var msg apiproto.SetPwdRequest
	accountId, _, err := utils.ParseHttpRequest(w, r, &msg)
	if err != nil {
		return
	}

	pwd := msg.GetPassword()
	if len(pwd) < 6 {
		log.Println(pwd)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "无效密码", nil)
		return
	}

	paySettings, err := srv.dbClient.QueryAccountPaySettings(context.Background(), &dbproto.IntValue{Value: accountId})
	if err != nil {
		log.Println(err)
		http.Error(w, "db：QueryAccountPaySettings!", http.StatusInternalServerError)
		return
	}

	token := msg.GetToken()
	if paySettings.GetPassword() != "" && !srv.VerifySetPayPwdToken(accountId, token) {
		log.Println(accountId, token)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "无效设置", nil)
		return
	}

	dbArg := &dbproto.PayPasswordArg{AccountId: accountId, Password: pwd}
	_, err = srv.dbClient.SetAccountPayPassword(context.Background(), dbArg)
	if err != nil {
		log.Println(err)
		http.Error(w, "db：SetAccountPayPassword!", http.StatusInternalServerError)
		return
	}
	result := srv.queryUserInfo(accountId)
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", result)
}

func (srv *UcServer) HandleResetPayPassword(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleResetPayPassword")
	var msg apiproto.ResetPwdRequest
	accountId, _, err := utils.ParseHttpRequest(w, r, &msg)
	if err != nil {
		return
	}

	log.Printf("%d --> %+v\n", accountId, msg)
	pwd := msg.GetPassword()
	newPwd := msg.GetNewPassword()
	if len(pwd) < 6 || len(newPwd) < 6 {
		log.Println(pwd, newPwd)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "无效密码", nil)
		return
	}

	paySettings, err := srv.dbClient.QueryAccountPaySettings(context.Background(), &dbproto.IntValue{Value: accountId})
	if err != nil {
		log.Println(err)
		http.Error(w, "db：QueryAccountPaySettings!", http.StatusInternalServerError)
		return
	}

	if paySettings.GetPassword() != pwd {
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "旧密码输入错误", nil)
		return
	}

	dbArg := &dbproto.PayPasswordArg{AccountId: accountId, Password: newPwd}
	_, err = srv.dbClient.SetAccountPayPassword(context.Background(), dbArg)
	if err != nil {
		log.Println(err)
		http.Error(w, "db：SetAccountPayPassword!", http.StatusInternalServerError)
		return
	}
	result := srv.queryUserInfo(accountId)
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", result)
}

func (srv *UcServer) HandleOpenPayPassword(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleOpenPayPassword")
	var msg apiproto.SetBoolRequest
	accountId, _, err := utils.ParseHttpRequest(w, r, &msg)
	if err != nil {
		return
	}

	dbArg := &dbproto.PayPasswordArg{AccountId: accountId, Open: msg.GetValue()}
	_, err = srv.dbClient.OpenAccountPayPassword(context.Background(), dbArg)
	if err != nil {
		log.Println(err)
		http.Error(w, "db：OpenAccountPayPassword!", http.StatusInternalServerError)
		return
	}
	result := srv.queryUserInfo(accountId)
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", result)
}

func (srv *UcServer) HandleVerifyPayPwd(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleVerifyPayPwd")
	var msg apiproto.VerifyPayPwdRequest
	accountId, ip, err := utils.ParseHttpRequest(w, r, &msg)
	if err != nil {
		return
	}

	log.Printf("%d --> %+v\n", accountId, msg)
	paySettings, err := srv.dbClient.QueryAccountPaySettings(context.Background(), &dbproto.IntValue{Value: accountId})
	if err != nil {
		log.Println(err)
		http.Error(w, "db：QueryAccountPaySettings!", http.StatusInternalServerError)
		return
	}
	if paySettings.Password != msg.GetPassword() {
		log.Println(err)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "密码错误", nil)
		return
	}

	deviceId := r.Header.Get("deviceId")
	log.Println("verify 1")
	srv.RecordPayPwdVerify(accountId, ip.String(), deviceId, r.URL.Path, time.Now().Unix())
	log.Println("verify 2")
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", nil)
}

func (srv *UcServer) HandleForgotPayPwd(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleForgotPayPwd")
	var msg apiproto.ForgotPayPwdRequest
	accountId, _, err := utils.ParseHttpRequest(w, r, &msg)
	if err != nil {
		return
	}

	phone := msg.GetPhone()
	loginPwd := msg.GetLoginPassword()
	code := msg.GetCode()
	if !srv.verifySmsCode(phone, code) {
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "验证失败", nil)
		return
	}

	// TODO: 密码验证
	if len(loginPwd) < 6 {
		log.Println("密码无效:", accountId, loginPwd)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "密码无效!", nil)
		return
	}

	ret, err := srv.dbClient.QueryPhoneUser(context.Background(), &dbproto.StringValue{Value: phone})
	if err != nil {
		http.Error(w, "错误：查询手机用户!", http.StatusInternalServerError)
		return
	}

	// TODO: md5, salt
	if accountId != ret.GetAccountId() || loginPwd != ret.GetPassword() {
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "密码错误", nil)
		return
	}

	token := srv.GenSetPayPwdToken(accountId)
	result := &apiproto.StringValue{Value: token}
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", result)
}

func (srv *UcServer) HandleAuthRealname(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleAuthRealname")
	var msg apiproto.AuthRealnameRequest
	accountId, _, err := utils.ParseHttpRequest(w, r, &msg)
	if err != nil {
		return
	}

	realname := strings.TrimSpace(msg.GetRealname())
	idno := strings.TrimSpace(msg.GetIdno())

	// TODO: 正则表达式验证
	if len(realname) < 2 || len(idno) != 18 {
		log.Printf("realname: %v, idno: %v\n", realname, idno)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "无效参数", nil)
		return
	}

	idcard, err := srv.dbClient.QueryAccountIdcard(context.Background(), &dbproto.IntValue{Value: accountId})
	if err != nil {
		log.Println(err)
		http.Error(w, "db：QueryAccountIdcard!", http.StatusInternalServerError)
		return
	}

	if idcard.GetAccountId() == accountId {
		log.Println("身份证已认证:", idcard.GetAccountId(), accountId, realname, idno)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "已实名认证", nil)
		return
	}

	verifyResult := srv.verifyIdcard(idno, realname)
	if verifyResult == nil {
		log.Println("身份证核验失败:", accountId, idno, realname)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "核验失败", nil)
		return
	}

	card := &dbproto.AccountIdcard{
		AccountId: accountId,
		Idno:      idno,
		Realname:  realname,
	}
	_, err = srv.dbClient.InsertAccountIdcard(context.Background(), card)
	if err != nil {
		log.Println(err)
		http.Error(w, "db：InsertAccountIdcard!", http.StatusInternalServerError)
		return
	}

	utils.AddCredits(srv.dbClient, accountId, apiproto.CreditsTask_AuthRealname, idno)
	userInfo := srv.queryUserInfo(accountId)
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", userInfo)
}

func (srv *UcServer) HandleValidTickets(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleValidTickets")
	var msg apiproto.ValidTicketsRequest
	accountId, _, err := utils.ParseHttpRequest(w, r, &msg)
	if err != nil {
		return
	}

	stream, err := srv.dbClient.QueryValidBuycaiTickets(context.Background(), &dbproto.IntValue{Value: accountId})
	if err != nil {
		log.Println(err)
		http.Error(w, "db: QueryValidBuycaiTickets", http.StatusInternalServerError)
		return
	}

	lotteryId := msg.GetLotteryId()
	info := &apiproto.BuycaiTicketInfo{}
	for {
		v, err := stream.Recv()
		if err == nil {
			if !utils.IsBuycaiTicketCanUse(accountId, int32(lotteryId), msg.GetSumMoney(), v) {
				continue
			}
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
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", info)
}

type Gift struct {
	Credits int64                   `json:"credits"`
	Tickets []*dbproto.BuycaiTicket `json:"tickets"`
}

//完成输入正确cdkey的活动.
func (srv *UcServer) HandleRedeemCdkey(w http.ResponseWriter, r *http.Request) {

	//获取客户端传递过来的兑换码.
	var params apiproto.ExchangeRequest
	accountId, _, err := utils.ParseHttpRequest(w, r, &params)

	log.Println("获取到的account_id：", accountId)
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}

	log.Printf("%+v\n", params)
	//验证兑换码长度.
	if len(params.Code) != 8 {
		log.Println("兑换码长度不够", params.Code)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "兑换码错误", nil)
		return
	}

	//验证兑换码.
	batch, _, status := utils.VerifyCdkey(params.Code)
	if !status {
		log.Println("兑换码错误", params.Code)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "兑换码错误", nil)
		return
	}

	//根据批次来获取当前批次兑换码的数量.
	//判断兑换码是否超出当前序列.
	cdkeyBatch, err := srv.dbClient.QueryCdkeyBatchByBatch(context.Background(), &dbproto.IntValue{Value: int64(batch)})
	if err != nil {
		log.Printf("%+v\n", err)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "兑换码错误", nil)
		return
	}

	//判断活动是否开始.
	now := time.Now().Unix()
	if now < int64(cdkeyBatch.GetValidStart()) {
		log.Println("活动还没开始")
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "活动尚未开始", nil)
		return
	}

	//判断兑换码是否过期.
	if now > int64(cdkeyBatch.GetValidEnd()) {
		log.Println("兑换码已经过期")
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "兑换码已过期", nil)
		return
	}

	//验证当前兑换码是否已经被兑换.
	boolValue, err := srv.dbClient.QueryCdkeyStatus(context.Background(), &dbproto.StringValue{Value: params.Code})
	if err != nil {
		log.Printf("%+v\n", err)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "兑换失败", nil)
		return
	}
	if boolValue.GetValue() {
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "该验证码已经被兑换过了", nil)
		return
	}

	//根据批次来获取礼包模板.
	//giftTemplate, err := srv.dbClient.QueryGiftTemplateByBatch(context.Background(), &dbproto.IntValue{Value:int64(batch)})
	//if err != nil {
	//	log.Printf("%+v\n", err)
	//	utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "兑换失败", nil)
	//	return
	//}
	//
	//userGift := &dbproto.UserGiftPackage{Gift:new(dbproto.UserGiftPackageContent)}
	//credits := giftTemplate.GetContent().GetCredits()
	////验证当前礼包模板中赠送积分是否随机.
	//if credits.GetRandomCredits() {
	//	//积分随机.
	//	//获取积分随机的最大上限和最少下限.
	//	upper_limit := credits.GetUpperLimit()
	//	lower_limit := credits.GetLowerLimit()
	//	userGift.Gift.Credits = RandNum(lower_limit, upper_limit)
	//
	//} else {
	//	//积分不随机.
	//	userGift.Gift.Credits = credits.GetCredits()
	//}
	//
	////验证当前礼包模板中赠送的购彩券是否随机.
	//ticket := giftTemplate.GetContent().GetTickets()
	//if ticket.GetRandomTickets() {
	//	//购彩券随机.
	//	upper_limit := ticket.GetUpperLimit()
	//	lower_limit  := ticket.GetLowerLimit()
	//	randNum := RandNum(lower_limit, upper_limit)
	//	log.Println("upper_limit:", upper_limit, "lower_limit:", lower_limit, "随机数：", randNum)
	//	for k, v := range ticket.GetTickets() {
	//		if k + 1 <= int(randNum) {
	//			userGift.Gift.Tickets = append(userGift.Gift.Tickets , v)
	//		}
	//	}
	//
	//} else {
	//	//购彩券不随机.
	//	userGift.Gift.Tickets = ticket.GetTickets()
	//}
	//userGift.AccountId = accountId
	//userGift.TaskId = int64(proto2.TaskId_CDKEY_ID)
	//userGift.GiftTemplateId = cdkeyBatch.GetGiftTemplateId()
	//userGift.Cdkey = params.GetCode()
	//
	//_, err = srv.dbClient.InsertUserGiftPackage(context.Background(), userGift)
	//if err != nil {
	//	log.Printf("%+v\n", err)
	//	utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "兑换失败", nil)
	//	return
	//}

	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "完成任务", nil)
}

func RandNum(min, max int32) int32 {
	if min >= max || min== 0 || max== 0{
		return max
	}
	return rand.Int31n(max-min)+min
}

type InviteGiftResp struct {
	Credits int64                   `json:"credits"`
	Tickets []*dbproto.BuycaiTicket `json:"tickets"`
}
type InviteGiftPackageResp struct {
	Credits int64                    `json:"credits"`
	Tickets []*apiproto.BuycaiTicket `json:"tickets"`
}

//邀请用户.
func (srv *UcServer) HandleInvite(w http.ResponseWriter, r *http.Request) {

	var params apiproto.ExchangeRequest
	accountId, _, err := utils.ParseHttpRequest(w, r, &params)
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}

	//验证是否通过实名认证.
	boolValue, err := srv.dbClient.QueryUserAuthenticateByAccountId(context.Background(), &dbproto.IntValue{Value: accountId})
	if err != nil {
		log.Printf("%+v\n", err)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "未实名认证", nil)
		return
	}
	if !boolValue.GetValue() {
		log.Println("未实名认证")
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "未实名认证", nil)
		return
	}

	code := params.Code
	//验证邀请码的长度.
	log.Println("Invite:", code)
	if len(code) < 8 {
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "邀请码错误", nil)
		return
	}

	//验证邀请码是否存在.
	intValue, err := srv.dbClient.QueryUserInfoByInvitationCode(context.Background(), &dbproto.StringValue{Value: code})
	if err != nil {
		log.Printf("%+v\n", err)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "服务器异常", nil)
		return
	}
	//邀请者id 如果小于 0 则表示没有,该邀请码不存在.
	inviterId := intValue.GetValue()
	if inviterId < 1 {
		log.Println("邀请码不存在")
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "邀请码错误", nil)
		return
	}

	//拦截防止自己邀请自己.
	if inviterId == accountId {
		log.Println("不能自己邀请自己：", "inviterId:", inviterId, "accountId:", accountId)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "不能自己邀请自己", nil)
		return
	}

	//验证受邀者是否已经存在邀请关系.
	boolValue, err = srv.dbClient.QueryUserInviteRelation(context.Background(), &dbproto.IntValue{accountId})
	if err != nil {
		log.Printf("%+v\n", err)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "服务器异常", nil)
		return
	}
	if boolValue.GetValue() {
		log.Println("已经存在邀请关系，不能够重复被邀请")
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "不能够重复被邀请", nil)
		return
	}

	//拦截邀请的人是否为自己的被邀请人.
	inviteRelation := &dbproto.InviteRelationArg{
		Invitee: accountId,
		Inviter: inviterId,
	}
	boolValue, err = srv.dbClient.QueryInviteRelation(context.Background(), inviteRelation)
	if err != nil {
		log.Printf("%+v\n", err)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "服务器异常", nil)
		return
	} else {
		if !boolValue.GetValue() {
			log.Println("不能邀请自己的邀请人")
			utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "不能邀请自己的邀请人", nil)
			return
		}
	}

	//获取邀请任务的礼包模板.

	//获取邀请礼包.
	//stringValue, err := srv.dbClient.QueryInviteGift(context.Background(), &dbproto.Nil{})
	//if err != nil {
	//	log.Printf("获取邀请礼包%+v\n", err)
	//	utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "服务器异常", nil)
	//	return
	//}
	//
	////返回礼包.
	//giftPackageStr := stringValue.GetValue()
	//log.Println("json:", giftPackageStr)
	//
	//gift := &InviteGiftResp{}
	//err = json.Unmarshal([]byte(giftPackageStr), gift)
	//if err != nil {
	//	log.Printf("%+v\n", err)
	//	utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "服务器异常", nil)
	//	return
	//}
	//
	////添加礼包给邀请者.
	//gifts := &dbproto.Gift{
	//	Credits:    gift.Credits,
	//	TicketList: gift.Tickets,
	//	Inviter:    inviterId,
	//	Invitee:    accountId,
	//}
	//
	////添加礼包给被邀请者.
	//_, err = srv.dbClient.SetUserInviteRelation(context.Background(), gifts)
	//if err != nil {
	//	log.Printf("%+v\n", err)
	//	utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "服务器异常", nil)
	//	return
	//}
	//
	//log.Printf("%+v\n", gifts.GetTicketList())
	//
	//resp := &InviteGiftPackageResp{}
	//
	//err = json.Unmarshal([]byte(giftPackageStr), resp)
	//if err != nil {
	//	log.Printf("%+v\n", err)
	//	utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "服务器异常", nil)
	//	return
	//}
	//
	//userInfo := srv.queryUserInfo(accountId)
	//reply := &apiproto.ExchangeReply{
	//	TicketNum: int64(len(resp.Tickets)),
	//	Credits:   resp.Credits,
	//	UserInfo:  userInfo,
	//}

	//log.Printf("%+v\n", userInfo)
	//utils.AddCredits(srv.dbClient, gifts.Inviter, apiproto.CreditsTask_InviteFriend, fmt.Sprint(gifts.Invitee))
	//utils.AddCredits(srv.dbClient, gifts.Invitee, apiproto.CreditsTask_InviteFriend, fmt.Sprint(gifts.Inviter))
	//utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "受邀成功", reply)
}

//用户的邀请数据.
func (srv *UcServer) HandleInviteInfo(w http.ResponseWriter, r *http.Request) {

	accountId, _, err := utils.ParseHttpRequest(w, r, nil)

	userInviteInfo, err := srv.dbClient.QueryUserInviteInfo(context.Background(), &dbproto.IntValue{Value: accountId})
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}

	inviteInfo := &apiproto.UserInviteInfoReply{
		InviteNum:    userInviteInfo.GetInviteNum(),
		Credits:      userInviteInfo.GetCredits(),
		TicketsNum:   userInviteInfo.GetInviteNum(),
		TicketsMoney: userInviteInfo.GetTicketsMoney(),
	}

	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", inviteInfo)
}
