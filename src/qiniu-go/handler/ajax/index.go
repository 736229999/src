package ajax

import (
	//"io/ioutil"
	"log"
	//"net/http"
	"github.com/henrylee2cn/faygo"
	"qiniu-go/util"

	//"strings"
)

//直播主界面的ajax


/**
	微信调用相机上传头像
*/
var UploadImg = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	info := util.WxUploads(ctx.URL().Query().Get("serverId"))
	log.Println(info)
	return ctx.String(200, info)
})


var IsLogin = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	info := ctx.URL().Query().Get("id")
	log.Println(info)
	return ctx.String(200, info)
})
