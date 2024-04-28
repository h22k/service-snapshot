package middleware

import (
	"fmt"
	"net/http"
	"os"
)

func ClientIPMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RemoteAddr != os.Getenv("API_KEY") { //TODO:: find better solution
			fmt.Println(r.RemoteAddr)
			w.WriteHeader(http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
