package handler

import (
	"encoding/json"
	"net/http"

	"github.com/aialok/tiny-url-shortner/internal/service"
)

func Shorten(svc *service.ShortenerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		originalURL := r.URL.Query().Get("original_url")
		if originalURL == "" {
			http.Error(w, "original_url is required", http.StatusBadRequest)
			return
		}

		url := svc.Shorten(originalURL)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(url)
	}
}
