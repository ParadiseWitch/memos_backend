package main

import (
	"memos/server/config"
	"memos/server/ioc"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	ioc.Provide(NewLogger)
	l := ioc.Invoke[*config.Logger]()
	print(l)
}

func NewLogger() *config.Logger {
	return &config.Logger{
		Level: config.LOGLEVEL_DEBUG,
	}
}
