package router

import (
	"memos/common/global"
	"memos/common/global/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	routerRole = append(routerRole, userInfoRouter)
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
		"msg2": config.Conf.Application.Name,
	})
}
