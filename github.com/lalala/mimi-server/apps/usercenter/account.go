package main

import (
	"log"
	"net/http"

	apiproto "github.com/caojunxyz/mimi-api/proto"
	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"github.com/caojunxyz/mimi-server/auth"
	"github.com/caojunxyz/mimi-server/utils"
	"golang.org/x/net/context"
)

func (srv *UcServer) isPhoneRegisted(phone string) (bool, error) {
	ret, err := srv.dbClient.QueryPhoneUser(context.Background(), &dbproto.StringValue{Value: phone})
	if err != nil {
		return false, err
	}

	return (ret.GetAccountId() > 0), nil
}

func (srv *UcServer) HandleSendSmsCode(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleSendSmsCode")
	var msg apiproto.SmsRequest
	_, _, err := utils.ParseHttpRequest(w, r, &msg)
	if err != nil {
		return
	}
	log.Printf("%+v\n", msg)

	phone := msg.GetPhone()
	typ := msg.GetType()
	b, err := srv.isPhoneRegisted(phone)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	var errMsg string
	var code apiproto.RespCode = apiproto.RespCode_Success

	switch typ {
	case apiproto.SmsType_Regist:
		if b {
			code = apiproto.RespCode_Fail
			errMsg = "手机号已注册"
		}
	case apiproto.SmsType_ResetPwd, apiproto.SmsType_ForgotPayPwd:
		if !b {
			code = apiproto.RespCode_Fail
			errMsg = "手机号未注册"
		}
	case apiproto.SmsType_BindPhone:
		if b {
			code = apiproto.RespCode_Fail
			errMsg = "手机号已绑定"
		}
	}

	if code != apiproto.RespCode_Success {
		utils.WriteHttpResponse(w, r, code, errMsg, nil)
		return
	}

	b = srv.sendSmsCode(phone)
	if !b {
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "发送失败", nil)
	}
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", nil)
}

func (srv *UcServer) HandleRegist(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleRegist")
	var msg apiproto.RegistRequest
	_, ip, err := utils.ParseHttpRequest(w, r, &msg)
	if err != nil {
		return
	}

	// TODO: 校验手机号和密码
	phone := msg.GetPhone()
	code := msg.GetCode()
	pwd := msg.GetPassword()
	if !srv.verifySmsCode(phone, code) {
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "验证失败", nil)
		return
	}

	phoneUser, err := srv.dbClient.QueryPhoneUser(context.Background(), &dbproto.StringValue{Value: phone})
	if err != nil {
		http.Error(w, "错误：查询手机用户!", http.StatusInternalServerError)
		return
	}
	if phoneUser.GetAccountId() > 0 {
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "手机号已注册", nil)
		return
	}

	user := &dbproto.PhoneUser{Phone: phone, Password: pwd}
	createAccountArg := &dbproto.CreateAccountArg{
		UserType: dbproto.UserType_Phone,
		User:     &dbproto.CreateAccountArg_PhoneUser{PhoneUser: user},
		Ip:       ip.String(),
	}
	ret, err := srv.dbClient.CreateAccount(context.Background(), createAccountArg)
	if err != nil {
		http.Error(w, "错误：验证短信验证码!", http.StatusInternalServerError)
		return
	}
	accountId := ret.GetValue()
	auth.SetHeader(w, accountId, "")
	result := &apiproto.LoginReply{
		User:     srv.queryUserInfo(accountId),
		Fund:     srv.queryFundInfo(accountId),
		Bankcard: srv.queryBankcard(accountId),
	}
	utils.AddCredits(srv.dbClient, accountId, apiproto.CreditsTask_RegistAccount, phone)
	utils.AddCredits(srv.dbClient, accountId, apiproto.CreditsTask_BindPhone, phone)
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", result)
}

