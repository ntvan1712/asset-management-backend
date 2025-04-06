package entity

import (
	"time"
)

type AssetFileEntity struct {
	ID        int       `json:"id"`
	FileType  string    `json:"file_type"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}
