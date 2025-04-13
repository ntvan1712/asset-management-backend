package repository

import (
	"asset_management_backend/common/error_app"
	"asset_management_backend/common/service"
	"asset_management_backend/infras"
	datasource "asset_management_backend/module/label_task/data/data_source"
	"asset_management_backend/module/label_task/data/model"
	"asset_management_backend/module/label_task/domain/entity"
	"context"
	"encoding/json"
	"time"
)

type labelTaskRepositoryImpl struct {
	labelTaskDS         datasource.LabelTaskDataSource
	fileStorageService  *service.FileStorageService
	messageQueueService *service.MessageQueueService
}

// CreateTask implements LabelTaskRepository.
func (a *labelTaskRepositoryImpl) CreateTask(
	ctx context.Context,
	request entity.LabelTaskRequest,
	creatorID int,
) (<-chan *entity.LabelTaskStreamResponse, error) {
	newTaskModel, err := a.labelTaskDS.Insert(ctx, model.NewAssetLabelTaskFromRequest(request, creatorID))
	if err != nil {
		return nil, err
	}

	bodyData, parseErr := json.Marshal(newTaskModel.ToTaskPublish())
	if parseErr != nil {
		return nil, err
	}

	completedTaskIdsCh, err := datasource.GetTaskStream()
	if err != nil {
		return nil, err
	}
	resultCh := make(chan *entity.LabelTaskStreamResponse, 1)

	if err := a.messageQueueService.SafePublish(ctx, bodyData, a.messageQueueService.LabelTaskQueueName); err != nil {
		close(resultCh)
		return nil, err
	}

	go func() {
		defer close(resultCh)
		defer datasource.Unsubscribe()

		for id := range completedTaskIdsCh {
			if id != nil && *id == newTaskModel.ID {
				labelTaskResult, err := a.labelTaskDS.FindByID(ctx, newTaskModel.ID)
				if err != nil {
					resultCh <- &entity.LabelTaskStreamResponse{Error: err}
				} else {
					resultCh <- &entity.LabelTaskStreamResponse{LabelTask: labelTaskResult.ToEntity(), Error: nil}
				}
				return
			}
		}
		resultCh <- &entity.LabelTaskStreamResponse{Error: error_app.ErrUnknown}

	}()

	return resultCh, nil
}

// CreateTasks implements LabelTaskRepository.
func (a *labelTaskRepositoryImpl) CreateTasks(ctx context.Context, requests []entity.LabelTaskRequest) {
	panic("unimplemented")
}

// CreateLabelTasksPresignedUrls implements AssetRepository.
func (a *labelTaskRepositoryImpl) CreateLabelTasksPresignedUrls(
	ctx context.Context,
	fileNames []string,
) ([]service.PresignedResponse, error) {
	return a.fileStorageService.CreatePresignedUrls(ctx, fileNames, "labels", time.Minute*5)
}

func NewLabelTaskRepository() LabelTaskRepository {
	return &labelTaskRepositoryImpl{
		labelTaskDS:         datasource.NewLabelTaskDataSource(infras.GetDbProvider()),
		fileStorageService:  service.NewFileStorageService(),
		messageQueueService: service.NewMessageQueueService(),
	}
}
