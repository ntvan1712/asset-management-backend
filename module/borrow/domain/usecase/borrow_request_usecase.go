package usecase

import (
	sharedmodel "asset_management_backend/common/shared_model"
	"asset_management_backend/module/borrow/domain/entity"
	"context"
)

type BorrowUsecase interface {
	CreateBorrowRequest(ctx context.Context, newBorrowRequest entity.BorrowRequestBody) (*entity.BorrowRequestEntity, error)
	UpdateBorrowRequestByID(ctx context.Context, borrowRequestID int, updateData map[string]interface{}) error
	CancelBorrowRequestByID(ctx context.Context, borrowRequestID int, requestorID int) error

	GetPendingRequests(
		ctx context.Context,
		paginateQuery sharedmodel.PaginateQuery,
	) ([]entity.BorrowRequestEntity, error)

	GetBorrowRequestsByRequestorID(
		ctx context.Context,
		requestorID int,
		paginateQuery sharedmodel.PaginateQuery,
	) ([]entity.BorrowRequestEntity, error)

	RejectBorrowRequest(
		ctx context.Context,
		borrowRequestID int,
		respondentID int,
		responseDescription *string,
	) error

	ApproveBorrowRequest(
		ctx context.Context,
		borrowRequestID int,
		respondentID int,
		assetID int,
		response *string,
	) error
}
