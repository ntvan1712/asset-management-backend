package model

import (
	"asset_management_backend/module/category/domain/entity"

	"github.com/uptrace/bun"
)

type AssetType struct {
	bun.BaseModel `bun:"table:asset_types"`

	ID          int    `bun:",pk,autoincrement" json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}


func (a *AssetType) ToEntity() *entity.AssetTypeEntity {
	if a == nil {
		return nil
	}
	return &entity.AssetTypeEntity{
		ID:          a.ID,
		Code:        a.Code,
		Name:        a.Name,
		Description: a.Description,
	}
}

func AssetTypeModelsToEntities(assetTypes []AssetType) []entity.AssetTypeEntity {
	var assetTypeEntities []entity.AssetTypeEntity

	for _, assetType := range assetTypes {
		assetTypeEntities = append(assetTypeEntities, *assetType.ToEntity())
	}

	return assetTypeEntities
}
