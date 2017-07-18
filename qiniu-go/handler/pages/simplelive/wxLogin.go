package simplelive

import (
	"github.com/henrylee2cn/faygo"
	"encoding/json"
	"qiniu-go/util"
	"github.com/henrylee2cn/faygo/errors"
	"log"
	"net/http"
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
		return errors.New("没有登录")
	}
	step2Url:= "https://api.weixin.qq.com/sns/userinfo?access_token="+step1Info.Access_token+"&openid="+step1Info.Openid+"&lang=zh_CN"
	wxUserInfo,_:= util.CurlGetWithNoParams(step2Url)
	log.Println(wxUserInfo)
	c := make(chan bool)
	go hasLoginByWx(c,ctx)

	c <- true

	log.Println(c)
	return nil
})

func hasLoginByWx(c chan bool,ctx *faygo.Context)  {
	log.Println("有channel")
	http.Redirect(ctx.W,ctx.R,"http://qiniu-go.jiditv.com/simpleLive/index.html",302)
	<- c


}
