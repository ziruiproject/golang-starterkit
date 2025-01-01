package repositories

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"technical-test-go/models/domain"
)

type OrderRepositoryImpl struct {
}

func NewOrderRepository() *OrderRepositoryImpl {
	return &OrderRepositoryImpl{}
}

func (repository *OrderRepositoryImpl) Save(ctx context.Context, tx *sqlx.Tx, order domain.Order) (domain.Order, error) {
	SQL := `INSERT INTO orders (user_id) VALUES ($1) RETURNING id, created_at, updated_at`
	err := tx.GetContext(ctx, &order, SQL, order.UserID)
	if err != nil {
		return domain.Order{}, fmt.Errorf("failed to save order: %w", err)
	}
	return order, nil
}

func (repository *OrderRepositoryImpl) SaveDetail(ctx context.Context, tx *sqlx.Tx, detail domain.OrderDetail) (domain.OrderDetail, error) {
	orderDetailsSQL := `INSERT INTO order_details (order_id, product_id, quantity, price)  
						VALUES ($1, $2, $3, $4) RETURNING *`
	_, err := tx.ExecContext(ctx, orderDetailsSQL, detail.OrderID, detail.ProductID, detail.Quantity, detail.Price)
	if err != nil {
		return domain.OrderDetail{}, fmt.Errorf("failed to create order details: %w", err)
	}
	return detail, nil
}

func (repository *OrderRepositoryImpl) FindById(ctx context.Context, tx *sqlx.Tx, orderId int) (domain.OrderWithDetails, error) {
	var order domain.Order
	orderSQL := `SELECT * FROM orders WHERE id = $1`
	err := tx.GetContext(ctx, &order, orderSQL, orderId)
	if err != nil {
		return domain.OrderWithDetails{}, fmt.Errorf("failed to find order: %w", err)
	}

	var orderDetails []domain.OrderDetail
	detailsSQL := `SELECT * FROM order_details WHERE order_id = $1`
	err = tx.SelectContext(ctx, &orderDetails, detailsSQL, orderId)
	if err != nil {
		return domain.OrderWithDetails{}, fmt.Errorf("failed to fetch order details: %w", err)
	}

	return domain.OrderWithDetails{
		Order:        order,
		OrderDetails: orderDetails,
	}, nil
}

func (repository *OrderRepositoryImpl) Delete(ctx context.Context, tx *sqlx.Tx, orderId int) error {
	SQL := `DELETE FROM orders WHERE id = $1`
	_, err := tx.ExecContext(ctx, SQL, orderId)
	if err != nil {
		return fmt.Errorf("failed to delete order: %w", err)
	}
	return nil
}
