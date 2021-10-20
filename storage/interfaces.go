package storage

import "github.com/abrorbeksoft/e-wallet.git/api/models"

type StorageI interface {
	Hello(message string) string
	GetWallet(id string) *models.Wallet
	PaymentHistory(id string) *models.Payments
	AddMoney(id string,amount int64) (bool, string)
	RemoveMoney(id string,amount int64) (bool, string)
	AllWallets(id string) []models.Wallet
}


