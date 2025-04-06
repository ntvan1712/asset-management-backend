package model

import (
	"asset_management_backend/module/category/domain/entity"

	"github.com/uptrace/bun"
)

type Currency struct {
	bun.BaseModel `bun:"table:currencies"`

	ID     int    `bun:",pk,autoincrement" json:"id"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}


func (c *Currency) ToEntity() *entity.CurrencyEntity {
	if c == nil {
		return nil
	}
	return &entity.CurrencyEntity{
		ID:     c.ID,
		Code:   c.Code,
		Name:   c.Name,
		Symbol: c.Symbol,
	}
}

func CurrencyModelsToEntities(currencies []Currency) []entity.CurrencyEntity {
	var currencyEntities []entity.CurrencyEntity

	for _, currency := range currencies {
		currencyEntities = append(currencyEntities, *currency.ToEntity())
	}

	return currencyEntities
}
