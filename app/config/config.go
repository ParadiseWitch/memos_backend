package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf *Config

type Config struct {
	Application Application `mapstructure:"application"`
	Db          Db          `mapstructure:"db"`
	Log         Log         `mapstructure:"log"`
}

func InitConf() {
	defaultConfig()
	viper.SetConfigName("settings")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AddConfigPath("./config/")
	viper.AddConfigPath(".")
	// watch conf file changed
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// callback if config file changed
		fmt.Println("config file changed:", e.Name)
		viper.Unmarshal(&Conf)
	})
	err := viper.ReadInConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		fmt.Println("config file not found")
		panic(fmt.Errorf("fatal error config file:%s", err))
	}
	viper.Unmarshal(&Conf)
}

func defaultConfig() {
	// set default config file
	viper.SetDefault("config.application.mode", APPMODE_DEV)
	viper.SetDefault("config.application.port", 8080)
}
