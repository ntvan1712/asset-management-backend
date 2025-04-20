package model

import (
	"asset_management_backend/app_config"
	"asset_management_backend/module/label_task/domain/entity"
	"time"

	"github.com/uptrace/bun"
)

type AssetLabelTask struct {
	bun.BaseModel `bun:"table:asset_label_tasks" json:"-"`

	ID                int        `bun:"id,pk,autoincrement" json:"id"`
	Path              string     `bun:"path,notnull" json:"path"`
	SerialNumber      *string    `bun:"serial_number" json:"serial_number,omitempty"`
	ModelNumber       *string    `bun:"model_number" json:"model_number,omitempty"`
	Manufacturer      *string    `bun:"manufacturer" json:"manufacturer,omitempty"`
	MadeIn            *string    `bun:"made_in" json:"made_in,omitempty"`
	AllWords          *[]string  `bun:",array" json:"all_words,omitempty"`
	CreatedAt         time.Time  `bun:"created_at,default:current_timestamp" json:"created_at"`
	CompletedAt       *time.Time `bun:"completed_at" json:"completed_at,omitempty"`
	ErrorCode         *string    `bun:"error_code" json:"error_code,omitempty"`
	TaskType          string     `bun:"task_type" json:"task_type,omitempty"`
	AssetLabelImageID *int       `bun:"asset_label_image_id" json:"asset_label_image_id,omitempty"`
	CreatorID         *int       `bun:"creator_id" json:"creator_id,omitempty"`
}

func (t *AssetLabelTask) ToEntity() *entity.AssetLabelTaskEntity {
	return &entity.AssetLabelTaskEntity{
		ID:           t.ID,
		Url:          app_config.GetAppConfig().MinioConfig.GetFullAssetUrl(t.Path),
		SerialNumber: t.SerialNumber,
		ModelNumber:  t.ModelNumber,
		Manufacturer: t.Manufacturer,
		MadeIn:       t.MadeIn,
		AllWords:     t.AllWords,
		CreatedAt:    t.CreatedAt,
		CompletedAt:  t.CompletedAt,
		ErrorCode:    t.ErrorCode,
		TaskType:     t.TaskType,
	}
}

func NewAssetLabelTaskFromRequest(request entity.LabelTaskRequest, creatorID int) AssetLabelTask {
	return AssetLabelTask{
		Path:              request.LabelImagePath,
		CreatedAt:         time.Now(),
		TaskType:          request.TaskType,
		AssetLabelImageID: nil,
		CreatorID:         &creatorID,
	}
}

func NewAssetLabelTasksFromRequests(requests []entity.LabelTaskRequest, creatorID int) []AssetLabelTask {
	var labelTasks []AssetLabelTask

	for _, request := range requests {
		newLabelTask := NewAssetLabelTaskFromRequest(request, creatorID)
		labelTasks = append(labelTasks, newLabelTask)
	}

	return labelTasks
}

func (t *AssetLabelTask) ToTaskPublish() LabelTaskPublishModel {
	return LabelTaskPublishModel{
		ID:   t.ID,
		Path: t.Path,
	}
}
