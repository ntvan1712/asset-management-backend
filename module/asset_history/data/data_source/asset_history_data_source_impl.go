package datasource

import (
	"asset_management_backend/common/error_app"
	"asset_management_backend/infras"
	"asset_management_backend/module/asset_history/data/model"
	"context"
)

type assetHistoryDataSourceImpl struct {
	dbProvider *infras.DbProvider
}

// DeleteByID implements AssetHistoryDataSource.
func (a assetHistoryDataSourceImpl) DeleteByID(ctx context.Context, historyID int) error {
	res, err := a.dbProvider.Instance.NewDelete().
		Table(model.TableAssetHistory).
		Where("id = ?", historyID).
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

// FindByAssetID implements AssetHistoryDataSource.
func (a assetHistoryDataSourceImpl) FindByAssetID(ctx context.Context, assetID int, page int, limit int) ([]model.AssetHistory, error) {
	offset := (page - 1) * limit

	var histories []model.AssetHistory
	err := a.dbProvider.Instance.NewSelect().
		Model(&histories).
		Relation("Creator").
		Relation("Borrower").
		Where("asset_id = ?", assetID).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Scan(ctx)

	if err != nil {
		return nil, err
	}
	return histories, nil
}

// Insert implements AssetHistoryDataSource.
func (a assetHistoryDataSourceImpl) Insert(ctx context.Context, newHistory model.AssetHistory) (*model.AssetHistory, error) {
	_, err := a.dbProvider.Instance.NewInsert().Model(&newHistory).Returning("*").Exec(ctx)
	if err != nil {
		if error_app.IsUniqueViolation(err) {
			return nil, error_app.ErrDuplicateKey
		}
		return nil, err
	}

	return &newHistory, nil
}

// UpdateByID implements AssetHistoryDataSource.
func (a assetHistoryDataSourceImpl) UpdateByID(ctx context.Context, historyID int, updateData map[string]interface{}) error {
	if len(updateData) == 0 {
		return nil
	}
	query := a.dbProvider.Instance.NewUpdate().
		Table(model.TableAssetHistory).
		Where("id = ?", historyID)

	infras.BuildUpdateQueryByMap(query, updateData)
	_, err := query.Exec(ctx)
	return err
}

func NewAssetDataSource(dbProvider *infras.DbProvider) AssetHistoryDataSource {
	return assetHistoryDataSourceImpl{dbProvider: dbProvider}
}
