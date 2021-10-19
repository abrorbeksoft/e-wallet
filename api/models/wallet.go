package models

type Wallet struct {
	Id string `json:"id"`
	Type string `json:"type"`
	Amount int64 `json:"amount"`
	UserId string `json:"user_id"`
	Identified bool `json:"identified"`
}
