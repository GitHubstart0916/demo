package api

import (
	"database/sql"
	"github.com/FREE-WE-1/backend/global"
	"github.com/FREE-WE-1/backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
	"strconv"
)

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
		c.JSON(http.StatusBadRequest, LoginResponse{})
		return
	}
	//TODO:获取Nodes和文件
}

type ModifyMapRequest struct {
	MapId   string `json:"mapid" binding:"required"`
	MapName string `json:"mapname" binding:"required"`
}

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
		c.JSON(http.StatusBadRequest, LoginResponse{})
		return
	}
	_, err = global.DatabaseConnection.Exec("UPDATE Map SET name = ? WHERE id = ?", mmap.MapName, mmap.MapId)
	switch err {
	case nil:
		c.JSON(http.StatusOK, RegisterResponse{
			Code: 0,
			Text: "修改成功",
		})
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
		c.JSON(http.StatusOK, RegisterResponse{
			Code: 0,
			Text: "删除成功",
		})
		mapPath := MapData.Path
		mapId := MapData.Id
		_ = exec.Command("rm", "-f", mapPath+"/"+strconv.Itoa(mapId)+".pgm")
		_ = exec.Command("rm", "-f", mapPath+"/"+strconv.Itoa(mapId)+".yaml")
	default:
		panic(err)
	}
}
