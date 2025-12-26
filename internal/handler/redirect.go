package handler

import (
	"net/http"
	"strings"

	"github.com/aialok/tiny-url-shortner/internal/service"
)

func Redirect(svc *service.ShortenerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, "/r/") {
			http.NotFound(w, r)
			return
		}

		short := strings.TrimPrefix(r.URL.Path, "/r/")
		url, ok := svc.Resolve(short)
		if !ok {
			http.NotFound(w, r)
			return
		}

		http.Redirect(w, r, url.OriginalUrl, http.StatusFound)
	}
}
