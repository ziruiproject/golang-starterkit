package services

import (
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
	"technical-test-go/auth"
	"technical-test-go/models/domain"
	"technical-test-go/models/web"
	"technical-test-go/repositories"
)

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
	DB             *sqlx.DB
}

func NewUserService(db *sqlx.DB, userRepository repositories.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, request web.UserCreateRequest) (web.UserResponse, error) {
	securedPassword, err := auth.HashPassword(request.Password)
	if err != nil {
		return web.UserResponse{}, err
	}

	user := domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: securedPassword,
	}

	user, err = service.UserRepository.Save(ctx, service.DB, user)
	if err != nil {
		return web.UserResponse{}, err
	}

	response := web.UserResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}

	return response, nil
}

func (service *UserServiceImpl) Update(ctx context.Context, request web.UserUpdateRequest) (web.UserResponse, error) {
	user, err := service.UserRepository.FindById(ctx, service.DB, request.Id)
	if err != nil {
		return web.UserResponse{}, err
	}

	if request.Name != "" {
		user.Name = request.Name
	}
	if request.Email != "" {
		user.Email = request.Email
	}

	user, err = service.UserRepository.Update(ctx, service.DB, user)
	if err != nil {
		return web.UserResponse{}, err
	}

	response := web.UserResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}

	return response, nil
}

func (service *UserServiceImpl) Delete(ctx context.Context, userId int) error {
	user, err := service.UserRepository.FindById(ctx, service.DB, userId)
	if err != nil {
		return err
	}

	err = service.UserRepository.Delete(ctx, service.DB, user)
	if err != nil {
		return err
	}

	return nil
}

func (service *UserServiceImpl) FindAll(ctx context.Context) ([]web.UserResponse, error) {
	users, err := service.UserRepository.FindAll(ctx, service.DB)
	if err != nil {
		return nil, err
	}

	var response []web.UserResponse
	for _, user := range users {
		response = append(response, web.UserResponse{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return response, nil
}

func (service *UserServiceImpl) FindById(ctx context.Context, userId int) (web.UserResponse, error) {
	user, err := service.UserRepository.FindById(ctx, service.DB, userId)
	if err != nil {
		if err.Error() == "user not found" {
			return web.UserResponse{}, errors.New("user not found")
		}
		return web.UserResponse{}, err
	}

	response := web.UserResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}

	return response, nil
}

func (service *UserServiceImpl) FindByEmail(ctx context.Context, userEmail string) (web.UserResponse, error) {
	user, err := service.UserRepository.FindByEmail(ctx, service.DB, userEmail)
	if err != nil {
		if err.Error() == "user not found" {
			return web.UserResponse{}, errors.New("user not found")
		}
		return web.UserResponse{}, err
	}

	response := web.UserResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}

	return response, nil
}
