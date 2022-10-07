package dto

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
