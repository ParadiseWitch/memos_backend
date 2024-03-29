package service

import (
	"fmt"
	"memos/server/logger"
	"memos/server/pkg/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	ContextUserIDKey = "uid"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			logger.Warnf("The request header lacks Authorization")
			c.JSON(http.StatusUnauthorized, CommonFailRes("权限不足"))
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			logger.Warnf("Token format error, token: %v", authHeader)
			c.JSON(http.StatusUnauthorized, CommonFailRes("权限不足"))
			c.Abort()
			return
		}

		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			fmt.Println(err)
			logger.Warnf("InvaildToken, err:%+v", err)
			c.JSON(http.StatusUnauthorized, CommonFailRes("权限不足"))
			c.Abort()
			return
		}
		c.Set(ContextUserIDKey, mc.UserID)
		c.Next()
	}
}
