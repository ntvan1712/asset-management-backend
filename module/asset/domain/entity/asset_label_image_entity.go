package entity

import (
	"time"
)

type AssetLabelImageEntity struct {
	ID        int       `json:"id"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}
