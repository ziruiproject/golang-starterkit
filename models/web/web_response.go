package web

import (
	"net/http"
)

type WebResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NotFoundResponse(message string) *WebResponse {
	return &WebResponse{
		Code:    http.StatusNotFound,
		Message: message,
		Data:    []interface{}{},
	}
}

func InternalServerResponse(message string) *WebResponse {
	return &WebResponse{
		Code:    http.StatusInternalServerError,
		Message: message,
		Data:    []interface{}{},
	}
}

func BadRequestResponse(message string) *WebResponse {
	return &WebResponse{
		Code:    http.StatusBadRequest,
		Message: message,
		Data:    []interface{}{},
	}
}
