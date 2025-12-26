package model

import "time"

type URL struct {
	Id          string    `json:"id"`
	OriginalUrl string    `json:"original_url"`
	ShortUrl    string    `json:"short_url"`
	Visits      int       `json:"visits"`
	CreatedAt   time.Time `json:"created_at"`
}
