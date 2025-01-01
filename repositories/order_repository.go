package repositories

import (
	"context"
	"github.com/jmoiron/sqlx"
	"technical-test-go/models/domain"
)

type OrderRepository interface {
	Save(ctx context.Context, tx *sqlx.Tx, order domain.Order) (domain.Order, error)
	FindById(ctx context.Context, tx *sqlx.Tx, orderId int) (domain.OrderWithDetails, error)
	Delete(ctx context.Context, tx *sqlx.Tx, orderId int) error
	SaveDetail(ctx context.Context, tx *sqlx.Tx, detail domain.OrderDetail) (domain.OrderDetail, error)
}
