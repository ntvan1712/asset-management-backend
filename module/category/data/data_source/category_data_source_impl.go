package datasource

import (
	"asset_management_backend/infras"
	"asset_management_backend/module/category/data/model"
	"context"
)

type categoryDataSourceImpl struct {
	dbProvider *infras.DbProvider
}

func (ds *categoryDataSourceImpl) FindAllCategories(ctx context.Context) (*model.AllCategories, error) {

	var result struct {
		AllCategories model.AllCategories `json:"all_categories"`
	}
	err := ds.dbProvider.Instance.NewRaw(`
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
	err := ds.dbProvider.Instance.NewSelect().Model(&assetTypes).Scan(ctx)
	return assetTypes, err
}

func (ds *categoryDataSourceImpl) DeleteAssetTypeByID(ctx context.Context, id int) error {
	return infras.DeleteByID(ctx, ds.dbProvider.Instance, model.TableAssetType, id)
}

func (ds *categoryDataSourceImpl) UpdateAssetType(ctx context.Context, id int, updateData map[string]interface{}) error {
	if len(updateData) == 0 {
		return nil
	}
	query := ds.dbProvider.Instance.NewUpdate().
		Table(model.TableAssetType)

	infras.BuildUpdateQueryByMap(query, updateData)

	_, err := query.Where("id = ?", id).Exec(ctx)
	return err
}

func (ds *categoryDataSourceImpl) InsertAssetType(ctx context.Context, assetType *model.AssetType) (*model.AssetType, error) {
	_, err := ds.dbProvider.Instance.NewInsert().Model(assetType).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return assetType, nil
}

func (ds *categoryDataSourceImpl) FindAllAssetQualities(ctx context.Context) ([]model.AssetQuality, error) {
	var assetQualities []model.AssetQuality
	err := ds.dbProvider.Instance.NewSelect().Model(&assetQualities).Scan(ctx)
	return assetQualities, err
}

func (ds *categoryDataSourceImpl) InsertAssetQuality(ctx context.Context, assetQuality *model.AssetQuality) (*model.AssetQuality, error) {
	_, err := ds.dbProvider.Instance.NewInsert().Model(assetQuality).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return assetQuality, nil
}

func (ds *categoryDataSourceImpl) UpdateAssetQuality(ctx context.Context, id int, updateData map[string]interface{}) error {
	if len(updateData) == 0 {
		return nil
	}
	query := ds.dbProvider.Instance.NewUpdate().
		Table(model.TableAssetQuality)
		

	infras.BuildUpdateQueryByMap(query, updateData)

	_, err := query.Where("id = ?", id).Exec(ctx)
	return err
}

func (ds *categoryDataSourceImpl) DeleteAssetQualityByID(ctx context.Context, id int) error {
	return infras.DeleteByID(ctx, ds.dbProvider.Instance, model.TableAssetQuality, id)
}

// Currency methods
func (ds *categoryDataSourceImpl) FindAllCurrencies(ctx context.Context) ([]model.Currency, error) {
	var currencies []model.Currency
	err := ds.dbProvider.Instance.NewSelect().Model(&currencies).Scan(ctx)
	return currencies, err
}

func (ds *categoryDataSourceImpl) FindAllLocations(ctx context.Context) ([]model.Location, error) {
	var locations []model.Location
	err := ds.dbProvider.Instance.NewSelect().Model(&locations).Scan(ctx)
	return locations, err
}

func (ds *categoryDataSourceImpl) InsertLocation(ctx context.Context, location *model.Location) (*model.Location, error) {
	_, err := ds.dbProvider.Instance.NewInsert().Model(location).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return location, nil
}

func (ds *categoryDataSourceImpl) UpdateLocation(ctx context.Context, id int, updateData map[string]interface{}) error {
	if len(updateData) == 0 {
		return nil
	}
	query := ds.dbProvider.Instance.NewUpdate().
		Table(model.TableLocation)
		

	infras.BuildUpdateQueryByMap(query, updateData)

	_, err := query.Where("id = ?", id).Exec(ctx)
	return err
}

func (ds *categoryDataSourceImpl) DeleteLocationByID(ctx context.Context, id int) error {
	return infras.DeleteByID(ctx, ds.dbProvider.Instance, model.TableLocation, id)
}

func NewCategoryDataSource(dbProvider *infras.DbProvider) *categoryDataSourceImpl {
	return &categoryDataSourceImpl{dbProvider: dbProvider}
}
