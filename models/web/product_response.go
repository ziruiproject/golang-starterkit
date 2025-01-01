package web

import "time"

type ProductResponse struct {
	Id          int          `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Price       int          `json:"price"`
	User        UserResponse `json:"user"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}
