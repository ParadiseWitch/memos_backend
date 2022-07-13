package router

import (
	"memos/app/config"
	"memos/common/global"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Routers = make([]func(*gin.RouterGroup), 0)

func InitRouter() {
	r := global.Engine
	v1 := r.Group("/api/v1")
	for _, f := range Routers {
		f(v1)
	}
	addr := config.Conf.Application.Host + ":" + strconv.Itoa(config.Conf.Application.Port)
	r.Run(addr)
}
