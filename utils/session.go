package utils

import (
	"context"
	"encoding/json"
	"github.com/FREE-WE-1/backend/global"
	"github.com/FREE-WE-1/backend/models"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type RoleType int

const (
	Visitor RoleType = 0
	User             = 1
)

// GenerateToken TODO: 添加权限判断
func GenerateToken() string {
	return "token" + strconv.Itoa(rand.Int())
}

func GetSessionData(ctx context.Context, token string) []byte {

	r, err := global.RedisClient.Get(ctx, token).Bytes()
	if err != nil {
		log.Printf("Failed to read session data %s: %s", token, err.Error())
		return nil
	}
	return r
}

const TokenHeaderName = "Token"

func AuthMiddleware(c *gin.Context) {
	token := c.Request.Header.Get(TokenHeaderName)

	if token == "" {
		c.Set("Role", Visitor)
		c.Next()
		return
	}
	rawDat := GetSessionData(c, token)
	if rawDat == nil {
		c.String(http.StatusUnauthorized, "Token 无效，请注销后重新登录")
		c.Abort()
		return
	}

	c.Set("Token", token)
	var userData models.AuthUser
	err := json.Unmarshal(rawDat, &userData)
	if err != nil {
		c.String(http.StatusUnauthorized, "Token解析错误")
		c.Abort()
		return
	}

	c.Set("UserId", userData.Id)
	c.Set("Role", User)
	c.Set("UserData", userData)
	c.Set("UserName", userData.UserName)
	c.Next()

}

func AuthRequired(c *gin.Context) {
	r, exists := c.Get("Role")
	if exists && r != Visitor {
		c.Next()
	} else {
		c.String(http.StatusUnauthorized, "请登录后继续")
		c.Abort()
	}
}

func LoginSession(c context.Context, token string, data interface{}) {
	var jsonStr []byte
	jsonStr, err := json.Marshal(data)
	if err != nil {
		panic(err) // never happens
	}
	err = global.RedisClient.Set(c, token, jsonStr, 0).Err()
	if err != nil {
		panic(err)
	}
}

func LogoutSession(c context.Context, token string) error {
	err := global.RedisClient.Del(c, token).Err()
	return err
}
