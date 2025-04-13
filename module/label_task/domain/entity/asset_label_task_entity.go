package entity

import "time"

type AssetLabelTaskEntity struct {
	ID           int        `json:"id"`
	Url          string     `json:"url"`
	SerialNumber *string    `json:"serial_number,omitempty"`
	ModelNumber  *string    `json:"model_number,omitempty"`
	Manufacturer *string    `json:"manufacturer,omitempty"`
	MadeIn       *string    `json:"made_in,omitempty"`
	AllWords     *[]string  `json:"all_words,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	CompletedAt  *time.Time `json:"completed_at,omitempty"`
	ErrorCode    *string    `json:"error_code,omitempty"`
	TaskType     string     `json:"task_type,omitempty"`
}
