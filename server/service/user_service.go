package service

import (
	"memos/server/db"
	"memos/server/dto"
	"memos/server/logger"
	"memos/server/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userTable = db.DB.Table("user")

func GetUserById(c *gin.Context) {
	var u dto.User
	id, ok := c.GetQuery("id")
	if !ok {
		logger.Errorf("GetUserById err, id: %s", id)
		c.JSON(http.StatusBadRequest,
			util.CommonFailRes(http.StatusBadRequest, "id is required"))
	}
	userTable.Where("id = ?", id).Find(&u)
	c.JSON(http.StatusOK, util.CommonSuccRes(u))
}

func AddUser(c *gin.Context) {
	var u dto.User
	err := c.BindJSON(&u)
	if err != nil {
		logger.Errorf("AddUser err, err: %s", err)
		c.JSON(http.StatusBadRequest,
			util.CommonFailRes(http.StatusBadRequest, "invalid request"))
	}
	c.JSON(http.StatusOK, util.CommonSuccRes(u))
}
