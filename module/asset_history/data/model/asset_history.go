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
