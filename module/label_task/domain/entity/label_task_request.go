package entity

type LabelTaskRequest struct {
	LabelImagePath string `json:"label_image_path" validate:"required"`
	TaskType       string `json:"task_type" validate:"required,oneof=add_asset search_asset"`
}
