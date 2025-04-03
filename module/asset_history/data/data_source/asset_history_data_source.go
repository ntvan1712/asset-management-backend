package datasource

import (
	"asset_management_backend/module/asset_history/data/model"
	"context"
)

type AssetHistoryDataSource interface {
	Insert(ctx context.Context, newHistory *model.AssetHistory) (*model.AssetHistory, error)
	UpdateByID(ctx context.Context, historyID int, updateData map[string]interface{}) (*model.AssetHistory, error)
	DeleteByID(ctx context.Context, historyID int) error

	FindByAssetID(
		ctx context.Context,
		assetID int,
		page int,
		limit int,
	) ([]model.AssetHistory, error)
}
