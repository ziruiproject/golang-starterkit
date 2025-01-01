package controllers

import (
	"net/http"
)

type ProductController interface {
	Create(writer http.ResponseWriter, request *http.Request)
	Update(writer http.ResponseWriter, request *http.Request)
	Delete(writer http.ResponseWriter, request *http.Request)
	FindAll(writer http.ResponseWriter, request *http.Request)
	FindById(writer http.ResponseWriter, request *http.Request)
	FindBySearch(writer http.ResponseWriter, request *http.Request)
	FindByUserId(writer http.ResponseWriter, request *http.Request)
}
