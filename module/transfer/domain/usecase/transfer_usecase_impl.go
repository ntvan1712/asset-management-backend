package usecase

import (
	sharedmodel "asset_management_backend/common/shared_model"
	"asset_management_backend/module/transfer/data/repository"
	"asset_management_backend/module/transfer/domain/entity"
	"context"
)

type transferUsecaseImpl struct {
	transferRepo repository.TransferRepository
}

// Create implements TransferUsecase.
func (t *transferUsecaseImpl) Create(ctx context.Context, newRequest entity.CreateTransferRequestEntity) (*entity.TransferRequestEntity, error) {
	return t.transferRepo.Insert(ctx, newRequest)
}

// CancelTransferRequest implements TransferUsecase.
func (t *transferUsecaseImpl) CancelTransferRequest(ctx context.Context, requestID int, requestorID int) error {
	return t.transferRepo.CancelTransferRequest(ctx, requestID, requestorID)
}

// ApproveTransferRequest implements TransferUsecase.
func (t *transferUsecaseImpl) ApproveTransferRequest(ctx context.Context, requestID int, respondentID int, responseDescription *string) error {
	return t.transferRepo.ApproveTransferRequest(ctx, requestID, respondentID, responseDescription)
}

// RejectTransferRequest implements TransferUsecase.
func (t *transferUsecaseImpl) RejectTransferRequest(ctx context.Context, requestID int, respondentID int, responseDescription *string) error {
	return t.transferRepo.RejectTransferRequest(ctx, requestID, respondentID, responseDescription)
}

// DeleteByID implements TransferUsecase.
func (t *transferUsecaseImpl) DeleteByID(ctx context.Context, requestID int, requestorID int) error {
	return t.transferRepo.DeleteByID(ctx, requestID, requestorID)
}

// GetAll implements TransferUsecase.
func (t *transferUsecaseImpl) GetAll(ctx context.Context, paginateQuery sharedmodel.PaginateQuery) ([]entity.TransferRequestEntity, error) {
	return t.transferRepo.FindAll(ctx, paginateQuery)
}

// GetMyTransferRequests implements TransferUsecase.
func (t *transferUsecaseImpl) GetMyTransferRequests(ctx context.Context, respondentID int, paginateQuery sharedmodel.PaginateQuery) ([]entity.TransferRequestEntity, error) {
	return t.transferRepo.FindByRespondentID(ctx, respondentID, paginateQuery)
}

// UpdateByID implements TransferUsecase.
func (t *transferUsecaseImpl) UpdateByID(ctx context.Context, requestID int, updateData map[string]interface{}) error {
	return t.transferRepo.UpdateByID(ctx, requestID, updateData)
}

func NewTransferUsecase() TransferUsecase {
	return &transferUsecaseImpl{
		transferRepo: repository.NewTransferRepository(),
	}
}
