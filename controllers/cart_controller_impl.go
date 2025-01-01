package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
	"technical-test-go/helpers"
	"technical-test-go/models/web"
	"technical-test-go/services"
)

type CartControllerImpl struct {
	cartService services.CartService
}

// Constructor for CartController
func NewCartController(cartService services.CartService) *CartControllerImpl {
	return &CartControllerImpl{
		cartService: cartService,
	}
}

func (controller *CartControllerImpl) Create(writer http.ResponseWriter, request *http.Request) {
	createRequest := web.CartCreateRequest{
		Quantity: 1,
	}

	err := json.NewDecoder(request.Body).Decode(&createRequest)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	response, err := controller.cartService.Create(request.Context(), createRequest)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helpers.WriteResponse(writer, http.StatusOK, "Cart item created successfully", response)
}

func (controller *CartControllerImpl) Update(writer http.ResponseWriter, request *http.Request) {
	cartId := chi.URLParam(request, "cartId")
	id, err := strconv.Atoi(cartId)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusBadRequest, "Invalid cart ID", nil)
		return
	}

	var updateRequest web.CartUpdateRequest
	updateRequest.Id = id

	err = json.NewDecoder(request.Body).Decode(&updateRequest)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	response, err := controller.cartService.Update(request.Context(), updateRequest)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helpers.WriteResponse(writer, http.StatusOK, "Cart item updated successfully", response)
}

func (controller *CartControllerImpl) Delete(writer http.ResponseWriter, request *http.Request) {
	cartId := chi.URLParam(request, "cartId")
	id, err := strconv.Atoi(cartId)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusBadRequest, "Invalid cart ID", nil)
		return
	}

	err = controller.cartService.Delete(request.Context(), id)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusNotFound, "Cart no found", nil)
		return
	}

	helpers.WriteResponse(writer, http.StatusOK, "Cart item deleted successfully", nil)
}

func (controller *CartControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request) {
	cartsResponse, err := controller.cartService.FindAll(request.Context())
	if err != nil {
		helpers.WriteResponse(writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helpers.WriteResponse(writer, http.StatusOK, "Cart items fetched successfully", cartsResponse)
}

func (controller *CartControllerImpl) FindById(writer http.ResponseWriter, request *http.Request) {
	productId := chi.URLParam(request, "cartId")
	id, err := strconv.Atoi(productId)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response, err := controller.cartService.FindById(request.Context(), id)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusNotFound, err.Error(), nil)
		return
	}

	helpers.WriteResponse(writer, http.StatusOK, "Cart item fetched successfully", response)
}

func (controller *CartControllerImpl) FindByUserId(writer http.ResponseWriter, request *http.Request) {
	userId := chi.URLParam(request, "userId")
	id, err := strconv.Atoi(userId)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}

	responses, err := controller.cartService.FindByUserId(request.Context(), id)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helpers.WriteResponse(writer, http.StatusOK, "Cart items for user fetched successfully", responses)
}

func (controller *CartControllerImpl) Checkout(writer http.ResponseWriter, request *http.Request) {
	userId := chi.URLParam(request, "userId")
	id, err := strconv.Atoi(userId)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}

	response, err := controller.cartService.Checkout(request.Context(), id)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusInternalServerError, fmt.Sprintf("Failed to checkout: %s", err.Error()), nil)
		return
	}

	helpers.WriteResponse(writer, http.StatusOK, "Checkout successful", response)
}
