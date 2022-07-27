package service

import (
	"memos/server/db"
	"memos/server/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserInfo(c *gin.Context) {
	logger.Error("userinforouter")
	db.InitDB()
	c.JSON(http.StatusOK, gin.H{
		"name": "Maid",
		"age":  14,
	})
}
