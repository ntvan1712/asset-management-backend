package repository

import (
	sharedmodel "asset_management_backend/common/shared_model"
	"asset_management_backend/infras"
	assetHistoryDataSource "asset_management_backend/module/asset_history/data/data_source"
	historyM "asset_management_backend/module/asset_history/data/model"
	transferDataSource "asset_management_backend/module/transfer/data/data_source"
	transferM "asset_management_backend/module/transfer/data/model"
	"asset_management_backend/module/transfer/domain/entity"
	"context"
)

type transferRepositoryImpl struct {
	transferDS transferDataSource.TransferDataSource
	historyDS  assetHistoryDataSource.AssetHistoryDataSource
}

// Insert implements TransferRepository.
func (t *transferRepositoryImpl) Insert(ctx context.Context, newRequest entity.CreateTransferRequestEntity) (*entity.TransferRequestEntity, error) {
	model, err := t.transferDS.Insert(ctx, transferM.NewTransferRequestFromBody(newRequest))
	if err != nil {
		return nil, err
	}
	return model.ToEntity(), nil
}

// CancelTransferRequest implements TransferRepository.
func (t *transferRepositoryImpl) CancelTransferRequest(ctx context.Context, requestID int, requestorID int) error {
	request, err := t.transferDS.CancelTransferRequest(ctx, requestID, requestorID)
	if err != nil {
		return err
	}

	go func() {
		t.historyDS.Insert(context.Background(), historyM.NewCancelTransferHistory(request.AssetID, &requestID))
	}()

	return nil
}

// ApproveTransferRequest implements TransferRepository.
func (t *transferRepositoryImpl) ApproveTransferRequest(ctx context.Context, requestID int, respondentID int, responseDescription *string) error {
	request, err := t.transferDS.ApproveTransferRequest(ctx, requestID, respondentID, responseDescription)
	if err != nil {
		return err
	}

	go func() {
		t.historyDS.Insert(context.Background(), historyM.NewOnBorrowHistory(request.AssetID, respondentID))
	}()

	return nil
}

// RejectTransferRequest implements TransferRepository.
func (t *transferRepositoryImpl) RejectTransferRequest(ctx context.Context, requestID int, respondentID int, responseDescription *string) error {
	request, err := t.transferDS.RejectTransferRequest(ctx, requestID, respondentID, responseDescription)
	if err != nil {
		return err
	}

	go func() {
		t.historyDS.Insert(context.Background(), historyM.NewRejectTransferHistory(request.AssetID, &requestID))
	}()

	return nil
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
	return transferM.TransferRequestModelsToEntities(models), nil
}

// FindByRespondentID implements TransferRepository.
func (t *transferRepositoryImpl) FindByRespondentID(ctx context.Context, respondentID int, paginateQuery sharedmodel.PaginateQuery) ([]entity.TransferRequestEntity, error) {
	models, err := t.transferDS.FindByRespondentID(ctx, respondentID, paginateQuery)
	if err != nil {
		return nil, err
	}
	return transferM.TransferRequestModelsToEntities(models), nil
}

// UpdateByID implements TransferRepository.
func (t *transferRepositoryImpl) UpdateByID(ctx context.Context, requestID int, updateData map[string]interface{}) error {
	return t.transferDS.UpdateByID(ctx, requestID, updateData)
}

func NewTransferRepository() TransferRepository {
	dbIns := infras.GetDbProvider()
	return &transferRepositoryImpl{
		transferDS: transferDataSource.NewTransferDataSource(dbIns),
		historyDS:  assetHistoryDataSource.NewAssetDataSource(dbIns),
	}
}
