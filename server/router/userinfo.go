package router

import (
	"memos/server/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	Routers = append(Routers, userInfoRouter)
}

func userInfoRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/userinfo")
	{
		r.GET("", home)
	}
}

func home(c *gin.Context) {
	logger.Error("userinforouter")
	c.JSON(http.StatusOK, gin.H{
		"name": "Maid",
		"age":  14,
	})
}
