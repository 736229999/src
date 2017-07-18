package core

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/caojunxyz/mimi-admin/backend/auth"
	"github.com/caojunxyz/mimi-admin/dbadminagent/proto"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

type Service struct {
	dbClient   dbproto.DbAdminAgentClient
	dbInstance dbproto.DbAdminAgentClient
}

func NewServer() *Service {
	return &Service{}
}

func (srv *Service) ConnectDb() dbproto.DbAdminAgentClient {
	conn, err := grpc.Dial("127.0.0.1:10000", grpc.WithInsecure())
	if err != nil {
		log.Panic(err)
	}
	srv.dbClient = dbproto.NewDbAdminAgentClient(conn)
	return srv.dbClient
}

func (srv *Service) Db() dbproto.DbAdminAgentClient {

	if srv.dbInstance == nil {
		srv.dbInstance = NewServer().ConnectDb()
	}
	return srv.dbInstance
}

//返回json到页面的方法.
func (srv *Service) Json(msg interface{}, code int, c *gin.Context) {

	dist := make(map[string]interface{})
	dist["msg"] = msg

	//json.
	// jsonStr,_ := srv.ToJson(dist)

	//设置响应头，允许跨域.
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
	c.Writer.Header().Set("Access-Control-Max-Age", "3600")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization,DNT,User-Agent,Keep-Alive,Content-Type,accept,origin,X-Requested-With")

	c.JSON(code, dist)
}

func (srv *Service) MD5(str string) string {

	h := md5.New()
	h.Write([]byte(str))
	b := h.Sum(nil)
	str = hex.EncodeToString(b)
	return str
}

//转换成json.
func (srv *Service) ToJson(dist interface{}) (string, error) {

	b, err := json.Marshal(dist)
	if err != nil {
		log.Printf("%+v\n", err)
		return "", err
	}
	return string(b), nil
}

//记录日志.
func (srv *Service) Log(c *gin.Context, params interface{}, message string, operating int) {

	userId := srv.GetUserInfo(c, "id")

	path := c.Request.URL.Path

	dist := make(map[string]interface{})
	dist["method"] = c.Request.Method
	dist["ip"] = c.ClientIP()
	dist["params"] = params
	dist["message"] = message

	jsonStr, err := srv.ToJson(dist)
	if err != nil {
		log.Printf("%+v\n", err)
	} else {
		logInfo := &dbproto.Log{
			UserId:    int64(userId.(float64)),
			Path:      path,
			Operating: int64(operating),
			Message:   message,
			Params:    jsonStr,
		}

		_, err = srv.Db().InsertLog(context.Background(), logInfo)
		if err != nil {
			log.Printf("%+v\n", err)
		}
	}
}

//获取用户信息.
func (srv *Service) GetUserInfo(c *gin.Context, key string) interface{} {

	err, dist := auth.Validate(c)
	if err != nil {
		return nil
	}

	return dist[key]
}
