package model

import (
	"time"

	"github.com/google/uuid"
)

type URL struct {
	ID          uuid.UUID `json:"id"`
	ShortCode   string    `json:"short_code"`
	OriginalURL string    `json:"original_url"`
	CreatedAt   time.Time `json:"created_at"`
}
