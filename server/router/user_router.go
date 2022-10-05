package router

import (
	"memos/server/service"

	"github.com/gin-gonic/gin"
)

func init() {
	Routers = append(Routers, userInfoRouter)
}

func userInfoRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/user")
	{
		r.POST("/register", service.CommonHandle(service.Register))
		r.POST("/login", service.CommonHandle(service.Login))
		r.GET("/info", service.CommonHandle(service.GetUserById))
		r.POST("/add", service.CommonHandle(service.AddUser))
		r.POST("/update", service.CommonHandle(service.UpdateUserById))
	}
}
