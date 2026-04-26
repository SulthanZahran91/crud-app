package main

import (
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"crud-app/handlers"
	"crud-app/repository"
	"crud-app/services"
)

func main() {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = filepath.Join("data", "app.db")
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

	frontendHandler, err := buildFrontendHandler()
	if err != nil {
		log.Printf("frontend disabled: %v", err)
	} else {
		mux.Handle("/", frontendHandler)
	}

	log.Printf("listening on :%s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("server: %v", err)
	}
}

func buildFrontendHandler() (http.Handler, error) {
	frontendRoot := filepath.Join("frontend", "dist")
	indexPath := filepath.Join(frontendRoot, "index.html")
	if _, err := os.Stat(indexPath); err != nil {
		return nil, err
	}

	frontendFS := os.DirFS(frontendRoot)
	fileServer := http.FileServer(http.FS(frontendFS))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/")
		if path == "" {
			http.ServeFile(w, r, indexPath)
			return
		}

		if _, err := fs.Stat(frontendFS, path); err == nil {
			fileServer.ServeHTTP(w, r)
			return
		}

		http.ServeFile(w, r, indexPath)
	}), nil
}
