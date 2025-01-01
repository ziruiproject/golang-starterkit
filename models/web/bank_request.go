package web

import "time"

type BankTransferRequest struct {
	FromAccountId int       `json:"from_account_id"`
	ToAccountId   int       `json:"to_account_id"`
	Amount        int64     `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
}

type BankCreateAccountRequest struct {
	UserId int `json:"user_id"`
}

type BankUpdateRequest struct {
	Balance int64 `json:"balance"`
}
