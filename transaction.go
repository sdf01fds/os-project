package project

import "time"

type Transaction struct {
	Id         int       `json:"id"`
	SenderId   int       `json:"sender_id"`
	ReceiverId int       `json:"receiver_id"`
	WalletId   int       `json:"wallet_id"`
	CreatedAt  time.Time `json:"created_at"`
	Amount     float32   `json:"amount"`
}
