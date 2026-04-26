package main

import (
	"log"
	"net/http"
	"os"

	"crud-app/handlers"
	"crud-app/repository"
	"crud-app/services"
)

func main() {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "../data/app.db"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	repo, err := repository.New(dbPath)
	if err != nil {
		log.Fatalf("init db: %v", err)
	}

	svc := services.New(repo)
	handler := handlers.New(svc)

	mux := http.NewServeMux()
	mux.Handle("/todos", handler)
	mux.Handle("/todos/", handler)

	log.Printf("listening on :%s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("server: %v", err)
	}
}
