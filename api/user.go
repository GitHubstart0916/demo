package api

import (
	"database/sql"
	"fmt"
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

type RegisterResponse struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

// ShowAccount godoc
// @Summary 用户注册
// @Description 用户注册
// @ID user-register
// @Accept  json
// @Produce json
// @Param  responseInfo body RegisterRequest true "待添加信息"
// @Success 200 {object} RegisterResponse
// @Failure default {object} RegisterResponse
// @Router /user/register [post]
// @Security ApiKeyAuth
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
			c.JSON(http.StatusBadRequest, RegisterResponse{
				Code: 1,
				Text: "用户名已存在",
			})
			return
		} else {
			panic(err)
		}
	}

	//c.String(http.StatusOK, "注册成功")
	c.JSON(http.StatusOK, RegisterResponse{
		Code: 0,
		Text: "注册成功",
	})
}

type LoginRequest struct {
	UserName string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Code  int    `json:"code"`
	Token string `json:"token"`
}

// ShowAccount godoc
// @Summary 用户登录
// @Description 用户登录
// @ID user-login
// @Accept  json
// @Produce json
// @Param  responseInfo body LoginRequest true "待添加信息"
// @Success 200 {object} LoginResponse
// @Failure default {object} LoginResponse
// @Router /user/login [post]
// @Security ApiKeyAuth
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

	fmt.Println(err)

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
		c.JSON(http.StatusBadRequest, LoginResponse{
			Code:  1,
			Token: "error",
		})
		return
	}

	token := utils.GenerateToken()
	//fmt.Println(UserData)
	//global.RedisClient.Set(c, token, UserData, 0)
	//fmt.Println(global.RedisClient.Get(c, token))
	utils.LoginSession(c, token, UserData)

	c.JSON(http.StatusOK, LoginResponse{
		Code:  0,
		Token: token,
	})
}

// ShowAccount godoc
// @Summary 用户登出
// @Description 用户登出
// @ID user-logout
// @Accept  json
// @Produce json
// @Success 200 {string} string "todo"
// @Failure default {string} string "todo"
// @Router /user/logout [post]
// @Security ApiKeyAuth
func logoutEndpoint(c *gin.Context) {
	//TODO:是否需要logout接口
	token := c.GetString("Token")
	utils.LogoutSession(c, token)
	c.String(http.StatusOK, "logout")

}

type StateResponse struct {
	UserId string `json:"userid" binding:"required"`
	State  int    `json:"state" binding:"required"`
}

func get_user_state(c *gin.Context) {
	c.String(http.StatusOK, "1:正常登录")
}

type ForgetPswRequest struct {
	UserId string `json:"userid" binding:"required"`
}

// ShowAccount godoc
// @Summary 忘记密码
// @Description 忘记密码
// @ID forget_password
// @Accept  json
// @Produce json
// @Param  responseInfo body LoginRequest true "待添加信息"
// @Success 200 "成功发送验证码"
// @Failure default {string} string "错误信息"
// @Router /user/forget-password [post]
// @Security ApiKeyAuth
func forget_password(c *gin.Context) {
	var fPsw ForgetPswRequest
	if err := c.ShouldBindJSON(&fPsw); err != nil {
		c.String(http.StatusBadRequest, "解析出错："+err.Error())
		return
	}
	hasUserId := false
	var UserData models.AuthUser
	err := global.DatabaseConnection.Get(&UserData, "SELECT * FROM User WHERE userName = ?", fPsw.UserId)
	switch err {
	case nil:
		hasUserId = true
	case sql.ErrNoRows:
		hasUserId = false
	default:
		panic(err)
	}
	if !hasUserId {
		c.String(utils.FindNoMapErr, "用户ID未找到")
		return
	}
	models.GetValidateCode([]string{UserData.Email.String})
	c.String(http.StatusOK, "成功发送验证码")
}

type UserInfo struct {
	Id       int    `json:"id"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

// ShowAccount godoc
// @Summary 获取用户所有信息
// @Description 获取用户所有信息
// @ID get-user-info
// @Accept  json
// @Produce json
// @Success 200 {object} UserInfo
// @Failure default {string} string "错误信息"
// @Router /user/get_user_info [post]
// @Security ApiKeyAuth
func get_user_info(c *gin.Context) {
	var userInfo UserInfo
	userData, exists := c.Get("UserData")
	if !exists {
		c.JSON(http.StatusBadRequest, "用户信息不存在")
	}
	userInfo.UserName = userData.(*models.AuthUser).UserName
	userInfo.Email = userData.(*models.AuthUser).Email.String
	userInfo.Id = userData.(*models.AuthUser).Id
	c.JSON(http.StatusOK, userInfo)
}
