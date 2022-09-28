package service

import (
	"memos/server/db"
	"memos/server/dto"
	"memos/server/logger"
	"memos/server/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

const TABLENAME = "sys_conf"

func GetSysConfById(c *gin.Context) {
	var s dto.SysConf
	id, ok := c.GetQuery("id")
	if !ok {
		logger.Errorf("GetUserById err, id: %s", id)
		c.JSON(http.StatusBadRequest,
			util.CommonFailRes(http.StatusBadRequest, "id is required"))
	}
	db.DB.Table(TABLENAME).Where("id = ?", id).Find(&s)
	c.JSON(http.StatusOK, util.CommonSuccRes(s))
}

func UpdateSysConfById(c *gin.Context) {
	var s dto.SysConf
	err := c.BindJSON(&s)
	if err != nil {
		logger.Errorf("UpdateUserById err, err: %s", err)
		c.JSON(http.StatusBadRequest,
			util.CommonFailRes(http.StatusBadRequest, "invalid request"))
	}
	db.DB.Table(TABLENAME).Save(&s)
}

func AddSysConf(c *gin.Context) {
	var s dto.SysConf
	err := c.BindJSON(&s)
	if err != nil {
		logger.Errorf("AddUser err, err: %s", err)
		c.JSON(http.StatusBadRequest,
			util.CommonFailRes(http.StatusBadRequest, "invalid request"))
	}
	db.DB.Table(TABLENAME).Create(&s)
	c.JSON(http.StatusOK, util.CommonSuccRes(s))
}
