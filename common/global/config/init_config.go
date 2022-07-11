package config

import (
	"fmt"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	Conf *Config
	l    sync.Mutex
)

func initViper() {
	defaultConfig()
	viper.SetConfigName("settings")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AddConfigPath("./config/")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		fmt.Println("config file not found")
		panic(fmt.Errorf("fatal error config file:%s", err))
	}
	// watch conf file changed
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		l.Lock()
		// callback if config file changed
		fmt.Println("config file changed:", e.Name)
		InitConfig()
		l.Unlock()
	})
}

func defaultConfig() {
	// set default config file
	viper.SetDefault("config.application.mode", APPMODE_DEV)
	viper.SetDefault("config.application.port", 8080)
}

func InitConfig() {
	initViper()
	application := Application{
		Mode: AppMode(viper.GetString("config.application.mode")),
		Host: viper.GetString("config.application.host"),
		Name: viper.GetString("config.application.name"),
		Port: viper.GetInt("config.application.port"),
	}
	db := Db{
		Host:     viper.GetString("config.db.host"),
		Port:     viper.GetInt("config.db.port"),
		User:     viper.GetString("config.db.user"),
		Name:     viper.GetString("config.db.name"),
		Password: viper.GetString("config.db.password"),
	}
	log := Log{
		Level: LogLevel(viper.GetString("config.log.level")),
		Path:  viper.GetString("config.log.path"),
	}
	Conf.Application = application
	Conf.Db = db
	Conf.Log = log
}

func init() {
	Conf = &Config{
		Application: Application{},
		Db:          Db{},
		Log:         Log{},
	}
	InitConfig()
}
