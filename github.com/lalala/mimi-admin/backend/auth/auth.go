package auth

import (
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

var KEY = []byte("Fa888666$$$a+099x")

const AUTH_HEADER_FIELD = "Authorization"

type Claims struct {
	UserInfo map[string]interface{}
	jwt.StandardClaims
}

func CreateToekn(userInfo map[string]interface{}) (string, error) {
	claims := Claims{
		userInfo,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "admin.cp.kxkr.com",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(KEY)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return signedToken, nil
}

//func Validate(c *gin.Context) bool {
//
//	authHeader := c.Request.Header.Get(AUTH_HEADER_FIELD)
//	if authHeader == "" {
//		log.Printf("%+v\n", "缺少token")
//		return false
//	}
//
//	claims := Claims{}
//	token, err := jwt.ParseWithClaims(authHeader, &claims, func(token *jwt.Token) (interface{}, error) {
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, fmt.Errorf("Unexpected signing method %v", token.Header["alg"])
//		}
//		return KEY, nil
//	})
//
//	if err != nil {
//		log.Println(err)
//		return false
//	}
//
//	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
//		log.Println("userInfo:", claims)
//		return true
//	} else {
//		return false
//	}
//}

func Validate(c *gin.Context) (error, map[string]interface{}) {

	authHeader := c.Request.Header.Get(AUTH_HEADER_FIELD)
	if authHeader == "" {
		log.Printf("%+v\n", "缺少token")
		return errors.New("缺少token"), nil
	}

	claims := Claims{}
	token, err := jwt.ParseWithClaims(authHeader, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", token.Header["alg"])
		}
		return KEY, nil
	})

	if err != nil {
		log.Println(err)
		return errors.New("获取失败"), nil
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return nil, claims.UserInfo
	} else {
		return errors.New("获取失败"), nil
	}
}
