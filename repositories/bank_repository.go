package repositories

import (
	"context"
	"github.com/jmoiron/sqlx"
	"technical-test-go/models/domain"
)

type BankRepository interface {
	Save(ctx context.Context, tx *sqlx.Tx, bank domain.Bank) (domain.Bank, error)
	FindById(ctx context.Context, tx *sqlx.Tx, id int) (domain.Bank, error)
	FindAll(ctx context.Context, tx *sqlx.Tx) ([]domain.Bank, error)
	Update(ctx context.Context, tx *sqlx.Tx, bank domain.Bank) (domain.Bank, error)
	Delete(ctx context.Context, tx *sqlx.Tx, id int) error
	Transfer(ctx context.Context, tx *sqlx.Tx, transfer domain.BankTransfer) error
}
