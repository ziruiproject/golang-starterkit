package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"template-go/middlewares"
)

func GetUserIDFromJWT(r *http.Request) (int, error) {
	userClaims, ok := r.Context().Value(middlewares.UserContextKey).(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("invalid or missing user claims")
	}

	userID, ok := userClaims["user_id"].(float64)
	if !ok {
		return 0, fmt.Errorf("user_id not found in claims")
	}

	return int(userID), nil
}
