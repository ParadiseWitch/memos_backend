package dto

type Role struct {
	ID    int64  `gorm:"primary_key"`
	Name  string `gorm:"size:32"`
	Right string `gorm:"size:8"`
}

var (
	AdminRole Role = Role{
		ID:    1,
		Name:  "admin",
		Right: "11111111",
	}
	ManagerRole Role = Role{
		ID:    2,
		Name:  "manager",
		Right: "01111111",
	}
	UserRole Role = Role{
		ID:    3,
		Name:  "user",
		Right: "00001111",
	}
)
