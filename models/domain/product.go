package domain

import "time"

type Product struct {
	Id          int       `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Price       int       `db:"price"`
	UserID      int       `db:"user_id"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
