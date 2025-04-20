package datasource

import (
	"asset_management_backend/module/label_task/data/model"
	"context"
)

type LabelTaskDataSource interface {
	Insert(ctx context.Context, newLabelTask model.AssetLabelTask) (*model.AssetLabelTask, error)
	InsertMany(ctx context.Context, newLabelTasks []model.AssetLabelTask) ([]model.AssetLabelTask, error)
	UpdateByID(ctx context.Context, id int, updateData map[string]interface{}) error
	DeleteByID(ctx context.Context, id int) error

	FindByID(ctx context.Context, id int) (*model.AssetLabelTask, error)
}
