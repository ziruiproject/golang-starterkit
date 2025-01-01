package repositories

import (
	"context"
	"github.com/jmoiron/sqlx"
	"technical-test-go/models/domain"
)

type CartRepository interface {
	Save(ctx context.Context, db *sqlx.DB, Cart domain.Cart) (domain.Cart, error)
	Update(ctx context.Context, db *sqlx.DB, Cart domain.Cart) (domain.Cart, error)
	Delete(ctx context.Context, db *sqlx.DB, cartId int) error
	FindAll(ctx context.Context, db *sqlx.DB) ([]domain.Cart, error)
	FindByUserId(ctx context.Context, db *sqlx.DB, userId int) ([]domain.Cart, error)
	FindById(ctx context.Context, db *sqlx.DB, cartId int) (domain.Cart, error)
}
