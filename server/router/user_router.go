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
		r.GET("/info", service.GetUserById)
		r.POST("/add", service.AddUser)
		r.POST("/update", service.UpdateUserById)
	}
}
