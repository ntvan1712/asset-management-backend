package entity

import "time"

type UpdateAssetRequest struct {
	CreatorID            *int             `json:"creator_id,omitempty"`
	UpdateData           *UpdateAssetInfo `json:"update_data,omitempty" validate:"omitempty"`
	ShouldCreateNewLabel *bool            `json:"should_create_new_label,omitempty"`
	NewAssetFilePaths    []string         `json:"new_asset_file_paths,omitempty" validate:"omitempty"`
	RemoveAssetFileIDs   []int            `json:"remove_asset_file_ids,omitempty"`
}

func (r *UpdateAssetRequest) IsEmpty() bool {
	return (r.UpdateData == nil || r.UpdateData.IsEmpty()) &&
		(r.ShouldCreateNewLabel == nil || !*r.ShouldCreateNewLabel) &&
		len(r.NewAssetFilePaths) == 0 &&
		len(r.RemoveAssetFileIDs) == 0
}

type UpdateAssetInfo struct {
	ModelNumber        *string    `json:"model_number,omitempty" validate:"omitempty"`
	Manufacturer       *string    `json:"manufacturer,omitempty" validate:"omitempty"`
	MadeIn             *string    `json:"made_in,omitempty" validate:"omitempty"`
	Supplier           *string    `json:"supplier,omitempty" validate:"omitempty"`
	PurchasePrice      *float32   `json:"purchase_price,omitempty" validate:"omitempty"`
	Status             *string    `json:"status,omitempty" validate:"omitempty,oneof=available on_borrow awaiting_transfer unusable"`
	PriceUnitId        *int       `json:"price_unit_id,omitempty" validate:"omitempty"`
	AssetTypeId        *int       `json:"asset_type_id,omitempty" validate:"omitempty"`
	AssetQualityId     *int       `json:"asset_quality_id,omitempty" validate:"omitempty"`
	LocationId         *int       `json:"location_id,omitempty" validate:"omitempty"`
	PurchaseDate       *time.Time `json:"purchase_date,omitempty" validate:"omitempty"`
	WarrantyExpiryDate *time.Time `json:"warranty_expiry_date,omitempty" validate:"omitempty"`
	Description        *string    `json:"description,omitempty" validate:"omitempty"`
}

func (u *UpdateAssetInfo) IsEmpty() bool {
	return u.ModelNumber == nil &&
		u.Manufacturer == nil &&
		u.Status == nil &&
		u.MadeIn == nil &&
		u.Supplier == nil &&
		u.PurchasePrice == nil &&
		u.PriceUnitId == nil &&
		u.AssetTypeId == nil &&
		u.AssetQualityId == nil &&
		u.LocationId == nil &&
		u.PurchaseDate == nil &&
		u.WarrantyExpiryDate == nil &&
		u.Description == nil
}
