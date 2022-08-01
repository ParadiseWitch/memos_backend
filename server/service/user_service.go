package service

import (
	"memos/server/db"
	"memos/server/dto"
	"memos/server/logger"
	"memos/server/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var userTable *gorm.DB

func getUserTable() *gorm.DB {
	if userTable == nil {
		userTable = db.DB.Table("user")
	}
	return userTable
}

func GetUserById(c *gin.Context) {
	var u dto.User
	id, ok := c.GetQuery("id")
	if !ok {
		logger.Errorf("GetUserById err, id: %s", id)
		c.JSON(http.StatusBadRequest,
			util.CommonFailRes(http.StatusBadRequest, "id is required"))
	}
	getUserTable().Where("id = ?", id).Find(&u)
	c.JSON(http.StatusOK, util.CommonSuccRes(u))
}

func UpdateUserById(c *gin.Context) {
	var u dto.User
	err := c.BindJSON(&u)
	if err != nil {
		logger.Errorf("UpdateUserById err, err: %s", err)
		c.JSON(http.StatusBadRequest,
			util.CommonFailRes(http.StatusBadRequest, "invalid request"))
	}
	// oldU := getUserTable().Where("id = ?", u.ID)
	getUserTable().Save(&u)
}

func AddUser(c *gin.Context) {
	var u dto.User
	err := c.BindJSON(&u)
	if err != nil {
		logger.Errorf("AddUser err, err: %s", err)
		c.JSON(http.StatusBadRequest,
			util.CommonFailRes(http.StatusBadRequest, "invalid request"))
	}
	getUserTable().Create(&u)
	c.JSON(http.StatusOK, util.CommonSuccRes(u))
}
