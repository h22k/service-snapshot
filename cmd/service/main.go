package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"snapshot/internal/durable"
	"snapshot/internal/model"
	"snapshot/internal/server"
)

func init() {
	// setup logger
	durable.SetupLogger()

	// load .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := durable.ConnectDB(os.Getenv("DB_DSN")); err != nil {
		log.Fatal("Error connecting to database")
	}

	//migrate database
	if err := durable.Connection().AutoMigrate(
		&model.Snapshot{}); err != nil {
		log.Fatal(err)
	}
}

func main() {
	mux := http.NewServeMux()
	server.SetupRoutes(mux)

	middlewareMux := server.SetupMiddleware(mux)
	server.StartServer(middlewareMux, os.Getenv("SERVER_PORT"))
}
