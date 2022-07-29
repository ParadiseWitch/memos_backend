package dto

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	Role     Role   `json:"role"`
}
