package repositories

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"technical-test-go/models/domain"
)

type UserRepositoryImpl struct {
	DB *sqlx.DB
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (u *UserRepositoryImpl) Save(ctx context.Context, db *sqlx.DB, user domain.User) (domain.User, error) {
	SQL := "INSERT INTO users(name, email, password) VALUES ($1, $2, $3) RETURNING id"
	err := db.QueryRowxContext(ctx, SQL, user.Name, user.Email, user.Password).Scan(&user.Id)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (u *UserRepositoryImpl) Update(ctx context.Context, db *sqlx.DB, user domain.User) (domain.User, error) {
	SQL := "UPDATE users SET name = $1, email = $2, password = $3 WHERE id = $4"
	_, err := db.ExecContext(ctx, SQL, user.Name, user.Email, user.Password, user.Id)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (u *UserRepositoryImpl) Delete(ctx context.Context, db *sqlx.DB, user domain.User) error {
	SQL := "DELETE FROM users WHERE id = $1"
	_, err := db.ExecContext(ctx, SQL, user.Id)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepositoryImpl) FindAll(ctx context.Context, db *sqlx.DB) ([]domain.User, error) {
	SQL := "SELECT id, name, email, password FROM users"
	var users []domain.User
	err := db.SelectContext(ctx, &users, SQL)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserRepositoryImpl) FindById(ctx context.Context, db *sqlx.DB, userId int) (domain.User, error) {
	SQL := "SELECT id, name, email, password FROM users WHERE id = $1"
	var user domain.User
	err := db.GetContext(ctx, &user, SQL, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, errors.New("user not found")
		}
		return domain.User{}, err
	}
	return user, nil
}

func (u *UserRepositoryImpl) FindByEmail(ctx context.Context, db *sqlx.DB, userEmail string) (domain.User, error) {
	SQL := "SELECT id, name, email, password FROM users WHERE email = $1"
	var user domain.User
	err := db.GetContext(ctx, &user, SQL, userEmail)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, errors.New("user not found")
		}
		return domain.User{}, err
	}
	return user, nil
}
