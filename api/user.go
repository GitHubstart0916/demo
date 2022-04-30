package api

import (
	"database/sql"
	"github.com/FREE-WE-1/backend/global"
	"github.com/FREE-WE-1/backend/models"
	"github.com/FREE-WE-1/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"net/http"
)

type RegisterRequest struct {
	UserName string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func registerEndpoint(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.String(http.StatusBadRequest, "解析出错："+err.Error())
		return
	}

	_, err := global.DatabaseConnection.Exec("INSERT INTO User(userName, password) VALUES (?, ?)", req.UserName, req.Password)

	if err != nil {
		errInfo, _ := err.(*mysql.MySQLError)
		if errInfo.Number == utils.MysqlDuplicateErr {
			c.String(http.StatusBadRequest, "用户名已存在")
			return
		} else {
			panic(err)
		}
	}

	c.String(http.StatusOK, "注册成功")
}

type LoginRequest struct {
	UserName string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func loginEndpoint(c *gin.Context) {
	// c.String(http.StatusAccepted, "login")
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.String(http.StatusBadRequest, "解析出错："+err.Error())
		return
	}

	//fmt.Println(req.UserName)
	//fmt.Println(req.Password)

	var UserData models.AuthUser
	err := global.DatabaseConnection.Get(&UserData, "SELECT * FROM User WHERE userName = ?", req.UserName)

	hasUser := false
	pwdTrue := false

	switch err {
	case nil:
		hasUser = true
		pwdTrue = utils.ComparePassword(req.Password, UserData.Password)
	case sql.ErrNoRows:
		hasUser = false
	default:
		panic(err)
	}

	if !hasUser || (hasUser && !pwdTrue) {
		c.String(http.StatusBadRequest, "用户名或密码错误")
		return
	}

	c.String(http.StatusAccepted, "登录成功")
}

func logoutEndpoint(c *gin.Context) {
	//TODO:是否需要logout接口
	c.String(http.StatusAccepted, "logout")
}
