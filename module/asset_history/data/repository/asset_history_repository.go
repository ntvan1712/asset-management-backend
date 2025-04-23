package repository

import (
	"asset_management_backend/module/asset_history/domain/entity"
	"context"
)

type AssetHistoryRepository interface {
	Insert(ctx context.Context, request entity.CreateAssetHistoryRequest) (*entity.AssetHistoryEntity, error)
	UpdateByID(ctx context.Context, historyID int, updateData map[string]interface{}) error
	DeleteByID(ctx context.Context, historyID int) error

	FindByAssetID(
		ctx context.Context,
		assetID int,
		page int,
		limit int,
	) ([]entity.AssetHistoryEntity, error)
}
