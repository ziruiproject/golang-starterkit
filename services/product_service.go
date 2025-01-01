package services

import (
	"context"
	"technical-test-go/models/web"
)

type ProductService interface {
	Create(ctx context.Context, request web.ProductCreateRequest) (web.ProductResponse, error)
	Update(ctx context.Context, request web.ProductUpdateRequest) (web.ProductResponse, error)
	Delete(ctx context.Context, ProductId int) error
	FindAll(ctx context.Context) ([]web.ProductResponse, error)
	FindById(ctx context.Context, ProductId int) (web.ProductResponse, error)
	FindBySearch(ctx context.Context, search string) ([]web.ProductResponse, error)
	FindByUserId(ctx context.Context, ProductId int) ([]web.ProductResponse, error)
}
