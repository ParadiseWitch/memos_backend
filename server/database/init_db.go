package database

import (
	"memos/server/dto"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitDB() {
	//url := config.Conf.Db.Url
	db, err := gorm.Open("mysql", "root:123456@tcp(maiiiiiid.fun:23306)/memos?charset=utf8&parseTime=True&loc=Local&timeout=1000ms")
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxIdleConns(10)
	db.CreateTable(&dto.User{})
	db.Table("users").Create(&dto.User{})

	if err := db.Close(); err != nil {
		return
	}
}
