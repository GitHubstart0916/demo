package api

import "github.com/FREE-WE-1/backend/global"

func InitRoutes() {

	user := global.Router.Group("/user")
	// TODO:使用中间件
	// user.Use()
	{
		user.POST("/login", loginEndpoint)
		user.POST("/logout", logoutEndpoint)
		user.POST("/register", registerEndpoint)
	}
}
