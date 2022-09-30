package service

import (
	"memos/server/db"
	"memos/server/dto"
	"memos/server/logger"
	"memos/server/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

const USER_TABLENAME = "user"

func GetUserById(c *gin.Context) {
	var u dto.User
	id, ok := c.GetQuery("id")
	if !ok {
		logger.Errorf("GetUserById err, id: %s", id)
		c.JSON(http.StatusBadRequest, util.CommonFailRes("id is required"))
		return
	}
	if err := db.DB.Table(USER_TABLENAME).Where("id = ?", id).Find(&u).Error; err != nil {
		logger.Infof("The user does not exist. uid: %d", id)
		c.JSON(http.StatusOK, util.CommonFailRes("The user does not exist"))
		return
	}
	c.JSON(http.StatusOK, util.CommonSuccRes(u))

}

func Login(c *gin.Context) {
	var paramU dto.User
	// TODO encryption password
	if err := c.BindJSON(&paramU); err != nil {
		logger.Errorf("Login err, err: %s", err)
		c.JSON(http.StatusBadRequest, util.CommonFailRes("invalid request"))
		return
	}

	var resultU dto.User

	if err := db.DB.Table(USER_TABLENAME).
		Where("name = ?", paramU.Name).
		Find(&resultU).Error; err != nil {
		logger.Infof("The user not exist. name: %d", paramU.Name)
		c.JSON(http.StatusOK, util.CommonFailRes("The user not exist"))
		return
	}

	if resultU.Password != paramU.Password {
		logger.Infof("Password mistake, name: %s", paramU.Name)
		c.JSON(http.StatusOK, util.CommonFailRes("Password mistake"))
		return
	}
	c.JSON(http.StatusOK, util.CommonSuccRes(nil))
}

func UpdateUserById(c *gin.Context) {
	var u dto.User
	err := c.BindJSON(&u)
	if err != nil {
		logger.Errorf("UpdateUserById err, err: %s", err)
		c.JSON(http.StatusBadRequest, util.CommonFailRes("invalid request"))
		return
	}
	if err := db.DB.Table(USER_TABLENAME).Save(&u).Error; err != nil {
		logger.Warnf("UpdateUserById err, err: %s", err)
		c.JSON(http.StatusBadRequest, util.CommonFailRes("UpdateUserById err"))
	}
}

func AddUser(c *gin.Context) {
	var u dto.User
	err := c.BindJSON(&u)
	if err != nil {
		logger.Errorf("AddUser err, err: %s", err)
		c.JSON(http.StatusBadRequest, util.CommonFailRes("invalid request"))
		return
	}
	if err := db.DB.Table(USER_TABLENAME).Create(&u).Error; err != nil {
		logger.Errorf("AddUser err, err: %s", err)
		c.JSON(http.StatusBadRequest, util.CommonFailRes("AddUser err"))
		return
	}
	c.JSON(http.StatusOK, util.CommonSuccRes(u))
}
