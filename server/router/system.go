package router

import (
	"memos/server/config"
	"memos/server/logger"
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
	logger.GinLogger()
	c.JSON(http.StatusOK, gin.H{
		"msg": config.Conf,
	})
}
