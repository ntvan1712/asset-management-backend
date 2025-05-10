package usecase

import (
	"asset_management_backend/common/enums"
	sharedmodel "asset_management_backend/common/shared_model"
	"asset_management_backend/module/borrow/data/repository"
	"asset_management_backend/module/borrow/domain/entity"
	"context"
	"time"
)

type borrowRequestUsecaseImpl struct {
	borrowRepo repository.BorrowRepository
}

// ApproveBorrowRequest implements BorrowUsecase.
func (b *borrowRequestUsecaseImpl) ApproveBorrowRequest(ctx context.Context, borrowRequestID int, respondentID int, assetID int, response *string) error {
	return b.borrowRepo.ApproveBorrowRequest(ctx, borrowRequestID, respondentID, assetID, response)
}

// DeclineBorrowRequest implements BorrowUsecase.
func (b *borrowRequestUsecaseImpl) RejectBorrowRequest(
	ctx context.Context,
	borrowRequestID int,
	respondentID int,
	responseDescription *string,
) error {
	updateMap := map[string]interface{}{
		"status":               enums.BorrowRequestEnum.Rejected,
		"respondent_id":        respondentID,
		"response_description": responseDescription,
		"response_at":          time.Now(),
	}
	return b.borrowRepo.UpdateByID(ctx, borrowRequestID, updateMap)
}

// GetBorrowRequestsByRequestorID implements BorrowUsecase.
func (b *borrowRequestUsecaseImpl) GetBorrowRequestsByRequestorID(ctx context.Context, requestorID int, paginateQuery sharedmodel.PaginateQuery) ([]entity.BorrowRequestEntity, error) {
	return b.borrowRepo.FindByRequestorID(ctx, requestorID, paginateQuery)
}

// CreateBorrowRequest implements BorrowUsecase.
func (b *borrowRequestUsecaseImpl) CreateBorrowRequest(ctx context.Context, newBorrowRequest entity.BorrowRequestBody) (*entity.BorrowRequestEntity, error) {
	return b.borrowRepo.Insert(ctx, newBorrowRequest)
}

// CancelBorrowRequestByID implements BorrowUsecase.
func (b *borrowRequestUsecaseImpl) CancelBorrowRequestByID(ctx context.Context, borrowRequestID int, requestorID int) error {
	return b.borrowRepo.DeleteByID(ctx, borrowRequestID, requestorID)
}

// GetPendingRequests implements BorrowUsecase.
func (b *borrowRequestUsecaseImpl) GetPendingRequests(ctx context.Context, paginateQuery sharedmodel.PaginateQuery) ([]entity.BorrowRequestEntity, error) {
	return b.borrowRepo.FindPendingRequests(ctx, paginateQuery)
}

// UpdateBorrowRequestByID implements BorrowUsecase.
func (b *borrowRequestUsecaseImpl) UpdateBorrowRequestByID(ctx context.Context, borrowRequestID int, updateData map[string]interface{}) error {
	panic("unimplemented")
}

func NewBorrowRequestUsecase() BorrowUsecase {
	return &borrowRequestUsecaseImpl{
		borrowRepo: repository.NewBorrowRepository(),
	}
}
