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
		user.POST("/forget-password", forget_password)
		//user.POST("/reset-password", reset_password)
		//user.GET("/get-user-state", utils.AuthMiddleware, get_user_state)
	}

	test := global.Router.Group("test")
	test.Use(utils.AuthRequired)
	{
		test.GET("/get-user-id", getUserIdEndpoint)
	}

	Map := global.Router.Group("/map")
	Map.Use(utils.AuthRequired)
	{
		Map.GET("/get-user-id", getUserIdEndpoint)
		Map.GET("/create-map", create_map)
		Map.GET("/get-map-data", get_map_data)
		Map.POST("/open-map", open_map)
		Map.POST("/modify-map", modify_map)
		Map.POST("/get-goods", get_goods)
		Map.GET("/delet-map", delet_map)
	}
	System := global.Router.Group("/system")
	System.Use(utils.AuthRequired)
	{
		//System.GET("/reset", reset)
		//System.GET("/update", update)
		//System.GET("/get-update-list", get_update_list)
		//System.GET("/update-part", update_part)
	}
}
