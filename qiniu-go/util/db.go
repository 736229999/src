package util
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"

	"github.com/xormplus/core"
)

/***********************************因为数据库这儿很容易联想到初始化数据，所以一些数据都包括微信相关配置都在这个文件里面进行初始化***************************/

var Engine *xorm.Engine
var ApiEngine *xorm.Engine

var Wxconfig *WxConfig
var JidiPhpApi string
var NewChatPhpApi string


func init() {
	//初始化了数据库
	if Engine == nil{
		var err error
		Engine, err = xorm.NewEngine("mysql", "root:root@(127.0.0.1:3306)/db_bill?charset=utf8")
		CheckError(err)
	}
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "sh_")
	Engine.SetTableMapper(tbMapper)
	Engine.ShowSQL(false)



	if ApiEngine == nil{
		var err error
		//ApiEngine, err = xorm.NewEngine("mysql", "root:root@(127.0.0.1:3306)/db_api?charset=utf8")
		ApiEngine, err = xorm.NewEngine("mysql", "root:caipiao88888@(10.9.4.28:3306)/imicha-api?charset=utf8")
		CheckError(err)
	}
	tMapper := core.NewPrefixMapper(core.SnakeMapper{}, "api_")
	ApiEngine.SetTableMapper(tMapper)
	ApiEngine.ShowSQL(false)


	//初始化微信的一些配置
	Wxconfig = &WxConfig{
		"wx3e847b7b25c374a0",
		"bfbe0d2e5cf8151afc0e95ac7a3da3e3",
	}
	JidiPhpApi = "http://wechat.jiditv.com/Public/index.php/"
	//JidiPhpApi = "http://demo.api.com/Public/index.php/"
	NewChatPhpApi = "http://www.jiditv.com/OpenAPI/"
}
