package db

import (
	"fmt"
	"memos/server/config"
	"memos/server/logger"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDB() {
	logger.Infof("Connecting to database %s:%d...", config.Conf.Db.Host, config.Conf.Db.Port)

	url := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/memos?charset=utf8&parseTime=True&loc=Local&timeout=1000ms",
		config.Conf.Db.User,
		config.Conf.Db.Password,
		config.Conf.Db.Host,
		config.Conf.Db.Port,
	)
	DB, err := gorm.Open("mysql", url)
	if err != nil {
		logger.Panic(err)
	}
	DB.DB().SetMaxIdleConns(10)
	DB.SingularTable(true)
	// if !db.HasTable("users") {
	// 	db.CreateTable(&dto.User{})
	// }
	// db.Table("users").Create(&dto.User{})
	// if err := db.Close(); err != nil {
	// 	return
	// }
}
