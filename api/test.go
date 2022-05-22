package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetUserIdResponse struct {
	Id int `json:"id"`
}

// ShowAccount godoc
// @Summary 获取用户id
// @Description 获取用户id
// @ID test-get-user-id
// @Success 200 {string} string "todo"
// @Failure default {string} string "todo"
// @Router /user/get-user-id [get]
// @Security ApiKeyAuth
func getUserIdEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, GetUserIdResponse{
		Id: c.GetInt("UserId"),
	})
}
