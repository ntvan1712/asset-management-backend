package model

import (
	"asset_management_backend/module/category/domain/entity"

	"github.com/uptrace/bun"
)

const TableAssetQuality = "asset_qualities"

type AssetQuality struct {
	bun.BaseModel `bun:"table:asset_qualities"`

	ID int `bun:",pk,autoincrement" json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (a *AssetQuality) ToEntity() *entity.AssetQualityEntity {
	if a == nil {
		return nil
	}
	return &entity.AssetQualityEntity{
		ID:          a.ID,
		Code:        a.Code,
		Name:        a.Name,
		Description: a.Description,
	}
}

func AssetQualityModelsToEntities(assetQualities []AssetQuality) []entity.AssetQualityEntity {
	var assetQualityEntities []entity.AssetQualityEntity

	for _, assetQuality := range assetQualities {
		assetQualityEntities = append(assetQualityEntities, *assetQuality.ToEntity())
	}

	return assetQualityEntities
}
