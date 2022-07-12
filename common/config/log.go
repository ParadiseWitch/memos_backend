package config

type Log struct {
	Path  string   `mapstructure:"path"`
	Level LogLevel `mapstructure:"level"`
}

type LogLevel string

const (
	LOGLEVEL_TRACE LogLevel = "trace"
	LOGLEVEL_DEBUG LogLevel = "debug"
	LOGLEVEL_INFO  LogLevel = "info"
	LOGLEVEL_WARN  LogLevel = "warn"
	LOGLEVEL_ERROR LogLevel = "error"
	LOGLEVEL_FATAL LogLevel = "fatal"
)
