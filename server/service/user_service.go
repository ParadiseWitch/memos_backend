package service

import (
	"github.com/gin-gonic/gin"
	"memos/server/database"
	"memos/server/logger"
	"net/http"
)

func UserInfo(c *gin.Context) {
	logger.Error("userinforouter")
	database.InitDB()
	c.JSON(http.StatusOK, gin.H{
		"name": "Maid",
		"age":  14,
	})
}
