package middleware

import "github.com/henrylee2cn/faygo"
import (
	"github.com/henrylee2cn/faygo/errors"
	"net/http"
	"log"
	"strconv"
	"qiniu-go/model/store"
)

/**************************如果用户没有进行登录，比如在中间某个页面直接分享链接，然后进来，则跳转到首页进行登录*************************/
/**
	通过查看store里面有没有StoreUser来看判断是否登录
 */
var HasLogin = faygo.HandlerFunc(func(ctx *faygo.Context) error {

	if ctx.URL().Query().Get("uid") == "" {
		log.Println("没有传入uid")
		http.Redirect(ctx.W,ctx.R,"https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx3e847b7b25c374a0&redirect_uri=http%3A%2F%2Fqiniu-go.jiditv.com%2Findex.html&response_type=code&scope=snsapi_userinfo&state=STATE#wechat_redirect",302)
		return errors.New("跳转到首页登录")
	}
	uidStr := ctx.URL().Query().Get("uid")
	uid,_ := strconv.Atoi(uidStr)
	log.Println(uid)
	if store.StoreUserMap[uid] == nil {
		//如果没有说明不是从主页注册进来的
		log.Println("没有传入uid")
		http.Redirect(ctx.W,ctx.R,"https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx3e847b7b25c374a0&redirect_uri=http%3A%2F%2Fqiniu-go.jiditv.com%2Findex.html&response_type=code&scope=snsapi_userinfo&state=STATE#wechat_redirect",302)
		return errors.New("跳转到首页登录")
	}
	return nil
})
