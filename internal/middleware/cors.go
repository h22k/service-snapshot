package middleware

import (
	"net/http"
	"snapshot/internal/durable"
)

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		durable.EnableCors(&w)
		next.ServeHTTP(w, r)
	})
}
