package repository

import (
	"github.com/jmoiron/sqlx"
	project "github.com/sdf0106/os-project"
)

type Authorization interface {
	CreateUser(user project.User) (project.User, error)
	GetUser(email, password string) (project.User, error)
}

type Transactions interface{}

type Wallet interface{}

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
