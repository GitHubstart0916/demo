package main

import (
	"fmt"
	"github.com/FREE-WE-1/backend/api"
	_ "github.com/FREE-WE-1/backend/docs"
	"github.com/FREE-WE-1/backend/global"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ReadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func ConnectDB() {
	mysqlStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		viper.GetString("DatabaseUser"),
		viper.GetString("DatabasePassword"),
		viper.GetString("DatabaseHost"),
		viper.GetInt("DatabasePort"),
		viper.GetString("DatabaseName"))
	DB, err := sqlx.Open("mysql", mysqlStr)
	if err != nil {
		panic(err)
	}
	global.DatabaseConnection = DB
}

func main() {
	ReadConfig()
	ConnectDB()

	global.Router = gin.Default()

	// Swagger configuration
	url := ginSwagger.URL("/swagger/doc.json") // The url pointing to API definition
	global.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	api.InitRoutes()

	global.Router.Run()
}
