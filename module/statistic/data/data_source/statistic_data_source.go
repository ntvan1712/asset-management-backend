package datasource

import (
	"asset_management_backend/module/statistic/data/model"
	"context"
)

type StatisticDataSource interface {
	GetAssetStatisticByType(ctx context.Context) (*model.CategoryStatisticResponse, error)
	GetAssetStatisticByQuality(ctx context.Context) (*model.CategoryStatisticResponse, error)
	GetAssetStatisticByLocation(ctx context.Context) (*model.CategoryStatisticResponse, error)
	GetAssetStatisticByStatus(ctx context.Context) (*model.CategoryStatisticResponse, error)
}