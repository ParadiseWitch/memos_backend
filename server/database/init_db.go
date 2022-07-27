package database

import (
	"fmt"
	"memos/server/config"
	"memos/server/dto"
	"memos/server/logger"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitDB() {
	logger.Infof("Connecting to database %s:%d...", config.Conf.Db.Host, config.Conf.Db.Port)

	url := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/memos?charset=utf8&parseTime=True&loc=Local&timeout=1000ms",
		config.Conf.Db.User,
		config.Conf.Db.Password,
		config.Conf.Db.Host,
		config.Conf.Db.Port,
	)
	db, err := gorm.Open("mysql", url)
	if err != nil {
		logger.Panic(err)
		panic(err)
	}
	db.DB().SetMaxIdleConns(10)
	if !db.HasTable("users") {
		db.CreateTable(&dto.User{})
	}
	db.Table("users").Create(&dto.User{})

	if err := db.Close(); err != nil {
		return
	}
}
