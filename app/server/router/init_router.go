package router

import (
	"memos/common/global"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := global.Engine
	if r == nil {
		r = gin.New()
	}
	r.Use(gin.Recovery())
}
