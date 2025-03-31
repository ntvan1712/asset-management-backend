package repository

import (
	"asset_management_backend/module/category/domain/entity"
	"context"
)

type CategoryRepository interface {
	// AssetTypes
	FindAllAssetTypes(ctx context.Context) ([]entity.AssetTypeEntity, error)
	DeleteAssetTypeByID(ctx context.Context, id int) error
	UpdateAssetType(ctx context.Context, id int, updateData map[string]interface{}) error
	InsertAssetType(ctx context.Context, assetType *entity.AssetTypeEntity) (*entity.AssetTypeEntity, error)

	// AssetQualities
	FindAllAssetQualities(ctx context.Context) ([]entity.AssetQualityEntity, error)
	DeleteAssetQualityByID(ctx context.Context, id int) error
	UpdateAssetQuality(ctx context.Context, id int, updateData map[string]interface{}) error
	InsertAssetQuality(ctx context.Context, assetQuality *entity.AssetQualityEntity) (*entity.AssetQualityEntity, error)

	// Locations
	FindAllLocations(ctx context.Context) ([]entity.LocationEntity, error)
	DeleteLocationByID(ctx context.Context, id int) error
	UpdateLocation(ctx context.Context, id int, updateData map[string]interface{}) error
	InsertLocation(ctx context.Context, location *entity.LocationEntity) (*entity.LocationEntity, error)

	// Currencies
	FindAllCurrencies(ctx context.Context) ([]entity.CurrencyEntity, error)
}
