package controller

import (
	"asset_management_backend/common/error_app"
	"asset_management_backend/module/borrowed_asset/domain/usecase"

	"github.com/gofiber/fiber/v2"
)

type BorrowedAssetController struct {
	borrowedAssetUsecase usecase.BorrowedAssetUsecase
}

func (b *BorrowedAssetController) ReturnAssetHandler(c *fiber.Ctx) error {
	borrowedAssetID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(error_app.BadRequestErrorResponse(err.Error()))
	}
	response, err := b.borrowedAssetUsecase.ReturnAsset(c.Context(), borrowedAssetID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse(err.Error()))
	}
	return c.JSON(response)
}

func NewBorrowedAssetController() *BorrowedAssetController {
	return &BorrowedAssetController{
		borrowedAssetUsecase: usecase.NewBorrowedAssetUsecase(),
	}
}
