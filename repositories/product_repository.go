package repositories

import (
	"context"
	"github.com/jmoiron/sqlx"
	"technical-test-go/models/domain"
)

type ProductRepository interface {
	Save(ctx context.Context, db *sqlx.DB, Product domain.Product) (domain.Product, error)
	Update(ctx context.Context, db *sqlx.DB, Product domain.Product) (domain.Product, error)
	Delete(ctx context.Context, db *sqlx.DB, Product domain.Product) error
	FindAll(ctx context.Context, db *sqlx.DB) ([]domain.Product, error)
	FindByUserId(ctx context.Context, db *sqlx.DB, userId int) ([]domain.Product, error)
	FindBySearch(ctx context.Context, db *sqlx.DB, search string) ([]domain.Product, error)
	FindById(ctx context.Context, db *sqlx.DB, userId int) (domain.Product, error)
}
