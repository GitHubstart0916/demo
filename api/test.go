package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetUserIdResponse struct {
	Id int `json:"id"`
}

func getUserIdEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, GetUserIdResponse{
		Id: c.GetInt("UserId"),
	})
}
