package model

import (
	"asset_management_backend/common/enums"
	sharedmodel "asset_management_backend/common/shared_model"
	app_utils "asset_management_backend/common/utils"
	assetModel "asset_management_backend/module/asset/data/model"
	categoryModel "asset_management_backend/module/category/data/model"
	"asset_management_backend/module/transfer/domain/entity"
	"time"

	"github.com/uptrace/bun"
)

const TableTransferRequest = "transfer_requests"

type TransferRequest struct {
	bun.BaseModel `bun:"table:transfer_requests"`

	ID                  int        `bun:",pk,autoincrement" json:"id"`
	RequestAt           *time.Time `json:"request_at,omitempty"`
	ReturnDate          *time.Time `json:"return_date,omitempty"`
	Reason              *string    `json:"reason,omitempty"`
	Status              string     `json:"status"`
	ResponseAt          *time.Time `json:"response_at,omitempty"`
	ResponseDescription *string    `json:"response_description,omitempty"`

	RequestorID         int  `json:"requestor_id"`
	UseAtLocationID     int  `json:"use_at_location_id"`
	AssetID             int  `json:"asset_id"`
	RespondentID        *int `json:"respondent_id,omitempty"`
	FromBorrowRequestID *int `json:"from_borrow_request_id,omitempty"`

	Requestor     *sharedmodel.UserBasicInfo `bun:"rel:belongs-to,join:requestor_id=id"`
	UseAtLocation *categoryModel.Location    `bun:"rel:belongs-to,join:use_at_location_id=id"`
	Asset         *assetModel.Asset          `bun:"rel:belongs-to,join:asset_id=id"`
	Respondent    *sharedmodel.UserBasicInfo `bun:"rel:belongs-to,join:respondent_id=id"`
}

func (transfer *TransferRequest) ToEntity() *entity.TransferRequestEntity {
	if transfer == nil {
		return nil
	}
	return &entity.TransferRequestEntity{
		ID:                  transfer.ID,
		RequestAt:           transfer.RequestAt,
		ReturnDate:          transfer.ReturnDate,
		Reason:              transfer.Reason,
		Status:              transfer.Status,
		ResponseAt:          transfer.ResponseAt,
		ResponseDescription: transfer.ResponseDescription,
		RequestorID:         transfer.RequestorID,
		UseAtLocationID:     transfer.UseAtLocationID,
		AssetID:             transfer.AssetID,
		RespondentID:        transfer.RespondentID,
		Requestor:           transfer.Requestor,
		UseAtLocation:       transfer.UseAtLocation.ToEntity(),
		Asset:               transfer.Asset.ToEntity(),
		Respondent:          transfer.Respondent,
		FromBorrowRequestID: transfer.FromBorrowRequestID,
	}
}

func NewTransferRequestFromBody(body entity.CreateTransferRequestEntity) TransferRequest {
	return TransferRequest{
		RequestAt:           app_utils.TimeNowPtr(),
		ReturnDate:          body.ReturnDate,
		Reason:              body.Reason,
		Status:              enums.TransferRequestEnum.Pending,
		ResponseAt:          nil,
		ResponseDescription: nil,
		RequestorID:         *body.RequestorID,
		UseAtLocationID:     body.UseAtLocationID,
		AssetID:             body.AssetID,
		RespondentID:        &body.RespondentID,
		FromBorrowRequestID: nil,
	}
}

func TransferRequestModelsToEntities(models []TransferRequest) []entity.TransferRequestEntity {
	var entities []entity.TransferRequestEntity

	for _, asset := range models {
		entities = append(entities, *asset.ToEntity())
	}

	return entities
}
