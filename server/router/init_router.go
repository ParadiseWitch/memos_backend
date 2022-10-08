package router

import (
	"memos/server/config"
	"memos/server/logger"
	"memos/server/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Routers = make([]func(*gin.RouterGroup), 0)
var R *gin.Engine

func InitRouter() {
	R = gin.Default()
	v1 := R.Group("/api/v1")
	v1.POST("/user/register", service.CommonHandle(service.Register))
	v1.POST("/user/login", service.CommonHandle(service.Login))
	v1.Use(service.JWTAuthMiddleware())
	for _, f := range Routers {
		f(v1)
	}
	R.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, service.CommonFailRes("not found"))
	})
}

func RunServer() {
	addr := config.Conf.Application.Host + ":" + strconv.Itoa(config.Conf.Application.Port)
	R.Use(logger.GinLogger(), logger.GinRecovery(true))
	if err := R.Run(addr); err != nil {
		panic(err)
	}
}
