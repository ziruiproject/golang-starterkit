package web

import "time"

type BankResponse struct {
	Id        int          `json:"id"`
	User      UserResponse `json:"user"`
	Balance   int64        `json:"balance"`
	UpdatedAt time.Time    `json:"updated_at"`
}

type BankTransfer struct {
	Id          int          `json:"id"`
	FromAccount UserResponse `json:"from_account"`
	ToAccount   UserResponse `json:"to_account"`
	Amount      int64        `json:"amount"`
	CreatedAt   time.Time    `json:"created_at"`
}
