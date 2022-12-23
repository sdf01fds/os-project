package project

import "time"

type Wallet struct {
	Id         int       `json:"id" db:"id"`
	UserId     int       `json:"user_id" db:"user_id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
	PrivateKey string    `json:"private_key" db:"private_key"`
	PublicKey  string    `json:"public_key" db:"public_key"`
	Balance    float32   `json:"balance" db:"balance"`
}

type UpdateWallet struct {
	Amount float32 `json:"amount"`
}
