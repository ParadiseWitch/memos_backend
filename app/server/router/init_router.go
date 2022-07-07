package router

import (
	"memos/common/global"

	"github.com/gin-gonic/gin"
)

var (
	routerRole = make([]func(*gin.RouterGroup), 0)
)

func InitRouter() {
	r := global.Engine
	v1 := r.Group("/api/v1")
	for _, f := range routerRole {
		f(v1)
	}
	r.Run()
}
