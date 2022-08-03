package dto

import "errors"

type Role struct {
	ID    int64  `gorm:"primary_key"`
	Name  string `gorm:"size:32"`
	Right string `gorm:"size:8"`
}

var ROLES = map[string]*Role{
	"ROLE_ADMIN": {
		ID:    1,
		Name:  "admin",
		Right: "11111111",
	},
	"ROLE_USER": {
		ID:    2,
		Name:  "user",
		Right: "00000001",
	},
}

func GetRoleById(id int64) (*Role, error) {
	for _, v := range ROLES {
		if v.ID == id {
			return v, nil
		}
	}
	return nil, errors.New("not found Role, id: " + string(rune(id)))
}
