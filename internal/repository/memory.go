package repository

import (
	"sync"

	"github.com/aialok/url-shortner/internal/model"
)

type URLRepository struct {
	mu sync.Mutex
	db map[string]model.URL
}

func NewURLRepository() *URLRepository {
	return &URLRepository{
		db: make(map[string]model.URL),
	}
}

func (r *URLRepository) Save(url model.URL) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.db[url.ShortUrl] = url
}

func (r *URLRepository) Get(shortUrl string) (model.URL, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	url, exists := r.db[shortUrl]
	if !exists {
		return model.URL{}, false
	}

	url.Visits++
	r.db[shortUrl] = url
	return url, true
}
