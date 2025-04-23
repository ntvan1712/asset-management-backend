package entity

import (
	sharedmodel "asset_management_backend/common/shared_model"
	"time"
)

type AssetHistoryEntity struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Content     *string   `json:"content,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	HistoryType string    `json:"history_type"`

	AssetID    int  `json:"asset_id"`
	CreatorID  *int `json:"creator_id"`
	BorrowerID *int `json:"borrower_id"`

	Creator  *sharedmodel.UserBasicInfo `json:"creator,omitempty"`
	Borrower *sharedmodel.UserBasicInfo `json:"borrower,omitempty"`
}
