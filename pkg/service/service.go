package service

import (
	project "github.com/sdf0106/os-project"
	"github.com/sdf0106/os-project/pkg/repository"
)

type Authorization interface {
	CreateUser(user project.User) (project.User, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}
type Transactions interface {
	CreateTransaction(userId int, transaction project.Transaction) (project.Transaction, error)
	GetAllTransactions(userId int) ([]project.Transaction, error)
	GetTransactionById(userId, transactionId int) (project.Transaction, error)
	DeleteTransaction(userId, transactionId int) error
}

type Wallet interface {
	CreateWallet(userId int, wallet project.Wallet) (int, error)
	GetAllWallets(userId int) ([]project.Wallet, error)
	GetWalletById(userId, walletId int) (project.Wallet, error)
	UpdateWalletBalance(userId, walletId int, amount float32) error
}

type Service struct {
	Authorization
	Transactions
	Wallet
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Transactions:  NewTransactionService(repos.Transactions),
	}
}
