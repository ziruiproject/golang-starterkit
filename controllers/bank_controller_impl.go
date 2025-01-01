package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
	"technical-test-go/helpers"
	"technical-test-go/models/web"
	"technical-test-go/services"
)

type BankControllerImpl struct {
	bankService services.BankService
}

func NewBankController(bankService services.BankService) *BankControllerImpl {
	return &BankControllerImpl{bankService: bankService}
}

func (controller *BankControllerImpl) CreateAccount(writer http.ResponseWriter, request *http.Request) {
	var createRequest web.BankCreateAccountRequest
	err := json.NewDecoder(request.Body).Decode(&createRequest)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	response, err := controller.bankService.CreateAccount(request.Context(), createRequest)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusNotFound, err.Error(), nil)
		return
	}

	helpers.WriteResponse(writer, http.StatusCreated, "Account created successfully", response)
}

func (controller *BankControllerImpl) UpdateAccount(writer http.ResponseWriter, request *http.Request) {
	accountId := chi.URLParam(request, "id")
	id, err := strconv.Atoi(accountId)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusBadRequest, "Invalid account ID", nil)
		return
	}

	var updateRequest web.BankUpdateRequest
	err = json.NewDecoder(request.Body).Decode(&updateRequest)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	response, err := controller.bankService.UpdateAccount(request.Context(), id, updateRequest)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helpers.WriteResponse(writer, http.StatusOK, "Account updated successfully", response)
}

func (controller *BankControllerImpl) DeleteAccount(writer http.ResponseWriter, request *http.Request) {
	accountId := chi.URLParam(request, "id")
	id, err := strconv.Atoi(accountId)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusBadRequest, "Invalid account ID", nil)
		return
	}

	err = controller.bankService.DeleteAccount(request.Context(), id)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helpers.WriteResponse(writer, http.StatusOK, "Account deleted successfully", nil)
}

func (controller *BankControllerImpl) GetAccountById(writer http.ResponseWriter, request *http.Request) {
	accountId := chi.URLParam(request, "id")
	id, err := strconv.Atoi(accountId)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusBadRequest, "Invalid account ID", nil)
		return
	}

	response, err := controller.bankService.GetAccountById(request.Context(), id)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusNotFound, err.Error(), nil)
		return
	}

	helpers.WriteResponse(writer, http.StatusOK, "Account fetched successfully", response)
}

func (controller *BankControllerImpl) GetAllAccounts(writer http.ResponseWriter, request *http.Request) {
	responses, err := controller.bankService.GetAllAccounts(request.Context())
	if err != nil {
		helpers.WriteResponse(writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helpers.WriteResponse(writer, http.StatusOK, "Accounts fetched successfully", responses)
}

func (controller *BankControllerImpl) Transfer(writer http.ResponseWriter, request *http.Request) {
	var transferRequest web.BankTransferRequest
	err := json.NewDecoder(request.Body).Decode(&transferRequest)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	err = controller.bankService.Transfer(request.Context(), transferRequest)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helpers.WriteResponse(writer, http.StatusOK, "Transfer completed successfully", nil)
}
