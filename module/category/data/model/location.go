package model

import (
	"asset_management_backend/module/category/domain/entity"

	"github.com/uptrace/bun"
)

const TableLocation = "locations"

type Location struct {
	bun.BaseModel `bun:"table:locations"`

	ID          int    `bun:",pk,autoincrement" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (l *Location) ToEntity() *entity.LocationEntity {
	if l == nil {
		return nil
	}
	return &entity.LocationEntity{
		ID:          l.ID,
		Name:        l.Name,
		Description: l.Description,
	}
}

func LocationModelsToEntities(locations []Location) []entity.LocationEntity {
	var locationEntities []entity.LocationEntity

	for _, location := range locations {
		locationEntities = append(locationEntities, *location.ToEntity())
	}

	return locationEntities
}
