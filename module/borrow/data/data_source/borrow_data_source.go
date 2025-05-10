package datasource

import (
	sharedmodel "asset_management_backend/common/shared_model"
	"asset_management_backend/module/borrow/data/model"
	"context"
)

type BorrowDataSource interface {
	Insert(ctx context.Context, newBorrowRequest model.BorrowRequest) (*model.BorrowRequest, error)
	UpdateByID(ctx context.Context, borrowRequestID int, updateData map[string]interface{}) error
	DeleteByID(
		ctx context.Context,
		borrowRequestID int,
		requestorID int,
	) error

	FindPendingRequests(
		ctx context.Context,
		paginateQuery sharedmodel.PaginateQuery,
	) ([]model.BorrowRequest, error)

	FindByRequestorID(
		ctx context.Context,
		requestorID int,
		paginateQuery sharedmodel.PaginateQuery,
	) ([]model.BorrowRequest, error)

	ApproveBorrowRequest(
		ctx context.Context,
		borrowRequestID int,
		respondentID int,
		assetID int,
		response *string,
	) error
}
