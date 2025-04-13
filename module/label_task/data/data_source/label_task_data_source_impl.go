package datasource

import (
	"asset_management_backend/common/error_app"
	"asset_management_backend/infras"
	"asset_management_backend/module/label_task/data/model"
	"context"
	"database/sql"
	"errors"
)

type labelTaskDataSourceImpl struct {
	dbProvider *infras.DbProvider
}

// DeleteByID implements AssetDataSource.
func (a labelTaskDataSourceImpl) DeleteByID(ctx context.Context, id int) error {
	panic("unimplemented")
}

// FindByID implements AssetDataSource.
func (a labelTaskDataSourceImpl) FindByID(ctx context.Context, id int) (*model.AssetLabelTask, error) {
	var result model.AssetLabelTask

	err := a.dbProvider.Instance.NewSelect().
		Model(&result).
		Where("id = ?", id).
		Scan(ctx)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, error_app.ErrDocumentNotFound
		}
		return nil, err
	}

	return &result, nil
}

// Insert implements AssetDataSource.
func (a labelTaskDataSourceImpl) Insert(ctx context.Context, labelTask model.AssetLabelTask) (*model.AssetLabelTask, error) {
	// Sử dụng Returning để lấy các trường của bản ghi đã được insert
	_, err := a.dbProvider.Instance.NewInsert().Model(&labelTask).Returning("*").Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &labelTask, nil
}

// UpdateByID implements AssetDataSource.
func (a labelTaskDataSourceImpl) UpdateByID(ctx context.Context, id int, updateData map[string]interface{}) (*model.AssetLabelTask, error) {
	panic("unimplemented")
}

func NewLabelTaskDataSource(dbProvider *infras.DbProvider) LabelTaskDataSource {
	return labelTaskDataSourceImpl{dbProvider: dbProvider}
}
