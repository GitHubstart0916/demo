package api

import (
	"github.com/gin-gonic/gin"
)

type SystemResetRequest struct {
	UserId int `json:"userid"`
}

// ShowAccount godoc
// @Summary 恢复初始设置
// @Description 恢复初始设置
// @ID reset
// @Success 200 "恢复初始设置成功"
// @Failure default {string} string "错误信息"
// @Router /system/reset [get]
// @Security ApiKeyAuth
func reset(c *gin.Context) {

}
