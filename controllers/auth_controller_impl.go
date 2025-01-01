package controllers

import (
	"encoding/json"
	"net/http"
	"technical-test-go/helpers"
	"technical-test-go/models/web"
	"technical-test-go/services"
)

type AuthControllerImpl struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) AuthController {
	return &AuthControllerImpl{
		authService: authService,
	}
}

func (controller *AuthControllerImpl) Register(writer http.ResponseWriter, request *http.Request) {
	var userCreateRequest web.UserCreateRequest

	err := json.NewDecoder(request.Body).Decode(&userCreateRequest)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	userResponse, err := controller.authService.Register(request.Context(), userCreateRequest)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helpers.WriteResponse(writer, http.StatusCreated, "User registered successfully", userResponse)
}

func (controller *AuthControllerImpl) Login(writer http.ResponseWriter, request *http.Request) {
	var loginRequest web.LoginRequest

	err := json.NewDecoder(request.Body).Decode(&loginRequest)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	token, err := controller.authService.Login(request.Context(), loginRequest)
	if err != nil {
		helpers.WriteResponse(writer, http.StatusUnauthorized, err.Error(), nil)
		return
	}

	helpers.WriteResponse(writer, http.StatusOK, "Login successful", map[string]string{"token": token})
}

func (controller *AuthControllerImpl) Logout(writer http.ResponseWriter, request *http.Request) {
	helpers.WriteResponse(writer, http.StatusOK, "Logout successful", nil)
}
