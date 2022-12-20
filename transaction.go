package project

import "time"

type Transaction struct {
	Id         int       `json:"id" db:"id"`
	SenderId   int       `json:"sender_id" db:"sender_id"`
	ReceiverId int       `json:"receiver_id" db:"receiver_id"`
	WalletId   int       `json:"wallet_id" db:"wallet_id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	Amount     float32   `json:"amount" db:"amount"`
}
