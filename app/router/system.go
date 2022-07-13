package router

import (
	"memos/app/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	Routers = append(Routers, sysInfoRouter)
}

func sysInfoRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/sysinfo")
	{
		r.GET("", sysInfoService)
	}
}

func sysInfoService(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": config.Conf,
	})
}
