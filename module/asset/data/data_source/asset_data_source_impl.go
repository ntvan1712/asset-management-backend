package datasource

import (
	"asset_management_backend/common/error_app"
	"asset_management_backend/infras"
	"asset_management_backend/module/asset/data/model"
	"context"
	"database/sql"
	"errors"

	"github.com/uptrace/bun"
)

type assetDataSourceImpl struct {
	dbProvider *infras.DbProvider
}

// GetSerialNumberByID implements AssetDataSource.
func (a assetDataSourceImpl) GetSerialNumberByID(ctx context.Context, assetID int) (*string, error) {
	var asset model.Asset
	err := a.dbProvider.Instance.NewSelect().
		Model(&asset).
		Column("serial_number").
		Where("id = ?", assetID).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return &asset.SerialNumber, nil
}

// UpdateLabelImageByAssetID implements AssetDataSource.
func (a assetDataSourceImpl) UpdateLabelImageByAssetID(ctx context.Context, assetID int, updateData map[string]interface{}) error {
	query := a.dbProvider.Instance.NewUpdate().
		Model(&model.AssetLabelImage{}).
		Where("asset_id = ?", assetID)

	infras.BuildUpdateQueryByMap(query, updateData)

	_, err := query.Exec(ctx)
	return err
}

// UpdateLabelImageByID implements AssetDataSource.
func (a assetDataSourceImpl) UpdateLabelImageByID(ctx context.Context, id int, updateData map[string]interface{}) error {
	query := a.dbProvider.Instance.NewUpdate().
		Model(&model.AssetLabelImage{}).
		Where("id = ?", id)

	infras.BuildUpdateQueryByMap(query, updateData)

	_, err := query.Exec(ctx)
	return err
}

// DeleteAssetFilesByIDs implements AssetDataSource.
func (a assetDataSourceImpl) DeleteAssetFilesByIDs(ctx context.Context, ids []int) error {
	_, err := a.dbProvider.Instance.NewDelete().
		Model(&model.AssetFile{}).
		Where("id IN (?)", bun.In(ids)).
		Exec(ctx)
	return err
}

// InsertLabelImage implements AssetDataSource.
func (a assetDataSourceImpl) InsertLabelImage(ctx context.Context, newLabelImage model.AssetLabelImage) (*model.AssetLabelImage, error) {
	// Sử dụng Returning để lấy các trường của bản ghi đã được insert
	_, err := a.dbProvider.Instance.NewInsert().Model(&newLabelImage).Returning("*").Exec(ctx)
	if err != nil {
		if error_app.IsUniqueViolation(err) {
			return nil, error_app.ErrDuplicateKey
		}
		return nil, err
	}

	return &newLabelImage, nil
}

// InsertAssetFiles implements AssetDataSource.
func (a assetDataSourceImpl) InsertAssetFiles(ctx context.Context, assetFiles []model.AssetFile) ([]model.AssetFile, error) {
	_, err := a.dbProvider.Instance.NewInsert().Model(&assetFiles).Returning("*").Exec(ctx)
	if err != nil {
		return nil, err
	}

	return assetFiles, nil
}

// DeleteByID implements AssetDataSource.
func (a assetDataSourceImpl) DeleteByID(ctx context.Context, assetID int) error {
	panic("unimplemented")
}

// FindByFilter implements AssetDataSource.
func (a assetDataSourceImpl) FindByFilter(
	ctx context.Context,
	filter model.AssetFilterModel,
) ([]model.Asset, error) {
	var assets []model.Asset

	query := a.dbProvider.Instance.NewSelect().Model(&assets)
	model.LoadAllAssetRelationQuery(query)

	if filter.SerialNumber != nil {
		query = query.Where("serial_number ILIKE ?", "%"+*filter.SerialNumber+"%")
	}
	if filter.Status != nil {
		query = query.Where("status = ?", *filter.Status)
	}
	if filter.LocationID != nil {
		query = query.Where("location_id = ?", *filter.LocationID)
	}
	if filter.AssetQualityID != nil {
		query = query.Where("asset_quality_id = ?", *filter.AssetQualityID)
	}
	if filter.AssetTypeID != nil {
		query = query.Where("asset_type_id = ?", *filter.AssetTypeID)
	}
	if filter.FromAddedAt != nil {
		query = query.Where("added_at >= ?", *filter.FromAddedAt)
	}
	if filter.ToAddedAt != nil {
		query = query.Where("added_at <= ?", *filter.ToAddedAt)
	}

	query = query.Offset(filter.Offset).Limit(filter.Limit)

	query = query.OrderExpr(filter.OrderExpr)

	err := query.Scan(ctx)
	return assets, err
}

// FindByID implements AssetDataSource.
func (a assetDataSourceImpl) FindByID(ctx context.Context, assetID int) (*model.Asset, error) {
	var result model.Asset
	selectModelQuery := a.dbProvider.Instance.NewSelect().Model(&result)
	model.LoadAllAssetRelationQuery(selectModelQuery)
	err := selectModelQuery.Where("asset.id = ?", assetID).Scan(ctx)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, error_app.ErrDocumentNotFound
		}
		return nil, err
	}

	return &result, nil
}

// Insert implements AssetDataSource.
func (a assetDataSourceImpl) Insert(ctx context.Context, newAsset model.Asset) (*model.Asset, error) {
	// Sử dụng Returning để lấy các trường của bản ghi đã được insert
	_, err := a.dbProvider.Instance.NewInsert().Model(&newAsset).Returning("*").Exec(ctx)
	if err != nil {
		if error_app.IsUniqueViolation(err) {
			return nil, error_app.ErrDuplicateKey
		}
		return nil, err
	}

	return &newAsset, nil
}

// UpdateByID implements AssetDataSource.
func (a assetDataSourceImpl) UpdateByID(ctx context.Context, assetID int, updateData map[string]interface{}) error {
	panic("unimplemented")
}

func NewAssetDataSource(dbProvider *infras.DbProvider) AssetDataSource {
	return assetDataSourceImpl{dbProvider: dbProvider}
}
