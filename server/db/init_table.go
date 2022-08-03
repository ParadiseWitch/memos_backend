package db

import (
	"memos/server/dto"
	"memos/server/logger"
)

/**
 * 初始化用户表及其数据
 */
func InitTableUser() {
	if DB.HasTable(dto.User{}) {
		return
	}
	logger.Info(`models.role: creating table "role"`)
	if err := DB.CreateTable(dto.User{}).Error; err != nil {
		logger.Panic(err)
	}
	DB.Create(&dto.User{
		ID:       1,
		Name:     "admin",
		Email:    "admin@memos.com",
		Password: "admin",
		Avatar:   "",
		Role:     1,
	})
	DB.Create(&dto.User{
		ID:       2,
		Name:     "maid",
		Email:    "maid@memos.com",
		Password: "maid",
		Avatar:   "",
		Role:     2,
	})
	DB.Create(&dto.User{
		ID:       3,
		Name:     "user",
		Email:    "user@memos.com",
		Password: "user",
		Avatar:   "",
		Role:     3,
	})
}
