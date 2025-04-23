package usecase

import (
	"asset_management_backend/module/asset_history/data/repository"
	"asset_management_backend/module/asset_history/domain/entity"
	"context"
)

type assetHistoryUsecaseImpl struct {
	assetRepo repository.AssetHistoryRepository
}

// DeleteByID implements AssetHistoryDataSource.
func (a assetHistoryUsecaseImpl) DeleteByID(ctx context.Context, historyID int) error {
	return a.assetRepo.DeleteByID(ctx, historyID)
}

// FindByAssetID implements AssetHistoryDataSource.
func (a assetHistoryUsecaseImpl) GetByAssetID(ctx context.Context, assetID int, page int, limit int) ([]entity.AssetHistoryEntity, error) {
	return a.assetRepo.FindByAssetID(ctx, assetID, page, limit)
}

// Insert implements AssetHistoryDataSource.
func (a assetHistoryUsecaseImpl) Create(ctx context.Context, request entity.CreateAssetHistoryRequest) (*entity.AssetHistoryEntity, error) {
	return a.assetRepo.Insert(ctx, request)
}

// UpdateByID implements AssetHistoryDataSource.
func (a assetHistoryUsecaseImpl) UpdateByID(ctx context.Context, historyID int, updateData map[string]interface{}) error {
	return a.assetRepo.UpdateByID(ctx, historyID, updateData)
}

func NewAssetHistoryUsecase() AssetHistoryUsecase {
	return assetHistoryUsecaseImpl{
		assetRepo: repository.NewAssetHistoryRepository(),
	}
}
