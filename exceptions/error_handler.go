package exceptions

import (
	"net/http"
	"technical-test-go/models/web"
)

type ResponseError struct {
	Code    int
	Message string
}

func (responseError *ResponseError) Error() int {
	return responseError.Code
}

func NewInternalServerError(err string) ResponseError {
	return ResponseError{
		Code:    http.StatusInternalServerError,
		Message: err,
	}
}

func NewNotFoundError() ResponseError {
	return ResponseError{
		Code:    http.StatusNotFound,
		Message: "Not Found",
	}
}

func NewOK() ResponseError {
	return ResponseError{
		Code:    http.StatusOK,
		Message: "Operation successful",
	}
}

func CheckHTTPError(requestError ResponseError, response *web.WebResponse) bool {
	switch requestError.Code {
	case http.StatusNotFound:
		response = web.NotFoundResponse(requestError.Message)
	case http.StatusBadRequest:
		response = web.BadRequestResponse(requestError.Message)
	case http.StatusInternalServerError:
		response = web.InternalServerResponse(requestError.Message)
	default:
		return false
	}

	return true
}
