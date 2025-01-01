package services

import (
	"context"
	"technical-test-go/models/web"
)

type CartService interface {
	Create(ctx context.Context, request web.CartCreateRequest) (web.CartResponse, error)
	Update(ctx context.Context, request web.CartUpdateRequest) (web.CartResponse, error)
	Delete(ctx context.Context, cartId int) error
	FindAll(ctx context.Context) ([]web.CartResponse, error)
	FindById(ctx context.Context, cartId int) (web.CartResponse, error)
	FindByUserId(ctx context.Context, userId int) (web.CartResponse, error)
	Checkout(ctx context.Context, userId int) (web.CartResponse, error)
}
