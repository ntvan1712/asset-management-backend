package repository

import (
	sharedmodel "asset_management_backend/common/shared_model"
	"asset_management_backend/infras"
	datasource "asset_management_backend/module/transfer/data/data_source"
	"asset_management_backend/module/transfer/data/model"
	"asset_management_backend/module/transfer/domain/entity"
	"context"
)

type transferRepositoryImpl struct {
	transferDS datasource.TransferDataSource
}

// CancelTransferRequest implements TransferRepository.
func (t *transferRepositoryImpl) CancelTransferRequest(ctx context.Context, requestID int, requestorID int) error {
	return t.transferDS.CancelTransferRequest(ctx, requestID, requestorID)
}

// ApproveTransferRequest implements TransferRepository.
func (t *transferRepositoryImpl) ApproveTransferRequest(ctx context.Context, requestID int, respondentID int, responseDescription *string) error {
	return t.transferDS.ApproveTransferRequest(ctx, requestID, respondentID, responseDescription)
}

// RejectTransferRequest implements TransferRepository.
func (t *transferRepositoryImpl) RejectTransferRequest(ctx context.Context, requestID int, respondentID int, responseDescription *string) error {
	return t.transferDS.RejectTransferRequest(ctx, requestID, respondentID, responseDescription)
}

// DeleteByID implements TransferRepository.
func (t *transferRepositoryImpl) DeleteByID(ctx context.Context, requestID int, requestorID int) error {
	return t.transferDS.DeleteByID(ctx, requestID, requestorID)
}

// FindAll implements TransferRepository.
func (t *transferRepositoryImpl) FindAll(ctx context.Context, paginateQuery sharedmodel.PaginateQuery) ([]entity.TransferRequestEntity, error) {
	models, err := t.transferDS.FindAll(ctx, paginateQuery)
	if err != nil {
		return nil, err
	}
	return model.TransferRequestModelsToEntities(models), nil
}

// FindByRespondentID implements TransferRepository.
func (t *transferRepositoryImpl) FindByRespondentID(ctx context.Context, respondentID int, paginateQuery sharedmodel.PaginateQuery) ([]entity.TransferRequestEntity, error) {
	models, err := t.transferDS.FindByRespondentID(ctx, respondentID, paginateQuery)
	if err != nil {
		return nil, err
	}
	return model.TransferRequestModelsToEntities(models), nil
}

// UpdateByID implements TransferRepository.
func (t *transferRepositoryImpl) UpdateByID(ctx context.Context, requestID int, updateData map[string]interface{}) error {
	return t.transferDS.UpdateByID(ctx, requestID, updateData)
}

func NewTransferRepository() TransferRepository {
	return &transferRepositoryImpl{
		transferDS: datasource.NewTransferDataSource(infras.GetDbProvider()),
	}
}
