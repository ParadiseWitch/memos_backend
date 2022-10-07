package config

import "memos/server/dto"

type Application struct {
	Mode    dto.AppMode `mapstructure:"mode"`
	Host    string      `mapstructure:"host"`
	Name    string      `mapstructure:"name"`
	Port    int         `mapstructure:"port"`
	Version string      `mapstructure:"version"`
}
