package service

import (
	"github.com/gin-gonic/gin"
	"memos/server/logger"
	"net/http"
)

func UserInfo(c *gin.Context) {
	logger.Error("userinforouter")
	c.JSON(http.StatusOK, gin.H{
		"name": "Maid",
		"age":  14,
	})
}
