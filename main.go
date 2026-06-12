package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"interactive_course/server"
)

func main() {
	contentDir := filepath.Join(".", "content")
	progressDir := filepath.Join(".", "progress")
	staticDir := filepath.Join(".", "static")

	os.MkdirAll(progressDir, 0755)

	srv := server.New(contentDir, progressDir)

	mux := http.NewServeMux()
	srv.RegisterRoutes(mux)

	mux.Handle("/", http.FileServer(http.Dir(staticDir)))

	port := "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	fmt.Printf("Interactive Go Tutorial running at http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
