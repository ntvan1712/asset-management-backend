package controller

import (
	"asset_management_backend/common/error_app"
	app_utils "asset_management_backend/common/utils"
	"asset_management_backend/common/validator_app"
	"asset_management_backend/module/category/domain/entity"
	"asset_management_backend/module/category/domain/usecase"

	"github.com/gofiber/fiber/v2"
)

type AssetQualityController struct {
	categoryUsecase usecase.CategoryUsecase
}

func NewAssetQualityController() *AssetQualityController {
	return &AssetQualityController{
		categoryUsecase: usecase.NewCategoryUsecase(),
	}
}

func (aq *AssetQualityController) GetAllHandler(c *fiber.Ctx) error {
	response, err := aq.categoryUsecase.GetAllAssetQualities(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse(err.Error()))
	}
	return c.JSON(response)
}

func (aq *AssetQualityController) DeleteByIDHandler(c *fiber.Ctx) error {
	assetQualityID, err := c.ParamsInt("asset_quality_id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(error_app.BadRequestErrorResponse("Id phải là số nguyên"))
	}

	err = aq.categoryUsecase.DeleteAssetQualityByID(c.Context(), assetQualityID)
	if err != nil {
		if err == error_app.ErrDocumentNotFound {
			return c.Status(fiber.StatusNotFound).JSON(error_app.NotFoundErrorResponse("Không tìm thấy chất loại chất tượng này"))
		}
		if err == error_app.ErrPermissionDenied {
			return c.Status(fiber.StatusForbidden).JSON(error_app.NotFoundErrorResponse("Không có quyền"))

		}
		return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse(err.Error()))
	}
	return c.SendString("Delete success")
}

func (aq *AssetQualityController) UpdateHandler(c *fiber.Ctx) error {
	assetQualityID, err := c.ParamsInt("asset_quality_id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(error_app.BadRequestErrorResponse("Id phải là số nguyên"))
	}

	type updateAssetQualityRequest struct {
		Code        *string `json:"code,omitempty"`
		Name        *string `json:"name,omitempty"`
		Description *string `json:"description,omitempty"`
	}
	var request *updateAssetQualityRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(error_app.BadRequestErrorResponse(err.Error()))
	}
	if request.Name == nil && request.Description == nil && request.Code == nil {
		return c.Status(fiber.StatusBadRequest).JSON(error_app.BadRequestErrorResponse("Vui lòng cung cấp thông tin cập nhật"))
	}

	err = aq.categoryUsecase.UpdateAssetQuality(c.Context(), assetQualityID, app_utils.StructToUpdateMap(request))
	if err != nil {
		if err == error_app.ErrDocumentNotFound {
			return c.Status(fiber.StatusNotFound).JSON(error_app.NotFoundErrorResponse("Không tìm thấy asset quality này"))
		}
		if err == error_app.ErrPermissionDenied {
			return c.Status(fiber.StatusForbidden).JSON(error_app.NotFoundErrorResponse("Không có quyền"))

		}
		return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse(err.Error()))
	}
	return c.SendString("Update success")
}

func (aq *AssetQualityController) CreateHandler(c *fiber.Ctx) error {
	type createAssetQualityRequest struct {
		Code        string `json:"code" validate:"required"`
		Name        string `json:"name" validate:"required"`
		Description string `json:"description" validate:"required"`
	}
	var request *createAssetQualityRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(error_app.BadRequestErrorResponse(err.Error()))
	}
	if err := validator_app.ValidateStruct(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	assetQualityEntity, err := aq.categoryUsecase.CreateAssetQuality(c.Context(), &entity.AssetQualityEntity{
		Code:        request.Code,
		Name:        request.Name,
		Description: request.Description,
	})

	if err != nil {
		if err == error_app.ErrPermissionDenied {
			return c.Status(fiber.StatusForbidden).JSON(error_app.NotFoundErrorResponse("Không có quyền"))

		}
		return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse(err.Error()))
	}
	return c.JSON(assetQualityEntity)
}
