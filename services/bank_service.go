package services

import (
	"context"
	"technical-test-go/models/web"
)

type BankService interface {
	CreateAccount(ctx context.Context, request web.BankCreateAccountRequest) (web.BankResponse, error)
	UpdateAccount(ctx context.Context, id int, request web.BankUpdateRequest) (web.BankResponse, error)
	DeleteAccount(ctx context.Context, id int) error
	GetAccountById(ctx context.Context, id int) (web.BankResponse, error)
	GetAllAccounts(ctx context.Context) ([]web.BankResponse, error)
	Transfer(ctx context.Context, request web.BankTransferRequest) error
}
