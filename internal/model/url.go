package model

import (
	"time"

	"github.com/google/uuid"
)

type URL struct {
    ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
    ShortCode   string    `json:"short_code" gorm:"uniqueIndex"`
    OriginalURL string    `json:"original_url"`
    CreatedAt   time.Time `json:"created_at"`
}
