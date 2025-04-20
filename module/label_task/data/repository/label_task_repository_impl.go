package repository

import (
	"asset_management_backend/common/logger"
	"asset_management_backend/common/service"
	"asset_management_backend/infras"
	datasource "asset_management_backend/module/label_task/data/data_source"
	"asset_management_backend/module/label_task/data/model"
	"asset_management_backend/module/label_task/domain/entity"
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
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

	streamID := int(uuid.New().ID())
	completedTaskIdsCh, err := datasource.GetTaskStream(streamID)
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
		defer datasource.Unsubscribe(streamID)

		for {
			select {
			case <-ctx.Done():
				// Context bị huỷ (timeout hoặc cancel) thì thoát goroutine
				return
			case id, ok := <-completedTaskIdsCh:
				if !ok {
					return
				}
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
		}
	}()

	return resultCh, nil
}

// CreateTasks implements LabelTaskRepository.
func (a *labelTaskRepositoryImpl) CreateTasks(
	ctx context.Context,
	requests []entity.LabelTaskRequest,
	creatorID int,
) (<-chan *entity.LabelTaskStreamResponse, error) {
	newTaskModels, err := a.labelTaskDS.InsertMany(ctx, model.NewAssetLabelTasksFromRequests(requests, creatorID))
	if err != nil {
		return nil, err
	}

	streamID := int(uuid.New().ID())
	completedTaskIdsCh, err := datasource.GetTaskStream(streamID)
	if err != nil {
		return nil, err
	}
	resultCh := make(chan *entity.LabelTaskStreamResponse)
	expectedTaskIDs := make(map[int]struct{})

	for _, task := range newTaskModels {
		bodyData, err := json.Marshal(task.ToTaskPublish())
		if err != nil {
			logger.Error("[LabelTaskRepository.CreateTasks]", "Marshal error", err)
			continue
		}

		if err := a.messageQueueService.SafePublish(ctx, bodyData, a.messageQueueService.LabelTaskQueueName); err != nil {
			logger.Error("[LabelTaskRepository.CreateTasks]", "Publish task error", err)
			continue
		}
		expectedTaskIDs[task.ID] = struct{}{}
	}

	go func() {
		defer close(resultCh)
		defer datasource.Unsubscribe(streamID)
		completedCount := 0
		publishedTasksCount := len(expectedTaskIDs)
		for {
			select {
			case <-ctx.Done():
				return
			case idPtr, ok := <-completedTaskIdsCh:
				if !ok {
					continue
				}
				if idPtr == nil {
					continue
				}
				id := *idPtr
				if _, exists := expectedTaskIDs[id]; exists {
					labelTaskResult, err := a.labelTaskDS.FindByID(ctx, id)
					if err != nil {
						resultCh <- &entity.LabelTaskStreamResponse{Error: err}
					} else {
						resultCh <- &entity.LabelTaskStreamResponse{LabelTask: labelTaskResult.ToEntity()}
					}
					delete(expectedTaskIDs, id)
					completedCount++
					if completedCount == publishedTasksCount {
						return
					}
				}
			}
		}
	}()

	return resultCh, nil
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
