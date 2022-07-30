package service

import (
  "github.com/gin-gonic/gin"
  "memos/server/config"
  "memos/server/logger"
  "net/http"
)

func SysInfo(c *gin.Context) {
  logger.Info("sysinforouter")
  c.JSON(http.StatusOK, gin.H{
    "msg": config.Conf,
  })
}
