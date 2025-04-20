package entity

import (
	sharedmodel "asset_management_backend/common/shared_model"
)

type AssetFilterQuery struct {
	SerialNumber   *string `query:"serial_number"`
	Status         *string `query:"status"`
	LocationID     *int    `query:"location_id"`
	AssetTypeID    *int    `query:"asset_type_id"`
	AssetQualityID *int    `query:"asset_quality_id"`

	FromAddedAt *int64 `query:"from_added_at,omitempty"`
	ToAddedAt   *int64 `query:"to_added_at,omitempty"`

	SortBy    string `query:"sort_by" validate:"omitempty,oneof=added_at purchase_date"`
	SortOrder string `query:"sort_order" validate:"omitempty,oneof=asc desc"`

	sharedmodel.PaginateQuery
}
