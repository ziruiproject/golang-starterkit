package repositories

import (
	"context"
	"github.com/jmoiron/sqlx"
	"technical-test-go/models/domain"
)

type UserRepository interface {
	Save(ctx context.Context, db *sqlx.DB, user domain.User) (domain.User, error)
	Update(ctx context.Context, db *sqlx.DB, user domain.User) (domain.User, error)
	Delete(ctx context.Context, db *sqlx.DB, user domain.User) error
	FindAll(ctx context.Context, db *sqlx.DB) ([]domain.User, error)
	FindById(ctx context.Context, db *sqlx.DB, userId int) (domain.User, error)
	FindByEmail(ctx context.Context, db *sqlx.DB, userEmail string) (domain.User, error)
}
