package server

import (
	"fmt"
	"log"
	"net/http"
	"snapshot/internal/middleware"
	"snapshot/internal/routes"
)

func SetupRoutes(mux *http.ServeMux) {
	routes.Health(mux)
	routes.Get(mux)
	routes.Post(mux)
}

func SetupMiddleware(handler http.Handler) http.Handler {
	panicHandler := middleware.RecoverPanic(handler)
	corsHandler := middleware.CorsMiddleware(panicHandler)
	loggingHandler := middleware.LoggerMiddleware(corsHandler)
	apiKeyHandler := middleware.ApiKeyMiddleware(loggingHandler)
	clientIPHandler := middleware.ClientIPMiddleWare(apiKeyHandler)
	return clientIPHandler
}

func StartServer(handler http.Handler, port string) {
	fmt.Printf("Server is starting on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
