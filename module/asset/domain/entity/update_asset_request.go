package entity

import "time"

type UpdateAssetRequest struct {
	UpdateData           *UpdateAssetInfo `json:"update_data" validate:"omitempty"`
	ShouldCreateNewLabel *bool            `json:"should_create_new_label,omitempty"`
	NewAssetFilePaths    []string         `json:"new_asset_file_paths" validate:"omitempty"`
	RemoveAssetFileIDs   []int            `json:"remove_asset_file_ids,omitempty"`
}

type UpdateAssetInfo struct {
	ModelNumber        *string    `json:"model_number" validate:"omitempty"`
	Manufacturer       *string    `json:"manufacturer" validate:"omitempty"`
	MadeIn             *string    `json:"made_in" validate:"omitempty"`
	Supplier           *string    `json:"supplier" validate:"omitempty"`
	PurchasePrice      *float32   `json:"purchase_price" validate:"omitempty"`
	PriceUnitId        *int       `json:"price_unit_id" validate:"omitempty"`
	AssetTypeId        *int       `json:"asset_type_id" validate:"omitempty"`
	AssetQualityId     *int       `json:"asset_quality_id" validate:"omitempty"`
	LocationId         *int       `json:"location_id" validate:"omitempty"`
	PurchaseDate       *time.Time `json:"purchase_date" validate:"omitempty"`
	WarrantyExpiryDate *time.Time `json:"warranty_expiry_date" validate:"omitempty"`
	Description        *string    `json:"description" validate:"omitempty"`
}
