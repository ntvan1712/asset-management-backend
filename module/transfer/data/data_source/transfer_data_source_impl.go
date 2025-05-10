package datasource

import (
	"asset_management_backend/common/enums"
	"asset_management_backend/common/error_app"
	sharedmodel "asset_management_backend/common/shared_model"
	app_utils "asset_management_backend/common/utils"
	"asset_management_backend/infras"
	assetM "asset_management_backend/module/asset/data/model"
	"asset_management_backend/module/borrow/data/model"
	borrowedAssetM "asset_management_backend/module/borrowed_asset/data/model"
	transferM "asset_management_backend/module/transfer/data/model"
	"context"
	"time"

	"github.com/uptrace/bun"
)

type transferDataSourceImpl struct {
	dbProvider *infras.DbProvider
}

// CancelTransferRequest implements TransferDataSource.
func (t transferDataSourceImpl) CancelTransferRequest(ctx context.Context, requestID int, requestorID int) error {
	tx, err := t.dbProvider.Instance.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback() // rollback nếu có lỗi

	// Xóa yêu cầu luân chuyển
	var transferRequest transferM.TransferRequest
	_, err = tx.NewDelete().
		Model(&transferRequest).
		Where("id = ? AND requestor_id = ?", requestID, requestorID).
		Returning("*").
		Exec(ctx)
	if err != nil {
		return err
	}
	// Cập nhật trạng thái asset
	if err := updateAssetStatus(ctx, tx, enums.AssetStatusEnum.Available, transferRequest.AssetID); err != nil {
		return err
	}

	// Cập nhật trạng thái yêu cầu mượn
	if transferRequest.FromBorrowRequestID != nil {
		updateMap := map[string]interface{}{
			"status":               enums.BorrowRequestEnum.Pending,
			"respondent_id":        nil,
			"response_at":          nil,
			"response_description": nil,
		}

		query := tx.NewUpdate().
			Table(model.TableBorrowRequest)

		infras.BuildUpdateQueryByMap(query, updateMap)

		if _, err := query.Where("id = ?", transferRequest.FromBorrowRequestID).Exec(ctx); err != nil {
			return err
		}
	}

	return tx.Commit()
}

