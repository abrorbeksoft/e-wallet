package models

type Payment struct {
	Id string `json:"id"`
	Amount int64 `json:"amount"`
	Type string `json:"type"`
	WalletId string `json:"wallet_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Payments struct {
	Count int `json:"count"`
	All []Payment `json:"all"`
}

