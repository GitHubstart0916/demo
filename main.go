package main

import (
	"fmt"
	"github.com/FREE-WE-1/backend/api"
	"github.com/FREE-WE-1/backend/global"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
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
	//mysqlStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
	//	viper.GetString("DatabaseUser"),
	//	viper.GetString("DatabasePassword"),
	//	viper.GetString("DatabaseHost"),
	//	viper.GetInt("DatabasePort"),
	//	viper.GetString("DatabaseName"))
	mysqlStr := "root:lyw002mysql@tcp(101.43.145.90:3306)/demo"
	DB, err := sqlx.Open("mysql", mysqlStr)
	if err != nil {
		panic(err)
	}
	global.DatabaseConnection = DB
}

func main() {
	//ReadConfig()
	ConnectDB()

	global.Router = gin.Default()
	api.InitRoutes()

	global.Router.Run()
}
