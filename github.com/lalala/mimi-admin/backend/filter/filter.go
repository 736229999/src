package filter

import (
	"github.com/caojunxyz/mimi-admin/backend/auth"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func FilterLogin(c *gin.Context) {

	log.Println("path:", c.Request.URL.Path, "method:", c.Request.Method)
	if c.Request.Method == "OPTIONS" {
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization,DNT,User-Agent,Keep-Alive,Content-Type,accept,origin,X-Requested-With")
		c.JSON(http.StatusOK, gin.H{"msg": nil})
		c.Abort()
		return
	}

	if err, _ := auth.Validate(c); err != nil {
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))

		c.JSON(http.StatusUnauthorized, gin.H{"msg": "登录失败"})
		c.Abort()
		return
	}

	c.Next()
}
