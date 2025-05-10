package sharedmodel

import (
	"asset_management_backend/common/error_app"
	"asset_management_backend/common/validator_app"

	"github.com/gofiber/fiber/v2"
)

type PaginateQuery struct {
	Limit int `query:"limit" validate:"required,gte=1,lte=50"`
	Page  int `query:"page" validate:"required,gte=1"`
}

func GetPaginateQuery(c *fiber.Ctx) (*PaginateQuery, error) {
	var filterQuery PaginateQuery
	if err := c.QueryParser(&filterQuery); err != nil {
		return nil, c.Status(fiber.StatusBadRequest).JSON(error_app.BadRequestErrorResponse(err.Error()))
	}
	if err := validator_app.ValidateStruct(filterQuery); err != nil {
		return nil, c.Status(fiber.StatusBadRequest).JSON(err)
	}
	return &filterQuery, nil
}
