package api

import (
	"database/sql"
	"fmt"
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

type FinishedCreatMapRequest struct {
	MapName int `json:"mapname" binding:"required"`
}

func finished_creat_map(c *gin.Context) {
	var fcmap FinishedCreatMapRequest
	if err := c.ShouldBindJSON(&fcmap); err != nil {
		c.String(http.StatusBadRequest, "解析出错："+err.Error())
		return
	}
	var mapid int
	err := global.DatabaseConnection.Get(&mapid, "SELECT id FROM (SELECT * FROM Map WHERE )")
	var Max_id int
	err = global.DatabaseConnection.Get(&Max_id, "SELECT max(id) FROM Map")
	if err != nil {
		panic(err)
	}
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
	MapName string `json:"mapName" binding:"required"`
}

type OpenMapResponse struct {
	Url string `json:"url"`
}

func open_map(c *gin.Context) {
	//TODO 返回图片和节点信息
	//var omap OpenMapRequest
	//if err := c.ShouldBindJSON(&omap); err != nil {
	//	c.String(http.StatusBadRequest, "解析出错："+err.Error())
	//	return
	//}
	//hasMap := false
	//var MapData models.Map
	//err := global.DatabaseConnection.Get(&MapData, "SELECT * FROM Map WHERE id = ?", omap.MapId)
	//switch err {
	//case nil:
	//	hasMap = true
	//case sql.ErrNoRows:
	//	hasMap = false
	//default:
	//	panic(err)
	//}
	//if !hasMap {
	//	c.JSON(utils.FindNoMapErr, "地图未找到")
	//	return
	//}
	var req OpenMapRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.String(http.StatusBadRequest, "解析出错："+err.Error())
		return
	}

	var path string
	UserId := c.GetInt("UserId")
	err := global.DatabaseConnection.Get(&path, "SELECT path FROM Map WHERE name = ? and user_id = ?", req.MapName, UserId)

	if err != nil {
		panic(err)
	}
	//ret := "http://127.0.0.1:8114/iamge/" +
	//	c.JSON(http.StatusOK, gin.H{"message": "Files Uploaded Successfully"})
	c.JSON(http.StatusOK, OpenMapResponse{
		Url: path,
	})
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

type GetAllMapResponse struct {
	Images []string `json:"images"`
}

type ImageData struct {
	Name string
	Path string
}

func getAllMapEndPoint(c *gin.Context) {
	userId := c.GetInt("UserId")

	var images []string
	var allImageData []ImageData

	fmt.Println(userId)

	err := global.DatabaseConnection.Select(&allImageData, "SELECT name AS name, path AS path FROM Map WHERE user_id = ?", userId)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(allImageData); i++ {
		images = append(images, allImageData[i].Name)
	}

	fmt.Println(allImageData)
	fmt.Println(images)

	c.JSON(http.StatusOK, GetAllMapResponse{
		Images: images,
	})
}
