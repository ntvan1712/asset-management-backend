package model

import (
	sharedmodel "asset_management_backend/common/shared_model"
	assetModel "asset_management_backend/module/asset/data/model"
	"time"

	"github.com/uptrace/bun"
)

type AssetHistory struct {
	bun.BaseModel `bun:"table:asset_histories"`

	ID          int `bun:",pk,autoincrement" json:"id"`
	Title       string
	Content     string
	CreatedAt   time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`
	HistoryType string

	AssetID    int  `json:"asset_id"`
	CreatorID  *int `json:"creator_id"`
	BorrowerID *int `json:"borrower_id"`

	Asset assetModel.Asset `bun:"rel:belongs-to,join:asset_id=id"`
	Creator *sharedmodel.UserBasicInfo `bun:"rel:belongs-to,join:creator_id=id"`
	Borrower *sharedmodel.UserBasicInfo `bun:"rel:belongs-to,join:borrower_id=id"`
}
