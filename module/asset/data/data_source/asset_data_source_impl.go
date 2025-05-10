package datasource

import (
	"asset_management_backend/common/error_app"
	sharedmodel "asset_management_backend/common/shared_model"
	"asset_management_backend/infras"
	assetM "asset_management_backend/module/asset/data/model"
	"asset_management_backend/module/borrowed_asset/data/model"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/uptrace/bun"
)

type assetDataSourceImpl struct {
	dbProvider *infras.DbProvider
}

// FindAssetByBorrowerID implements AssetDataSource.
func (a assetDataSourceImpl) FindAssetByBorrowerID(ctx context.Context, borrowerID int,paginateQuery sharedmodel.PaginateQuery) ([]assetM.Asset, error) {
	var assets []assetM.Asset
	query := a.dbProvider.Instance.NewSelect().Model(&assets)
	assetM.LoadAllAssetRelationQuery(query)

	err := query.Join(fmt.Sprintf("JOIN %s ba ON ba.asset_id = asset.id", model.TableBorrowedAsset)).
		Where("ba.borrower_id = ?", borrowerID).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return assets, nil
}

// GetSerialNumberByID implements AssetDataSource.
func (a assetDataSourceImpl) GetSerialNumberByID(ctx context.Context, assetID int) (*string, error) {
	var asset assetM.Asset
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
	if len(updateData) == 0 {
		return nil
	}
	query := a.dbProvider.Instance.NewUpdate().
		Table(assetM.TableAssetLabelImage)

	infras.BuildUpdateQueryByMap(query, updateData)

	_, err := query.Where("asset_id = ?", assetID).Exec(ctx)
	return err
}

// UpdateLabelImageByID implements AssetDataSource.
func (a assetDataSourceImpl) UpdateLabelImageByID(ctx context.Context, id int, updateData map[string]interface{}) error {
	if len(updateData) == 0 {
		return nil
	}
	query := a.dbProvider.Instance.NewUpdate().
		Table(assetM.TableAssetLabelImage)

	infras.BuildUpdateQueryByMap(query, updateData)

	_, err := query.Where("id = ?", id).Exec(ctx)
	return err
}

// DeleteAssetFilesByIDs implements AssetDataSource.
func (a assetDataSourceImpl) DeleteAssetFilesByIDs(ctx context.Context, ids []int) error {
	_, err := a.dbProvider.Instance.NewDelete().
		Table("asset_files").
		Where("id IN (?)", bun.In(ids)).
		Exec(ctx)
	return err
}

// InsertLabelImage implements AssetDataSource.
func (a assetDataSourceImpl) InsertLabelImage(ctx context.Context, newLabelImage assetM.AssetLabelImage) (*assetM.AssetLabelImage, error) {
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
func (a assetDataSourceImpl) InsertAssetFiles(ctx context.Context, assetFiles []assetM.AssetFile) ([]assetM.AssetFile, error) {
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
	filter assetM.AssetFilterModel,
) ([]assetM.Asset, error) {
	var assets []assetM.Asset

	query := a.dbProvider.Instance.NewSelect().Model(&assets)
	assetM.LoadAllAssetRelationQuery(query)

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
func (a assetDataSourceImpl) FindByID(ctx context.Context, assetID int) (*assetM.Asset, error) {
	var result assetM.Asset
	selectModelQuery := a.dbProvider.Instance.NewSelect().Model(&result)
	assetM.LoadAllAssetRelationQuery(selectModelQuery)
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
func (a assetDataSourceImpl) Insert(ctx context.Context, newAsset assetM.Asset) (*assetM.Asset, error) {
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
	if len(updateData) == 0 {
		return nil
	}
	query := a.dbProvider.Instance.NewUpdate().
		Table(assetM.TableAsset)

	infras.BuildUpdateQueryByMap(query, updateData)

	_, err := query.Where("id = ?", assetID).Exec(ctx)
	return err
}

func NewAssetDataSource(dbProvider *infras.DbProvider) AssetDataSource {
	return assetDataSourceImpl{dbProvider: dbProvider}
}
