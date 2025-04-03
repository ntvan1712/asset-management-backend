package datasource

import (
	"asset_management_backend/module/asset/data/model"
	"context"
	"github.com/uptrace/bun"
)

type assetDataSourceImpl struct {
	dbInstance *bun.DB
}

// DeleteByID implements AssetDataSource.
func (a assetDataSourceImpl) DeleteByID(ctx context.Context, assetID int) error {
	panic("unimplemented")
}

// FindByFilter implements AssetDataSource.
func (a assetDataSourceImpl) FindByFilter(ctx context.Context, assetTypeID *int, assetQualityID *int, page int, limit int) ([]model.Asset, error) {
	panic("unimplemented")
}

// FindByID implements AssetDataSource.
func (a assetDataSourceImpl) FindByID(ctx context.Context, assetID int) (*model.Asset, error) {
	panic("unimplemented")
}

// Insert implements AssetDataSource.
func (a assetDataSourceImpl) Insert(ctx context.Context, newAsset *model.Asset) (*model.Asset, error) {
	panic("unimplemented")
}

// UpdateByID implements AssetDataSource.
func (a assetDataSourceImpl) UpdateByID(ctx context.Context, assetID int, updateData map[string]interface{}) (*model.Asset, error) {
	panic("unimplemented")
}

func NewAssetDataSource(dbInstance *bun.DB) AssetDataSource {
	return assetDataSourceImpl{dbInstance: dbInstance}
}
