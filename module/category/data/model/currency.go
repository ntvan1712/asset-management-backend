package model

import (
	"asset_management_backend/module/category/domain/entity"

	"github.com/uptrace/bun"
)

type Currency struct {
	bun.BaseModel `bun:"table:currencies"`

	ID     int    `bun:",pk,autoincrement" json:"id"`
	Code   string `bun:",unique,notnull,type:varchar(50)" json:"code"`
	Name   string `bun:",notnull,type:varchar(100)" json:"name"`
	Symbol string `bun:",notnull,type:varchar(10)" json:"symbol"`
}


func (c *Currency) ToEntity() *entity.CurrencyEntity {
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
