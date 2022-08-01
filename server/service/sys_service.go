package service

import (
	"memos/server/config"
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
