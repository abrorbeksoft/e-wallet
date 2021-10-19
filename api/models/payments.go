package models

type Payment struct {
	Id string `json:"id"`
	Amount int64 `json:"amount"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Payments struct {
	Count int `json:"count"`
	All []Payment `json:"all"`
}

