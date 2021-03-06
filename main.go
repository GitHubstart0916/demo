package main

import (
	"context"
	"fmt"
	"github.com/FREE-WE-1/backend/api"
	_ "github.com/FREE-WE-1/backend/docs"
	"github.com/FREE-WE-1/backend/global"
	"github.com/FREE-WE-1/backend/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
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

func initRedis() {
	global.RedisClient = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("RedisAddress"),
		Password: viper.GetString("RedisPassword"),
		DB:       viper.GetInt("RedisDatabase"),
	})
	status := global.RedisClient.Ping(context.Background())
	if status.Err() != nil {
		panic("Redis 无法连接：" + status.Err().Error())
	}
}

func main() {
	ReadConfig()
	ConnectDB()
	initRedis()
	//stringA := []string{"18375200@buaa.edu.cn"}
	//_, err := models.SendEmailValidate(stringA)
	//if err != nil {
	//	print(err.Error())
	//}

	global.Router = gin.Default()

	// Swagger configuration
	url := ginSwagger.URL("/swagger/doc.json") // The url pointing to API definition
	global.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, utils.TokenHeaderName)
	global.Router.Use(cors.New(corsConfig))
	global.Router.Use(utils.AuthMiddleware)

	global.Router.Static("/image", "./images")

	api.InitRoutes()

	http.ListenAndServe(":8114", global.Router)

	global.Router.Run()
}
