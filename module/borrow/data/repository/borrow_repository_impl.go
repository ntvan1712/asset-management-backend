package repository

import (
	sharedmodel "asset_management_backend/common/shared_model"
	"asset_management_backend/infras"
	assetHistoryDataSource "asset_management_backend/module/asset_history/data/data_source"
	historyM "asset_management_backend/module/asset_history/data/model"
	borrowDataSource "asset_management_backend/module/borrow/data/data_source"
	borrowM "asset_management_backend/module/borrow/data/model"
	"asset_management_backend/module/borrow/domain/entity"
	"context"
)

type borrowRepositoryImpl struct {
	borrowDS  borrowDataSource.BorrowDataSource
	historyDS assetHistoryDataSource.AssetHistoryDataSource
}

// ApproveBorrowRequest implements BorrowRepository.
func (b *borrowRepositoryImpl) ApproveBorrowRequest(
	ctx context.Context,
	borrowRequestID int,
	respondentID int,
	assetID int,
	response *string,
) error {

	if err := b.borrowDS.ApproveBorrowRequest(ctx, borrowRequestID, respondentID, assetID, response); err != nil {
		return err
	}
	go func() {
		b.historyDS.Insert(context.Background(), historyM.NewTransferHistory(assetID, &respondentID))
	}()

	return nil
}

// FindByRequestorID implements BorrowRepository.
func (b *borrowRepositoryImpl) FindByRequestorID(ctx context.Context, requestorID int, paginateQuery sharedmodel.PaginateQuery) ([]entity.BorrowRequestEntity, error) {
	borrowRequests, err := b.borrowDS.FindByRequestorID(ctx, requestorID, paginateQuery)
	if err != nil {
		return nil, err
	}

	return borrowM.BorrowRequestModelsToEntities(borrowRequests), nil
}

// DeleteByID implements BorrowDataSource.
func (b *borrowRepositoryImpl) DeleteByID(
	ctx context.Context,
	borrowRequestID int,
	requestorID int,
) error {
	return b.borrowDS.DeleteByID(ctx, borrowRequestID, requestorID)
}

// FindPendingRequests implements BorrowDataSource.
func (b *borrowRepositoryImpl) FindPendingRequests(ctx context.Context, paginateQuery sharedmodel.PaginateQuery) ([]entity.BorrowRequestEntity, error) {
	borrowRequests, err := b.borrowDS.FindPendingRequests(ctx, paginateQuery)
	if err != nil {
		return nil, err
	}

	return borrowM.BorrowRequestModelsToEntities(borrowRequests), nil
}

// Insert implements BorrowDataSource.
func (b *borrowRepositoryImpl) Insert(ctx context.Context, newBorrowRequest entity.BorrowRequestBody) (*entity.BorrowRequestEntity, error) {

	borrowRequest, err := b.borrowDS.Insert(ctx, borrowM.BorrowRequestModelFromBody(&newBorrowRequest))
	if err != nil {
		return nil, err
	}

	return borrowRequest.ToEntity(), nil
}

// UpdateByID implements BorrowDataSource.
func (b *borrowRepositoryImpl) UpdateByID(ctx context.Context, borrowRequestID int, updateData map[string]interface{}) error {
	return b.borrowDS.UpdateByID(ctx, borrowRequestID, updateData)
}

func NewBorrowRepository() BorrowRepository {
	dbInstance := infras.GetDbProvider()
	return &borrowRepositoryImpl{
		borrowDS:  borrowDataSource.NewBorrowDataSource(dbInstance),
		historyDS: assetHistoryDataSource.NewAssetDataSource(dbInstance),
	}
}
