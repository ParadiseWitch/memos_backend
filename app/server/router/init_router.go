package router

import (
	"memos/common/global"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := global.GetEngine()
	r.Use(gin.Recovery())
}
