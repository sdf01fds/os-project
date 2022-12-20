package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	project "github.com/sdf0106/os-project"
	"time"
)

type WalletPostgres struct {
	db *sqlx.DB
}

func NewWalletPostgres(db *sqlx.DB) *WalletPostgres {
	return &WalletPostgres{db: db}
}

func (r *WalletPostgres) CreateWallet(userId int, wallet project.Wallet) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (user_id, private_key, public_key, balance, created_at) values ($1, $2, $3, $4, $5)",
		walletsTable)
	row := r.db.QueryRow(query, userId, wallet.PrivateKey, wallet.PrivateKey, wallet.Balance, time.Now())
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *WalletPostgres) GetAllWallets(userId int) ([]project.Wallet, error) {
	var wallets []project.Wallet
	query := fmt.Sprintf("SELECT wt.id, wt.user_id, wt.private_key, wt.public_kay, wt.balance, wt.created_at, wt.updated_at FROM %s ut INNER JOIN %s wt ON ut.id = wt.user_id WHERE ut.id = $1",
		usersTable, walletsTable)
	if err := r.db.Select(&wallets, query, userId); err != nil {
		return nil, err
	}

	return wallets, nil
}

func (r *WalletPostgres) GetWalletById(userId, walletId int) (project.Wallet, error) {
	var wallet project.Wallet
	query := fmt.Sprintf("SELECT wt.id, wt.user_id, wt.private_key, wt.public_kay, wt.balance, wt.created_at, wt.updated_at FROM %s ut INNER JOIN %s wt ON ut.id = wt.user_id WHERE ut.id = $1 AND wt.id = $2",
		usersTable, walletsTable)
	err := r.db.Get(&wallet, query, userId, walletId)

	return wallet, err
}

func (r *WalletPostgres) UpdateWalletBalance(userId, walletId int, amount float32) error {
	query := fmt.Sprintf("UPDATE %s wt SET wi.balance = $1 FROM %s ut WHERE wt.id = $2 AND ut.id = $3 AND wt.user_id = ut.id", walletsTable, usersTable)
	_, err := r.db.Exec(query, amount, walletId, userId)
	return err
}
