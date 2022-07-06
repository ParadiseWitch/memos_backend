package global

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

const (
	Version = "0.0.1"
)

var (
	Engine *gin.Engine
	Viper  *viper.Viper
)

func GetEngine() *gin.Engine {
	if Engine == nil {
		Engine = gin.New()
	}
	return Engine
}

func GetViper() *viper.Viper {
	if Viper == nil {
		Viper = viper.New()
		initViper()
	}
	return Viper
}

func initViper() {
	// set default config file
	Viper.SetConfigName("settings")
	Viper.SetConfigType("yaml")
	Viper.AddConfigPath("./")
	Viper.AddConfigPath("./config/")
	Viper.AddConfigPath(".")
	err := Viper.ReadInConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		fmt.Println("config file not found")
		panic(fmt.Errorf("fatal error config file:%s", err))
	}
	// watch conf file changed
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// callback if config file changed
		fmt.Println("config file changed:", e.Name)
	})
}
