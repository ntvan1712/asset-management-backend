package model

import (
	"asset_management_backend/app_config"
	"asset_management_backend/common/enums"
	"asset_management_backend/module/asset/domain/entity"
	categoryModel "asset_management_backend/module/category/data/model"
	"time"

	"github.com/uptrace/bun"
)

type Asset struct {
	bun.BaseModel `bun:"table:assets"`

	ID                 int        `bun:",pk,autoincrement" json:"id"`
	SerialNumber       string     `bun:",notnull,type:varchar(100)" json:"serial_number"`
	ModelNumber        *string    `json:"model_number,omitempty"`
	Manufacturer       *string    `json:"manufacturer,omitempty"`
	MadeIn             *string    `json:"made_in,omitempty"`
	PurchasePrice      *float32   `json:"purchase_price,omitempty"`
	PurchaseDate       *time.Time `json:"purchase_date,omitempty"`
	WarrantyExpiryDate *time.Time `json:"warranty_expiry_date,omitempty"`
	Supplier           *string    `json:"supplier,omitempty"`
	Description        *string    `json:"description,omitempty"`
	AddedAt            *time.Time `json:"added_at,omitempty"`
	Status             string     `json:"status"`
	LabelImagePath     *string    `json:"label_image_path,omitempty"`

	PriceUnitID    *int `json:"price_unit_id,omitempty"`
	LocationID     int  `json:"location_id"`
	AssetQualityID int  `json:"asset_quality_id"`
	AssetTypeID    int  `json:"asset_type_id"`

	PriceUnit    *categoryModel.Currency    `bun:"rel:belongs-to,join:price_unit_id=id"`
	Location     categoryModel.Location     `bun:"rel:belongs-to,join:location_id=id"`
	AssetQuality categoryModel.AssetQuality `bun:"rel:belongs-to,join:asset_quality_id=id"`
	AssetType    categoryModel.AssetType    `bun:"rel:belongs-to,join:asset_type_id=id"`

	AssetFiles []AssetFile `bun:"rel:has-many,join:id=asset_id"`
}

func NewAssetModelFromRequest(request *entity.CreateAssetRequest) Asset {
	return Asset{
		SerialNumber:       *request.SerialNumber,
		ModelNumber:        request.ModelNumber,
		Manufacturer:       request.Manufacturer,
		MadeIn:             request.MadeIn,
		PurchasePrice:      request.PurchasePrice,
		PurchaseDate:       request.PurchaseDate,
		WarrantyExpiryDate: request.WarrantyExpiryDate,
		Supplier:           request.Supplier,
		Description:        request.Description,
		Status:             enums.AssetStatusEnum.Available,
		LabelImagePath:     request.AssetLabelPath,
		PriceUnitID:        request.PriceUnitId,
		LocationID:         request.LocationId,
		AssetQualityID:     request.AssetQualityId,
		AssetTypeID:        request.AssetTypeId,
	}
}

func (a *Asset) ToEntity() *entity.AssetEntity {
	var labelImageUrl *string
	if a.LabelImagePath != nil {
		url := app_config.GetAppConfig().MinioConfig.GetFullAssetUrl(*a.LabelImagePath)
		labelImageUrl = &url
	}
	var assetFiles []entity.AssetFileEntity
	if a.AssetFiles != nil {
		assetFiles = AssetFileModelsToEntities(a.AssetFiles)
	}
	return &entity.AssetEntity{
		ID:                 a.ID,
		SerialNumber:       a.SerialNumber,
		ModelNumber:        a.ModelNumber,
		Manufacturer:       a.Manufacturer,
		MadeIn:             a.MadeIn,
		PurchasePrice:      a.PurchasePrice,
		PurchaseDate:       a.PurchaseDate,
		WarrantyExpiryDate: a.WarrantyExpiryDate,
		Supplier:           a.Supplier,
		Description:        a.Description,
		AddedAt:            a.AddedAt,
		Status:             a.Status,
		LabelImageUrl:      labelImageUrl,
		PriceUnit:          a.PriceUnit.ToEntity(),
		Location:           a.Location.ToEntity(),
		AssetQuality:       a.AssetQuality.ToEntity(),
		AssetType:          a.AssetType.ToEntity(),
		AssetFiles:         assetFiles,
	}
}
