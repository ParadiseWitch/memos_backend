package router

import (
  "github.com/gin-gonic/gin"
  "memos/server/service"
)

func init() {
  Routers = append(Routers, userInfoRouter)
}

func userInfoRouter(v1 *gin.RouterGroup) {
  r := v1.Group("/userinfo")
  {
    r.GET("", service.UserInfo)
  }
}
