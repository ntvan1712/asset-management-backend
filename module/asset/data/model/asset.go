package model

import (
	categoryModel "asset_management_backend/module/category/data/model"
	"time"

	"github.com/uptrace/bun"
)

type Asset struct {
	bun.BaseModel `bun:"table:assets"`

	ID                 int `bun:",pk,autoincrement" json:"id"`
	SerialNumber       string
	ModelNumber        *string
	Manufacturer       *string
	MadeIn             *string
	PurchasePrice      *float32
	PurchaseDate       *time.Time
	WarrantyExpiryDate *time.Time
	Supplier           *string
	Description        *string
	AddedAt            *time.Time
	Status             string

	PriceUnitID  *int
	LocationID   int
	QualityID    int
	AssetTypeID  int
	LabelImageID int

	PriceUnit  *categoryModel.Currency
	Location   categoryModel.Location
	Quality    categoryModel.AssetQuality
	AssetType  categoryModel.AssetType
	LabelImage AssetLabelImage

	AssetImages []AssetImage `bun:"rel:has-many,join:id=asset_id"`
}
