package services

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"technical-test-go/models/domain"
	"technical-test-go/models/web"
	"technical-test-go/repositories"
	"time"
)

type BankServiceImpl struct {
	BankRepository repositories.BankRepository
	UserRepository repositories.UserRepository
	DB             *sqlx.DB
}

func NewBankService(db *sqlx.DB, bankRepository repositories.BankRepository, userRepository repositories.UserRepository) *BankServiceImpl {
	return &BankServiceImpl{
		BankRepository: bankRepository,
		UserRepository: userRepository,
		DB:             db,
	}
}

func (service *BankServiceImpl) CreateAccount(ctx context.Context, request web.BankCreateAccountRequest) (web.BankResponse, error) {
	// Start a transaction
	tx, err := service.DB.BeginTxx(ctx, nil)
	if err != nil {
		return web.BankResponse{}, fmt.Errorf("failed to start transaction: %w", err)
	}

	// Ensure the transaction is committed or rolled back
	defer func() {
		if err != nil {
			tx.Rollback() // Rollback on error
		} else {
			err = tx.Commit() // Commit if no errors
		}
	}()

	// Create the bank account
	bank := domain.Bank{
		UserId:    request.UserId,
		Balance:   0,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}

	// Fetch the user related to the bank account
	user, err := service.UserRepository.FindById(ctx, service.DB, request.UserId)
	if err != nil {
		return web.BankResponse{}, fmt.Errorf("failed to fetch user: %w", err)
	}

	createdBank, err := service.BankRepository.Save(ctx, tx, bank)
	if err != nil {
		return web.BankResponse{}, fmt.Errorf("failed to create account: %w", err)
	}

	// Return the response with the newly created bank account and user details
	return web.BankResponse{
		Id: createdBank.Id,
		User: web.UserResponse{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		},
		Balance:   createdBank.Balance,
		UpdatedAt: createdBank.UpdateAt,
	}, nil
}

func (service *BankServiceImpl) UpdateAccount(ctx context.Context, id int, request web.BankUpdateRequest) (web.BankResponse, error) {
	// Start a transaction
	tx, err := service.DB.BeginTxx(ctx, nil)
	if err != nil {
		return web.BankResponse{}, fmt.Errorf("failed to start transaction: %w", err)
	}

	// Ensure the transaction is committed or rolled back
	defer func() {
		if err != nil {
			tx.Rollback() // Rollback on error
		} else {
			err = tx.Commit() // Commit if no errors
		}
	}()

	// Fetch the bank account
	bank, err := service.BankRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.BankResponse{}, fmt.Errorf("account not found: %w", err)
	}

	// Update the bank account balance
	bank.Balance = (bank.Balance + request.Balance)
	bank.UpdateAt = time.Now()

	if bank.Balance < 0 {
		return web.BankResponse{}, fmt.Errorf("not enough balance")
	}

	// Update the bank account in the database
	updatedBank, err := service.BankRepository.Update(ctx, tx, bank)
	if err != nil {
		return web.BankResponse{}, fmt.Errorf("failed to update account: %w", err)
	}

	// Fetch the user related to the bank account
	user, err := service.UserRepository.FindById(ctx, service.DB, bank.UserId)
	if err != nil {
		return web.BankResponse{}, fmt.Errorf("failed to fetch user: %w", err)
	}

	// Return the response with the updated bank account and user details
	return web.BankResponse{
		Id: updatedBank.Id,
		User: web.UserResponse{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		},
		Balance:   updatedBank.Balance,
		UpdatedAt: updatedBank.UpdateAt,
	}, nil
}

func (service *BankServiceImpl) DeleteAccount(ctx context.Context, id int) error {
	// Start a transaction
	tx, err := service.DB.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	// Ensure the transaction is committed or rolled back
	defer func() {
		if err != nil {
			tx.Rollback() // Rollback on error
		} else {
			err = tx.Commit() // Commit if no errors
		}
	}()

	err = service.BankRepository.Delete(ctx, tx, id)
	if err != nil {
		return fmt.Errorf("failed to delete account: %w", err)
	}

	return nil
}

func (service *BankServiceImpl) GetAccountById(ctx context.Context, id int) (web.BankResponse, error) {
	// Start a transaction
	tx, err := service.DB.BeginTxx(ctx, nil)
	if err != nil {
		return web.BankResponse{}, fmt.Errorf("failed to start transaction: %w", err)
	}

	// Ensure the transaction is committed or rolled back
	defer func() {
		if err != nil {
			tx.Rollback() // Rollback on error
		} else {
			err = tx.Commit() // Commit if no errors
		}
	}()

	bank, err := service.BankRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.BankResponse{}, fmt.Errorf("account not found: %w", err)
	}

	user, err := service.UserRepository.FindById(ctx, service.DB, bank.UserId)
	if err != nil {
		return web.BankResponse{}, fmt.Errorf("failed to fetch user: %w", err)
	}

	return web.BankResponse{
		Id: bank.Id,
		User: web.UserResponse{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		},
		Balance:   bank.Balance,
		UpdatedAt: bank.UpdateAt,
	}, nil
}

func (service *BankServiceImpl) GetAllAccounts(ctx context.Context) ([]web.BankResponse, error) {
	// Start a transaction
	tx, err := service.DB.BeginTxx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}

	// Ensure the transaction is committed or rolled back
	defer func() {
		if err != nil {
			tx.Rollback() // Rollback on error
		} else {
			err = tx.Commit() // Commit if no errors
		}
	}()

	banks, err := service.BankRepository.FindAll(ctx, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch accounts: %w", err)
	}

	var responses []web.BankResponse
	for _, bank := range banks {
		user, err := service.UserRepository.FindById(ctx, service.DB, bank.UserId)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch user for account ID %d: %w", bank.Id, err)
		}

		responses = append(responses, web.BankResponse{
			Id: bank.Id,
			User: web.UserResponse{
				Id:    user.Id,
				Name:  user.Name,
				Email: user.Email,
			},
			Balance:   bank.Balance,
			UpdatedAt: bank.UpdateAt,
		})
	}

	return responses, nil
}

func (service *BankServiceImpl) Transfer(ctx context.Context, request web.BankTransferRequest) error {
	// Start a transaction
	tx, err := service.DB.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	// Ensure the transaction is committed or rolled back
	defer func() {
		if err != nil {
			tx.Rollback() // Rollback on error
		} else {
			err = tx.Commit() // Commit if no errors
		}
	}()

	transfer := domain.BankTransfer{
		FromAccountId: request.FromAccountId,
		ToAccountId:   request.ToAccountId,
		Amount:        request.Amount,
		CreatedAt:     time.Now(),
	}

	err = service.BankRepository.Transfer(ctx, tx, transfer)
	if err != nil {
		return fmt.Errorf("failed to process transfer: %w", err)
	}

	return nil
}
