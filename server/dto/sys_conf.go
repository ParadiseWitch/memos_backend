package dto

type SysConf struct {
	ID     int64  `gorm:"primary_key"`
	Name   string `gorm:"size:32"`
	Value  string `gorm:"size:32"`
	Type   string `gorm:"size:32"`
	Remark string `gorm:"size:32"`
}
