package router

import (
	"memos/common/global"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	global.Routers = append(global.Routers, userInfoRouter)
}

func userInfoRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/userinfo")
	{
		r.GET("", home)
	}
}

func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Maid",
		"age":  14,
		"msg":  global.Config.Application.Name,
	})
}
