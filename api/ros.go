package api

import (
	"bytes"
	"fmt"
	"github.com/FREE-WE-1/backend/global"
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

// CreatMapEndpoint ShowAccount godoc
// @Summary 用户注册
// @Description 用户注册
// @ID user-register
// @Accept  json
// @Produce json
// @Param  responseInfo body MakeDirRequest true "待添加信息"
// @Success 200 {object} RegisterResponse
// @Failure default {object} RegisterResponse
// @Router /user/register [get]
// @Security ApiKeyAuth
func CreatMapEndpoint(c *gin.Context) {
	//var req MakeDirRequest
	//if err := c.ShouldBindJSON(&req); err != nil {
	//	c.String(http.StatusBadRequest, "解析出错："+err.Error())
	//	return
	//}
	//UserId := c.GetInt("UserId")

	var stdout bytes.Buffer

	RunCMD("cd /home/start_0916/ROS/team08-proj/robot-ros",
		"source ./devel/setup.sh",
		"roslaunch robot_main run_gmapping.launch")

	RunCMD("cd /home/start_0916/ROS/team08-proj/robot-ros",
		"source ./devel/setup.sh",
		"rosrun wpr_simulation keyboard_vel_ctrl")

	fmt.Println(stdout.String())

	c.String(http.StatusOK, "成功创建文件夹")
}

type SaveMapRequest struct {
	Name string `json:"name" binding:"required"`
}

type SaveMapResponse struct {
	Name string `json:"name"`
}

// SaveMapEndpoint ShowAccount godoc
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
func SaveMapEndpoint(c *gin.Context) {
	var req SaveMapRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.String(http.StatusBadRequest, "解析出错："+err.Error())
		return
	}

	userId := c.GetInt("UserId")

	RunCMD("cd /home/start_0916/ROS/team08-proj/robot-ros",
		"source ./devel/setup.sh",
		"cd /home/start_0916/ROS/team08-proj/robot-ros/src/robot_main/maps",
		"mkdir "+req.Name,
		"rosrun map_server map_saver -f /home/start_0916/ROS/team08-proj/robot-ros/src/robot_main/maps/"+req.Name+"/map")

	//cmd := exec.Command("bash", "-c", "python3 Trans/trans.py "+req.Name)
	//cmd.Run()
	//
	//cmd = exec.Command("bash", "-c", "cp /home/start_0916/ROS/team08-proj/robot-ros/src/robot_main/maps/"+req.Name+"/"+req.Name+".jpg"+
	//	" /home/start_0916/SE/backend/demo/images/")
	//cmd.Run()

	RunCMD("cd /home/start_0916/ROS/team08-proj/robot-ros",
		"source ./devel/setup.sh",
		"rosnode kill -a")

	name := req.Name
	path := "http://127.0.0.1:8114/image/" + req.Name + ".jpg"
	_, err := global.DatabaseConnection.Exec(`insert into Map (user_id, name, path) values (?, ?, ?)`, userId, name, path)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, SaveMapResponse{
		Name: req.Name,
	})
}

func RunCMD(arg ...string) {
	var commands []string
	var ret string

	for i := 0; i < len(arg); i++ {
		commands = append(commands, arg[i])
	}

	//commands = append(commands, "cd /home/start_0916/SE/ROS/team08-proj/robot-ros")
	//commands = append(commands, "source ./devel/setup.sh")
	//commands = append(commands, "rosrun wpr_simulation keyboard_vel_ctrl")

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
}

type TransRequest struct {
	Name string `json:"name" binding:"required"`
}

// TransEndpoint SaveMapEndpoint ShowAccount godoc
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
func TransEndpoint(c *gin.Context) {

	var req TransRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.String(http.StatusBadRequest, "解析出错："+err.Error())
		return
	}

	cmd := exec.Command("bash", "-c", "python3 Trans/trans.py "+req.Name)
	cmd.Run()

	cmd = exec.Command("bash", "-c", "cp /home/start_0916/ROS/team08-proj/robot-ros/src/robot_main/maps/"+req.Name+"/"+req.Name+".jpg"+
		" /home/start_0916/SE/backend/demo/images/")
	cmd.Run()

	c.String(http.StatusOK, "OK")
}

type BeginMarkRequest struct {
	MapName string `json:"mapName" binding:"required"`
}

// BeginMarkEndpoint ShowAccount godoc
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
func BeginMarkEndpoint(c *gin.Context) {
	var req BeginMarkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.String(http.StatusBadRequest, "解析出错："+err.Error())
		return
	}

	RunCMD("cd /home/start_0916/ROS/team08-proj/robot-ros",
		"source ./devel/setup.sh",
		"roslaunch robot_main run_serving.launch map_yaml_path:=\""+req.MapName+"/map.yaml\" waypoints_xml_path:=\""+req.MapName+"/waypoints.xml\"")

	RunCMD("cd /home/start_0916/ROS/team08-proj/robot-ros",
		"source ./devel/setup.sh",
		"rosrun wpr_simulation keyboard_vel_ctrl")

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

