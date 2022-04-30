package global

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

var Router *gin.Engine
var DatabaseConnection *sqlx.DB
var RedisClient *redis.Client
