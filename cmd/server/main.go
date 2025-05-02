package main

import (
	"log"
	"net/http"
	"shortly/internal/api"
	"shortly/internal/storage"
)

func main() {
	err := storage.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize DB: %v", err)
	}

	http.HandleFunc("/shorten", api.ShortenURLHandler)
	http.HandleFunc("/", api.RedirectHandler)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