func (srv *UcServer) HandleLogin(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleLogin")
	var msg apiproto.LoginRequest
	_, ip, err := utils.ParseHttpRequest(w, r, &msg)
	if err != nil {
		return
	}

	log.Printf("login: %+v (%s)\n", msg, ip)
	accountId := int64(0)
	userType := msg.GetType()
	openid := msg.GetOpenid()
	if (userType == apiproto.UserType_QQ || userType == apiproto.UserType_Weixin) && openid == "" {
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "无效openid", nil)
		return
	}

	var createAccountArg *dbproto.CreateAccountArg
	switch userType {
	case apiproto.UserType_Phone:
		ret, err := srv.dbClient.QueryPhoneUser(context.Background(), &dbproto.StringValue{Value: msg.GetPhone()})
		if err != nil {
			http.Error(w, "错误：查询手机用户!", http.StatusInternalServerError)
			return
		}
		accountId = ret.GetAccountId()
		// TODO: md5, salt
		if accountId <= 0 || msg.GetPassword() != ret.GetPassword() {
			utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "用户名或密码错误", nil)
			return
		}
	case apiproto.UserType_QQ:
		ret, err := srv.dbClient.QueryQQUser(context.Background(), &dbproto.StringValue{Value: openid})
		if err != nil {
			http.Error(w, "错误：查询QQ用户!", http.StatusInternalServerError)
			return
		}
		accountId = ret.GetAccountId()
		if accountId == 0 {
			createAccountArg = &dbproto.CreateAccountArg{
				UserType: dbproto.UserType_QQ,
				User:     &dbproto.CreateAccountArg_QqUser{QqUser: &dbproto.QQUser{Openid: openid}},
				Ip:       ip.String(),
			}
		}
	case apiproto.UserType_Weixin:
		ret, err := srv.dbClient.QueryWeixinUser(context.Background(), &dbproto.StringValue{Value: openid})
		if err != nil {
			http.Error(w, "错误：查询微信用户!", http.StatusInternalServerError)
			return
		}
		accountId = ret.GetAccountId()
		if accountId == 0 {
			createAccountArg = &dbproto.CreateAccountArg{
				UserType: dbproto.UserType_Weixin,
				User:     &dbproto.CreateAccountArg_WxUser{WxUser: &dbproto.WeixinUser{Openid: openid}},
				Ip:       ip.String(),
			}
		}
	}

	if accountId == 0 && createAccountArg != nil {
		ret, err := srv.dbClient.CreateAccount(context.Background(), createAccountArg)
		if err != nil {
			http.Error(w, "错误：创建账户!", http.StatusInternalServerError)
			return
		}
		accountId = ret.GetValue()
		utils.AddCredits(srv.dbClient, accountId, apiproto.CreditsTask_RegistAccount, openid)
	}
	auth.SetHeader(w, accountId, "")
	result := &apiproto.LoginReply{
		User:     srv.queryUserInfo(accountId),
		Fund:     srv.queryFundInfo(accountId),
		Bankcard: srv.queryBankcard(accountId),
	}
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", result)
}

func (srv *UcServer) HandleForgotPwd(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleForgotPwd")
	var msg apiproto.ForgotPwdRequest
	_, ip, err := utils.ParseHttpRequest(w, r, &msg)
	if err != nil {
		return
	}

	phone := msg.GetPhone()
	pwd := msg.GetPassword()
	code := msg.GetCode()
	// TODO: 正则校验, md5+salt(password)

	if !srv.verifySmsCode(phone, code) {
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "验证失败", nil)
		return
	}

	_, err = srv.dbClient.SetPhonePassword(context.Background(), &dbproto.PhonePassword{
		Phone: phone, Password: pwd, Ip: ip.String(),
	})
	if err != nil {
		http.Error(w, "错误：设置密码!", http.StatusInternalServerError)
		return
	}
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", nil)
}

