package middlewares

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
	"technical-test-go/auth"
	"technical-test-go/helpers"
)

type contextKey string

const UserContextKey contextKey = "userClaims"

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			helpers.WriteResponse(w, http.StatusUnauthorized, "Missing Authorization header", nil)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			helpers.WriteResponse(w, http.StatusUnauthorized, "Invalid Authorization header format", nil)
			return
		}

		token, err := auth.ValidateJWT(tokenString)
		if err != nil || !token.Valid {
			helpers.WriteResponse(w, http.StatusUnauthorized, "Invalid or expired token", nil)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			ctx := context.WithValue(r.Context(), UserContextKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			helpers.WriteResponse(w, http.StatusUnauthorized, "Invalid token claims", nil)
		}
	})
}

func ProtectedRoute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		JWTMiddleware(next).ServeHTTP(w, r)
	})
}
