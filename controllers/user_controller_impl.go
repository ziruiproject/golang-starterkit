package controllers

import (
	"encoding/json"
	"github.com/ggicci/httpin"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
	"template-go/commons/config"
	"template-go/helpers"
	"template-go/models/web"
	"template-go/services"
)

type UserControllerImpl struct {
	userService services.UserService
	sanitation  config.Sanitation
}

func NewUserController(userService services.UserService, sanitation config.Sanitation) UserController {
	return &UserControllerImpl{
		userService: userService,
		sanitation:  sanitation,
	}
}

func (controller *UserControllerImpl) Create(writer http.ResponseWriter, request *http.Request) {
	var userCreateRequest web.UserCreateRequest

	err := json.NewDecoder(request.Body).Decode(&userCreateRequest)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	err = controller.sanitation.Validator.Struct(&userCreateRequest)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusUnprocessableEntity, "Validation error", nil)
		return
	}
	controller.sanitation.Sanitizer.Sanitize(&userCreateRequest)

	userResponse, err := controller.userService.Create(request.Context(), userCreateRequest)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	helpers.WriteResponse(writer, http.StatusCreated, "User created successfully", userResponse)
}

func (controller *UserControllerImpl) Update(writer http.ResponseWriter, request *http.Request) {
	var userUpdateRequest web.UserUpdateRequest

	err := json.NewDecoder(request.Body).Decode(&userUpdateRequest)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusBadRequest, err.Error(), nil)
		return
	}

	err = controller.sanitation.Validator.Struct(&userUpdateRequest)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}
	controller.sanitation.Sanitizer.Sanitize(&userUpdateRequest)

	userId := chi.URLParam(request, "userId")
	id, err := strconv.Atoi(userId)
	userUpdateRequest.Id = id

	userResponse, err := controller.userService.Update(request.Context(), userUpdateRequest)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	helpers.WriteResponse(writer, http.StatusOK, "User updated successfully", userResponse)
}

func (controller *UserControllerImpl) Delete(writer http.ResponseWriter, request *http.Request) {
	userId := chi.URLParam(request, "userId")
	id, err := strconv.Atoi(userId)

	err = controller.userService.Delete(request.Context(), id)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusNotFound, "No Record Found", []interface{}{})
		return
	}

	helpers.WriteResponse(writer, http.StatusOK, "User deleted successfully", nil)
}

func (controller *UserControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request) {
	params := request.Context().Value(httpin.Input).(*web.DefaultParams)
	err := controller.sanitation.Validator.Struct(params)

	if err != nil {
		helpers.WriteResponse(writer, http.StatusUnprocessableEntity, "Validation error", err)
	}
	usersResponse, err := controller.userService.FindAll(request.Context())
	if err != nil {
		helpers.WriteResponse(writer, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	helpers.WriteResponse(writer, http.StatusOK, "Users fetched successfully", usersResponse)
}

func (controller *UserControllerImpl) FindByIdentifier(writer http.ResponseWriter, request *http.Request) {
	identifier := chi.URLParam(request, "identifier")

	var userResponse web.UserResponse
	var err error

	userId, errId := strconv.Atoi(identifier)
	if errId != nil {
		userResponse, err = controller.userService.FindByEmail(request.Context(), identifier)
	} else {
		userResponse, err = controller.userService.FindById(request.Context(), userId)
	}

	if err != nil {
		helpers.WriteResponse(writer, http.StatusNotFound, err.Error(), []interface{}{})
		return
	}

	helpers.WriteResponse(writer, http.StatusOK, "User fetched successfully", userResponse)
}
