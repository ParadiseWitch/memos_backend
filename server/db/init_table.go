package db

import (
	"memos/server/dto"
	"memos/server/logger"
)

/**
 * 初始化表
 */
func InitTable() {
	if DB.HasTable(dto.SysConf{}) {
		return
	}
	logger.Info(`models.sysconf: creating table "sys_conf"`)
	if err := DB.CreateTable(dto.SysConf{}).Error; err != nil {
		logger.Panic(err)
	}

	if DB.HasTable(dto.User{}) {
		return
	}
	logger.Info(`models.user: creating table "user"`)
	if err := DB.CreateTable(dto.User{}).Error; err != nil {
		logger.Panic(err)
	}
}

/**
 * 初始化用户表数据
 */
func InitUserTableData() {
	if DB.HasTable(dto.User{}) {
		return
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
