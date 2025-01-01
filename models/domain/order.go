package domain

import (
	"time"
)

type Order struct {
	Id        int       `db:"id"`
	UserID    int       `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type OrderDetail struct {
	Id        int   `db:"id"`
	OrderID   int   `db:"order_id"`
	ProductID int   `db:"product_id"`
	Quantity  int   `db:"quantity"`
	Price     int64 `db:"price"`
}

type OrderWithDetails struct {
	Order        Order         `db:"order"`
	OrderDetails []OrderDetail `db:"order_details"`
}
