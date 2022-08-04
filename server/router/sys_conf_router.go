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
		r.GET("/info", service.GetSysConfById)
		r.POST("/update", service.UpdateSysConfById)
		r.POST("/add", service.AddSysConf)
	}
}
