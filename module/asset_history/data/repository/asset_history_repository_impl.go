package repository

import (
	"asset_management_backend/infras"
	datasource "asset_management_backend/module/asset_history/data/data_source"
	"asset_management_backend/module/asset_history/data/model"
	"asset_management_backend/module/asset_history/domain/entity"
	"context"
)

type assetHistoryRepositoryImpl struct {
	assetDS datasource.AssetHistoryDataSource
}

// DeleteByID implements AssetHistoryDataSource.
func (a assetHistoryRepositoryImpl) DeleteByID(ctx context.Context, historyID int) error {
	return a.assetDS.DeleteByID(ctx, historyID)
}

// FindByAssetID implements AssetHistoryDataSource.
func (a assetHistoryRepositoryImpl) FindByAssetID(ctx context.Context, assetID int, page int, limit int) ([]entity.AssetHistoryEntity, error) {
	models, err := a.assetDS.FindByAssetID(ctx, assetID, page, limit)
	if err != nil {
		return nil, err
	}
	return model.AssetHistoryEntitiesFromModels(models), nil
}

// Insert implements AssetHistoryDataSource.
func (a assetHistoryRepositoryImpl) Insert(ctx context.Context, request entity.CreateAssetHistoryRequest) (*entity.AssetHistoryEntity, error) {
	insertedModel, err := a.assetDS.Insert(ctx, model.NewAssetHistoryByRequest(request))
	if err != nil {
		return nil, err
	}
	return insertedModel.ToEntity(), nil
}

// UpdateByID implements AssetHistoryDataSource.
func (a assetHistoryRepositoryImpl) UpdateByID(ctx context.Context, historyID int, updateData map[string]interface{}) error {
	return a.assetDS.UpdateByID(ctx, historyID, updateData)
}

func NewAssetHistoryRepository() AssetHistoryRepository {
	return assetHistoryRepositoryImpl{
		assetDS: datasource.NewAssetDataSource(infras.GetDbProvider()),
	}
}
