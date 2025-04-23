package model

import (
	"asset_management_backend/common/enums"
	"asset_management_backend/module/asset/domain/entity"
	categoryModel "asset_management_backend/module/category/data/model"
	"context"
	"strings"
	"time"

	"github.com/uptrace/bun"
)

const TableAsset = "assets"

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

	PriceUnitID    *int `json:"price_unit_id,omitempty"`
	LocationID     int  `json:"location_id"`
	AssetQualityID int  `json:"asset_quality_id"`
	AssetTypeID    int  `json:"asset_type_id"`

	PriceUnit    *categoryModel.Currency    `bun:"rel:belongs-to,join:price_unit_id=id"`
	Location     categoryModel.Location     `bun:"rel:belongs-to,join:location_id=id"`
	AssetQuality categoryModel.AssetQuality `bun:"rel:belongs-to,join:asset_quality_id=id"`
	AssetType    categoryModel.AssetType    `bun:"rel:belongs-to,join:asset_type_id=id"`

	AssetFiles      []AssetFile      `bun:"rel:has-many,join:id=asset_id"`
	AssetLabelImage *AssetLabelImage `bun:"rel:has-one,join:id=asset_id"`
}

func LoadAllAssetRelationQuery(selectModelQuery *bun.SelectQuery) {
	selectModelQuery.Relation("PriceUnit").
		Relation("Location").
		Relation("AssetQuality").
		Relation("AssetType").
		Relation("AssetFiles").
		Relation("AssetLabelImage")
}

func NewAssetModelFromRequest(request *entity.CreateAssetRequest) Asset {
	// Đảm bảo serial number luôn là
	upperSerialNumber := strings.ToUpper(*request.SerialNumber)
	request.SerialNumber = &upperSerialNumber
	addedAt := time.Now().UTC()
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
		PriceUnitID:        request.PriceUnitId,
		LocationID:         request.LocationId,
		AssetQualityID:     request.AssetQualityId,
		AssetTypeID:        request.AssetTypeId,
		AddedAt: &addedAt,
	}
}

func (a *Asset) ToEntity() *entity.AssetEntity {
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
		PriceUnit:          a.PriceUnit.ToEntity(),
		Location:           a.Location.ToEntity(),
		AssetQuality:       a.AssetQuality.ToEntity(),
		AssetType:          a.AssetType.ToEntity(),
		AssetFiles:         assetFiles,
		AssetLabelImage:    a.AssetLabelImage.ToEntity(),
	}
}

func AssetModelsToEntities(assetModels []Asset) []entity.AssetEntity {
	var entities []entity.AssetEntity

	for _, asset := range assetModels {
		entities = append(entities, *asset.ToEntity())
	}

	return entities
}


var _ bun.AfterDeleteHook = (*Asset)(nil)

func (*Asset) AfterDelete(ctx context.Context, query *bun.DeleteQuery) error { return nil }
