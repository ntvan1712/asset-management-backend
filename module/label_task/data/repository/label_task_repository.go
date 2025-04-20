package repository

import (
	"asset_management_backend/common/service"
	"asset_management_backend/module/label_task/domain/entity"
	"context"
)

type LabelTaskRepository interface {
	// Tạo presigned url cho client upload file
	CreateLabelTasksPresignedUrls(ctx context.Context, fileNames []string) ([]service.PresignedResponse, error)

	CreateTasks(
		ctx context.Context,
		requests []entity.LabelTaskRequest,
		creatorID int,
	) (<-chan *entity.LabelTaskStreamResponse, error)

	CreateTask(
		ctx context.Context,
		request entity.LabelTaskRequest,
		creatorID int,
	) (<-chan *entity.LabelTaskStreamResponse, error)
}
