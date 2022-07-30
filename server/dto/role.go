package dto

type Role struct {
  ID    int64  `gorm:"primary_key"`
  Name  string `gorm:"size:32"`
  Right string `gorm:"size:8"`
}
