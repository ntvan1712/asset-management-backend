package usecase

import (
	sharedmodel "asset_management_backend/common/shared_model"
	"asset_management_backend/module/transfer/domain/entity"
	"context"
)

type TransferUsecase interface {
	Create(ctx context.Context, newRequest entity.CreateTransferRequestEntity) (*entity.TransferRequestEntity, error)
	UpdateByID(ctx context.Context, requestID int, updateData map[string]interface{}) error
	DeleteByID(
		ctx context.Context,
		requestID int,
		requestorID int,
	) error

	GetAll(
		ctx context.Context,
		paginateQuery sharedmodel.PaginateQuery,
	) ([]entity.TransferRequestEntity, error)

	// Trả về những yêu cầu được luân chuyển tới `respondentID`
	GetMyTransferRequests(
		ctx context.Context,
		respondentID int,
		paginateQuery sharedmodel.PaginateQuery,
	) ([]entity.TransferRequestEntity, error)

	RejectTransferRequest(
		ctx context.Context,
		requestID int,
		respondentID int,
		responseDescription *string,
	) error

	ApproveTransferRequest(
		ctx context.Context,
		requestID int,
		respondentID int,
		responseDescription *string,
	) error

	CancelTransferRequest(
		ctx context.Context,
		requestID int,
		requestorID int,
	) error
}
