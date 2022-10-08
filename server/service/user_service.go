package service

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"memos/server/db"
	"memos/server/dto"
	"memos/server/ex"

	"github.com/gin-gonic/gin"
)

const USER_TABLENAME = "user"

const secret = "memos"

func encryptPassword(str string) (result string) {
	h := md5.New()
	h.Write([]byte(secret + str))
	return hex.EncodeToString(h.Sum(nil))
}

func Register(c *gin.Context) (data any, err *ex.HttpEX) {
	var paramU dto.User
	if err := c.BindJSON(&paramU); err != nil {
		return nil, ex.New(nil, ex.EC_INVALID_PARAM, "register err")
	}

	_, err = GetUserByName(paramU.Name)
	if err == nil || err.Code != ex.EC_DATA_NOT_EXIST {
		// 该用户已经注册
		return nil, ex.New(nil, ex.EC_USER_HAS_REGISTERED, "name="+paramU.Name)
	}

	paramU.Password = encryptPassword(paramU.Password)
	if err := db.DB.Create(&paramU).Error; err != nil {
		return nil, ex.New(err, ex.EC_DB_OP_ERROR, "name="+paramU.Name)
	}
	return nil, nil
}

func Login(c *gin.Context) (data any, e *ex.HttpEX) {
	var paramU dto.User
	if err := c.BindJSON(&paramU); err != nil {
		return nil, ex.New(nil, ex.EC_INVALID_PARAM, "login err")
	}

	resultU, err := GetUserByName(paramU.Name)
	if err != nil {
		return nil, err
	}

	if resultU.Password != encryptPassword(paramU.Password) {
		return nil, ex.New(nil, ex.EC_PSD_MISTAKE, "name="+paramU.Name)
	}
	return nil, nil
}

func GetUserById(c *gin.Context) (data any, e *ex.HttpEX) {
	var u dto.User
	id, ok := c.GetQuery("id")
	if !ok {
		return nil, ex.New(nil, ex.EC_INVALID_PARAM, "id is require, id="+id)
	}
	rset := db.DB.Table(USER_TABLENAME).Where("id = ?", id).Find(&u)
	if rset.RecordNotFound() {
		return nil, ex.New(rset.Error, ex.EC_DATA_NOT_EXIST, "getuserbyid: the user does not exist. uid="+id)
	}
	if rset.Error != nil {
		return nil, ex.New(rset.Error, ex.EC_DB_OP_ERROR, "getuserbyid err, uid="+id)
	}
	return u, nil
}

func GetUserByName(name string) (*dto.User, *ex.HttpEX) {
	var resultU dto.User
	rset := db.DB.Table(USER_TABLENAME).Where("name = ?", name).Find(&resultU)
	if rset.RecordNotFound() {
		return nil, ex.New(rset.Error, ex.EC_DATA_NOT_EXIST, "the user not exist, name: "+name)
	}
	if rset.Error != nil {
		return nil, ex.New(rset.Error, ex.EC_DB_OP_ERROR, "getuserbyname err, name="+name)
	}
	return &resultU, nil
}

func UpdateUserById(c *gin.Context) (data any, e *ex.HttpEX) {
	var u dto.User
	err := c.BindJSON(&u)
	if err != nil {
		return nil, ex.New(err, ex.EC_INVALID_PARAM, "")
	}
	if err := db.DB.Table(USER_TABLENAME).Save(&u).Error; err != nil {
		return nil, ex.New(err, ex.EC_DB_OP_ERROR, fmt.Sprintf("updateuserbyid err, data submitted:%+v", u))
	}
	return nil, nil
}

func UpdateUserByName(name string) (*dto.User, *ex.HttpEX) {
	var resultU dto.User
	if err := db.DB.Table(USER_TABLENAME).Save(&resultU).Error; err != nil {
		return nil, ex.New(err, ex.EC_DB_OP_ERROR, "UpdateUserByName err, name="+name)
	}
	return &resultU, nil
}

func AddUser(c *gin.Context) (data any, e *ex.HttpEX) {
	var u dto.User
	err := c.BindJSON(&u)
	if err != nil {
		return nil, ex.New(err, ex.EC_INVALID_PARAM, "")
	}
	if err := db.DB.Table(USER_TABLENAME).Create(&u).Error; err != nil {
		return nil, ex.New(err, ex.EC_DB_OP_ERROR, fmt.Sprintf("adduser err, data submitted:%+v", u))
	}
	return u, nil
}
