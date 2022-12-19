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
type Transactions interface{}

type Wallet interface{}

type Service struct {
	Authorization
	Transactions
	Wallet
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
