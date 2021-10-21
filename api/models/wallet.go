package models

type Wallet struct {
	Id string `json:"id"`
	Type string `json:"type"`
	Amount int64 `json:"amount"`
	UserId string `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateWallet struct {
	Type string `json:"type"`
}