func (srv *UcServer) HandleSetPwd(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleSetPwd")
	var msg apiproto.SetPwdRequest
	accountId, ip, err := utils.ParseHttpRequest(w, r, &msg)
	if err != nil {
		return
	}

	// 绑定了手机，且没有设置密码的账号才能设置密码
	userInfo, err := srv.dbClient.QueryUserInfo(context.Background(), &dbproto.IntValue{Value: accountId})
	if err != nil {
		log.Println(err)
		http.Error(w, "错误：QueryUserInfo!", http.StatusInternalServerError)
		return
	}

	phoneUser := userInfo.GetPhone()
	if phoneUser == nil {
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "还未绑定手机", nil)
		return
	}

	if phoneUser.GetPassword() != "" {
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "已经设置了密码", nil)
		return
	}

	pwd := msg.GetPassword()
	phone := phoneUser.GetPhone()
	// TODO: 正则校验, md5+salt(password)

	_, err = srv.dbClient.SetPhonePassword(context.Background(), &dbproto.PhonePassword{
		Phone: phone, Password: pwd, Ip: ip.String(),
	})
	if err != nil {
		http.Error(w, "db：SetPhonePassword!", http.StatusInternalServerError)
		return
	}
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", srv.queryUserInfo(accountId))
}

func (srv *UcServer) HandleResetPwd(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleResetPwd")
	var msg apiproto.ResetPwdRequest
	accountId, ip, err := utils.ParseHttpRequest(w, r, &msg)
	if err != nil {
		return
	}

	// 绑定了手机，且旧密码验证通过
	userInfo, err := srv.dbClient.QueryUserInfo(context.Background(), &dbproto.IntValue{Value: accountId})
	if err != nil {
		http.Error(w, "错误：QueryUserInfo!", http.StatusInternalServerError)
		return
	}

	phoneUser := userInfo.GetPhone()
	if phoneUser == nil {
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "还未绑定手机", nil)
		return
	}

	pwd := msg.GetPassword()
	phone := phoneUser.GetPhone()
	//  TODO: md5
	if phoneUser.GetPassword() != pwd {
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "旧密码错误", nil)
		return
	}

	newPwd := msg.GetNewPassword()
	_, err = srv.dbClient.SetPhonePassword(context.Background(), &dbproto.PhonePassword{
		Phone: phone, Password: newPwd, Ip: ip.String(),
	})
	if err != nil {
		http.Error(w, "错误：重置密码!", http.StatusInternalServerError)
		return
	}
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", srv.queryUserInfo(accountId))
}

func (srv *UcServer) HandleBindPhone(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleBindPhone")
	var msg apiproto.BindPhoneRequest
	accountId, ip, err := utils.ParseHttpRequest(w, r, &msg)
	if err != nil {
		return
	}

	phone := msg.GetPhone()
	code := msg.GetCode()
	log.Println(accountId, phone)
	// 检查是否已绑定其它账号
	ret, err := srv.dbClient.QueryPhoneUser(context.Background(), &dbproto.StringValue{Value: phone})
	if err != nil {
		http.Error(w, "db：QueryPhoneUser!", http.StatusInternalServerError)
		return
	}

	log.Printf("%+v\n", ret)
	if ret.GetAccountId() > 0 {
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "手机号已绑定", nil)
		return
	}

	if !srv.verifySmsCode(phone, code) {
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "验证失败", nil)
		return
	}

	dbArg := &dbproto.SetUserArg{
		UserType: dbproto.UserType_Phone,
		SetType:  dbproto.AccountChangeType_Bind,
		Ip:       ip.String(),
		User:     &dbproto.SetUserArg_PhoneUser{&dbproto.PhoneUser{Phone: phone, AccountId: accountId}},
	}
	_, err = srv.dbClient.SetAccountUser(context.Background(), dbArg)
	if err != nil {
		http.Error(w, "db：SetAccountUser!", http.StatusInternalServerError)
		return
	}
	utils.AddCredits(srv.dbClient, accountId, apiproto.CreditsTask_BindPhone, phone)
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", srv.queryUserInfo(accountId))
}

