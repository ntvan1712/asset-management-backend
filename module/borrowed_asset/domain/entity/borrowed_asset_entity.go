package entity

import (
	sharedmodel "asset_management_backend/common/shared_model"
	"time"
)

type BorrowedAssetEntity struct {
	ID          int        `bun:",pk,autoincrement" json:"id"`
	BorrowDate  *time.Time `json:"borrow_date,omitempty"`
	ReturnDate  *time.Time `json:"return_date,omitempty"`
	Description *string    `json:"description,omitempty"`

	BorrowerID int `json:"borrower_id"`
	AcceptorID int `json:"acceptor_id"`
	AssetID    int `json:"asset_id"`

	Borrower *sharedmodel.UserBasicInfo `json:"borrower,omitempty"`
	Acceptor *sharedmodel.UserBasicInfo `json:"acceptor,omitempty"`
	// Asset    *assetModel.Asset          `bun:"rel:belongs-to,join:asset_id=id"`
}
