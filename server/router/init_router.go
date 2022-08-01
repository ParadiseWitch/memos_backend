package router

import (
	"memos/server/config"
	"memos/server/logger"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Routers = make([]func(*gin.RouterGroup), 0)
var R *gin.Engine

func InitRouter() {
	R = gin.Default()
	v1 := R.Group("/api/v1")
	for _, f := range Routers {
		f(v1)
	}
}

func RunServer() {
	addr := config.Conf.Application.Host + ":" + strconv.Itoa(config.Conf.Application.Port)
	R.Use(logger.GinLogger(), logger.GinRecovery(true))
	if err := R.Run(addr); err != nil {
		panic(err)
	}
}
