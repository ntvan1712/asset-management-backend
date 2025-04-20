package model

import (
	"asset_management_backend/common/enums"
	app_utils "asset_management_backend/common/utils"
	"asset_management_backend/module/asset/domain/entity"
	"fmt"
	"strings"
	"time"
)

type AssetFilterModel struct {
	SerialNumber   *string
	Status         *string
	LocationID     *int
	AssetTypeID    *int
	AssetQualityID *int

	FromAddedAt *time.Time
	ToAddedAt   *time.Time

	Offset    int
	Limit     int
	OrderExpr string
}

func NewAssetFilterModelFromQuery(query entity.AssetFilterQuery) AssetFilterModel {

	sortBy := query.SortBy
	if sortBy == "" {
		sortBy = enums.AssetSortByEnum.AddedAt
	}
	sortOrder := strings.ToLower(query.SortOrder)
	if sortOrder != enums.SortOrderEnum.Asc {
		sortOrder = enums.SortOrderEnum.Desc
	}
	var serialNumber *string
	if query.SerialNumber != nil && *query.SerialNumber != "" {
		serialNumberValue := strings.ToUpper(*query.SerialNumber)
		serialNumber = &serialNumberValue
	}

	// Sắp xếp null xuống cuối
	orderExpr := fmt.Sprintf("%s IS NULL, %s %s", sortBy, sortBy, sortOrder)
	offset := (query.Page - 1) * query.Limit

	return AssetFilterModel{
		SerialNumber:   serialNumber,
		Status:         app_utils.FormatStringPtr(query.Status),
		LocationID:     app_utils.FormatIntPtr(query.LocationID),
		AssetTypeID:    app_utils.FormatIntPtr(query.AssetTypeID),
		AssetQualityID: app_utils.FormatIntPtr(query.AssetQualityID),
		FromAddedAt:    app_utils.FormatInt64ToUnixTimePtr(query.FromAddedAt),
		ToAddedAt:      app_utils.FormatInt64ToUnixTimePtr(query.ToAddedAt),
		Offset:         offset,
		OrderExpr:      orderExpr,
		Limit:          query.Limit,
	}
}
