package main

import (
	"log"
	"strings"
	"time"

	apiproto "github.com/caojunxyz/mimi-api/proto"
	"github.com/caojunxyz/mimi-server/proto"
	"golang.org/x/net/context"
)

var banknameTable = map[string]apiproto.Bankname{
	"工商": apiproto.Bankname_IcbcBank,
	"光大": apiproto.Bankname_CebBank,
	"广发": apiproto.Bankname_CgbChinaBank,
	"华夏": apiproto.Bankname_HxbBank,
	"建设": apiproto.Bankname_CcbBank,
	"交通": apiproto.Bankname_CommBank,
	"民生": apiproto.Bankname_CmbcBank,
	"农业": apiproto.Bankname_AbcChinaBank,
	"平安": apiproto.Bankname_PinganBank,
	"浦发": apiproto.Bankname_SpdBank,
	"兴业": apiproto.Bankname_CibBank,
	"邮政": apiproto.Bankname_PsbcBank,
	"招商": apiproto.Bankname_CmbChinaBank,
	"中国": apiproto.Bankname_BocBank,
	"中信": apiproto.Bankname_CityBank,
}

func convertBankname(name string) apiproto.Bankname {
	for k, v := range banknameTable {
		if strings.Contains(name, k) {
			return v
		}
	}
	log.Panicf("无效银行: %s", name)
	return apiproto.Bankname_UnknownBank
}

var reverseBanknameTable = map[apiproto.Bankname]string{
	apiproto.Bankname_IcbcBank:     "工商",
	apiproto.Bankname_CebBank:      "光大",
	apiproto.Bankname_CgbChinaBank: "广发",
	apiproto.Bankname_HxbBank:      "华夏",
	apiproto.Bankname_CcbBank:      "建设",
	apiproto.Bankname_CommBank:     "交通",
	apiproto.Bankname_CmbcBank:     "民生",
	apiproto.Bankname_AbcChinaBank: "农业",
	apiproto.Bankname_PinganBank:   "平安",
	apiproto.Bankname_SpdBank:      "浦发",
	apiproto.Bankname_CibBank:      "兴业",
	apiproto.Bankname_PsbcBank:     "邮政",
	apiproto.Bankname_CmbChinaBank: "招商",
	apiproto.Bankname_BocBank:      "中国",
	apiproto.Bankname_CityBank:     "中信",
}

func convertReverseBankname(name apiproto.Bankname) string {
	return reverseBanknameTable[name]
}

func (srv *UcServer) sendSmsCode(phone string) bool {
	arg := &proto.SmsRequest{PhoneList: []string{phone}, CodeValidDur: int64(time.Minute * 5)}
	result, _ := srv.thirdapiClient.SendSmsCode(context.Background(), arg)
	return result.GetValue()
}

func (srv *UcServer) verifySmsCode(phone, code string) bool {
	// 校验短信验证码
	arg := &proto.SmsRequest{PhoneList: []string{phone}, Code: code}
	result, _ := srv.thirdapiClient.VerifySmsCode(context.Background(), arg)
	return result.GetValue()
}

func (srv *UcServer) verifyIdcard(idno, realname string) *proto.Idcard {
	arg := &proto.InfoVerifyRequest{IdcardNo: idno, Realname: realname}
	result, _ := srv.thirdapiClient.VerifyIdcard(context.Background(), arg)
	return result
}

func (srv *UcServer) verifyBankcard(idno, realname, cardno, phone string) *proto.Bankcard {
	arg := &proto.InfoVerifyRequest{IdcardNo: idno, Realname: realname, BankcardNo: cardno, Phone: phone}
	result, _ := srv.thirdapiClient.VerifyBankcard(context.Background(), arg)
	return result
}
