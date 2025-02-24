package services

import (
	"context"
	"template-go/models/web"
)

type AuthService interface {
	Register(ctx context.Context, request web.UserCreateRequest) (web.UserResponse, error)
	Login(ctx context.Context, request web.LoginRequest) (string, error)
}
