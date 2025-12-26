package service

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/aialok/url-shortner/internal/model"
	"github.com/aialok/url-shortner/internal/repository"
)

type ShortenerService struct {
	repo *repository.URLRepository
}

func NewShortenerService(repo *repository.URLRepository) *ShortenerService {
	return &ShortenerService{
		repo: repo,
	}
}

func (s *ShortenerService) Shorten(originalURL string) model.URL {
	short := generateShortUrl(originalURL)

	url := model.URL{
		Id:          short,
		OriginalUrl: originalURL,
		ShortUrl:    short,
		Visits:      0,
		CreatedAt:   time.Now(),
	}

	s.repo.Save(url)
	return url
}

func (s *ShortenerService) Resolve(shortUrl string) (model.URL, bool) {
	return s.repo.Get(shortUrl)
}

func generateShortUrl(url string) string {
	hash := sha256.Sum256([]byte(url))
	return hex.EncodeToString(hash[:])[:6]
}
