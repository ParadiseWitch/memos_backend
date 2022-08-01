package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CommonSuccRes(data interface{}) gin.H {
	return gin.H{
		"code":   http.StatusOK,
		"status": "success",
		"data":   data,
	}
}

func CommonFailRes(code int, err string) gin.H {
	return gin.H{
		"code":   code,
		"status": "fail",
		"error":  err,
	}
}
