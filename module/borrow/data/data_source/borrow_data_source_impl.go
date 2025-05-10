package datasource

import (
	"asset_management_backend/common/enums"
	"asset_management_backend/common/error_app"
	sharedmodel "asset_management_backend/common/shared_model"
	app_utils "asset_management_backend/common/utils"
	"asset_management_backend/infras"
	assetM "asset_management_backend/module/asset/data/model"
	borrowM "asset_management_backend/module/borrow/data/model"
	transferM "asset_management_backend/module/transfer/data/model"
	"context"
	"time"
)

type borrowDataSourceImpl struct {
	dbProvider *infras.DbProvider
}

// ApproveBorrowRequest implements BorrowDataSource.
func (b borrowDataSourceImpl) ApproveBorrowRequest(ctx context.Context, borrowRequestID int, respondentID int, assetID int, response *string) error {
	tx, err := b.dbProvider.Instance.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback() // rollback nếu có lỗi

	var assetStatus string
	err = tx.NewSelect().
		Table(assetM.TableAsset).
		Column("status").
		Where("id = ?", assetID).
		Scan(ctx, &assetStatus)
	if err != nil {
		return err
	}

	if assetStatus != enums.AssetStatusEnum.Available {
		return error_app.ErrDuplicateKey
	}

	_, err = tx.NewUpdate().
		Table(borrowM.TableBorrowRequest).
		Set("status = ?", enums.BorrowRequestEnum.Approved).
		Set("respondent_id = ?", respondentID).
		Set("response_at = ?", time.Now()).
		Set("response_description = ?", response).
		Where("id = ?", borrowRequestID).
		Exec(ctx)
	if err != nil {
		return err
	}

	// Lấy borrow request để tạo transfer request
	var borrowRequest borrowM.BorrowRequest
	err = tx.NewSelect().
		Model(&borrowRequest).
		Where("id = ?", borrowRequestID).
		Scan(ctx)
	if err != nil {
		return err
	}

	// Insert transfer_requests
	transfer := &transferM.TransferRequest{
		RequestAt:           app_utils.TimeNowPtr(),
		ReturnDate:          borrowRequest.ReturnDate,
		Status:              enums.TransferRequestEnum.Pending,
		RequestorID:         respondentID,
		UseAtLocationID:     borrowRequest.UseAtLocationID,
		AssetID:             assetID,
		RespondentID:        &borrowRequest.RequestorID,
		FromBorrowRequestID: &borrowRequestID,
	}

	_, err = tx.NewInsert().
		Model(transfer).
		Exec(ctx)
	if err != nil {
		return err
	}

	// Update asset status
	_, err = tx.NewUpdate().
		Table(assetM.TableAsset).
		Set("status = ?", enums.AssetStatusEnum.AwaitingTransfer).
		Where("id = ?", assetID).
		Exec(ctx)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// FindByRequestorID implements BorrowDataSource.
func (b borrowDataSourceImpl) FindByRequestorID(ctx context.Context, requestorID int, paginateQuery sharedmodel.PaginateQuery) ([]borrowM.BorrowRequest, error) {
	offset := (paginateQuery.Page - 1) * paginateQuery.Limit

	var requests []borrowM.BorrowRequest
	err := b.dbProvider.Instance.NewSelect().
		Model(&requests).
		Relation("Requestor").
		Relation("UseAtLocation").
		Relation("AssetType").
		Relation("Respondent").
		Where("requestor_id = ?", requestorID).
		Order("request_at DESC").
		Offset(offset).
		Limit(paginateQuery.Limit).
		Scan(ctx)

	if err != nil {
		return nil, err
	}
	return requests, nil
}

// DeleteByID implements BorrowDataSource.
func (b borrowDataSourceImpl) DeleteByID(
	ctx context.Context,
	borrowRequestID int,
	requestorID int,
) error {
	res, err := b.dbProvider.Instance.NewDelete().
		Table(borrowM.TableBorrowRequest).
		Where("id = ? AND requestor_id = ?", borrowRequestID, requestorID).
		Exec(ctx)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return error_app.ErrDocumentNotFound
	}
	return nil
}

// FindPendingRequests implements BorrowDataSource.
func (b borrowDataSourceImpl) FindPendingRequests(ctx context.Context, paginateQuery sharedmodel.PaginateQuery) ([]borrowM.BorrowRequest, error) {
	offset := (paginateQuery.Page - 1) * paginateQuery.Limit

	var requests []borrowM.BorrowRequest
	err := b.dbProvider.Instance.NewSelect().
		Model(&requests).
		Relation("Requestor").
		Relation("UseAtLocation").
		Relation("AssetType").
		Relation("Respondent").
		Where("status = ?", enums.BorrowRequestEnum.Pending).
		Order("request_at DESC").
		Offset(offset).
		Limit(paginateQuery.Limit).
		Scan(ctx)

	if err != nil {
		return nil, err
	}
	return requests, nil
}

// Insert implements BorrowDataSource.
func (b borrowDataSourceImpl) Insert(ctx context.Context, newBorrowRequest borrowM.BorrowRequest) (*borrowM.BorrowRequest, error) {
	_, err := b.dbProvider.Instance.NewInsert().Model(&newBorrowRequest).Returning("*").Exec(ctx)
	if err != nil {
		if error_app.IsUniqueViolation(err) {
			return nil, error_app.ErrDuplicateKey
		}
		return nil, err
	}

	return &newBorrowRequest, nil
}

// UpdateByID implements BorrowDataSource.
func (b borrowDataSourceImpl) UpdateByID(ctx context.Context, borrowRequestID int, updateData map[string]interface{}) error {
	if len(updateData) == 0 {
		return nil
	}
	query := b.dbProvider.Instance.NewUpdate().
		Table(borrowM.TableBorrowRequest).
		Where("id = ?", borrowRequestID)

	infras.BuildUpdateQueryByMap(query, updateData)

	_, err := query.Exec(ctx)
	return err
}

func NewBorrowDataSource(dbProvider *infras.DbProvider) BorrowDataSource {
	return borrowDataSourceImpl{dbProvider: dbProvider}
}
