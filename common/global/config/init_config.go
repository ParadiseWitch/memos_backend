package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	Conf *Config
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
		// callback if config file changed
		fmt.Println("config file changed:", e.Name)
		InitConfig()
	})
}

func defaultConfig() {
	// set default config file
	viper.SetDefault("config.application.mode", APPMODE_DEV)
	viper.SetDefault("config.application.port", 8080)
}

func InitConfig() {
	initViper()
	application := &Application{
		Mode: AppMode(viper.GetString("config.application.mode")),
		Host: viper.GetString("config.application.host"),
		Name: viper.GetString("config.application.name"),
		Port: viper.GetInt("config.application.port"),
	}
	db := &Db{
		Host:     viper.GetString("config.db.host"),
		Port:     viper.GetInt("config.db.port"),
		User:     viper.GetString("config.db.user"),
		Name:     viper.GetString("config.db.name"),
		Password: viper.GetString("config.db.password"),
	}
	log := &Log{
		Level: LogLevel(viper.GetString("config.log.level")),
		Path:  viper.GetString("config.log.path"),
	}
	Conf = &Config{
		Application: *application,
		Db:          *db,
		Log:         *log,
	}
}

func init() {
	InitConfig()
}
