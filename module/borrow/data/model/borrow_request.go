package model

import (
	"asset_management_backend/common/enums"
	sharedmodel "asset_management_backend/common/shared_model"
	"asset_management_backend/module/borrow/domain/entity"
	categoryModel "asset_management_backend/module/category/data/model"
	"time"

	"github.com/uptrace/bun"
)

const TableBorrowRequest = "borrow_requests"

type BorrowRequest struct {
	bun.BaseModel `bun:"table:borrow_requests"`

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

	Requestor     *sharedmodel.UserBasicInfo `bun:"rel:belongs-to,join:requestor_id=id"`
	UseAtLocation *categoryModel.Location    `bun:"rel:belongs-to,join:use_at_location_id=id"`
	AssetType     *categoryModel.AssetType   `bun:"rel:belongs-to,join:asset_type_id=id"`
	Respondent    *sharedmodel.UserBasicInfo `bun:"rel:belongs-to,join:respondent_id=id"`
}

func (b *BorrowRequest) ToEntity() *entity.BorrowRequestEntity {
	if b == nil {
		return nil
	}
	return &entity.BorrowRequestEntity{
		ID:                  b.ID,
		RequestAt:           b.RequestAt,
		ReturnDate:          b.ReturnDate,
		Reason:              b.Reason,
		Status:              b.Status,
		ResponseAt:          b.ResponseAt,
		ResponseDescription: b.ResponseDescription,
		RequestorID:         b.RequestorID,
		UseAtLocationID:     b.UseAtLocationID,
		AssetTypeID:         b.AssetTypeID,
		RespondentID:        b.RespondentID,
		Requestor:           b.Requestor,
		UseAtLocation:       b.UseAtLocation.ToEntity(),
		AssetType:           b.AssetType.ToEntity(),
		Respondent:          b.Respondent,
	}
}

func BorrowRequestModelsToEntities(models []BorrowRequest) []entity.BorrowRequestEntity {
	var entities []entity.BorrowRequestEntity

	for _, asset := range models {
		entities = append(entities, *asset.ToEntity())
	}

	return entities
}

func BorrowRequestModelFromBody(requestBody *entity.BorrowRequestBody) BorrowRequest {
	var requestAt time.Time
	if requestBody.RequestAt == nil {
		requestAt = time.Now().UTC()
	} else {
		requestAt = *requestBody.RequestAt
	}
	return BorrowRequest{
		RequestAt:       &requestAt,
		ReturnDate:      requestBody.ReturnDate,
		Reason:          requestBody.Reason,
		Status:          enums.BorrowRequestEnum.Pending,
		RequestorID:     *requestBody.RequestorID,
		UseAtLocationID: requestBody.UseAtLocationID,
		AssetTypeID:     requestBody.AssetTypeID,
	}
}
