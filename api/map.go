package api

import (
	"database/sql"
	"github.com/FREE-WE-1/backend/global"
	"github.com/FREE-WE-1/backend/models"
	"github.com/FREE-WE-1/backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
	"strconv"
)

type StandardResponse struct {
	Code  int    `json:"code"`
	Token string `json:"token"`
}

// ShowAccount godoc
// @Summary 建立地图
// @Description 建立地图
// @ID create-map
// @Accept  json
// @Produce json
// @Success 200 {string} string "服务器错误信息"
// @Failure default {string} string "服务器错误信息"
// @Router /map/create-map [post]
// @Security ApiKeyAuth
func create_map(c *gin.Context) {
	//todo redis
	//调用ROS api
	user_id := c.GetInt("UserId")

	var Max_id int
	err := global.DatabaseConnection.Get(&Max_id, "SELECT max(id) FROM Map")

	if err != nil {
		panic(err)
	}

	global.DatabaseConnection.Exec("INSERT INTO Map(path, user_id) VALUES (?, ?)", "/map/"+strconv.Itoa(Max_id+1), user_id)
	//这里似乎不会出错

	c.String(http.StatusOK, "create map successfully")
}

type GetMapdataRequest struct {
	MapId int `json:"mapid" binding:"required"`
}

type GetMapdataResponse struct {
	//CreateAt string `json:"createat" `
	//UpdateAt string `json:"updateat" `
	Path  string `json:"path" binding:"required"`
	Count string `json:"count" binding:"required"`
}

// ShowAccount godoc
// @Summary 查询地图状态
// @Description 查询地图状态
// @ID get-map-data
// @Accept  json
// @Produce json
// @Param  responseInfo body GetMapdataRequest true "DO"
// @Success 200 {Object} GetMapdataResponse
// @Failure default {Object} GetMapdataResponse
// @Router /map/get-map-data [post]
// @Security ApiKeyAuth
func get_map_data(c *gin.Context) {

	user_id := c.GetInt("UserId")

	var gmdata GetMapdataRequest
	if err := c.ShouldBindJSON(&gmdata); err != nil {
		c.String(http.StatusBadRequest, "解析出错："+err.Error())
		return
	}

	var MapData models.Map
	err := global.DatabaseConnection.Get(&MapData, "SELECT* FROM Map WHERE id = ?", gmdata.MapId)

	IsUser := false
	hasMap := false

	switch err {
	case nil:
		hasMap = true
		IsUser = (user_id == MapData.User_id)
	case sql.ErrNoRows:
		IsUser = false
	default:
		panic(err)
	}

	if !hasMap || !IsUser {
		c.JSON(http.StatusBadRequest, LoginResponse{})
		return
	}

	c.JSON(http.StatusOK, GetMapdataResponse{
		Path:  MapData.Path,
		Count: MapData.Count,
	})
}

type OpenMapRequest struct {
	MapId string `json:"mapid" binding:"required"`
}

type OpenMapResponse struct {
	NodeList [...]int32 `json:"nodelist" binding:"required"`
}

func open_map(c *gin.Context) {
	//TODO 返回图片和节点信息
	var omap OpenMapRequest
	if err := c.ShouldBindJSON(&omap); err != nil {
		c.String(http.StatusBadRequest, "解析出错："+err.Error())
		return
	}
	hasMap := false
	var MapData models.Map
	err := global.DatabaseConnection.Get(&MapData, "SELECT * FROM Map WHERE id = ?", omap.MapId)
	switch err {
	case nil:
		hasMap = true
	case sql.ErrNoRows:
		hasMap = false
	default:
		panic(err)
	}
	if !hasMap {
		c.JSON(utils.FindNoMapErr, "地图未找到")
		return
	}
	//TODO:获取Nodes和文件
}

type ModifyMapRequest struct {
	MapId   string `json:"mapid" binding:"required"`
	MapName string `json:"mapname" binding:"required"`
}

// ShowAccount godoc
// @Summary 修改地图信息
// @Description 修改地图信息
// @ID modify_map
// @Accept  json
// @Produce json
// @Param  responseInfo body ModifyMapRequest true "待修改信息"
// @Success 200 "修改地图成功"
// @Failure default {string} string "错误信息"
// @Router /map/modify-map [post]
// @Security ApiKeyAuth
func modify_map(c *gin.Context) {
	//TODO 修改地图信息
	var mmap ModifyMapRequest
	if err := c.ShouldBindJSON(&mmap); err != nil {
		c.String(http.StatusBadRequest, "解析出错："+err.Error())
		return
	}
	hasMap := false
	var MapData models.Map
	err := global.DatabaseConnection.Get(&MapData, "SELECT * FROM Map WHERE id = ?", mmap.MapId)
	switch err {
	case nil:
		hasMap = true
	case sql.ErrNoRows:
		hasMap = false
	default:
		panic(err)
	}
	if !hasMap {
		c.String(utils.FindNoMapErr, "地图未找到")
		return
	}
	_, err = global.DatabaseConnection.Exec("UPDATE Map SET name = ? WHERE id = ?", mmap.MapName, mmap.MapId)
	switch err {
	case nil:
		c.String(http.StatusOK, "修改地图成功")
	default:
		panic(err)
	}
}

type GetGoodsRequest struct {
	Goodname string `json:"goodname" binding:"required"`
}

func get_goods(c *gin.Context) {
	var req GetGoodsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.String(http.StatusBadRequest, "解析出错："+err.Error())
		return
	}

	//TODO ROS
}

type DeletMapRequest struct {
	MapId string `json:"mapid" binding:"required"`
}

// ShowAccount godoc
// @Summary 删除地图
// @Description 删除地图
// @ID delet_map
// @Accept  json
// @Produce json
// @Param  responseInfo body DeletMapRequest true "地图ID"
// @Success 200 "地图删除成功"
// @Failure default {string} string "服务器错误信息"
// @Router /map/delet-map [post]
// @Security ApiKeyAuth
func delet_map(c *gin.Context) {
	var dmap DeletMapRequest
	if err := c.ShouldBindJSON(&dmap); err != nil {
		c.String(http.StatusBadRequest, "解析出错："+err.Error())
		return
	}

	var MapData models.Map
	err := global.DatabaseConnection.Get(&MapData, "SELECT * FROM Map WHERE id = ?", dmap.MapId)
	hasMap := false
	switch err {
	case nil:
		hasMap = true
	case sql.ErrNoRows:
		hasMap = false
	default:
		panic(err)
	}
	if !hasMap {
		c.JSON(http.StatusBadRequest, LoginResponse{
			Code:  1,
			Token: "error",
		})
		return
	}
	_, err = global.DatabaseConnection.Exec("DELETE FROM Map WHERE id = ?", dmap.MapId)
	switch err {
	case nil:
		c.String(http.StatusOK, "地图删除成功")
		mapPath := MapData.Path
		mapId := MapData.Id
		_ = exec.Command("rm", "-f", mapPath+"/"+strconv.Itoa(mapId)+".pgm")
		_ = exec.Command("rm", "-f", mapPath+"/"+strconv.Itoa(mapId)+".yaml")
	default:
		panic(err)
	}
}
