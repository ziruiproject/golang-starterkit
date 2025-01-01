package domain

import (
	"time"
)

// Order represents the 'orders' table in the database
type Order struct {
	Id        int       `db:"id"`
	UserID    int       `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// OrderDetail represents the 'order_details' table in the database
type OrderDetail struct {
	Id        int   `db:"id"`
	OrderID   int   `db:"order_id"`
	ProductID int   `db:"product_id"`
	Quantity  int   `db:"quantity"`
	Price     int64 `db:"price"`
}

// OrderWithDetails represents an order along with its details
type OrderWithDetails struct {
	Order        Order         `db:"order"`
	OrderDetails []OrderDetail `db:"order_details"`
}
