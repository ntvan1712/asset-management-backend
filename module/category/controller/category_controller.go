package controller

import (
	"asset_management_backend/common/error_app"
	"asset_management_backend/module/category/domain/usecase"

	"github.com/gofiber/fiber/v2"
)

type CategoryController struct {
	categoryUsecase usecase.CategoryUsecase
}


func (aq *CategoryController) GetAllHandler(c *fiber.Ctx) error {
	response, err := aq.categoryUsecase.GetAllCategories(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse(err.Error()))
	}
	return c.JSON(response)
}

func NewCategoryController() *CategoryController {
	return &CategoryController{
		categoryUsecase: usecase.NewCategoryUsecase(),
	}
}