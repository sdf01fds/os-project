package service

import (
	"github.com/google/uuid"
	project "github.com/sdf0106/os-project"
	"github.com/sdf0106/os-project/pkg/repository"
)

type WalletService struct {
	repo repository.Wallet
}

func (s *WalletService) CreateWallet(userId int, wallet project.Wallet) (int, error) {
	wallet.PrivateKey = uuid.New().String()
	wallet.PublicKey = uuid.New().String()
	return s.repo.CreateWallet(userId, wallet)
}

func (s *WalletService) GetAllWallets(userId int) ([]project.Wallet, error) {
	return s.repo.GetAllWallets(userId)
}

func (s *WalletService) GetWalletById(userId, walletId int) (project.Wallet, error) {
	return s.repo.GetWalletById(userId, walletId)
}

func (s *WalletService) UpdateWalletBalance(userId, walletId int, amount float32) error {
	return s.repo.UpdateWalletBalance(userId, walletId, amount)
}
