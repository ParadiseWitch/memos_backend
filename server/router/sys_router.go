package router

import (
	"github.com/gin-gonic/gin"
	"memos/server/service"
)

func init() {
	Routers = append(Routers, sysInfoRouter)
}

func sysInfoRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/sysinfo")
	{
		r.GET("", service.SysInfo)
	}
}
