package config

import "memos/server/dto"

type Logger struct {
	Level      dto.Level `mapstructure:"level"`
	Filename   string    `mapstructure:"filename"`
	MaxSize    int       `mapstructure:"max_size"`
	MaxAge     int       `mapstructure:"max_age"`
	MaxBackups int       `mapstructure:"max_backups"`
}
