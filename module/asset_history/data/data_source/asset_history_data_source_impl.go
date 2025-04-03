package datasource

import (
	"asset_management_backend/module/asset_history/data/model"
	"context"
	"github.com/uptrace/bun"
)

type assetHistoryDataSourceImpl struct {
	dbInstance *bun.DB
}

// DeleteByID implements AssetHistoryDataSource.
func (a assetHistoryDataSourceImpl) DeleteByID(ctx context.Context, historyID int) error {
	panic("unimplemented")
}

// FindByAssetID implements AssetHistoryDataSource.
func (a assetHistoryDataSourceImpl) FindByAssetID(ctx context.Context, assetID int, page int, limit int) ([]model.AssetHistory, error) {
	panic("unimplemented")
}

// Insert implements AssetHistoryDataSource.
func (a assetHistoryDataSourceImpl) Insert(ctx context.Context, newHistory *model.AssetHistory) (*model.AssetHistory, error) {
	panic("unimplemented")
}

// UpdateByID implements AssetHistoryDataSource.
func (a assetHistoryDataSourceImpl) UpdateByID(ctx context.Context, historyID int, updateData map[string]interface{}) (*model.AssetHistory, error) {
	panic("unimplemented")
}

func NewAssetDataSource(dbInstance *bun.DB) AssetHistoryDataSource {
	return assetHistoryDataSourceImpl{dbInstance: dbInstance}
}
