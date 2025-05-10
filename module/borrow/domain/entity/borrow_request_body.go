package entity

import "time"

type BorrowRequestBody struct {
	RequestAt  *time.Time `json:"request_at,omitempty"`
	ReturnDate *time.Time `json:"return_date"`
	Reason     *string    `json:"reason,omitempty"`

	RequestorID     *int `json:"requestor_id"`
	UseAtLocationID int  `json:"use_at_location_id" validate:"required"`
	AssetTypeID     int  `json:"asset_type_id" validate:"required"`
}
