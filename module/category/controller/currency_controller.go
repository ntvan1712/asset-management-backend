package controller

import (
	"asset_management_backend/common/error_app"
	"asset_management_backend/module/category/domain/usecase"

	"github.com/gofiber/fiber/v2"
)

type CurrencyController struct {
	categoryUsecase usecase.CategoryUsecase
}

func NewCurrencyController() *CurrencyController {
	return &CurrencyController{
		categoryUsecase: usecase.NewCategoryUsecase(),
	}
}

func (aq *CurrencyController) GetAllHandler(c *fiber.Ctx) error {
	response, err := aq.categoryUsecase.GetAllCurrencies(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse(err.Error()))
	}
	return c.JSON(response)
}
