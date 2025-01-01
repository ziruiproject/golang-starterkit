package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"strconv"
	"technical-test-go/helpers"
	"technical-test-go/models/web"
	"technical-test-go/services"
	"technical-test-go/utils"
)

type ProductControllerImpl struct {
	productService services.ProductService
}

func NewProductController(productService services.ProductService) ProductController {
	return &ProductControllerImpl{
		productService: productService,
	}
}

func (controller *ProductControllerImpl) Create(writer http.ResponseWriter, request *http.Request) {
	var productRequest web.ProductCreateRequest

	err := json.NewDecoder(request.Body).Decode(&productRequest)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	userId, err := utils.GetUserIDFromJWT(request)
	productRequest.UserID = userId

	productResponse, err := controller.productService.Create(request.Context(), productRequest)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helpers.WriteResponse(writer, http.StatusCreated, "Product created successfully", productResponse)
}

func (controller *ProductControllerImpl) Update(writer http.ResponseWriter, request *http.Request) {
	var productRequest web.ProductUpdateRequest
	productId := chi.URLParam(request, "productId")

	id, err := strconv.Atoi(productId)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusBadRequest, "Invalid product ID", nil)
		return
	}

	productRequest.Id = id

	err = json.NewDecoder(request.Body).Decode(&productRequest)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	productResponse, err := controller.productService.Update(request.Context(), productRequest)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helpers.WriteResponse(writer, http.StatusOK, "Product updated successfully", productResponse)
}

func (controller *ProductControllerImpl) Delete(writer http.ResponseWriter, request *http.Request) {
	productId := chi.URLParam(request, "productId")
	id, err := strconv.Atoi(productId)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusBadRequest, "Invalid product ID", nil)
		return
	}

	err = controller.productService.Delete(request.Context(), id)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusNotFound, "No Record Found", []interface{}{})
		return
	}

	helpers.WriteResponse(writer, http.StatusOK, "Product deleted successfully", nil)
}

func (controller *ProductControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request) {
	products, err := controller.productService.FindAll(request.Context())
	if err != nil {
		helpers.WriteResponse(writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helpers.WriteResponse(writer, http.StatusOK, "Products fetched successfully", products)
}

func (controller *ProductControllerImpl) FindById(writer http.ResponseWriter, request *http.Request) {
	productId := chi.URLParam(request, "productId")
	log.Println(productId)
	id, err := strconv.Atoi(productId)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusBadRequest, "Invalid product ID", nil)
		return
	}

	productResponse, err := controller.productService.FindById(request.Context(), id)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusNotFound, "No Record Found", []interface{}{})
		return
	}

	helpers.WriteResponse(writer, http.StatusOK, "Product fetched successfully", productResponse)
}

func (controller *ProductControllerImpl) FindBySearch(writer http.ResponseWriter, request *http.Request) {
	search := request.URL.Query().Get("product")

	if search == "" {
		helpers.WriteResponse(writer, http.StatusBadRequest, "Search query is required", nil)
		return
	}

	products, err := controller.productService.FindBySearch(request.Context(), search)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusNotFound, err.Error(), []interface{}{})
		return
	}

	helpers.WriteResponse(writer, http.StatusOK, "Products fetched successfully", products)
}

func (controller *ProductControllerImpl) FindByUserId(writer http.ResponseWriter, request *http.Request) {
	userId := chi.URLParam(request, "userId")
	id, err := strconv.Atoi(userId)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}

	products, err := controller.productService.FindByUserId(request.Context(), id)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusNotFound, err.Error(), []interface{}{})
		return
	}

	helpers.WriteResponse(writer, http.StatusOK, "Products fetched successfully", products)
}
