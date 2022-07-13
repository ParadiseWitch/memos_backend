package router

import (
	"memos/server/config"
	global2 "memos/server/global"
	"memos/server/logger"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Routers = make([]func(*gin.RouterGroup), 0)

func InitRouter() {
	r := global2.Engine
	v1 := r.Group("/api/v1")
	for _, f := range Routers {
		f(v1)
	}
	addr := config.Conf.Application.Host + ":" + strconv.Itoa(config.Conf.Application.Port)
	if err := r.Run(addr); err != nil {
		panic(err)
	}
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
}