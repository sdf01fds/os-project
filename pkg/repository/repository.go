package repository

import (
	"github.com/jmoiron/sqlx"
	project "github.com/sdf0106/os-project"
)

type Authorization interface {
	CreateUser(user project.User) (project.User, error)
	GetUser(email, password string) (project.User, error)
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

type Repository struct {
	Authorization
	Transactions
	Wallet
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