func (srv *UcServer) HandleChangePhone(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleChangePhone")
	var msg apiproto.ChangePhoneRequest
	accountId, ip, err := utils.ParseHttpRequest(w, r, &msg)
	if err != nil {
		return
	}

	phone := msg.GetPhone()
	newPhone := msg.GetNewPhone()
	code := msg.GetCode()
	ret, err := srv.dbClient.QueryPhoneUser(context.Background(), &dbproto.StringValue{Value: phone})
	if err != nil {
		http.Error(w, "db：QueryPhoneUser!", http.StatusInternalServerError)
		return
	}
	if ret.GetAccountId() != accountId {
		log.Println("账号不一致:", accountId, ret.GetAccountId())
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "旧手机号未绑定", nil)
		return
	}

	ret, err = srv.dbClient.QueryPhoneUser(context.Background(), &dbproto.StringValue{Value: newPhone})
	if err != nil {
		http.Error(w, "db：QueryPhoneUser!", http.StatusInternalServerError)
		return
	}
	if ret.GetAccountId() > 0 {
		log.Println("新手机号已绑定:", accountId, phone, ret.GetAccountId(), newPhone)
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "新手机号已绑定", nil)
		return
	}

	if !srv.verifySmsCode(newPhone, code) {
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "验证失败", nil)
		return
	}

	_, err = srv.dbClient.ChangePhone(context.Background(), &dbproto.ChangePhoneArg{Phone: phone, NewPhone: newPhone, Ip: ip.String()})
	if err != nil {
		http.Error(w, "db：ChangePhone!", http.StatusInternalServerError)
		return
	}
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", srv.queryUserInfo(accountId))
}

func (srv *UcServer) HandleBindWeixin(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleBindWeixin")
	var msg apiproto.BindWeixinRequest
	accountId, ip, err := utils.ParseHttpRequest(w, r, &msg)
	if err != nil {
		return
	}

	log.Printf("%d --> %+v\n", accountId, msg)
	openid := msg.GetOpenid()
	if openid == "" {
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "无效参数", nil)
		return
	}
	// 检查是否已绑定其它账号
	ret, err := srv.dbClient.QueryWeixinUser(context.Background(), &dbproto.StringValue{Value: openid})
	if err != nil {
		http.Error(w, "db：QueryWeixinUser!", http.StatusInternalServerError)
		return
	}

	if ret.GetAccountId() > 0 {
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "微信已绑定其它账号", nil)
		return
	}

	dbArg := &dbproto.SetUserArg{
		UserType: dbproto.UserType_Weixin,
		SetType:  dbproto.AccountChangeType_Bind,
		Ip:       ip.String(),
		User:     &dbproto.SetUserArg_WxUser{&dbproto.WeixinUser{Openid: openid, AccountId: accountId}},
	}
	_, err = srv.dbClient.SetAccountUser(context.Background(), dbArg)
	if err != nil {
		http.Error(w, "db：SetAccountUser!", http.StatusInternalServerError)
		return
	}
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", srv.queryUserInfo(accountId))
}

