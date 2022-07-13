package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var Conf *Config

type Config struct {
	Application Application `mapstructure:"application"`
	Db          Db          `mapstructure:"db"`
	Logger      Logger      `mapstructure:"logger"`
}

func InitConf() {
	// TODO: get viper instance
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
		//fmt.Println("config file changed:", e.Name)
		zap.L().Info("config file changed" + e.Name)
		if err := viper.Unmarshal(&Conf); err != nil {
			return
		}
	})
	err := viper.ReadInConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		panic(fmt.Errorf("fatal error config file:%s", err))
	}
	if err := viper.Unmarshal(&Conf); err != nil {
		return
	}
}

func defaultConfig() {
	// set default config file
	viper.SetDefault("config.application.mode", APPMODE_DEV)
	viper.SetDefault("config.application.port", 8080)
}
