package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// URL Struct
type URL struct {
	Id          string    `json:"id"`
	OriginalUrl string    `json:"original_url"`
	ShortUrl    string    `json:"short_url"`
	Visits      int       `json:"visits"`
	CreatedAt   time.Time `json:"created_at"`
}

// InMemoryDB to store the URL data
var db = make(map[string]URL)

// GenerateShortUrl function
func GenerateShortUrl(url string) string {
	hasher := sha256.New()
	hasher.Write([]byte(url))
	hashBytes := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString[:6]
}

// Store URL in DB
func SaveInDB(OriginalUrl string) {
	shortUrl := GenerateShortUrl(OriginalUrl)
	url := URL{
		Id:          shortUrl,
		OriginalUrl: OriginalUrl,
		ShortUrl:    shortUrl,
		Visits:      0,
		CreatedAt:   time.Now(),
	}
	db[shortUrl] = url
}

// Get URL from DB
func GetFromDB(shortUrl string) (URL, bool) {
	url, exists := db[shortUrl]
	fmt.Println("Before:", url.Visits)
	if !exists {
		return URL{}, false
	}
	url.Visits++
	db[shortUrl] = url
	fmt.Println("After:", url.Visits)
	return url, true
}

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
	url, exists := GetFromDB(shortURL)
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
	SaveInDB(originalURL)
	url, exists := GetFromDB(GenerateShortUrl(originalURL))
	if !exists {
		http.Error(w, "url not found", http.StatusNotFound)
		return
	}
	urlJSON := url
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(urlJSON); err != nil {
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

	url, exists := GetFromDB(shortID)
	if !exists {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, url.OriginalUrl, http.StatusFound)
}

func main() {
	url := "https://www.google.com"
	SaveInDB(url)
	fmt.Println(GetFromDB(GenerateShortUrl(url)))

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
