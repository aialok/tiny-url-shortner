package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
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
	if !exists {
		return URL{}, false
	}
	url.Visits++
	return url, true
}

func main() {
	url := "https://www.google.com"
	SaveInDB(url)
	fmt.Println(GetFromDB(GenerateShortUrl(url)))
}