// ApproveTransferRequest implements TransferDataSource.
func (t transferDataSourceImpl) ApproveTransferRequest(
	ctx context.Context,
	requestID int,
	respondentID int,
	responseDescription *string,
) error {
	tx, err := t.dbProvider.Instance.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback() // rollback nếu có lỗi

	// Cập nhật phản hồi transfer request
	updateMap := map[string]interface{}{
		"status":               enums.TransferRequestEnum.Approved,
		"respondent_id":        respondentID,
		"response_at":          time.Now(),
		"response_description": responseDescription,
	}
	var transferRequest transferM.TransferRequest
	query := tx.NewUpdate().Model(&transferRequest)
	infras.BuildUpdateQueryByMap(query, updateMap)
	if _, err := query.Where("id = ?", requestID).Returning("*").Exec(ctx); err != nil {
		return err
	}

	// Cập nhật trạng thái asset thành đã cho mượn
	if err := updateAssetStatus(ctx, tx, enums.AssetStatusEnum.OnBorrow, transferRequest.AssetID); err != nil {
		return err
	}

	//Insert new borrowed asset
	transfer := &borrowedAssetM.BorrowedAsset{
		ID:          requestID,
		BorrowDate:  app_utils.TimeNowPtr(),
		ReturnDate:  transferRequest.ReturnDate,
		Description: nil,
		BorrowerID:  respondentID,
		AcceptorID:  transferRequest.RequestorID,
		AssetID:     transferRequest.AssetID,
	}

	_, err = tx.NewInsert().
		Model(transfer).
		Exec(ctx)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// RejectTransferRequest implements TransferDataSource.
func (t transferDataSourceImpl) RejectTransferRequest(
	ctx context.Context,
	requestID int,
	respondentID int,
	responseDescription *string,
) error {
	tx, err := t.dbProvider.Instance.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback() // rollback nếu có lỗi

	// Cập nhật phản hồi transfer request
	updateMap := map[string]interface{}{
		"status":               enums.TransferRequestEnum.Rejected,
		"respondent_id":        respondentID,
		"response_at":          time.Now(),
		"response_description": responseDescription,
	}
	var assetID int
	query := tx.NewUpdate().
		Table(transferM.TableTransferRequest)

	infras.BuildUpdateQueryByMap(query, updateMap)
	if err := query.Where("id = ?", requestID).Returning("asset_id").Scan(ctx, &assetID); err != nil {
		return err
	}

	// Cập nhật trạng thái asset thành có thể cho mượn
	if err := updateAssetStatus(ctx, tx, enums.AssetStatusEnum.Available, assetID); err != nil {
		return err
	}

	return tx.Commit()
}

// DeleteByID implements TransferDataSource.
func (t transferDataSourceImpl) DeleteByID(ctx context.Context, requestID int, requestorID int) error {
	res, err := t.dbProvider.Instance.NewDelete().
		Table(transferM.TableTransferRequest).
		Where("id = ? AND requestor_id = ?", requestID, requestorID).
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

// FindByRespondentID implements TransferDataSource.
func (t transferDataSourceImpl) FindByRespondentID(ctx context.Context, respondentID int, paginateQuery sharedmodel.PaginateQuery) ([]transferM.TransferRequest, error) {
	offset := (paginateQuery.Page - 1) * paginateQuery.Limit

	var requests []transferM.TransferRequest
	query := t.dbProvider.Instance.NewSelect().Model(&requests)
	loadTransferRelation(query)

	err := query.Where("respondent_id = ?", respondentID).
		Order("request_at DESC").
		Offset(offset).
		Limit(paginateQuery.Limit).
		Scan(ctx)

	if err != nil {
		return nil, err
	}
	return requests, nil
}

// FindAll implements TransferDataSource.
func (t transferDataSourceImpl) FindAll(ctx context.Context, paginateQuery sharedmodel.PaginateQuery) ([]transferM.TransferRequest, error) {
	offset := (paginateQuery.Page - 1) * paginateQuery.Limit

	var requests []transferM.TransferRequest

	query := t.dbProvider.Instance.NewSelect().Model(&requests)
	loadTransferRelation(query)

	err := query.Order("request_at DESC").
		Offset(offset).
		Limit(paginateQuery.Limit).
		Scan(ctx)

	if err != nil {
		return nil, err
	}
	return requests, nil
}

// Insert implements TransferDataSource.
func (t transferDataSourceImpl) Insert(ctx context.Context, newRequest transferM.TransferRequest) (*transferM.TransferRequest, error) {
	_, err := t.dbProvider.Instance.NewInsert().Model(&newRequest).Returning("*").Exec(ctx)
	if err != nil {
		if error_app.IsUniqueViolation(err) {
			return nil, error_app.ErrDuplicateKey
		}
		return nil, err
	}

	return &newRequest, nil
}

// UpdateByID implements TransferDataSource.
func (t transferDataSourceImpl) UpdateByID(ctx context.Context, requestID int, updateData map[string]interface{}) error {
	if len(updateData) == 0 {
		return nil
	}
	query := t.dbProvider.Instance.NewUpdate().
		Table(transferM.TableTransferRequest)

	infras.BuildUpdateQueryByMap(query, updateData)

	_, err := query.Where("id = ?", requestID).Exec(ctx)
	return err
}

func loadTransferRelation(query *bun.SelectQuery) *bun.SelectQuery {
	return query.
		Relation("Requestor").
		Relation("UseAtLocation").
		Relation("Asset").
		Relation("Asset.PriceUnit").
		Relation("Asset.Location").
		Relation("Asset.AssetQuality").
		Relation("Asset.AssetType").
		Relation("Asset.AssetFiles").
		Relation("Asset.AssetLabelImage").
		Relation("Asset.BorrowedAsset").
		Relation("Asset.BorrowedAsset.Borrower").
		Relation("Asset.BorrowedAsset.Acceptor").
		Relation("Respondent")
}

func updateAssetStatus(
	ctx context.Context,
	tx bun.Tx,
	status string,
	assetID int,
) error {
	_, err := tx.NewUpdate().
		Table(assetM.TableAsset).
		Set("status = ?", status).
		Where("id = ?", assetID).
		Exec(ctx)
	return err

}

func NewTransferDataSource(dbProvider *infras.DbProvider) TransferDataSource {
	return transferDataSourceImpl{dbProvider: dbProvider}
}
