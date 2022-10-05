package router

import (
	"memos/server/service"

	"github.com/gin-gonic/gin"
)

func init() {
	Routers = append(Routers, sysConfInfoRouter)
}

func sysConfInfoRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/sys")
	{
		r.GET("/info", service.CommonHandle(service.GetSysConfById))
		r.POST("/update", service.CommonHandle(service.UpdateSysConfById))
		r.POST("/add", service.CommonHandle(service.AddSysConf))
	}
}
