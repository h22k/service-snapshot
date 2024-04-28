package middleware

import (
	"fmt"
	"net/http"
	"time"
)

type statusResponseWriter struct {
	http.ResponseWriter
	status int
}

func (srw *statusResponseWriter) WriteHeader(code int) {
	srw.status = code
	srw.ResponseWriter.WriteHeader(code)
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		srw := &statusResponseWriter{ResponseWriter: w}
		started := time.Now()
		next.ServeHTTP(srw, r)
		ended := time.Now()
		fmt.Printf("%s	%s	%s	%d	%s TOOK: [%s sec]\n", time.Now().Format("2006/01/02 15:04:05"), r.Method, r.URL, srw.status, r.Header.Get("x-forwarded-for"), ended.Sub(started))
	})
}
