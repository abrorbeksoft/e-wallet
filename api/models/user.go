package models

type User struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Username string `json:"username"`
	Email string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}