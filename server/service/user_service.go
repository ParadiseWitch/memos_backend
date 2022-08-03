package service

import (
	"memos/server/db"
	"memos/server/dto"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var userTable *gorm.DB

func getUserTable() *gorm.DB {
	if userTable == nil {
		userTable = db.DB.Table("user")
	}
	return userTable
}

func GetUserById(c *gin.Context) {
	var u dto.User
	GetById("user", u, c)
}

func UpdateUserById(c *gin.Context) {
	var u dto.User
	UpdateById("user", u, c)
}

func AddUser(c *gin.Context) {
	var u dto.User
	Add("user", u, c)
}
