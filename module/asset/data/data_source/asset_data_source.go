package datasource

import (
	"asset_management_backend/module/asset/data/model"
	"context"
)

type AssetDataSource interface {
	Insert(ctx context.Context, newAsset model.Asset) (*model.Asset, error)
	UpdateByID(ctx context.Context, assetID int, updateData map[string]interface{}) (*model.Asset, error)
	DeleteByID(ctx context.Context, assetID int) error

	FindByID(ctx context.Context, assetID int) (*model.Asset, error)

	FindByFilter(
		ctx context.Context,
		assetTypeID *int,
		assetQualityID *int,
		page int,
		limit int,
	) ([]model.Asset, error)

	InsertAssetFiles(ctx context.Context, assetFiles []model.AssetFile) ([]model.AssetFile, error)
}
