package controllers

import (
	"net/http"
)

type UserController interface {
	Create(writer http.ResponseWriter, request *http.Request)
	Update(writer http.ResponseWriter, request *http.Request)
	Delete(writer http.ResponseWriter, request *http.Request)
	FindAll(writer http.ResponseWriter, request *http.Request)
	FindByIdentifier(writer http.ResponseWriter, request *http.Request)
}
