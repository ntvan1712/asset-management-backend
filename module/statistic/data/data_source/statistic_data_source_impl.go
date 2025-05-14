package datasource

import (
	"asset_management_backend/infras"
	assetM "asset_management_backend/module/asset/data/model"
	categoryM "asset_management_backend/module/category/data/model"
	statisticM "asset_management_backend/module/statistic/data/model"
	"context"
	"fmt"
)

type statisticDataSourceImpl struct {
	dbProvider *infras.DbProvider
}

// GetAssetStatisticByStatus implements StatisticDataSource.
func (s *statisticDataSourceImpl) GetAssetStatisticByStatus(ctx context.Context) (*statisticM.CategoryStatisticResponse, error) {
	var items []statisticM.CategoryStatisticModel
	err := s.dbProvider.Instance.NewSelect().
		Table(assetM.TableAsset).
		ColumnExpr("status AS name").
		ColumnExpr("COUNT(id) AS count").
		GroupExpr("status").
		OrderExpr("count DESC").
		Scan(ctx, &items)

	if err != nil {
		return nil, err
	}

	totalAssetCount := 0
	for _, item := range items {
		totalAssetCount += item.Count
	}

	return &statisticM.CategoryStatisticResponse{
		TotalAssetCount: totalAssetCount,
		Items:           items,
	}, nil
}

// GetAssetStatisticByLocation implements StatisticDataSource.
func (s *statisticDataSourceImpl) GetAssetStatisticByLocation(ctx context.Context) (*statisticM.CategoryStatisticResponse, error) {
	return s.getAssetStatisticByCategory(ctx, categoryM.TableLocation, "location_id")
}

// GetAssetStatisticByQuality implements StatisticDataSource.
func (s *statisticDataSourceImpl) GetAssetStatisticByQuality(ctx context.Context) (*statisticM.CategoryStatisticResponse, error) {
	return s.getAssetStatisticByCategory(ctx, categoryM.TableAssetQuality, "asset_quality_id")
}

// GetAssetStatisticByType implements StatisticDataSource.
func (s *statisticDataSourceImpl) GetAssetStatisticByType(ctx context.Context) (*statisticM.CategoryStatisticResponse, error) {
	return s.getAssetStatisticByCategory(ctx, categoryM.TableAssetType, "asset_type_id")
}

func (s *statisticDataSourceImpl) getAssetStatisticByCategory(
	ctx context.Context,
	categoryTable string,
	categoryColumnName string,
) (*statisticM.CategoryStatisticResponse, error) {
	var items []statisticM.CategoryStatisticModel

	err := s.dbProvider.Instance.NewSelect().
		Table(assetM.TableAsset).
		ColumnExpr(fmt.Sprintf("%s AS id", categoryColumnName)).
		ColumnExpr(fmt.Sprintf("%s.name AS name", categoryTable)).
		ColumnExpr("COUNT(assets.id) AS count").
		Join(fmt.Sprintf("LEFT JOIN %s ON %s.id = assets.%s", categoryTable, categoryTable, categoryColumnName)).
		GroupExpr(fmt.Sprintf("%s, %s.name", categoryColumnName, categoryTable)).
		OrderExpr("count DESC").
		Scan(ctx, &items)
	if err != nil {
		return nil, err
	}
	totalAssetCount := 0
	for _, item := range items {
		totalAssetCount += item.Count
	}

	return &statisticM.CategoryStatisticResponse{
		TotalAssetCount: totalAssetCount,
		Items:           items,
	}, nil
}

func NewStatisticDataSource(dbProvider *infras.DbProvider) StatisticDataSource {
	return &statisticDataSourceImpl{
		dbProvider: dbProvider,
	}
}
