package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/caojunxyz/mimi-server/auth"
)

func (srv *UcServer) HandleTest(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleTest")
	w.Write([]byte("I am uc server!"))
	// b := srv.sendSmsCode("18628339914")
	// w.Write([]byte(fmt.Sprintf("验证结果:%v", b)))
}

func (srv *UcServer) HandleDev(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleDev")
}

func (srv *UcServer) ServeHTTP() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	log.Println("ServeHTTP on port ", *httpPort)
	mux := http.NewServeMux()
	mux.HandleFunc("/user/sendcode", srv.HandleSendSmsCode)

	mux.HandleFunc("/user/test", srv.HandleTest)
	mux.HandleFunc("/user/regist", srv.HandleRegist)
	mux.HandleFunc("/user/login", srv.HandleLogin)
	mux.HandleFunc("/user/forgotpwd", srv.HandleForgotPwd)
	mux.HandleFunc("/user/setpwd", auth.Validate(srv.HandleSetPwd))
	mux.HandleFunc("/user/resetpwd", auth.Validate(srv.HandleResetPwd))

	mux.HandleFunc("/user/bindphone", auth.Validate(srv.HandleBindPhone))
	mux.HandleFunc("/user/changephone", auth.Validate(srv.HandleChangePhone))
	mux.HandleFunc("/user/bindweixin", auth.Validate(srv.HandleBindWeixin))
	mux.HandleFunc("/user/unbindweixin", auth.Validate(srv.HandleUnbindWeixin))
	mux.HandleFunc("/user/bindqq", auth.Validate(srv.HandleBindQQ))
	mux.HandleFunc("/user/unbindqq", auth.Validate(srv.HandleUnbindQQ))

	mux.HandleFunc("/user/seticon", auth.Validate(srv.HandleSetIcon))
	mux.HandleFunc("/user/setnickname", auth.Validate(srv.HandleSetNickname))
	mux.HandleFunc("/user/setsex", auth.Validate(srv.HandleSetSex))
	mux.HandleFunc("/user/setpaypwd", auth.Validate(srv.HandleSetPayPassword))
	mux.HandleFunc("/user/resetpaypwd", auth.Validate(srv.HandleResetPayPassword))
	mux.HandleFunc("/user/openpaypwd", auth.Validate(srv.HandleOpenPayPassword))
	mux.HandleFunc("/user/forgotpaypwd", auth.Validate(srv.HandleForgotPayPwd))
	mux.HandleFunc("/user/verifypaypwd", auth.Validate(srv.HandleVerifyPayPwd))
	mux.HandleFunc("/user/info", auth.Validate(srv.HandleGetUserInfo))
	mux.HandleFunc("/user/fundinfo", auth.Validate(srv.HandleGetFundInfo))
	mux.HandleFunc("/user/withdraw/info", auth.Validate(srv.HandleGetWithdrawInfo))
	mux.HandleFunc("/user/withdraw/progress", auth.Validate(srv.HandleGetWithdrawProgress))
	mux.HandleFunc("/user/withdraw", auth.Validate(srv.HandleWithdraw))
	mux.HandleFunc("/user/bankcard", auth.Validate(srv.HandleGetBankcard))
	mux.HandleFunc("/user/deletebankcard", auth.Validate(srv.HandleDeleteBankcard))
	mux.HandleFunc("/user/ticketinfo", auth.Validate(srv.HandleGetBuycaiTicketInfo))
	mux.HandleFunc("/user/buycaitickets", auth.Validate(srv.HandleValidTickets))
	mux.HandleFunc("/user/authrealname", auth.Validate(srv.HandleAuthRealname))
	mux.HandleFunc("/user/addbankcard", auth.Validate(srv.HandleAddBankcard))
	mux.HandleFunc("/user/fundhistory", auth.Validate(srv.HandleGetFundHistory))
	mux.HandleFunc("/user/cdkey/redeem", auth.Validate(srv.HandleRedeemCdkey))
	mux.HandleFunc("/user/invite", auth.Validate(srv.HandleInvite))
	mux.HandleFunc("/user/invite/info", auth.Validate(srv.HandleInviteInfo))
	mux.HandleFunc("/user/dailycheck", auth.Validate(srv.HandleDailyCheck))
	mux.HandleFunc("/user/credits/taskinfo", auth.Validate(srv.HandleCreditsTaskInfo))
	mux.HandleFunc("/user/phoneregistgift/open", srv.HandleOpenPhoneRegistGift)
	mux.HandleFunc("/user/phoneregistgift/draw", auth.Validate(srv.HandleReceivePhoneRegistGift))
	mux.HandleFunc("/user/receive/gift", srv.HandleReceiveGift)


	httpServer := http.Server{
		Addr:         fmt.Sprintf(":%d", *httpPort),
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	if err := httpServer.ListenAndServe(); err != nil {
		panic(err)
	}
}
