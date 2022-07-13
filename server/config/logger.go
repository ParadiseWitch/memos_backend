package config

type Logger struct {
	Path       string `mapstructure:"path"`
	Level      Level  `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type Level string

const (
	LOGLEVEL_DEBUG  Level = "debug"
	LOGLEVEL_INFO   Level = "info"
	LOGLEVEL_WARN   Level = "warn"
	LOGLEVEL_ERROR  Level = "error"
	LOGLEVEL_DPANIC Level = "dpanic"
	LOGLEVEL_PANIC  Level = "panic"
	LOGLEVEL_FATAL  Level = "fatal"
)
