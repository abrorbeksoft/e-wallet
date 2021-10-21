package models

type User struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateUser struct {
	Name string `json:"name"`
	Surname string `json:"surname"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SuccessMessage struct {
	Message string `json:"message"`
	Token string `json:"token"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}