package db

import (
	"log"
	"memos/server/dto"
)

/*
前4位用于管理，后4位用户用户
目前都是预留功能
*/
func InitTableRole() {
	// Fixme
	if DB.HasTable(dto.Role{}) {
		return
	}
	log.Println(`models.role: creating table "role"`)
	if err := DB.CreateTable(dto.Role{}).Error; err != nil {
		log.Panic(err)
	}
	DB.Create(&dto.AdminRole)
	DB.Create(&dto.ManagerRole)
	DB.Create(&dto.UserRole)
}

func InitTableUser() {
	if DB.HasTable(dto.User{}) {
		return
	}
	DB.Create(&dto.User{
		ID:       1,
		Name:     "admin",
		Email:    "admin@memos.com",
		Password: "admin",
		Avatar:   "",
		Role:     dto.AdminRole,
	})
	DB.Create(&dto.User{
		ID:       2,
		Name:     "maid",
		Email:    "maid@memos.com",
		Password: "maid",
		Avatar:   "",
		Role:     dto.ManagerRole,
	})
	DB.Create(&dto.User{
		ID:       3,
		Name:     "user",
		Email:    "user@memos.com",
		Password: "user",
		Avatar:   "",
		Role:     dto.UserRole,
	})
}
