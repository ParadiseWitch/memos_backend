package service

import (
	"memos/server/db"
	"memos/server/dto"
	"memos/server/logger"
	"memos/server/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

const SYSCONF_TABLENAME = "sys_conf"

func GetSysConfById(c *gin.Context) {
	var s dto.SysConf
	id, ok := c.GetQuery("id")
	if !ok {
		logger.Errorf("GetSysConfById err, id: %s", id)
		c.JSON(http.StatusBadRequest, util.CommonFailRes("id is required"))
		return
	}
	if err := db.DB.Table(SYSCONF_TABLENAME).
		Where("id = ?", id).
		Find(&s).Error; err != nil {
		logger.Warnf("GetSysConfById err, id: %d", id)
		c.JSON(http.StatusOK, util.CommonFailRes("GetSysConfById err"))
		return
	}
	c.JSON(http.StatusOK, util.CommonSuccRes(s))
}

func UpdateSysConfById(c *gin.Context) {
	var s dto.SysConf
	err := c.BindJSON(&s)
	if err != nil {
		logger.Errorf("UpdateSysConfById err, err: %s", err)
		c.JSON(http.StatusBadRequest, util.CommonFailRes("invalid request"))
		return
	}
	if err := db.DB.Table(SYSCONF_TABLENAME).Save(&s).Error; err != nil {
		logger.Warnf("UpdateSysConfById err, old sysconf: %s", s)
		c.JSON(http.StatusOK, util.CommonFailRes("UpdateSysConfById err"))
	}
}

func AddSysConf(c *gin.Context) {
	var s dto.SysConf
	err := c.BindJSON(&s)
	if err != nil {
		logger.Errorf("AddSysConf err, err: %s", err)
		c.JSON(http.StatusBadRequest, util.CommonFailRes("invalid request"))
		return
	}
	if err := db.DB.Table(SYSCONF_TABLENAME).Create(&s).Error; err != nil {
		logger.Warnf("AddSysConf err, data submitted: %s", s)
		c.JSON(http.StatusOK, util.CommonFailRes("AddSysConf err"))
		return
	}
	c.JSON(http.StatusOK, util.CommonSuccRes(s))
}
