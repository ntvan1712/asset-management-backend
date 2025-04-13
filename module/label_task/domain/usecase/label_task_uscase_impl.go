package usecase

import (
	"asset_management_backend/common/error_app"
	"asset_management_backend/common/service"
	"asset_management_backend/module/label_task/data/repository"
	"asset_management_backend/module/label_task/domain/entity"
	"context"
)

type labelTaskRepositoryImpl struct {
	labelTaskRepo repository.LabelTaskRepository
}

// CreateTask implements LabelTaskUsecase.
func (a *labelTaskRepositoryImpl) CreateTask(ctx context.Context, request entity.LabelTaskRequest, creatorID int) (<-chan *entity.LabelTaskStreamResponse, error){
	return a.labelTaskRepo.CreateTask(ctx, request, creatorID)
}

// GetLabelTaskPresignedUrl implements LabelTaskUsecase.
func (a *labelTaskRepositoryImpl) GetLabelTaskPresignedUrl(ctx context.Context, fileName string) (*service.PresignedResponse, error) {
	result, err := a.labelTaskRepo.CreateLabelTasksPresignedUrls(ctx, []string{fileName})
	if err != nil {
		return nil, err
	}
	if len(result) != 1 {
		return nil, error_app.ErrUnknown
	}
	return &result[0], nil
}

// GetLabelTasksPresignedUrls implements AssetRepository.
func (a *labelTaskRepositoryImpl) GetLabelTasksPresignedUrls(
	ctx context.Context,
	fileNames []string,
) ([]service.PresignedResponse, error) {
	return a.labelTaskRepo.CreateLabelTasksPresignedUrls(ctx, fileNames)
}

func NewLabelTaskUsecase() LabelTaskUsecase {
	return &labelTaskRepositoryImpl{
		labelTaskRepo: repository.NewLabelTaskRepository(),
	}
}
