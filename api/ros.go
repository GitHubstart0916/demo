package api

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/posener/complete/cmd"
	"net/http"
	"os/exec"
)

const ShellToUse = "bash"
const Path = "~/ROS/"
const Shell = "./mk.sh"

type MakeDirRequest struct {
	Path string `json:"path" binding:"required"`
	Name string `json:"name" binding:"required"`
}

// ShowAccount godoc
// @Summary 用户注册
// @Description 用户注册
// @ID user-register
// @Accept  json
// @Produce json
// @Param  responseInfo body RegisterRequest true "待添加信息"
// @Success 200 {object} RegisterResponse
// @Failure default {object} RegisterResponse
// @Router /user/register [post]
// @Security ApiKeyAuth
func makeDirEndpoint(c *gin.Context) {
	var req MakeDirRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.String(http.StatusBadRequest, "解析出错："+err.Error())
		return
	}

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	ret := "mkdir " + req.Name
	ret = "'" + ret + "'"

	command := Shell + " " + Path + " " + ret
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Run()

	fmt.Println(stdout.String())

	c.String(http.StatusOK, "成功创建文件夹")
}
