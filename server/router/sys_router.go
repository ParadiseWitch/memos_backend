package router

import (
  "memos/server/service"

  "github.com/gin-gonic/gin"
)

func init() {
  Routers = append(Routers, sysInfoRouter)
}

func sysInfoRouter(v1 *gin.RouterGroup) {
  r := v1.Group("/sysinfo")
  {
    r.GET("", service.SysInfo)
  }
}
