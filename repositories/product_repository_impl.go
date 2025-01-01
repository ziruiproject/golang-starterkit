package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"strings"
	"technical-test-go/models/domain"
	"time"
)

type productRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &productRepositoryImpl{}
}

func (repository *productRepositoryImpl) Save(ctx context.Context, db *sqlx.DB, product domain.Product) (domain.Product, error) {

	log.Println(product)
	SQL := `INSERT INTO products (name, description, price, user_id, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := db.QueryRowxContext(ctx, SQL, product.Name, product.Description, product.Price, product.UserID, time.Now(), time.Now()).Scan(&product.Id)
	if err != nil {
		return domain.Product{}, fmt.Errorf("failed to save product: %w", err)
	}

	return product, nil
}

func (repository *productRepositoryImpl) Update(ctx context.Context, db *sqlx.DB, product domain.Product) (domain.Product, error) {
	SQL := `UPDATE products
			SET name = $1, description = $2, price = $3, updated_at = $4
			WHERE id = $5`
	_, err := db.ExecContext(ctx, SQL, product.Name, product.Description, product.Price, product.UpdatedAt, product.Id)
	if err != nil {
		return domain.Product{}, fmt.Errorf("failed to update product: %w", err)
	}
	return product, nil
}

func (repository *productRepositoryImpl) Delete(ctx context.Context, db *sqlx.DB, product domain.Product) error {
	SQL := `DELETE FROM products WHERE id = $1`
	_, err := db.ExecContext(ctx, SQL, product.Id)
	if err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}
	return nil
}

func (repository *productRepositoryImpl) FindAll(ctx context.Context, db *sqlx.DB) ([]domain.Product, error) {
	SQL := `SELECT * FROM products`
	var products []domain.Product
	err := db.SelectContext(ctx, &products, SQL)
	if err != nil {
		return nil, fmt.Errorf("failed to find all products: %w", err)
	}
	return products, nil
}

func (repository *productRepositoryImpl) FindByUserId(ctx context.Context, db *sqlx.DB, userId int) ([]domain.Product, error) {
	SQL := `SELECT id, name, description, price, user_id, created_at, updated_at FROM products WHERE user_id = $1`
	var products []domain.Product
	err := db.SelectContext(ctx, &products, SQL, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to find products by user id: %w", err)
	}
	return products, nil
}

func (repository *productRepositoryImpl) FindBySearch(ctx context.Context, db *sqlx.DB, search string) ([]domain.Product, error) {
	SQL := `SELECT id, name, description, price, user_id, created_at, updated_at
			FROM products
			WHERE name ILIKE $1 OR description ILIKE $2`

	var products []domain.Product

	search = strings.Trim(search, `"`)
	search = strings.TrimSpace(search)

	searchPattern := fmt.Sprintf("%%%s%%", search)
	log.Printf("Search query: %s", searchPattern)

	err := db.SelectContext(ctx, &products, SQL, searchPattern, searchPattern)
	if err != nil {
		return nil, fmt.Errorf("failed to search products: %w", err)
	}

	if len(products) == 0 {
		return nil, errors.New("Product Not Found")
	}

	log.Printf("Found %d products", len(products))
	return products, nil
}

func (repository *productRepositoryImpl) FindById(ctx context.Context, db *sqlx.DB, Id int) (domain.Product, error) {
	SQL := `SELECT id, name, description, price, user_id, created_at, updated_at FROM products WHERE id = $1`
	var product domain.Product
	err := db.GetContext(ctx, &product, SQL, Id)
	if err != nil {
		return domain.Product{}, fmt.Errorf("failed to find product by id: %w", err)
	}
	return product, nil
}
