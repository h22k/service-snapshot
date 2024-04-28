package middleware

import (
	"net/http"
	"os"
)

func ApiKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if apiKey := r.Header.Get("X-Api-Key"); apiKey != os.Getenv("API_KEY") {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		next.ServeHTTP(w, r)
	})
}
