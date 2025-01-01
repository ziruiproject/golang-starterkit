package repositories

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"technical-test-go/models/domain"
)

type CartRepositoryImpl struct {
	DB *sqlx.DB
}

func NewCartRepository() *CartRepositoryImpl {
	return &CartRepositoryImpl{}
}

func (repository *CartRepositoryImpl) Save(ctx context.Context, db *sqlx.DB, cart domain.Cart) (domain.Cart, error) {
	SQL := `INSERT INTO carts (user_id, product_id, quantity) 
			VALUES ($1, $2, $3)`
	_, err := db.ExecContext(ctx, SQL, cart.UserId, cart.ProductId, cart.Quantity)
	if err != nil {
		return domain.Cart{}, fmt.Errorf("failed to save cart: %w", err)
	}
	return cart, nil
}

func (repository *CartRepositoryImpl) Update(ctx context.Context, db *sqlx.DB, cart domain.Cart) (domain.Cart, error) {
	SQL := `UPDATE carts 
			SET quantity = $1 
			WHERE id = $2`
	_, err := db.ExecContext(ctx, SQL, cart.Quantity, cart.Id)
	if err != nil {
		return domain.Cart{}, fmt.Errorf("failed to update cart: %w", err)
	}
	return cart, nil
}

func (repository *CartRepositoryImpl) Delete(ctx context.Context, db *sqlx.DB, cartId int) error {
	SQL := `DELETE FROM carts 
			WHERE id = $1`
	_, err := db.ExecContext(ctx, SQL, cartId)
	if err != nil {
		return fmt.Errorf("failed to delete cart: %w", err)
	}
	return nil
}

func (repository *CartRepositoryImpl) FindAll(ctx context.Context, db *sqlx.DB) ([]domain.Cart, error) {
	SQL := `SELECT * 
			FROM carts`
	var carts []domain.Cart
	err := db.SelectContext(ctx, &carts, SQL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all carts: %w", err)
	}
	return carts, nil
}

func (repository *CartRepositoryImpl) FindByUserId(ctx context.Context, db *sqlx.DB, userId int) ([]domain.Cart, error) {
	SQL := `SELECT * FROM carts 
			WHERE user_id = $1`
	var carts []domain.Cart
	err := db.SelectContext(ctx, &carts, SQL, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch carts for user ID %d: %w", userId, err)
	}
	return carts, nil
}

func (repository *CartRepositoryImpl) FindById(ctx context.Context, db *sqlx.DB, cartId int) (domain.Cart, error) {
	SQL := `SELECT * 
			FROM carts 
			WHERE id = $1`
	var cart domain.Cart
	err := db.GetContext(ctx, &cart, SQL, cartId)
	if err != nil {
		return domain.Cart{}, fmt.Errorf("failed to find cart for user ID %d and product ID %d: %w", cartId, err)
	}
	return cart, nil
}
