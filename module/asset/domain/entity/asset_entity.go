package entity

import (
	borrowedAssetE "asset_management_backend/module/borrowed_asset/domain/entity"
	categoryEntity "asset_management_backend/module/category/domain/entity"
	"time"
)

type AssetEntity struct {
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

	PriceUnit    *categoryEntity.CurrencyEntity     `json:"price_unit,omitempty"`
	Location     *categoryEntity.LocationEntity     `json:"location,omitempty"`
	AssetQuality *categoryEntity.AssetQualityEntity `json:"asset_quality,omitempty"`
	AssetType    *categoryEntity.AssetTypeEntity    `json:"asset_type,omitempty"`

	AssetFiles      []AssetFileEntity           `json:"asset_files,omitempty"`
	AssetLabelImage *AssetLabelImageEntity      `json:"asset_label_image,omitempty"`
	BorrowedAsset   *borrowedAssetE.BorrowedAssetEntity `json:"borrowed_asset,omitempty"`
}
