package model

import (
	"asset_management_backend/module/category/domain/entity"

	"github.com/uptrace/bun"
)

type AssetQuality struct {
	bun.BaseModel `bun:"table:asset_qualities"`

	ID          int    `bun:",pk,autoincrement" json:"id"`
	Code        string `bun:",unique,notnull,type:varchar(50)" json:"code"`
	Name        string `bun:",notnull,type:varchar(100)" json:"name"`
	Description string `bun:",type:text" json:"description"`
}

func (a *AssetQuality) ToEntity() *entity.AssetQualityEntity {
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
