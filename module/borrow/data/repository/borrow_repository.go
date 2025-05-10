package repository

import (
	sharedmodel "asset_management_backend/common/shared_model"
	"asset_management_backend/module/borrow/domain/entity"
	"context"
)

type BorrowRepository interface {
	Insert(ctx context.Context, newBorrowRequest entity.BorrowRequestBody) (*entity.BorrowRequestEntity, error)
	UpdateByID(ctx context.Context, borrowRequestID int, updateData map[string]interface{}) error
	DeleteByID(
		ctx context.Context,
		borrowRequestID int,
		requestorID int,
	) error

	FindPendingRequests(
		ctx context.Context,
		paginateQuery sharedmodel.PaginateQuery,
	) ([]entity.BorrowRequestEntity, error)

	FindByRequestorID(
		ctx context.Context,
		requestorID int,
		paginateQuery sharedmodel.PaginateQuery,
	) ([]entity.BorrowRequestEntity, error)

	ApproveBorrowRequest(
		ctx context.Context,
		borrowRequestID int,
		respondentID int,
		assetID int,
		response *string,
	) error
}
