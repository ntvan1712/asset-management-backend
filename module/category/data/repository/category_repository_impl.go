package repository

import (
	"asset_management_backend/infras"
	datasource "asset_management_backend/module/category/data/data_source"
	"asset_management_backend/module/category/data/model"
	"asset_management_backend/module/category/domain/entity"
	"context"
)

type categoryRepositoryImpl struct {
	categoryDS datasource.CategoryDataSource
}

// UpdateAssetQuality implements CategoryRepository.
func (r *categoryRepositoryImpl) UpdateAssetQuality(ctx context.Context, id int, updateData map[string]interface{}) error {
	return r.categoryDS.UpdateAssetQuality(ctx, id, updateData)
}

// UpdateAssetType implements CategoryRepository.
func (r *categoryRepositoryImpl) UpdateAssetType(ctx context.Context, id int, updateData map[string]interface{}) error {
	return r.categoryDS.UpdateAssetType(ctx, id, updateData)
}

// UpdateLocation implements CategoryRepository.
func (r *categoryRepositoryImpl) UpdateLocation(ctx context.Context, id int, updateData map[string]interface{}) error {
	return r.categoryDS.UpdateLocation(ctx, id, updateData)
}

// AssetTypes

func (r *categoryRepositoryImpl) FindAllAssetTypes(ctx context.Context) ([]entity.AssetTypeEntity, error) {
	assetTypes, err := r.categoryDS.FindAllAssetTypes(ctx)
	if err != nil {
		return nil, err
	}
	return model.AssetTypeModelsToEntities(assetTypes), nil
}

func (r *categoryRepositoryImpl) DeleteAssetTypeByID(ctx context.Context, id int) error {
	return r.categoryDS.DeleteAssetTypeByID(ctx, id)
}

func (r *categoryRepositoryImpl) InsertAssetType(ctx context.Context, assetType *entity.AssetTypeEntity) (*entity.AssetTypeEntity, error) {
	assetTypeModel, err := r.categoryDS.InsertAssetType(ctx, &model.AssetType{
		ID:          assetType.ID,
		Code:        assetType.Code,
		Name:        assetType.Name,
		Description: assetType.Description,
	})
	if err != nil {
		return nil, err
	}
	return assetTypeModel.ToEntity(), nil
}

// AssetQualities

func (r *categoryRepositoryImpl) FindAllAssetQualities(ctx context.Context) ([]entity.AssetQualityEntity, error) {
	assetQualities, err := r.categoryDS.FindAllAssetQualities(ctx)
	if err != nil {
		return nil, err
	}
	return model.AssetQualityModelsToEntities(assetQualities), nil
}

func (r *categoryRepositoryImpl) DeleteAssetQualityByID(ctx context.Context, id int) error {
	return r.categoryDS.DeleteAssetQualityByID(ctx, id)
}

func (r *categoryRepositoryImpl) InsertAssetQuality(ctx context.Context, assetQuality *entity.AssetQualityEntity) (*entity.AssetQualityEntity, error) {
	assetQualityModel, err := r.categoryDS.InsertAssetQuality(ctx, &model.AssetQuality{
		Code:        assetQuality.Code,
		Name:        assetQuality.Name,
		Description: assetQuality.Description,
	})
	if err != nil {
		return nil, err
	}
	return assetQualityModel.ToEntity(), nil
}

// Currencies

func (r *categoryRepositoryImpl) FindAllCurrencies(ctx context.Context) ([]entity.CurrencyEntity, error) {
	currencies, err := r.categoryDS.FindAllCurrencies(ctx)
	if err != nil {
		return nil, err
	}
	return model.CurrencyModelsToEntities(currencies), nil
}

// Locations

func (r *categoryRepositoryImpl) FindAllLocations(ctx context.Context) ([]entity.LocationEntity, error) {
	locations, err := r.categoryDS.FindAllLocations(ctx)
	if err != nil {
		return nil, err
	}
	return model.LocationModelsToEntities(locations), nil
}

func (r *categoryRepositoryImpl) DeleteLocationByID(ctx context.Context, id int) error {
	return r.categoryDS.DeleteLocationByID(ctx, id)
}

func (r *categoryRepositoryImpl) InsertLocation(ctx context.Context, location *entity.LocationEntity) (*entity.LocationEntity, error) {
	locationModel, err := r.categoryDS.InsertLocation(ctx, &model.Location{
		Name:        location.Name,
		Description: location.Description,
	})
	if err != nil {
		return nil, err
	}
	return locationModel.ToEntity(), nil
}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepositoryImpl{categoryDS: datasource.NewCategoryDataSource(infras.GetDbInstance())}
}
