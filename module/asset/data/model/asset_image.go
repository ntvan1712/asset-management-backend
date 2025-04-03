package model

import (
	"time"

	"github.com/uptrace/bun"
)

type AssetImage struct {
	bun.BaseModel `bun:"table:asset_label_images"`

	ID        int       `bun:",pk,autoincrement" json:"id"`
	Path      string    `bun:",notnull,type:varchar(100)" json:"path"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`

	AssetID int   `json:"asset_id"`
	Asset   Asset `bun:"rel:belongs-to,join:asset_id=id"`
}
