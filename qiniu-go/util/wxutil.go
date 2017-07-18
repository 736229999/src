package util

import (
	"github.com/henrylee2cn/faygo"
	"encoding/json"
	"bytes"
	"net/http"
	"strings"
	"log"
	"qiniu-go/model/store"
)

var WxConfigMap  = make(map[string]string)
//用于解析json微信第一步获取access_token还有openid这些
type AccessToken struct {
	Access_token string
	Expires_in  int
	Refresh_token string
	Openid string
	Scope string
	Unionid string
}
//用于解析微信js-sdk的配置信息
type WxJsConfig struct {
	AppId string
	NonceStr string
	Timestamp int
	Url string
	Signature string
	RawString string
}
//用于解析微信支付的一些信息
type WxPayConfig struct {
	AppId string
	NonceStr string
	Package string
	SignType string
	TimeStamp string
	PaySign string
}

/**
	微信授权登录，获取用户的基本信息
 */
func  WxLogin(ctx *faygo.Context) (*store.StoreUser,*store.StoreWx) {

	code:=ctx.R.URL.Query().Get("code")

	step1Url := "https://api.weixin.qq.com/sns/oauth2/access_token?appid="+Wxconfig.Appid+"&secret="+Wxconfig.Appsecret+"&code="+code+"&grant_type=authorization_code"
	step1String := CurlGetReturnString(step1Url)
	step1Info := &AccessToken{}
	err:=json.Unmarshal([]byte(step1String),&step1Info)

	CheckError(err)
	if step1Info.Openid == "" {
		http.Redirect(ctx.W,ctx.R,"https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx3e847b7b25c374a0&redirect_uri=http%3A%2F%2Fqiniu-go.jiditv.com%2Findex.html&response_type=code&scope=snsapi_base&state=STATE#wechat_redirect",302)
	}
	step2Url:= "https://api.weixin.qq.com/sns/userinfo?access_token="+step1Info.Access_token+"&openid="+step1Info.Openid+"&lang=zh_CN"
	wxUserInfo,_:= CurlGetWithNoParams(step2Url)
	if wxUserInfo["nickname"] != nil {
		//存一个用户基本信息
		storeUser := new(store.StoreUser)
		storeUser.NickName = wxUserInfo["nickname"].(string)
		storeUser.Avatar = wxUserInfo["headimgurl"].(string)
		storeUser.Sex = 1
		//存一个用户的微信信息
		storeWx := new(store.StoreWx)
		storeWx.OpenId = wxUserInfo["openid"].(string)
		storeWx.UnionId = wxUserInfo["unionid"].(string)
		log.Println(storeUser)
		log.Println(storeWx)

		return storeUser,storeWx

	}

	return nil,nil

}


/**
	微信 js-sdk 配置
 */
func WxJsSdkConfig(ctx *faygo.Context) interface{}{

	scheme := "http://"
	url := strings.Join([]string{scheme, ctx.R.Host, ctx.R.RequestURI}, "")

	//微信分享给朋友	js-sdk
	wxSdkUrl := NewChatPhpApi+"Extension/wxjssdk/sample.php?appId="+Wxconfig.Appid+"&appSecret="+Wxconfig.Appsecret+"&url="+url
	wxSdkInfo := CurlGetReturnString(wxSdkUrl)
	strByte := []byte(wxSdkInfo)
	//这个地方是因为在解码过程中出现了一些开始和结束符
	strByte = bytes.TrimPrefix(strByte, []byte("\xef\xbb\xbf"))
	wxJsInfo := &WxJsConfig{}
	err2:=json.Unmarshal(strByte,&wxJsInfo)
	CheckError(err2)
	return wxJsInfo
}

/**
	微信支付
 */
func WxPay(openid string,money string) interface{}{
	wxPayApiUrl := JidiPhpApi+"Index/WxPayApi?openid="+openid+"&money="+money
	wxPayString:=CurlGetReturnString(wxPayApiUrl)
	strByte := []byte(wxPayString)
	strByte = bytes.TrimPrefix(strByte, []byte("\xef\xbb\xbf"))


	wxPayInfo := &WxPayConfig{}
	json.Unmarshal(strByte,&wxPayInfo)

	return wxPayInfo
}

/**
	微信上传本地图片或拍照图片
	@return 图片路径
 */
func WxUploads(media_id string) string {
	//先获取access_token
	url1 := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + Wxconfig.Appid + "&secret=" + Wxconfig.Appsecret
	access_token, _ := CurlGetWithNoParams(url1)

	//url2 := "http://file.api.weixin.qq.com/cgi-bin/media/get?access_token=" + access_token["access_token"].(string) + "&media_id=" + ctx.URL().Query().Get("serverId")
	url2 := JidiPhpApi+"Util/analysisWxImg?access_token="+ access_token["access_token"].(string) +"&media_id="+media_id+"&foldername=用户身份证"
	info := CurlGetReturnString(url2)
	return info
}