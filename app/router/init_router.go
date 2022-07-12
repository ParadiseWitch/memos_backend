package router

import (
	"memos/common/global"

	"github.com/gin-gonic/gin"
)

var Routers = make([]func(*gin.RouterGroup), 0)

func InitRouter() {
	r := global.Engine
	v1 := r.Group("/api/v1")
	for _, f := range Routers {
		f(v1)
	}
	r.Run()
}
