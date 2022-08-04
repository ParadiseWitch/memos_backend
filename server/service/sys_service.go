package service

import (
	"memos/server/dto"

	"github.com/gin-gonic/gin"
)

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
