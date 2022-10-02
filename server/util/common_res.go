package util

import (
	"github.com/gin-gonic/gin"
)

func CommonSuccRes(data interface{}) gin.H {
	return gin.H{
		"status": "success",
		"data":   data,
	}
}

func CommonFailRes(err string) gin.H {
	return gin.H{
		"status": "fail",
		"error":  err,
	}
}
