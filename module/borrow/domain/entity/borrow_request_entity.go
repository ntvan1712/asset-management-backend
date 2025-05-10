package entity

import (
	sharedmodel "asset_management_backend/common/shared_model"
	categoryEntity "asset_management_backend/module/category/domain/entity"
	"time"
)

type BorrowRequestEntity struct {
	ID                  int        `bun:",pk,autoincrement" json:"id"`
	RequestAt           *time.Time `json:"request_at,omitempty"`
	ReturnDate          *time.Time `json:"return_date,omitempty"`
	Reason              *string    `json:"reason,omitempty"`
	Status              string     `json:"status"`
	ResponseAt          *time.Time `json:"response_at,omitempty"`
	ResponseDescription *string    `json:"response_description,omitempty"`

	RequestorID     int  `json:"requestor_id"`
	UseAtLocationID int  `json:"use_at_location_id"`
	AssetTypeID     int  `json:"asset_type_id"`
	RespondentID    *int `json:"respondent_id,omitempty"`

	Requestor     *sharedmodel.UserBasicInfo      `json:"requestor,omitempty"`
	UseAtLocation *categoryEntity.LocationEntity  `json:"use_at_location,omitempty"`
	AssetType     *categoryEntity.AssetTypeEntity `json:"asset_type,omitempty"`
	Respondent    *sharedmodel.UserBasicInfo      `json:"respondent,omitempty"`
}
