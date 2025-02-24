package services

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"template-go/auth"
	"template-go/models/domain"
	"template-go/models/web"
	"template-go/repositories"
)

type AuthServiceImpl struct {
	userRepo repositories.UserRepository
	DB       *gorm.DB
}

func NewAuthService(db *gorm.DB, userRepo repositories.UserRepository) AuthService {
	return &AuthServiceImpl{
		userRepo: userRepo,
		DB:       db,
	}
}

func (service *AuthServiceImpl) Register(ctx context.Context, request web.UserCreateRequest) (web.UserResponse, error) {
	securedPassword, err := auth.HashPassword(request.Password)
	if err != nil {
		return web.UserResponse{}, err
	}

	user := domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: securedPassword,
	}

	user, err = service.userRepo.Save(ctx, service.DB, user)
	if err != nil {
		return web.UserResponse{}, err
	}

	return web.UserResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (service *AuthServiceImpl) Login(ctx context.Context, request web.LoginRequest) (string, error) {
	user, err := service.userRepo.FindByEmail(ctx, service.DB, request.Email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if !auth.VerifyPassword(request.Password, user.Password) {
		return "", errors.New("invalid email or password")
	}

	token, err := auth.GenerateJWT(user.Id, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
