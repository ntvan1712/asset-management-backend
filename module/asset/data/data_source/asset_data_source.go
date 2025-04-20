package datasource

import (
	"asset_management_backend/module/asset/data/model"
	"context"
)

type AssetDataSource interface {
	Insert(ctx context.Context, newAsset model.Asset) (*model.Asset, error)
	UpdateByID(ctx context.Context, assetID int, updateData map[string]interface{}) error
	DeleteByID(ctx context.Context, assetID int) error

	FindByID(ctx context.Context, assetID int) (*model.Asset, error)
	GetSerialNumberByID(ctx context.Context, assetID int) (*string, error)

	FindByFilter(
		ctx context.Context,
		filter model.AssetFilterModel,
	) ([]model.Asset, error)

	InsertAssetFiles(ctx context.Context, assetFiles []model.AssetFile) ([]model.AssetFile, error)
	DeleteAssetFilesByIDs(ctx context.Context, ids []int) error

	InsertLabelImage(ctx context.Context, newLabelImage model.AssetLabelImage) (*model.AssetLabelImage, error)
	UpdateLabelImageByID(ctx context.Context, id int, updateData map[string]interface{}) error
	UpdateLabelImageByAssetID(ctx context.Context, assetID int, updateData map[string]interface{}) error
}
