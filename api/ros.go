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
const Path = "/home/start_0916/SE/ROS/team08-proj/robot-ros"
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
// @Param  responseInfo body MakeDirRequest true "待添加信息"
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

// ShowAccount godoc
// @Summary 用户注册
// @Description 用户注册
// @ID user-register
// @Accept  json
// @Produce json
// @Param  responseInfo body RegisterRequest true "待添加信息"
// @Success 200 {object} RegisterResponse
// @Failure default {object} RegisterResponse
// @Router /user/register [get]
// @Security ApiKeyAuth
func openServeEndpoint(c *gin.Context) {
	//var req MakeDirRequest
	//if err := c.ShouldBindJSON(&req); err != nil {
	//	c.String(http.StatusBadRequest, "解析出错："+err.Error())
	//	return
	//}

	//TODO:执行命令
	var commands []string
	var ret string

	commands = append(commands, "cd /home/start_0916/SE/ROS/team08-proj/robot-ros")
	commands = append(commands, "source ./devel/setup.sh")
	commands = append(commands, "roslaunch wpb_home_apps innovation.launch")

	ret = "gnome-terminal -- bash -c "
	ret = ret + "\""
	for i := 0; i < len(commands); i++ {
		ret = ret + commands[i]
		if i != len(commands)-1 {
			ret = ret + ";"
		}
	}
	ret = ret + "\""
	fmt.Println(ret)
	cmd := exec.Command("bash", "-c", ret)
	//cmd.Stdin = in
	//cmd.Stdout = &out
	cmd.Run()

	c.String(http.StatusOK, "成功启动服务")
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
// @Router /user/register [get]
// @Security ApiKeyAuth
func moveEndpoint(c *gin.Context) {
	//var req MakeDirRequest
	//if err := c.ShouldBindJSON(&req); err != nil {
	//	c.String(http.StatusBadRequest, "解析出错："+err.Error())
	//	return
	//}

	//utils.RunCMD("rosrun wpr_simulation keyboard_vel_ctrl")

	//in := bytes.NewBuffer(nil)
	//var out bytes.Buffer
	//TODO:执行命令
	var commands []string
	var ret string

	commands = append(commands, "cd /home/start_0916/SE/ROS/team08-proj/robot-ros")
	commands = append(commands, "source ./devel/setup.sh")
	commands = append(commands, "rosrun wpr_simulation keyboard_vel_ctrl")

	ret = "gnome-terminal -- bash -c "
	ret = ret + "\""
	for i := 0; i < len(commands); i++ {
		ret = ret + commands[i]
		if i != len(commands)-1 {
			ret = ret + ";"
		}
	}
	ret = ret + "\""
	fmt.Println(ret)
	cmd := exec.Command("bash", "-c", ret)
	//cmd.Stdin = in
	//cmd.Stdout = &out
	cmd.Run()

	//go func() {
	//	//in.WriteString("cd /data/local/tmp\n")
	//	in.WriteString("ls\n")
	//	in.WriteString("exit\n")
	//}()
	//
	//if err := cmd.Run(); err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//fmt.Println(out.String())
	c.String(http.StatusOK, "成功启动服务")
}
