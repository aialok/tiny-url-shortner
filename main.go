package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/aialok/url-shortner/internal/repository"
	"github.com/aialok/url-shortner/internal/service"
)

var repo = repository.NewURLRepository()
var svc = service.NewShortenerService(repo)

type HealthResponse struct {
	Status string `json:"status"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	response := HealthResponse{
		Status: "ok",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func urlHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Query().Get("short_url")
	if shortURL == "" {
		http.Error(w, "short_url is required", http.StatusBadRequest)
		return
	}
	url, exists := svc.Resolve(shortURL)
	if !exists {
		http.Error(w, "url not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(url); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	originalURL := r.URL.Query().Get("original_url")
	if originalURL == "" {
		http.Error(w, "original_url is required", http.StatusBadRequest)
		return
	}
	url := svc.Shorten(originalURL)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(url); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if !strings.HasPrefix(path, "/r/") {
		http.NotFound(w, r)
		return
	}

	shortID := strings.TrimPrefix(path, "/r/")
	if shortID == "" {
		http.NotFound(w, r)
		return
	}

	url, exists := svc.Resolve(shortID)
	if !exists {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, url.OriginalUrl, http.StatusFound)
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/url", urlHandler)
	http.HandleFunc("/shorten", shortenHandler)
	http.HandleFunc("/r/", redirectHandler)
	// Start a server
	fmt.Println("Server running at port http://localhost:3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
