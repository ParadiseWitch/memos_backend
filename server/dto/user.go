package dto

type User struct {
	ID       int64  `gorm:"primary_key"`
	Name     string `gorm:"size:32"`
	Email    string `gorm:"size:32"`
	Password string `gorm:"size:32"`
	Avatar   string `gorm:"size:32"`
	Role     int8   `gorm:"size:32"`
}
