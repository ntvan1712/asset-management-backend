package controller

import (
	"asset_management_backend/common/error_app"
	"asset_management_backend/common/validator_app"
	"asset_management_backend/module/asset/domain/entity"
	"asset_management_backend/module/asset/domain/usecase"

	"github.com/gofiber/fiber/v2"
)

type AssetController struct {
	assetUsecase usecase.AssetUsecase
}

func (a *AssetController) GetAssetFilesPresignedUrlsHandler(c *fiber.Ctx) error {
	var fileNames []string
	if err := c.BodyParser(&fileNames); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(error_app.BadRequestErrorResponse(err.Error()))
	}
	response, err := a.assetUsecase.GetAssetFilesPresignedUrls(c.Context(), fileNames)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse(err.Error()))
	}
	return c.JSON(response)
}

func (a *AssetController) CreateAssetHandler(c *fiber.Ctx) error {
	var request *entity.CreateAssetRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(error_app.BadRequestErrorResponse(err.Error()))
	}
	if err := validator_app.ValidateStruct(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	response, err := a.assetUsecase.CreateAsset(c.Context(), request)
	if err != nil {
		if err == error_app.ErrDuplicateKey {
			return c.Status(fiber.StatusConflict).JSON(error_app.ConflictErrorResponse("Serial number đã tồn tại"))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse(err.Error()))
	}
	return c.JSON(response)
}

func (a *AssetController) GetAssetByIDHandler(c *fiber.Ctx) error {
	assetID, err := c.ParamsInt("asset_id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(error_app.BadRequestErrorResponse("Asset Id phải là số nguyên"))
	}
	response, err := a.assetUsecase.GetAssetByID(c.Context(), assetID)
	if err != nil {
		if err == error_app.ErrDocumentNotFound {
			return c.Status(fiber.StatusNotFound).JSON(error_app.NotFoundErrorResponse("Không tìm thấy tài sản này"))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse(err.Error()))
	}
	return c.JSON(response)
}
func NewAssetController() *AssetController {
	return &AssetController{
		assetUsecase: usecase.NewAssetUsecase(),
	}
}
