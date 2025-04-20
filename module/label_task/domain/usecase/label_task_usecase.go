package usecase

import (
	"asset_management_backend/common/service"
	"asset_management_backend/module/label_task/domain/entity"
	"context"
)

type LabelTaskUsecase interface {
	// Tạo presigned url cho client upload file
	GetLabelTasksPresignedUrls(ctx context.Context, fileNames []string) ([]service.PresignedResponse, error)
	GetLabelTaskPresignedUrl(ctx context.Context, fileName string) (*service.PresignedResponse, error)

	CreateTask(
		ctx context.Context,
		request entity.LabelTaskRequest,
		creatorID int,
	) (<-chan *entity.LabelTaskStreamResponse, error)

	CreateTasks(
		ctx context.Context,
		requests []entity.LabelTaskRequest,
		creatorID int,
	) (<-chan *entity.LabelTaskStreamResponse, error)
}
