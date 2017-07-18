package main

import "fmt"

type AlipayConfig struct {
	AppId      	  string
	Method	   	  string
	Charset	   	  string
	SignType  	  string
	Sign 	   	  string
	Timestamp    	  string
	Version	     	  string
	NotifyUrl         string
	PrivateKey        string
	PublicKey         string
	SellerId          string
}

func (cf AlipayConfig) GetNotifyUrl(gw string) string{
	return fmt.Sprintf("%s%s", gw, cf.NotifyUrl)
}

type WechatConfig struct {
	AppId      string `json:"app_id"`
	Package    string `json:"package"`
	MchId      string `json:"mch_id"`
	ApiKey     string `json:"api_key"`
	NotifyUrl  string
}

func (cf WechatConfig) GetNotifyUrl(gw string) string{
	return fmt.Sprintf("%s%s", gw, cf.NotifyUrl)
}

var wechatConfig = map[string]WechatConfig{
	"iOS": WechatConfig{
		ApiKey     : "fQC4WYiA48DOntcfX5s7eGHC6TVlIkZk",
		Package    : "Sign=WXPay",
		MchId      : "1456995702",
		AppId      : "wx0cdac55753fcd5a9",
		NotifyUrl  : "/recharge/wechat/notify",
	},
	"Android": WechatConfig{
		ApiKey     : "fQC4WYiA48DOntcfX5s7eGHC6TVlIkZk",
		Package    : "Sign=WXPay",
		MchId      : "1456995702",
		AppId      : "wx0cdac55753fcd5a9",
		NotifyUrl  : "/recharge/wechat/notify",
	},
}

var alipayConfig  = AlipayConfig{
	AppId       : "2017042106869800",
	Method      : "alipay.trade.app.pay",
	Charset     : "utf-8",
	SignType    : "RSA",
	Version     : "1.0",
	NotifyUrl   : "/recharge/alipay/notify",
	SellerId    : "2088621656612025",
	PrivateKey  : "MIIEowIBAAKCAQEAuau/6rEbYtNrdvetJpczizQqffWHmlskX9Tg8ppcImFcX4KigSF9llfjc7NpAWme4BDf8tP1+nRg/5y+3eyF3NrkEQh/b8Na8zfeKg6MK8LbKp/g5zetsYQp18iNCKhxGHESLHePsXYHSZdgfJVsy4ynO4NxmBnyyiutdTcj8PAfPPs4XDbKOJ7KKYvom7Wcn66IrYijknZebmfUkbUlg0iYhN2Zo/oRQifRr/h53gowEBo3mqQz4/Y+fqCc5maqZSWOUAox8djepwz63I8lo50Kem3VMViHraReQaMSXSMI+dl8cUPCF7uu2PTlJDkh3ogSSSBngBak6SqI/yGyBQIDAQABAoIBACPFNPowGMiXVrLa0J8IrkN5T15o+TEBsiZMSvIyTuHIxBUag9hA7YMTd5yIyggdoyAj4CFWOKOB3FYiipsSwE9mVF37tyF/D3ygHMsZdmPP7I1vs9KX5Xy3q6AI4TLz0KYW0puChUp5JLpG0a3u08D7XhSXJEFbbm0IsShN8TyvkcaO3o0Rm7HgbQu2Hqfj9Z9WKk7uRWbrlVQuaZm/jPGFAT8IgIpFtmbuIW3MHmlq/klJ16+NU/hm9NxoCDcFWh/oSaGxm9WwXAI5kgHGyYnCuA+Vd18cIK5ajKOpSIZdOgG5rI+ew2WK6olXfCTnCsinR+UvlvK/orq7OUC+/9kCgYEA8R3sPdhS+m7+KCrBUBDTrdmfF0L7HdO8DDmaZe8rEvSLb0Bq8OSnHVimxnVHVtjNU7xoPxW0OlPT+eOYe77SEYDHz59RbL7Fkb5H0xXfIdFoIbTf2IwD61kKgw3bqcna31/s3YqTo3kFKTcEwmnQ9PcERIvGsmh3ru+U1adkbPsCgYEAxSGuZ1DAZRuUPvdjLsjfeA3MvlHYafF6XePJTOSdUf+idoRH2IBDAKHW18/O20txqhQG8kljOy4vDnuLvwowiLok3F/MAEwCrjhr21mrck/pCqA3bmFgdLO/9EyZDf2xgdJMCC+8BWQRzR6C0ebTuUIlRIE5Lr04atYQFaO6LP8CgYByflL/ywkcAjiNuj+xVUwu8XeqBBRaRYC2DWRTow5BIf5UBbLCUNFKDh6jfm0xwCE/8edOjW/XgpbVwk1V9Xaq+Qvoey6fBLesTT6t7WTXGijoWIsyuMnjp5RmRt5X8ZINH+/KA1O3/G+G8qk/6B9hMyrDodvoO9MZLoruKV52oQKBgBTj4dh+/TR5vLcDflY54TyWscjFYZqwhVgnyHBGlclXvr6ye+6cgP0zKVJLak9g6lWss6O/VS2zUjdFrmR7TvNeToOv/y1U4L/XqM5g7UM65qZCnj0rQo57ce8QovORm9r3Dyma7WDg3uYOwqw9utgaA1sEGqwyTyU7NA8m/5a/AoGBAI68aldO1R8zWoyVnM8zbk6uOtkXcw1aF/IG0+K1wP/ghqgInGMMYbVMJiC9ZDJJ1x3MRaTMX6sHIlBaVoi2TQmgHC+vBiI8LnLnHuxDynX5Y5dzJDNEl+5k8yHYXvklDoA5BVL8dnzD5pS2J6Bxi1jZiMgRBw+IFBwFW1Z/DlHM",
	PublicKey   : "-----BEGIN PUBLIC KEY-----\n"+"MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDDI6d306Q8fIfCOaTXyiUeJHkrIvYISRcc73s3vF1ZT7XN8RNPwJxo8pWaJMmvyTn9N4HQ632qJBVHf8sxHi/fEsraprwCtzvzQETrNRwVxLO5jVmRGi60j8Ue1efIlzPXV9je9mkjzOmdssymZkh2QhUrCmZYI/FCEa3/cNMW0QIDAQAB"+"\n-----END PUBLIC KEY-----",
}

//获取结构体的方法.
func GetWechat(os string) *WechatConfig {
	cf, ok := wechatConfig[os]
	if ok {
		return &cf
	}
	return nil
}

func GetAlipay() *AlipayConfig {
	return &alipayConfig
}
