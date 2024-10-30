package models

import "time"

type URL struct {
	Id          int       `json:"id"`
	OriginalURL string    `json:"url"`
	ShortCode   string    `json:"shortCode"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	AccessCount int       `json:"accessCount,omitempty"`
}

type URLCreateOrUpdateReq struct {
	OriginalURL string `json:"url" validate:"required,notBlank"`
}
