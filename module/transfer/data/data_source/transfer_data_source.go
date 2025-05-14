package datasource

import (
	sharedmodel "asset_management_backend/common/shared_model"
	"asset_management_backend/module/transfer/data/model"
	"context"
)

type TransferDataSource interface {
	Insert(ctx context.Context, newRequest model.TransferRequest) (*model.TransferRequest, error)
	UpdateByID(ctx context.Context, requestID int, updateData map[string]interface{}) error
	DeleteByID(
		ctx context.Context,
		requestID int,
		requestorID int,
	) error

	// Trả về tất cả yêu cầu luân chuyển
	FindAll(
		ctx context.Context,
		paginateQuery sharedmodel.PaginateQuery,
	) ([]model.TransferRequest, error)

	// Trả về những yêu cầu được luân chuyển tới `respondentID`
	FindByRespondentID(
		ctx context.Context,
		respondentID int,
		paginateQuery sharedmodel.PaginateQuery,
	) ([]model.TransferRequest, error)

	RejectTransferRequest(
		ctx context.Context,
		requestID int,
		respondentID int,
		responseDescription *string,
	) (*model.TransferRequest, error)

	ApproveTransferRequest(
		ctx context.Context,
		requestID int,
		respondentID int,
		responseDescription *string,
	) (*model.TransferRequest, error)

	CancelTransferRequest(
		ctx context.Context,
		requestID int,
		requestorID int,
	) (*model.TransferRequest, error)
}
