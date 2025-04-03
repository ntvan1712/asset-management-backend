package model

import (
	"time"

	"github.com/uptrace/bun"
)

type AssetLabelImage struct {
	bun.BaseModel `bun:"table:asset_images"`

	ID        int       `bun:",pk,autoincrement" json:"id"`
	Path      string    `bun:",notnull,type:varchar(100)" json:"path"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`
}
