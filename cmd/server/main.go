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

	// Clean up
	if count, err := storage.CleanupExpiredURLs(); err == nil {
		log.Printf("[*] Cleaned up %d expired URLs", count)
	} else {
		log.Printf("Cleanup error: %v", err)
	}

	// Start
	http.HandleFunc("/shorten", api.ShortenURLHandler)
	http.HandleFunc("/", api.RedirectHandler)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
