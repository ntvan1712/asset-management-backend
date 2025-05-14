package entity

import "time"

type CreateTransferRequestEntity struct {
	ReturnDate *time.Time `json:"return_date,omitempty"`
	Reason     *string    `json:"reason,omitempty"`

	RequestorID     *int `json:"requestor_id"`
	UseAtLocationID int  `json:"use_at_location_id" validate:"required"`
	AssetID         int  `json:"asset_id" validate:"required"`
	RespondentID    int  `json:"respondent_id" validate:"required"`
}
