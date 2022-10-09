package service

import (
	"fmt"
	"memos/server/db"
	"memos/server/dto"
	"memos/server/ex"

	"github.com/gin-gonic/gin"
)

const SYSCONF_TABLENAME = "sys_conf"

func GetSysConfById(c *gin.Context) (data any, e *ex.HttpEX) {
	var s dto.SysConf
	id, ok := c.GetQuery("id")
	if !ok {
		return nil, ex.New(nil, ex.EC_INVALID_PARAM(), "id is require, id="+id)
	}

	rset := db.DB.Table(SYSCONF_TABLENAME).Where("id = ?", id).Find(&s)
	if rset.RecordNotFound() {
		return nil, ex.New(rset.Error, ex.EC_DATA_NOT_EXIST("系统配置"), "getsysconfbyid: the system config not exist. id: "+id)
	}
	if rset.Error != nil {
		return nil, ex.New(rset.Error, ex.EC_DB_OP_ERROR(), "GetSysConfById err, id: "+id)
	}
	return s, nil
}

func UpdateSysConfById(c *gin.Context) (data any, e *ex.HttpEX) {
	var s dto.SysConf
	err := c.BindJSON(&s)
	if err != nil {
		return nil, ex.New(err, ex.EC_INVALID_PARAM(), "UpdateSysConfById err")
	}
	if err := db.DB.Table(SYSCONF_TABLENAME).Save(&s).Error; err != nil {
		return nil, ex.New(err, ex.EC_DB_OP_ERROR(), fmt.Sprintf("updatesysconfbyid err, data submitted%+v", s))
	}
	return nil, nil
}

func AddSysConf(c *gin.Context) (data any, e *ex.HttpEX) {
	var s dto.SysConf
	err := c.BindJSON(&s)
	if err != nil {
		return nil, ex.New(err, ex.EC_INVALID_PARAM(), "addsysconf err")
	}
	if err := db.DB.Table(SYSCONF_TABLENAME).Create(&s).Error; err != nil {
		return nil, ex.New(err, ex.EC_DB_OP_ERROR(), fmt.Sprintf("addsysconf err, data submitted:%+v", e))
	}
	return s, nil
}
