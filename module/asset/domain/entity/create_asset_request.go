package entity

import "time"

type CreateAssetRequest struct {
	CreatorID          *int       `json:"creator_id,omitempty"`
	SerialNumber       *string    `json:"serial_number" validate:"omitempty"`
	AssetLabelPath     *string    `json:"asset_label_path" validate:"omitempty"`
	ModelNumber        *string    `json:"model_number" validate:"omitempty"`
	Manufacturer       *string    `json:"manufacturer" validate:"omitempty"`
	MadeIn             *string    `json:"made_in" validate:"omitempty"`
	Supplier           *string    `json:"supplier" validate:"omitempty"`
	PurchasePrice      *float32   `json:"purchase_price" validate:"omitempty"`
	PriceUnitId        *int       `json:"price_unit_id" validate:"omitempty"`
	AssetTypeId        int        `json:"asset_type_id" validate:"required"`
	AssetQualityId     int        `json:"asset_quality_id" validate:"required"`
	LocationId         int        `json:"location_id" validate:"required"`
	PurchaseDate       *time.Time `json:"purchase_date" validate:"omitempty"`
	WarrantyExpiryDate *time.Time `json:"warranty_expiry_date" validate:"omitempty"`
	Description        *string    `json:"description" validate:"omitempty"`
	AssetFilePaths     []string   `json:"asset_file_paths" validate:"omitempty"`
	LabelTaskID        *int       `json:"label_task_id" validate:"omitempty"`
}
