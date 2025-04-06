package datasource

import (
	"asset_management_backend/module/category/data/model"
	"context"
)

type CategoryDataSource interface {
	// AssetTypes
	FindAllAssetTypes(ctx context.Context) ([]model.AssetType, error)
	DeleteAssetTypeByID(ctx context.Context, id int) error
	UpdateAssetType(ctx context.Context, id int, updateData map[string]interface{}) error
	InsertAssetType(ctx context.Context, assetType *model.AssetType) (*model.AssetType, error)

	// AssetQualities
	FindAllAssetQualities(ctx context.Context) ([]model.AssetQuality, error)
	DeleteAssetQualityByID(ctx context.Context, id int) error
	UpdateAssetQuality(ctx context.Context, id int, updateData map[string]interface{}) error
	InsertAssetQuality(ctx context.Context, assetQuality *model.AssetQuality) (*model.AssetQuality, error)

	// Locations
	FindAllLocations(ctx context.Context) ([]model.Location, error)
	DeleteLocationByID(ctx context.Context, id int) error
	UpdateLocation(ctx context.Context, id int, updateData map[string]interface{}) error
	InsertLocation(ctx context.Context, location *model.Location) (*model.Location, error)

	// Currencies
	FindAllCurrencies(ctx context.Context) ([]model.Currency, error)

	FindAllCategories(ctx context.Context) (*model.AllCategories, error)
}
