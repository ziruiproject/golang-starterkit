package services

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"technical-test-go/models/domain"
	"technical-test-go/models/web"
	"technical-test-go/repositories"
	"time"
)

type productServiceImpl struct {
	ProductRepository repositories.ProductRepository
	UserRepository    repositories.UserRepository
	DB                *sqlx.DB
}

func NewProductService(db *sqlx.DB, productRepository repositories.ProductRepository, userRepository repositories.UserRepository) ProductService {
	return &productServiceImpl{
		ProductRepository: productRepository,
		UserRepository:    userRepository,
		DB:                db,
	}
}

func (service *productServiceImpl) Create(ctx context.Context, request web.ProductCreateRequest) (web.ProductResponse, error) {
	product := domain.Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		UserID:      request.UserID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	createdProduct, err := service.ProductRepository.Save(ctx, service.DB, product)
	if err != nil {
		return web.ProductResponse{}, err
	}

	return service.mapToProductResponse(ctx, createdProduct)
}

func (service *productServiceImpl) Update(ctx context.Context, request web.ProductUpdateRequest) (web.ProductResponse, error) {
	product := domain.Product{
		Id:          request.Id,
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		UpdatedAt:   time.Now(),
	}

	currentProduct, err := service.ProductRepository.FindById(ctx, service.DB, product.Id)
	if err != nil {
		return web.ProductResponse{}, err
	}

	product.UserID = currentProduct.UserID

	updatedProduct, err := service.ProductRepository.Update(ctx, service.DB, product)
	if err != nil {
		return web.ProductResponse{}, err
	}

	return service.mapToProductResponse(ctx, updatedProduct)
}

func (service *productServiceImpl) Delete(ctx context.Context, productId int) error {
	product := domain.Product{Id: productId}
	return service.ProductRepository.Delete(ctx, service.DB, product)
}

func (service *productServiceImpl) FindAll(ctx context.Context) ([]web.ProductResponse, error) {
	products, err := service.ProductRepository.FindAll(ctx, service.DB)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var responses []web.ProductResponse
	for _, product := range products {
		response, err := service.mapToProductResponse(ctx, product)
		if err != nil {
			return nil, err
		}
		responses = append(responses, response)
	}

	return responses, nil
}

func (service *productServiceImpl) FindByUserId(ctx context.Context, userId int) ([]web.ProductResponse, error) {
	products, err := service.ProductRepository.FindByUserId(ctx, service.DB, userId)
	if err != nil {
		return nil, err
	}

	var responses []web.ProductResponse
	for _, product := range products {
		response, err := service.mapToProductResponse(ctx, product)
		if err != nil {
			return nil, err
		}
		responses = append(responses, response)
	}

	return responses, nil
}

func (service *productServiceImpl) FindBySearch(ctx context.Context, search string) ([]web.ProductResponse, error) {
	products, err := service.ProductRepository.FindBySearch(ctx, service.DB, search)
	if err != nil {
		return nil, fmt.Errorf("failed to find products by search: %w", err)
	}

	var responses []web.ProductResponse
	for _, product := range products {
		response, err := service.mapToProductResponse(ctx, product)
		if err != nil {
			return nil, fmt.Errorf("failed to map product to response: %w", err)
		}
		responses = append(responses, response)
	}

	return responses, nil
}

func (service *productServiceImpl) FindById(ctx context.Context, productId int) (web.ProductResponse, error) {
	product, err := service.ProductRepository.FindById(ctx, service.DB, productId)
	if err != nil {
		return web.ProductResponse{}, err
	}

	return service.mapToProductResponse(ctx, product)
}

func (service *productServiceImpl) mapToProductResponse(ctx context.Context, product domain.Product) (web.ProductResponse, error) {
	log.Println("hii")
	user, err := service.UserRepository.FindById(ctx, service.DB, product.UserID)
	if err != nil {
		return web.ProductResponse{}, err
	}

	userResponse := web.UserResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}

	return web.ProductResponse{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		User:        userResponse,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}, nil
}
