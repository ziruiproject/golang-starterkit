package controllers

import (
	"net/http"
)

type BankController interface {
	CreateAccount(writer http.ResponseWriter, request *http.Request)
	UpdateAccount(writer http.ResponseWriter, request *http.Request)
	DeleteAccount(writer http.ResponseWriter, request *http.Request)
	GetAccountById(writer http.ResponseWriter, request *http.Request)
	GetAllAccounts(writer http.ResponseWriter, request *http.Request)
	Transfer(writer http.ResponseWriter, request *http.Request)
}
