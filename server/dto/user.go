package dto

type User struct {
	Id       int64     `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Avatar   string    `json:"avatar"`
	Role     Role `json:"role"`
}
