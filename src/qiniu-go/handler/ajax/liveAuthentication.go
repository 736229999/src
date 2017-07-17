package ajax

import (
	"github.com/henrylee2cn/faygo"
	"log"
	"qiniu-go/util"
	"qiniu-go/model/apiModel"
	"strconv"
	"time"
)

var UserLiveAuthMap  = make(map[string]*UserLiveAuthImg)

type UserLiveAuthImg struct {
	FrontImg string
	BackImg string
	HoldImg string
}
/**
	微信调用相机上传头像
*/
var LiveAuthUploadImg = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	if ctx.URL().Query().Get("serverId") == "" {
		return ctx.String(200, "请勿随便提交")
	}

	info := util.WxUploads(ctx.URL().Query().Get("serverId"))
	log.Println(info)

	uid := ctx.URL().Query().Get("uid")
	imgType := ctx.URL().Query().Get("imgType")

	if UserLiveAuthMap[uid] == nil {
		UserLiveAuthMap[uid] = new(UserLiveAuthImg)
	}

	switch imgType {
	case "front":
		UserLiveAuthMap[uid].FrontImg = info
	case "back":
		UserLiveAuthMap[uid].BackImg = info
	case "hold":
		UserLiveAuthMap[uid].HoldImg = info
	}

	return ctx.String(200, info)
})

/**
	将用户照片存到数据库
*/
var StoreIdCard = faygo.HandlerFunc(func(ctx *faygo.Context) error {

	liveAuth := new(apiModel.Liveauth)
	defer func() {
		liveAuth = nil
	}()
	uid := ctx.URL().Query().Get("uid")
	//UserLiveAuthMap[uid] = new(UserLiveAuthImg)
	if UserLiveAuthMap[uid] == nil {
		return ctx.String(200, "请传照片")
	}

	liveAuth.Uid,_ = strconv.Atoi(uid)
	liveAuth.Frontimg = UserLiveAuthMap[uid].FrontImg
	liveAuth.Backimg = UserLiveAuthMap[uid].BackImg
	liveAuth.Holdimg = UserLiveAuthMap[uid].HoldImg
	liveAuth.Addtime = time.Now().Unix()
	liveAuth.Lasttime = time.Now().Unix()

	_,err :=util.ApiEngine.Insert(liveAuth)
	util.CheckError(err)

	UserLiveAuthMap[uid] = nil	//解除引用释放内存
	return ctx.String(200, "已提交审核")
})