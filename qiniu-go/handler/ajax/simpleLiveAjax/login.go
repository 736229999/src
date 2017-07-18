package simpleLiveAjax

import (
	//"io/ioutil"
	"log"
	//"net/http"
	"github.com/henrylee2cn/faygo"

	"qiniu-go/util"
	"qiniu-go/model/apiModel"
	"encoding/json"
)

type AccessToken struct {
	Access_token string
	Expires_in  int
	Refresh_token string
	Openid string
	Scope string
	Unionid string
}
var WxLogin = faygo.HandlerFunc(func(ctx *faygo.Context) error {

	code:=ctx.R.URL.Query().Get("code")
	step1Url := "https://api.weixin.qq.com/sns/oauth2/access_token?appid="+util.Wxconfig.Appid+"&secret="+util.Wxconfig.Appsecret+"&code="+code+"&grant_type=authorization_code"
	step1String := util.CurlGetReturnString(step1Url)
	step1Info := &AccessToken{}
	err:=json.Unmarshal([]byte(step1String),&step1Info)

	util.CheckError(err)
	if step1Info.Openid == "" {
		return ctx.String(200,"false")
	}


	step2Url:= "https://api.weixin.qq.com/sns/userinfo?access_token="+step1Info.Access_token+"&openid="+step1Info.Openid+"&lang=zh_CN"
	wxUserInfo,_:= util.CurlGetWithNoParams(step2Url)

	user  := new(apiModel.User)
	defer func() {
		user = nil
	}()
	log.Println(wxUserInfo["unionid"].(string))
	flag,err := util.ApiEngine.Where("wx_unionid=?",wxUserInfo["unionid"].(string)).Cols("id").Get(user)

	util.CheckError(err)
	if flag {
		//就返回数据库中的该条数据
		log.Println(user)
		return ctx.JSON(200,user,true)

	}else {
		//插入一条新的数据再返回
		user.Name = wxUserInfo["nickname"].(string)
		user.Wx_unionid = wxUserInfo["unionid"].(string)
		_,er := util.ApiEngine.Insert(user)
		util.CheckError(er)
		return ctx.JSON(200,user,true)
	}

})