type addNodeRequest struct {
	MapName  string `json:"mapName" binding:"required"`
	NodeName string `json:"nodeName" binding:"required"`
}

// AddNodeEndpoint ShowAccount godoc
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
func AddNodeEndpoint(c *gin.Context) {
	var req addNodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.String(http.StatusBadRequest, "解析出错："+err.Error())
		return
	}

	UserId := c.GetInt("UserId")
	var mapId int

	err := global.DatabaseConnection.Get(&mapId, `select id AS mapId from Map where user_id = ? and name = ?`, UserId, req.MapName)

	if err != nil {
		panic(err)
	}

	_, err = global.DatabaseConnection.Exec(`insert into Node (map_id, node_name) VALUES (?, ?)`, mapId, req.NodeName)

	if err != nil {
		panic(err)
	}

	RunCMD("cd /home/start_0916/ROS/team08-proj/robot-ros",
		"source ./devel/setup.sh",
		"rosrun robot_main add_new_waypoint -n "+req.NodeName)

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

type FinishMarkRequest struct {
	MapName string `json:"MapName" binding:"required"`
}

// FinishMarkEndpoint ShowAccount godoc
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
func FinishMarkEndpoint(c *gin.Context) {
	var req FinishMarkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.String(http.StatusBadRequest, "解析出错："+err.Error())
		return
	}

	RunCMD("cd /home/start_0916/ROS/team08-proj/robot-ros",
		"source ./devel/setup.sh",
		"rosrun waterplus_map_tools wp_saver -f /home/start_0916/ROS/team08-proj/robot-ros/src/robot_main/maps/"+req.MapName+"/waypoints.xml")

	RunCMD("cd /home/start_0916/ROS/team08-proj/robot-ros",
		"source ./devel/setup.sh",
		"rosnode kill -a")

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

type GetAllNodeRequest struct {
	MapName string `json:"MapName" binding:"required"`
}

type nodeName struct {
	Name string
}

// GetAllNodeEndpoint ShowAccount godoc
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
func GetAllNodeEndpoint(c *gin.Context) {
	var req GetAllNodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.String(http.StatusBadRequest, "解析出错："+err.Error())
		return
	}

	UserId := c.GetInt("UserId")
	var mapId int

	err := global.DatabaseConnection.Get(&mapId, `select id AS mapId from Map where user_id = ? and name = ?`, UserId, req.MapName)

	if err != nil {
		panic(err)
	}

	var nodes []nodeName
	var ret []string

	err = global.DatabaseConnection.Select(&nodes, `select node_name AS name from Node where map_id = ?`, mapId)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(nodes); i++ {
		ret = append(ret, nodes[i].Name)
	}

	c.JSON(http.StatusOK, gin.H{
		"node_list": ret,
	})
}

type BeginServeRequest struct {
	MapName string `json:"mapName" binding:"required"`
}

// BeginServeEndpoint ShowAccount godoc
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
func BeginServeEndpoint(c *gin.Context) {
	var req BeginServeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.String(http.StatusBadRequest, "解析出错："+err.Error())
		return
	}

	RunCMD("cd /home/start_0916/ROS/team08-proj/robot-ros",
		"source ./devel/setup.sh",
		"roslaunch robot_main run_serving.launch map_yaml_path:=\""+req.MapName+"/map.yaml\" waypoints_xml_path:=\""+req.MapName+"/waypoints.xml\"")

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

type GotoRequest struct {
	Name string `json:"name" binding:"required"`
}

// GotoEndpoint ShowAccount godoc
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
func GotoEndpoint(c *gin.Context) {

	var req GotoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.String(http.StatusBadRequest, "解析出错："+err.Error())
		return
	}

	RunCMD("cd /home/start_0916/ROS/team08-proj/robot-ros",
		"source ./devel/setup.sh",
		"rosrun robot_main goto_waypoint -n "+req.Name)

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

type GetItRequest struct {
	Name string `json:"name" binding:"required"`
}

// GetItEndpoint ShowAccount godoc
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
func GetItEndpoint(c *gin.Context) {

	var req GetItRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.String(http.StatusBadRequest, "解析出错："+err.Error())
		return
	}

	RunCMD("cd /home/start_0916/ROS/team08-proj/robot-ros",
		"source ./devel/setup.sh",
		"rosrun robot_main fetch_object -n "+req.Name)

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

// FinishServeEndpoint ShowAccount godoc
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
func FinishServeEndpoint(c *gin.Context) {
	RunCMD("cd /home/start_0916/ROS/team08-proj/robot-ros",
		"source ./devel/setup.sh",
		"rosnode kill -a")

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
