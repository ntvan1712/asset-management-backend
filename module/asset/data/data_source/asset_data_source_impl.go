package datasource

import (
	"asset_management_backend/common/error_app"
	"asset_management_backend/module/asset/data/model"
	"context"
	"database/sql"
	"errors"

	"github.com/uptrace/bun"
)

type assetDataSourceImpl struct {
	dbInstance *bun.DB
}

// InsertAssetFiles implements AssetDataSource.
func (a assetDataSourceImpl) InsertAssetFiles(ctx context.Context, assetFiles []model.AssetFile) ([]model.AssetFile, error) {
	_, err := a.dbInstance.NewInsert().Model(&assetFiles).Returning("*").Exec(ctx)
	if err != nil {
		return nil, err
	}

	// Trả về các AssetFile sau khi insert
	return assetFiles, nil
}

// DeleteByID implements AssetDataSource.
func (a assetDataSourceImpl) DeleteByID(ctx context.Context, assetID int) error {
	panic("unimplemented")
}

// FindByFilter implements AssetDataSource.
func (a assetDataSourceImpl) FindByFilter(ctx context.Context, assetTypeID *int, assetQualityID *int, page int, limit int) ([]model.Asset, error) {
	panic("unimplemented")
}

// FindByID implements AssetDataSource.
func (a assetDataSourceImpl) FindByID(ctx context.Context, assetID int) (*model.Asset, error) {
	var result model.Asset
	err := a.dbInstance.NewSelect().
		Model(&result).
		Relation("PriceUnit").
		Relation("Location").
		Relation("AssetQuality").
		Relation("AssetType").
		Relation("AssetFiles").
		Where("asset.id = ?", assetID).
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
func (a assetDataSourceImpl) Insert(ctx context.Context, newAsset model.Asset) (*model.Asset, error) {
	// Sử dụng Returning để lấy các trường của bản ghi đã được insert
	_, err := a.dbInstance.NewInsert().Model(&newAsset).Returning("*").Exec(ctx)
	if err != nil {
		if error_app.IsUniqueViolation(err) {
			return nil, error_app.ErrDuplicateKey
		}
		return nil, err
	}

	return &newAsset, nil
}

// UpdateByID implements AssetDataSource.
func (a assetDataSourceImpl) UpdateByID(ctx context.Context, assetID int, updateData map[string]interface{}) (*model.Asset, error) {
	panic("unimplemented")
}

func NewAssetDataSource(dbInstance *bun.DB) AssetDataSource {
	return assetDataSourceImpl{dbInstance: dbInstance}
}
