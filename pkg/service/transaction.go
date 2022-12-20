package service

import (
	project "github.com/sdf0106/os-project"
	"github.com/sdf0106/os-project/pkg/repository"
)

type TransactionsService struct {
	repo repository.Transactions
}

func (s *TransactionsService) CreateTransaction(userId int, transaction project.Transaction) (project.Transaction, error) {
	return s.repo.CreateTransaction(userId, transaction)
}

func (s *TransactionsService) GetAllTransactions(userId int) ([]project.Transaction, error) {
	return s.repo.GetAllTransactions(userId)
}

func (s *TransactionsService) GetTransactionById(userId, transactionId int) (project.Transaction, error) {
	return s.repo.GetTransactionById(userId, transactionId)
}

func (s *TransactionsService) DeleteTransactionById(userId, transactionId int) error {
	return s.repo.DeleteTransaction(userId, transactionId)
}
