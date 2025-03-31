package usecase

import (
	"asset_management_backend/module/category/domain/entity"
	"context"
)

type CategoryUsecase interface {
	// AssetTypes
	GetAllAssetTypes(ctx context.Context) ([]entity.AssetTypeEntity, error)
	DeleteAssetTypeByID(ctx context.Context, id int) error
	UpdateAssetType(ctx context.Context, id int, updateData map[string]interface{}) error
	CreateAssetType(ctx context.Context, assetType *entity.AssetTypeEntity) (*entity.AssetTypeEntity, error)

	// AssetQualities
	GetAllAssetQualities(ctx context.Context) ([]entity.AssetQualityEntity, error)
	DeleteAssetQualityByID(ctx context.Context, id int) error
	UpdateAssetQuality(ctx context.Context, id int, updateData map[string]interface{}) error
	CreateAssetQuality(ctx context.Context, assetQuality *entity.AssetQualityEntity) (*entity.AssetQualityEntity, error)

	// Currencies
	GetAllCurrencies(ctx context.Context) ([]entity.CurrencyEntity, error)

	// Locations
	GetAllLocations(ctx context.Context) ([]entity.LocationEntity, error)
	DeleteLocationByID(ctx context.Context, id int) error
	UpdateLocation(ctx context.Context, id int, updateData map[string]interface{}) error
	CreateLocation(ctx context.Context, location *entity.LocationEntity) (*entity.LocationEntity, error)
}
