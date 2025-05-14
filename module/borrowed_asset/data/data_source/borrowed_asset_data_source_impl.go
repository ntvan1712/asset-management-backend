package datasource

import (
	"asset_management_backend/common/enums"
	"asset_management_backend/infras"
	assetM "asset_management_backend/module/asset/data/model"
	borrowedAssetM "asset_management_backend/module/borrowed_asset/data/model"
	"context"

	"github.com/uptrace/bun"
)

type borrowedAssetDataSourceImpl struct {
	dbProvider *infras.DbProvider
}

// ReturnAsset implements BorrowedAssetDataSource.
func (b borrowedAssetDataSourceImpl) ReturnAsset(ctx context.Context, borrowedAssetId int) (*borrowedAssetM.BorrowedAsset, error) {
	tx, err := b.dbProvider.Instance.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Xóa và lấy nguyên BorrowedAsset
	var borrowedAsset borrowedAssetM.BorrowedAsset
	_, err = tx.NewDelete().
		Model(&borrowedAsset).
		Where("id = ?", borrowedAssetId).
		Returning("*").
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	// Update asset status
	if err = updateAssetStatus(ctx, tx, enums.AssetStatusEnum.Available, borrowedAsset.AssetID); err != nil {
		return nil, err
	}

	// Commit
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &borrowedAsset, nil
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
func NewBorrowedAssetDataSource(dbProvider *infras.DbProvider) BorrowedAssetDataSource {
	return borrowedAssetDataSourceImpl{dbProvider: dbProvider}
}
