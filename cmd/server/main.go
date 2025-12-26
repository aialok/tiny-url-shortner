package main

import (
	"fmt"
	"net/http"

	"github.com/aialok/tiny-url-shortner/internal/handler"
	"github.com/aialok/tiny-url-shortner/internal/repository"
	"github.com/aialok/tiny-url-shortner/internal/service"
)

func main() {
	repo := repository.NewURLRepository()
	svc := service.NewShortenerService(repo)

	http.HandleFunc("/", handler.Health)
	http.HandleFunc("/shorten", handler.Shorten(svc))
	http.HandleFunc("/r/", handler.Redirect(svc))

	fmt.Println("Server running at http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
