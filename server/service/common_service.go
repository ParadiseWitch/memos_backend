package service

import (
	"memos/server/dto"
	"memos/server/ex"
	"memos/server/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type servicehander func(c *gin.Context) (data interface{}, e *ex.HttpEX)

func CommonSuccRes(data interface{}) gin.H {
	return gin.H{
		"status": "success",
		"data":   data,
	}
}

func CommonFailRes(msg string) gin.H {
	return gin.H{
		"status": "fail",
		"msg":    msg,
	}
}

func CommonHandle(hander servicehander) func(c *gin.Context) {
	return func(c *gin.Context) {
		data, ex := hander(c)
		if ex != nil {
			// 统一异常处理
			HandlerHttpEx(*ex, c)
			return
		}
		// 统一返回值处理
		c.JSON(http.StatusOK, CommonSuccRes(data))
	}
}

func HandlerHttpEx(ex ex.HttpEX, c *gin.Context) {
	exCode := ex.Code
	httpCode := exCode.HttpCode
	logMsg := exCode.LogMsg + ", " + ex.DetailMsg
	tipMsg := exCode.TipMsg
	LogByLevel(exCode.LogLevel, logMsg)
	c.JSON(httpCode, CommonFailRes(tipMsg))
}

func LogByLevel(l dto.Level, msg string) {
	if l == dto.LOGLEVEL_WARN {
		logger.Warn(msg)
	} else if l == dto.LOGLEVEL_ERROR {
		logger.Error(msg)
	}
}
