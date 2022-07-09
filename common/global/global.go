package global

import (
	"memos/common/global/config"

	"github.com/gin-gonic/gin"
)

const (
	Version = "0.0.1"
)

var (
	Engine *gin.Engine
	Config *config.Config
)

func init() {
	Engine = gin.New()
	Config = config.Conf
}
