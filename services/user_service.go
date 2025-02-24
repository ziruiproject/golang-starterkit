package services

import (
	"context"
	"template-go/models/web"
)

type UserService interface {
	Create(ctx context.Context, request web.UserCreateRequest) (web.UserResponse, error)
	Update(ctx context.Context, request web.UserUpdateRequest) (web.UserResponse, error)
	Delete(ctx context.Context, userId int) error
	FindAll(ctx context.Context) ([]web.UserResponse, error)
	FindById(ctx context.Context, userId int) (web.UserResponse, error)
	FindByEmail(ctx context.Context, userEmail string) (web.UserResponse, error)
}
