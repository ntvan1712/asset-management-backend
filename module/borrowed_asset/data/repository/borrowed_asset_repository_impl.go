package repository

import (
	"asset_management_backend/infras"
	historyDataSource "asset_management_backend/module/asset_history/data/data_source"
	"asset_management_backend/module/asset_history/data/model"
	borrowedAssetDataSource "asset_management_backend/module/borrowed_asset/data/data_source"
	"asset_management_backend/module/borrowed_asset/domain/entity"
	"context"
)

type borrowedAssetRepositoryImpl struct {
	borrowedAssetDS borrowedAssetDataSource.BorrowedAssetDataSource
	historyDS       historyDataSource.AssetHistoryDataSource
}

// ReturnAsset implements BorrowedAssetRepository.
func (b borrowedAssetRepositoryImpl) ReturnAsset(ctx context.Context, borrowedAssetId int) (*entity.BorrowedAssetEntity, error) {
	result, err := b.borrowedAssetDS.ReturnAsset(ctx, borrowedAssetId)
	if err != nil {
		return nil, err
	}

	go func() {
		b.historyDS.Insert(ctx, model.NewReturnHistory(result.AssetID, &result.AcceptorID, &result.BorrowerID))
	}()

	return result.ToEntity(), nil

}

func NewBorrowedAssetRepository() BorrowedAssetRepository {
	dbIns := infras.GetDbProvider()
	return borrowedAssetRepositoryImpl{
		borrowedAssetDS: borrowedAssetDataSource.NewBorrowedAssetDataSource(dbIns),
		historyDS:       historyDataSource.NewAssetDataSource(dbIns),
	}
}
