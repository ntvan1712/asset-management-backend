package model

import (
	sharedmodel "asset_management_backend/common/shared_model"
	"asset_management_backend/module/borrowed_asset/domain/entity"
	// assetModel "asset_management_backend/module/asset/data/model"
	"time"

	"github.com/uptrace/bun"
)

const TableBorrowedAsset = "borrowed_assets"

type BorrowedAsset struct {
	bun.BaseModel `bun:"table:borrowed_assets"`

	ID          int        `bun:",pk,autoincrement" json:"id"`
	BorrowDate  *time.Time `json:"borrow_date,omitempty"`
	ReturnDate  *time.Time `json:"return_date,omitempty"`
	Description *string    `json:"description,omitempty"`

	BorrowerID int `json:"borrower_id"`
	AcceptorID int `json:"acceptor_id"`
	AssetID    int `json:"asset_id"`

	Borrower *sharedmodel.UserBasicInfo `bun:"rel:belongs-to,join:borrower_id=id"`
	Acceptor *sharedmodel.UserBasicInfo `bun:"rel:belongs-to,join:acceptor_id=id"`
	// Asset    *assetModel.Asset          `bun:"rel:belongs-to,join:asset_id=id"`
}

func (b *BorrowedAsset) ToEntity() *entity.BorrowedAssetEntity {
	if b == nil {
		return nil
	}
	return &entity.BorrowedAssetEntity{
		ID:          b.ID,
		BorrowDate:  b.BorrowDate,
		ReturnDate:  b.ReturnDate,
		Description: b.Description,
		BorrowerID:  b.BorrowerID,
		AcceptorID:  b.AcceptorID,
		AssetID:     b.AssetID,
		Borrower:    b.Borrower,
		Acceptor:    b.Acceptor,
	}
}
