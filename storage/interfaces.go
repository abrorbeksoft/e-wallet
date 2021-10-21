package storage

import "github.com/abrorbeksoft/e-wallet.git/api/models"

type StorageI interface {
	Hello(message string) string
	GetWallet(id string) *models.Wallet
	PaymentHistory(id string) *models.Payments
	AddMoney(id string,amount int64) (bool, string)
	RemoveMoney(id string,amount int64) (bool, string)
	AllWallets(id string) []models.Wallet
	CreateWallet(wallet *models.CreateWallet, userId string) (string,error)
	// auth
	CreateUser(user *models.CreateUser) (string,error)
	GetUser(username string) (*models.User, error)
}
