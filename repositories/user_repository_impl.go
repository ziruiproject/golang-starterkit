package repositories

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"template-go/models/domain"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (u *UserRepositoryImpl) Save(ctx context.Context, db *gorm.DB, user domain.User) (domain.User, error) {
	err := db.WithContext(ctx).Create(&user).Error

	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (u *UserRepositoryImpl) Update(ctx context.Context, db *gorm.DB, user domain.User) (domain.User, error) {
	err := db.WithContext(ctx).Save(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.User{}, errors.New("user not found")
	}

	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (u *UserRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, user domain.User) error {
	err := db.WithContext(ctx).Delete(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("user not found")
	}

	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB) ([]domain.User, error) {
	var users []domain.User
	err := db.WithContext(ctx).Find(&users).Error

	if err != nil {
		return []domain.User{}, err
	}
	return users, nil
}

func (u *UserRepositoryImpl) FindById(ctx context.Context, db *gorm.DB, userId int) (domain.User, error) {
	var user domain.User
	err := db.WithContext(ctx).First(&user, "id = ?", userId).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.User{}, errors.New("user not found")
	}

	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (u *UserRepositoryImpl) FindByEmail(ctx context.Context, db *gorm.DB, userEmail string) (domain.User, error) {
	var user domain.User
	err := db.WithContext(ctx).First(&user, "email = ?", userEmail).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.User{}, errors.New("user not found")
	}

	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
