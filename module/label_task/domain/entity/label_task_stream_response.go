package entity

type LabelTaskStreamResponse struct {
	LabelTask *AssetLabelTaskEntity   `json:"label_task,omitempty"`
	Error     error `json:"error,omitempty"`
}
