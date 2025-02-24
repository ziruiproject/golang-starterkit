package helpers

import (
	"encoding/json"
	"net/http"
	"template-go/models/web"
)

func WriteResponse(writer http.ResponseWriter, code int, message string, data interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	response := web.WebResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		http.Error(writer, "Failed to encode response", http.StatusInternalServerError)
	}
}
