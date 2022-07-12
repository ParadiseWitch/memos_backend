package router

import (
	"memos/app/config"
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
	c.JSON(http.StatusOK, gin.H{
		"name": "Maid",
		"age":  14,
		"msg":  config.Conf.Application.Name,
	})
}
