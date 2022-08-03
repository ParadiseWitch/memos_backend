package service

import (
	"memos/server/db"
	"memos/server/logger"
	"memos/server/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetById(tableName string, module interface{}, c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		logger.Errorf("GetUserById err, id: %s", id)
		c.JSON(http.StatusBadRequest,
			util.CommonFailRes(http.StatusBadRequest, "id is required"))
	}
	db.DB.Table(tableName).Where("id = ?", id).Find(&module)
	c.JSON(http.StatusOK, util.CommonSuccRes(module))
}

func UpdateById(tableName string, module interface{}, c *gin.Context) {
	err := c.BindJSON(&module)
	if err != nil {
		logger.Errorf("UpdateUserById err, err: %s", err)
		c.JSON(http.StatusBadRequest,
			util.CommonFailRes(http.StatusBadRequest, "invalid request"))
	}
	db.DB.Table(tableName).Save(&module)
}

func Add(tableName string, module interface{}, c *gin.Context) {
	err := c.BindJSON(&module)
	if err != nil {
		logger.Errorf("AddUser err, err: %s", err)
		c.JSON(http.StatusBadRequest,
			util.CommonFailRes(http.StatusBadRequest, "invalid request"))
	}
	db.DB.Table(tableName).Create(&module)
	c.JSON(http.StatusOK, util.CommonSuccRes(module))
}
