package api

import (
	"encoding/json"
	"net/http"
	"shortly/internal/storage"
	"shortly/internal/utils"
	"strings"
)

type ShortenRequest struct {
	OriginalURL string `json:"original_url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

func ShortenURLHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ShortenRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.OriginalURL == "" {
		http.Error(w, "Invalid JSON or missing URL", http.StatusBadRequest)
		return
	}
	existingID, err := storage.FindShortIDByOriginalURL(req.OriginalURL)
	if err == nil {
		// URL already exists, return existing short URL
		resp := ShortenResponse{
			ShortURL: "http://localhost:8080/" + existingID,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}
	// TODO: generate short ID and store mapping
	shortID := utils.GenerateShortID(6)
	err = storage.SaveURL(shortID, req.OriginalURL)
	if err != nil {
		http.Error(w, "Failed to save URL", http.StatusInternalServerError)
		return
	}
	resp := ShortenResponse{
		ShortURL: "http://localhost:8080/" + shortID, // temporary
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	// Handle GET /{shortID} to redirect
	shortID := strings.TrimPrefix(r.URL.Path, "/")

	url, err := storage.GetOriginalURL(shortID)
	if err != nil {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url, http.StatusSeeOther)
}