func (srv *UcServer) HandleUnbindWeixin(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleUnbindWeixin")
	var msg apiproto.BindWeixinRequest
	accountId, ip, err := utils.ParseHttpRequest(w, r, &msg)
	if err != nil {
		return
	}

	ret, err := srv.dbClient.QueryUserInfo(context.Background(), &dbproto.IntValue{Value: accountId})
	if err != nil {
		log.Println(err)
		http.Error(w, "db：QueryUserInfo!", http.StatusInternalServerError)
		return
	}
	if ret.GetPhone() == nil && ret.GetQq() == nil {
		log.Println("唯一登陆方式不能解绑")
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "唯一登陆方式不能解绑", nil)
		return
	}

	log.Printf("%d --> %+v\n", accountId, msg)
	openid := msg.GetOpenid()
	if openid == "" {
		log.Println("无效Openid")
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "无效参数", nil)
		return
	}
	dbArg := &dbproto.SetUserArg{
		UserType: dbproto.UserType_Weixin,
		SetType:  dbproto.AccountChangeType_Unbind,
		Ip:       ip.String(),
		User:     &dbproto.SetUserArg_WxUser{&dbproto.WeixinUser{Openid: openid, AccountId: accountId}},
	}
	log.Printf("%+v\n", dbArg)
	_, err = srv.dbClient.SetAccountUser(context.Background(), dbArg)
	if err != nil {
		log.Println(err)
		http.Error(w, "db：SetAccountUser!", http.StatusInternalServerError)
		return
	}
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", srv.queryUserInfo(accountId))
}

func (srv *UcServer) HandleBindQQ(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleBindQQ")
	var msg apiproto.BindQQRequest
	accountId, ip, err := utils.ParseHttpRequest(w, r, &msg)
	if err != nil {
		return
	}

	log.Printf("%d --> %+v\n", accountId, msg)
	openid := msg.GetOpenid()
	if openid == "" {
		log.Println("无效Openid")
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "无效参数", nil)
		return
	}
	// 检查是否已绑定其它账号
	ret, err := srv.dbClient.QueryQQUser(context.Background(), &dbproto.StringValue{Value: openid})
	if err != nil {
		http.Error(w, "db：QueryQQUser!", http.StatusInternalServerError)
		return
	}

	if ret.GetAccountId() > 0 {
		log.Println(accountId, ret.GetAccountId())
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "QQ已绑定其它账号", nil)
		return
	}

	dbArg := &dbproto.SetUserArg{
		UserType: dbproto.UserType_QQ,
		SetType:  dbproto.AccountChangeType_Bind,
		Ip:       ip.String(),
		User:     &dbproto.SetUserArg_QqUser{&dbproto.QQUser{Openid: openid, AccountId: accountId}},
	}
	_, err = srv.dbClient.SetAccountUser(context.Background(), dbArg)
	if err != nil {
		http.Error(w, "db：SetAccountUser!", http.StatusInternalServerError)
		return
	}
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", srv.queryUserInfo(accountId))
}

func (srv *UcServer) HandleUnbindQQ(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleUnbindQQ")
	var msg apiproto.BindQQRequest
	accountId, ip, err := utils.ParseHttpRequest(w, r, &msg)
	if err != nil {
		return
	}

	ret, err := srv.dbClient.QueryUserInfo(context.Background(), &dbproto.IntValue{Value: accountId})
	if err != nil {
		log.Println(err)
		http.Error(w, "db：QueryUserInfo!", http.StatusInternalServerError)
		return
	}
	if ret.GetPhone() == nil && ret.GetWeixin() == nil {
		log.Println("唯一登陆方式不能解绑")
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "唯一登陆方式不能解绑", nil)
		return
	}

	log.Printf("%d --> %+v\n", accountId, msg)
	openid := msg.GetOpenid()
	if openid == "" {
		utils.WriteHttpResponse(w, r, apiproto.RespCode_Fail, "无效参数", nil)
		return
	}
	dbArg := &dbproto.SetUserArg{
		UserType: dbproto.UserType_QQ,
		SetType:  dbproto.AccountChangeType_Unbind,
		Ip:       ip.String(),
		User:     &dbproto.SetUserArg_QqUser{&dbproto.QQUser{Openid: openid, AccountId: accountId}},
	}
	_, err = srv.dbClient.SetAccountUser(context.Background(), dbArg)
	if err != nil {
		log.Println(err)
		http.Error(w, "db：SetAccountUser!", http.StatusInternalServerError)
		return
	}
	utils.WriteHttpResponse(w, r, apiproto.RespCode_Success, "", srv.queryUserInfo(accountId))
}
