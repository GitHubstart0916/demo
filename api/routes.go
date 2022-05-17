package api

import (
	"github.com/FREE-WE-1/backend/global"
	"github.com/FREE-WE-1/backend/utils"
)

func InitRoutes() {

	user := global.Router.Group("/user")
	// TODO:使用中间件
	// user.Use()
	{
		user.POST("/login", loginEndpoint)
		user.POST("/logout", logoutEndpoint)
		user.POST("/register", registerEndpoint)
	}

	test := global.Router.Group("test")
	test.Use(utils.AuthRequired)
	{
		test.GET("/get-user-id", getUserIdEndpoint)
	}
}
