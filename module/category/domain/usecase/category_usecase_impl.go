package usecase

import (
	"asset_management_backend/module/category/data/repository"
	"asset_management_backend/module/category/domain/entity"
	"context"
)

type categoryUsecaseImpl struct {
	categoryRepo repository.CategoryRepository
}

// GetAllCategories implements CategoryUsecase.
func (c *categoryUsecaseImpl) GetAllCategories(ctx context.Context) (*entity.AllCategoriesEntity, error) {
	return c.categoryRepo.FindAllCategories(ctx)
}

// UpdateAssetQuality implements CategoryUsecase.
func (c *categoryUsecaseImpl) UpdateAssetQuality(ctx context.Context, id int, updateData map[string]interface{}) error {
	return c.categoryRepo.UpdateAssetQuality(ctx, id, updateData)
}

// UpdateAssetType implements CategoryUsecase.
func (c *categoryUsecaseImpl) UpdateAssetType(ctx context.Context, id int, updateData map[string]interface{}) error {
	return c.categoryRepo.UpdateAssetType(ctx, id, updateData)
}

// UpdateLocation implements CategoryUsecase.
func (c *categoryUsecaseImpl) UpdateLocation(ctx context.Context, id int, updateData map[string]interface{}) error {
	return c.categoryRepo.UpdateLocation(ctx, id, updateData)
}

// DeleteAssetQualityByID implements CategoryUsecase.
func (c *categoryUsecaseImpl) DeleteAssetQualityByID(ctx context.Context, id int) error {
	return c.categoryRepo.DeleteAssetQualityByID(ctx, id)
}

// DeleteAssetTypeByID implements CategoryUsecase.
func (c *categoryUsecaseImpl) DeleteAssetTypeByID(ctx context.Context, id int) error {
	return c.categoryRepo.DeleteAssetTypeByID(ctx, id)
}

// DeleteLocationByID implements CategoryUsecase.
func (c *categoryUsecaseImpl) DeleteLocationByID(ctx context.Context, id int) error {
	return c.categoryRepo.DeleteLocationByID(ctx, id)
}

// GetAllAssetQualities implements CategoryUsecase.
func (c *categoryUsecaseImpl) GetAllAssetQualities(ctx context.Context) ([]entity.AssetQualityEntity, error) {
	return c.categoryRepo.FindAllAssetQualities(ctx)
}

// GetAllAssetTypes implements CategoryUsecase.
func (c *categoryUsecaseImpl) GetAllAssetTypes(ctx context.Context) ([]entity.AssetTypeEntity, error) {
	return c.categoryRepo.FindAllAssetTypes(ctx)
}

// GetAllCurrencies implements CategoryUsecase.
func (c *categoryUsecaseImpl) GetAllCurrencies(ctx context.Context) ([]entity.CurrencyEntity, error) {
	return c.categoryRepo.FindAllCurrencies(ctx)
}

// GetAllLocations implements CategoryUsecase.
func (c *categoryUsecaseImpl) GetAllLocations(ctx context.Context) ([]entity.LocationEntity, error) {
	return c.categoryRepo.FindAllLocations(ctx)
}

func (c *categoryUsecaseImpl) CreateAssetQuality(ctx context.Context, assetQuality *entity.AssetQualityEntity) (*entity.AssetQualityEntity, error) {
	return c.categoryRepo.InsertAssetQuality(ctx, assetQuality)
}

func (c *categoryUsecaseImpl) CreateAssetType(ctx context.Context, assetType *entity.AssetTypeEntity) (*entity.AssetTypeEntity, error) {
	return c.categoryRepo.InsertAssetType(ctx, assetType)
}

func (c *categoryUsecaseImpl) CreateLocation(ctx context.Context, location *entity.LocationEntity) (*entity.LocationEntity, error) {
	return c.categoryRepo.InsertLocation(ctx, location)
}
func NewCategoryUsecase() CategoryUsecase {
	return &categoryUsecaseImpl{categoryRepo: repository.NewCategoryRepository()}
}
