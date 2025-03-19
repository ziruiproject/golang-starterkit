package middlewares

import (
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

func ZerologMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)

		log.Info().
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Int("status", http.StatusOK).
			Dur("duration", duration).
			Msg("HTTP request")
	})
}
