package repositories

import (
	"context"
	"gorm.io/gorm"
	"template-go/models/domain"
)

type UserRepository interface {
	Save(ctx context.Context, db *gorm.DB, user domain.User) (domain.User, error)
	Update(ctx context.Context, db *gorm.DB, user domain.User) (domain.User, error)
	Delete(ctx context.Context, db *gorm.DB, user domain.User) error
	FindAll(ctx context.Context, db *gorm.DB) ([]domain.User, error)
	FindById(ctx context.Context, db *gorm.DB, userId int) (domain.User, error)
	FindByEmail(ctx context.Context, db *gorm.DB, userEmail string) (domain.User, error)
}
