package repository

import (
	sharedmodel "asset_management_backend/common/shared_model"
	"asset_management_backend/infras"
	datasource "asset_management_backend/module/borrow/data/data_source"
	"asset_management_backend/module/borrow/data/model"
	"asset_management_backend/module/borrow/domain/entity"
	"context"
)

type borrowRepositoryImpl struct {
	borrowDS datasource.BorrowDataSource
}

// ApproveBorrowRequest implements BorrowRepository.
func (b *borrowRepositoryImpl) ApproveBorrowRequest(ctx context.Context, borrowRequestID int, respondentID int, assetID int, response *string) error {
	return b.borrowDS.ApproveBorrowRequest(ctx, borrowRequestID, respondentID, assetID, response)
}

// FindByRequestorID implements BorrowRepository.
func (b *borrowRepositoryImpl) FindByRequestorID(ctx context.Context, requestorID int, paginateQuery sharedmodel.PaginateQuery) ([]entity.BorrowRequestEntity, error) {
	borrowRequests, err := b.borrowDS.FindByRequestorID(ctx, requestorID, paginateQuery)
	if err != nil {
		return nil, err
	}

	return model.BorrowRequestModelsToEntities(borrowRequests), nil
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

	return model.BorrowRequestModelsToEntities(borrowRequests), nil
}

// Insert implements BorrowDataSource.
func (b *borrowRepositoryImpl) Insert(ctx context.Context, newBorrowRequest entity.BorrowRequestBody) (*entity.BorrowRequestEntity, error) {

	borrowRequest, err := b.borrowDS.Insert(ctx, model.BorrowRequestModelFromBody(&newBorrowRequest))
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
	return &borrowRepositoryImpl{
		borrowDS: datasource.NewBorrowDataSource(infras.GetDbProvider()),
	}
}
