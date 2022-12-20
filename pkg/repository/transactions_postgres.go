package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	project "github.com/sdf0106/os-project"
	"time"
)

type TransactionsPostgres struct {
	db *sqlx.DB
}

func NewTransactionsPostgres(db *sqlx.DB) *TransactionsPostgres {
	return &TransactionsPostgres{db: db}
}

func (r *TransactionsPostgres) CreateTransaction(userId int, transaction project.Transaction) (project.Transaction, error) {
	query := fmt.Sprintf("INSERT INTO %s (sender_id, receiver_id, wallet_id, created_at, amount) values ($1, $2, $3, $4, $5) RETURNING ID", transactionsTable)
	row := r.db.QueryRow(query, userId, transaction.ReceiverId, transaction.WalletId, time.Now(), transaction.Amount)
	err := row.Scan(&transaction.Id)

	return transaction, err
}

func (r *TransactionsPostgres) GetAllTransactions(userId int) ([]project.Transaction, error) {
	var transactions []project.Transaction
	query := fmt.Sprintf("SELECT tt.id, ti.sender_id, ti.receiver_id, tt.wallet_id, tt.amount, tt.created_at FROM %s ut INNER JOIN %s tt ON ut.id = tt.sender_id WHERE ut.id = $1 ORDER BY tt.id",
		usersTable, transactionsTable)
	if err := r.db.Select(&transactions, query, userId); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *TransactionsPostgres) GetTransactionById(userId, transactionId int) (project.Transaction, error) {
	var transaction project.Transaction
	query := fmt.Sprintf("SELECT tt.id, ti.sender_id, ti.receiver_id, tt.wallet_id, tt.amount, tt.created_at FROM %s ut INNER JOIN %s tt ON ut.id = tt.sender_id WHERE ut.id = $1 AND tt.id = $2 ORDER BY tt.id ",
		usersTable, transactionsTable)

	if err := r.db.Get(&transaction, query, userId, transactionId); err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *TransactionsPostgres) DeleteTransaction(userId, transactionId int) error {
	query := fmt.Sprintf("DELETE FROM %s tt USING %s ut WHERE tt.sender_id=ut.id AND  ut.id = $1 AND tt.id  = $2",
		transactionsTable, usersTable)
	_, err := r.db.Exec(query, userId, transactionId)

	return err
}
