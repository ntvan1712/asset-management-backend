package datasource

import (
	"asset_management_backend/common/error_app"
	"asset_management_backend/module/category/data/model"
	"context"
	"fmt"

	"github.com/uptrace/bun"
)

type categoryDataSourceImpl struct {
	dbInstance *bun.DB
}

func (ds *categoryDataSourceImpl) FindAllCategories(ctx context.Context) (*model.AllCategories, error) {

	var result struct {
		AllCategories model.AllCategories `json:"all_categories"`
	}
	err := ds.dbInstance.NewRaw(`
        SELECT json_build_object(
            'asset_qualities', (SELECT json_agg(aq) FROM asset_qualities aq),
            'asset_types', (SELECT json_agg(at) FROM asset_types at),
            'locations', (SELECT json_agg(l) FROM locations l),
            'currencies', (SELECT json_agg(c) FROM currencies c)
        ) AS all_categories
    `).Scan(ctx, &result)
	if err != nil {
		return nil, err
	}

	return &result.AllCategories, nil
}

func (ds *categoryDataSourceImpl) FindAllAssetTypes(ctx context.Context) ([]model.AssetType, error) {
	var assetTypes []model.AssetType
	err := ds.dbInstance.NewSelect().Model(&assetTypes).Scan(ctx)
	return assetTypes, err
}

func (ds *categoryDataSourceImpl) DeleteAssetTypeByID(ctx context.Context, id int) error {
	res, err := ds.dbInstance.NewDelete().
		Model((*model.AssetType)(nil)).
		Where("id = ?", id).
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

func (ds *categoryDataSourceImpl) UpdateAssetType(ctx context.Context, id int, updateData map[string]interface{}) error {
	query := ds.dbInstance.NewUpdate().
		Model((*model.AssetType)(nil)).
		Where("id = ?", id)

	for key, value := range updateData {
		query = query.Set(fmt.Sprintf("%s = ?", key), value)
	}

	_, err := query.Exec(ctx)
	return err
}

func (ds *categoryDataSourceImpl) InsertAssetType(ctx context.Context, assetType *model.AssetType) (*model.AssetType, error) {
	_, err := ds.dbInstance.NewInsert().Model(assetType).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return assetType, nil
}

func (ds *categoryDataSourceImpl) FindAllAssetQualities(ctx context.Context) ([]model.AssetQuality, error) {
	var assetQualities []model.AssetQuality
	err := ds.dbInstance.NewSelect().Model(&assetQualities).Scan(ctx)
	return assetQualities, err
}

func (ds *categoryDataSourceImpl) InsertAssetQuality(ctx context.Context, assetQuality *model.AssetQuality) (*model.AssetQuality, error) {
	_, err := ds.dbInstance.NewInsert().Model(assetQuality).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return assetQuality, nil
}

func (ds *categoryDataSourceImpl) UpdateAssetQuality(ctx context.Context, id int, updateData map[string]interface{}) error {
	query := ds.dbInstance.NewUpdate().
		Model((*model.AssetQuality)(nil)).
		Where("id = ?", id)

	for key, value := range updateData {
		query = query.Set(fmt.Sprintf("%s = ?", key), value)
	}

	_, err := query.Exec(ctx)
	return err
}

func (ds *categoryDataSourceImpl) DeleteAssetQualityByID(ctx context.Context, id int) error {
	res, err := ds.dbInstance.NewDelete().
		Model((*model.AssetQuality)(nil)).
		Where("id = ?", id).
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

// Currency methods
func (ds *categoryDataSourceImpl) FindAllCurrencies(ctx context.Context) ([]model.Currency, error) {
	var currencies []model.Currency
	err := ds.dbInstance.NewSelect().Model(&currencies).Scan(ctx)
	return currencies, err
}

func (ds *categoryDataSourceImpl) FindAllLocations(ctx context.Context) ([]model.Location, error) {
	var locations []model.Location
	err := ds.dbInstance.NewSelect().Model(&locations).Scan(ctx)
	return locations, err
}

func (ds *categoryDataSourceImpl) InsertLocation(ctx context.Context, location *model.Location) (*model.Location, error) {
	_, err := ds.dbInstance.NewInsert().Model(location).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return location, nil
}

func (ds *categoryDataSourceImpl) UpdateLocation(ctx context.Context, id int, updateData map[string]interface{}) error {
	query := ds.dbInstance.NewUpdate().
		Model((*model.Location)(nil)).
		Where("id = ?", id)

	for key, value := range updateData {
		query = query.Set(fmt.Sprintf("%s = ?", key), value)
	}

	_, err := query.Exec(ctx)
	return err
}

func (ds *categoryDataSourceImpl) DeleteLocationByID(ctx context.Context, id int) error {
	res, err := ds.dbInstance.NewDelete().
		Model((*model.Location)(nil)).
		Where("id = ?", id).
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

func NewCategoryDataSource(dbInstance *bun.DB) *categoryDataSourceImpl {
	return &categoryDataSourceImpl{dbInstance: dbInstance}
}
