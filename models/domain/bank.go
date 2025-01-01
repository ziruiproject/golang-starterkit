package domain

import "time"

type Bank struct {
	Id        int       `db:"id"`
	UserId    int       `db:"user_id"`
	Balance   int64     `db:"balance"`
	CreatedAt time.Time `db:"created_at"`
	UpdateAt  time.Time `db:"updated_at"`
}

type BankTransfer struct {
	FromAccountId int       `db:"from_account_id"`
	ToAccountId   int       `db:"to_account_id"`
	Amount        int64     `db:"amount"`
	CreatedAt     time.Time `db:"created_at"`
}
