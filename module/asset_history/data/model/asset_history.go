package model

import (
	"asset_management_backend/common/enums"
	sharedmodel "asset_management_backend/common/shared_model"
	assetModel "asset_management_backend/module/asset/data/model"
	"asset_management_backend/module/asset_history/domain/entity"
	"time"

	"github.com/uptrace/bun"
)

const TableAssetHistory = "asset_histories"

type AssetHistory struct {
	bun.BaseModel `bun:"table:asset_histories"`

	ID          int `bun:",pk,autoincrement" json:"id"`
	Title       string
	Content     *string
	CreatedAt   time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`
	HistoryType string

	AssetID    int  `json:"asset_id"`
	CreatorID  *int `json:"creator_id"`
	BorrowerID *int `json:"borrower_id"`

	Asset    assetModel.Asset           `bun:"rel:belongs-to,join:asset_id=id"`
	Creator  *sharedmodel.UserBasicInfo `bun:"rel:belongs-to,join:creator_id=id"`
	Borrower *sharedmodel.UserBasicInfo `bun:"rel:belongs-to,join:borrower_id=id"`
}

func (a *AssetHistory) ToEntity() *entity.AssetHistoryEntity {
	if a == nil {
		return nil
	}
	return &entity.AssetHistoryEntity{
		ID:          a.ID,
		Title:       a.Title,
		Content:     a.Content,
		CreatedAt:   a.CreatedAt,
		HistoryType: a.HistoryType,
		AssetID:     a.AssetID,
		CreatorID:   a.CreatorID,
		BorrowerID:  a.BorrowerID,
		Creator:     a.Creator,
		Borrower:    a.Borrower,
	}
}

func AssetHistoryEntitiesFromModels(models []AssetHistory) []entity.AssetHistoryEntity {
	var entities []entity.AssetHistoryEntity

	for _, element := range models {
		entities = append(entities, *element.ToEntity())
	}

	return entities
}

func NewAssetCreationHistory(assetID int, creatorID *int) AssetHistory {
	contentValue := "Tài sản được thêm mới"
	return AssetHistory{
		Title:       "Thêm mới",
		Content:     &contentValue,
		CreatedAt:   time.Now().UTC(),
		HistoryType: enums.AssetHistoryType.Create,
		AssetID:     assetID,
		CreatorID:   creatorID,
		BorrowerID:  nil,
	}
}

func NewReturnHistory(assetID int, creatorID *int, borrowerID *int) AssetHistory {
	contentValue := "Tài sản được trả"
	return AssetHistory{
		Title:       "Trả tài sản",
		Content:     &contentValue,
		CreatedAt:   time.Now().UTC(),
		HistoryType: enums.AssetHistoryType.Return,
		AssetID:     assetID,
		CreatorID:   creatorID,
		BorrowerID:  borrowerID,
	}
}

func NewTransferHistory(assetID int, creatorID *int) AssetHistory {
	contentValue := "Tài sản được luân chuyển"
	return AssetHistory{
		Title:       "Luân chuyển",
		Content:     &contentValue,
		CreatedAt:   time.Now().UTC(),
		HistoryType: enums.AssetHistoryType.Transfer,
		AssetID:     assetID,
		CreatorID:   creatorID,
		BorrowerID:  nil,
	}
}

func NewCancelTransferHistory(assetID int, creatorID *int) AssetHistory {
	contentValue := "Hủy yêu cầu luân chuyển"
	return AssetHistory{
		Title:       "Hủy luân chuyển",
		Content:     &contentValue,
		CreatedAt:   time.Now().UTC(),
		HistoryType: enums.AssetHistoryType.CancelTransfer,
		AssetID:     assetID,
		CreatorID:   creatorID,
		BorrowerID:  nil,
	}
}

func NewRejectTransferHistory(assetID int, creatorID *int) AssetHistory {
	contentValue := "Từ chối yêu cầu luân chuyển"
	return AssetHistory{
		Title:       "Từ chối luân chuyển",
		Content:     &contentValue,
		CreatedAt:   time.Now().UTC(),
		HistoryType: enums.AssetHistoryType.RejectTransfer,
		AssetID:     assetID,
		CreatorID:   creatorID,
		BorrowerID:  nil,
	}
}

func NewOnBorrowHistory(assetID int, borrowID int) AssetHistory {
	contentValue := "Mượn tài sản"
	return AssetHistory{
		Title:       "Mượn",
		Content:     &contentValue,
		CreatedAt:   time.Now().UTC(),
		HistoryType: enums.AssetHistoryType.OnBorrow,
		AssetID:     assetID,
		CreatorID:   nil,
		BorrowerID:  &borrowID,
	}
}

func NewAssetUpdateHistory(assetID int, creatorID *int) AssetHistory {
	contentValue := "Tài sản được cập nhật"
	return AssetHistory{
		Title:       "Cập nhật",
		Content:     &contentValue,
		CreatedAt:   time.Now().UTC(),
		HistoryType: enums.AssetHistoryType.Update,
		AssetID:     assetID,
		CreatorID:   creatorID,
		BorrowerID:  nil,
	}
}

func NewAssetHistoryByRequest(request entity.CreateAssetHistoryRequest) AssetHistory {
	return AssetHistory{
		Title:       request.Title,
		Content:     request.Content,
		CreatedAt:   time.Now().UTC(),
		HistoryType: enums.AssetHistoryType.ByManager,
		AssetID:     *request.AssetID,
		CreatorID:   request.CreatorID,
		BorrowerID:  nil,
	}
}
