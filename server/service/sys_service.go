package service

import (
	"memos/server/config"
	"memos/server/dto"
	"memos/server/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SysInfo(c *gin.Context) {
	logger.Info("sysinforouter")
	c.JSON(http.StatusOK, gin.H{
		"msg": config.Conf,
	})
}

func GetSysConfById(c *gin.Context) {
	var s dto.SysConf
	GetById("sys_conf", s, c)
}

func UpdateSysConfById(c *gin.Context) {
	var s dto.SysConf
	UpdateById("sys_conf", s, c)
}

func AddSysConf(c *gin.Context) {
	var s dto.SysConf
	Add("sys_conf", s, c)
}
