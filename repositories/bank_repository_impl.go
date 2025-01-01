package repositories

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"technical-test-go/models/domain"
)

type BankRepositoryImpl struct {
}

func NewBankRepository() *BankRepositoryImpl {
	return &BankRepositoryImpl{}
}

func (repository BankRepositoryImpl) Save(ctx context.Context, tx *sqlx.Tx, bank domain.Bank) (domain.Bank, error) {
	SQL := `INSERT INTO bank_accounts(user_id, balance, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id`

	var id int
	err := tx.QueryRowContext(ctx, SQL, bank.UserId, bank.Balance, bank.CreatedAt, bank.UpdateAt).Scan(&id)
	if err != nil {
		return domain.Bank{}, fmt.Errorf("failed to create bank account: %w", err)
	}

	bank.Id = id
	return bank, nil
}

func (repository BankRepositoryImpl) FindById(ctx context.Context, tx *sqlx.Tx, id int) (domain.Bank, error) {
	SQL := `SELECT id, user_id, balance, created_at, updated_at FROM bank_accounts WHERE id = $1`

	var bank domain.Bank
	err := tx.GetContext(ctx, &bank, SQL, id)
	if err != nil {
		return domain.Bank{}, fmt.Errorf("failed to find bank account: %w", err)
	}

	log.Println(bank)
	return bank, nil
}

func (repository BankRepositoryImpl) FindAll(ctx context.Context, tx *sqlx.Tx) ([]domain.Bank, error) {
	SQL := `SELECT id, user_id, balance, created_at, updated_at FROM bank_accounts`

	var banks []domain.Bank
	err := tx.SelectContext(ctx, &banks, SQL)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve bank accounts: %w", err)
	}

	return banks, nil
}

func (repository BankRepositoryImpl) Update(ctx context.Context, tx *sqlx.Tx, bank domain.Bank) (domain.Bank, error) {
	SQL := `SELECT balance FROM bank_accounts WHERE id = $1 FOR UPDATE`
	var currentBalance int64
	err := tx.GetContext(ctx, &currentBalance, SQL, bank.Id)
	if err != nil {
		return domain.Bank{}, fmt.Errorf("failed to lock bank account: %w", err)
	}

	updateSQL := `UPDATE bank_accounts SET balance = $1, updated_at = $2 WHERE id = $3 RETURNING id, balance, updated_at`
	var updatedBank domain.Bank
	err = tx.GetContext(ctx, &updatedBank, updateSQL, bank.Balance, bank.UpdateAt, bank.Id)
	if err != nil {
		return domain.Bank{}, fmt.Errorf("failed to update bank account: %w", err)
	}

	return updatedBank, nil
}

func (repository BankRepositoryImpl) Delete(ctx context.Context, tx *sqlx.Tx, id int) error {
	SQL := `DELETE FROM bank_accounts WHERE id = $1`

	_, err := tx.ExecContext(ctx, SQL, id)
	if err != nil {
		return fmt.Errorf("failed to delete bank account: %w", err)
	}

	return nil
}

func (repository BankRepositoryImpl) Transfer(ctx context.Context, tx *sqlx.Tx, transfer domain.BankTransfer) error {
	var fromBalance int64
	err := tx.GetContext(ctx, &fromBalance, `SELECT balance FROM bank_accounts WHERE user_id = $1 FOR UPDATE`, transfer.FromAccountId)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to lock sender account: %w", err)
	}

	var toBalance int64
	err = tx.GetContext(ctx, &toBalance, `SELECT balance FROM bank_accounts WHERE user_id = $1 FOR UPDATE`, transfer.ToAccountId)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to lock receiver account: %w", err)
	}

	if fromBalance < transfer.Amount {
		tx.Rollback()
		return fmt.Errorf("insufficient balance in sender's account")
	}

	_, err = tx.ExecContext(ctx, `UPDATE bank_accounts SET balance = balance - $1 WHERE user_id = $2`, transfer.Amount, transfer.FromAccountId)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to debit amount: %w", err)
	}

	_, err = tx.ExecContext(ctx, `UPDATE bank_accounts SET balance = balance + $1 WHERE user_id = $2`, transfer.Amount, transfer.ToAccountId)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to credit amount: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
