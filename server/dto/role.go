package dto

import (
	"log"
	"memos/server/db"
)

type Role struct {
	ID    int64  `gorm:"primary_key"`
	Name  string `gorm:"size:32"`
	Right string `gorm:"size:8"`
}

/*
前4位用于管理，后4位用户用户
目前都是预留功能
*/
func InitTableRole() {
	if db.DB.HasTable(Role{}) {
		return
	}
	log.Println(`models.role: creating table "role"`)
	if err := db.DB.CreateTable(Role{}).Error; err != nil {
		log.Panic(err)
	}
	db.DB.Create(&Role{
		ID:    1,
		Name:  "admin",
		Right: "11111111",
	})
	db.DB.Create(&Role{
		ID:    2,
		Name:  "manager",
		Right: "01111111",
	})
	db.DB.Create(&Role{
		ID:    3,
		Name:  "user",
		Right: "00001111",
	})
}
