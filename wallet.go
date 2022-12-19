package project

import "time"

type Wallet struct {
	Id         int       `json:"id"`
	Userid     int       `json:"user_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at""`
	PrivateKey string    `json:"private_key"`
	PublicKey  string    `json:"public_key"`
	Balance    float32   `json:"balance"`
}
