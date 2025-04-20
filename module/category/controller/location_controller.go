package controller

import (
	"asset_management_backend/common/error_app"
	app_utils "asset_management_backend/common/utils"
	"asset_management_backend/common/validator_app"
	"asset_management_backend/module/category/domain/entity"
	"asset_management_backend/module/category/domain/usecase"

	"github.com/gofiber/fiber/v2"
)

type LocationController struct {
	categoryUsecase usecase.CategoryUsecase
}

func NewLocationController() *LocationController {
	return &LocationController{
		categoryUsecase: usecase.NewCategoryUsecase(),
	}
}

func (aq *LocationController) GetAllHandler(c *fiber.Ctx) error {
	response, err := aq.categoryUsecase.GetAllLocations(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse(err.Error()))
	}
	return c.JSON(response)
}

func (aq *LocationController) DeleteByIDHandler(c *fiber.Ctx) error {
	LocationID, err := c.ParamsInt("location_id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(error_app.BadRequestErrorResponse("Id phải là số nguyên"))
	}

	err = aq.categoryUsecase.DeleteLocationByID(c.Context(), LocationID)
	if err != nil {
		if err == error_app.ErrDocumentNotFound {
			return c.Status(fiber.StatusNotFound).JSON(error_app.NotFoundErrorResponse("Không tìm thấy vị trí này"))
		}
		if err == error_app.ErrPermissionDenied {
			return c.Status(fiber.StatusForbidden).JSON(error_app.NotFoundErrorResponse("Không có quyền"))

		}
		return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse(err.Error()))
	}
	return c.SendString("Delete success")
}

func (aq *LocationController) UpdateHandler(c *fiber.Ctx) error {
	locationID, err := c.ParamsInt("location_id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(error_app.BadRequestErrorResponse("Id phải là số nguyên"))
	}

	type updateLocationRequest struct {
		Name        *string `json:"name,omitempty"`
		Description *string `json:"description,omitempty"`
	}
	var request *updateLocationRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(error_app.BadRequestErrorResponse(err.Error()))
	}

	if request.Name == nil && request.Description == nil {
		return c.Status(fiber.StatusBadRequest).JSON(error_app.BadRequestErrorResponse("Vui lòng cung cấp thông tin cập nhật"))
	}

	err = aq.categoryUsecase.UpdateLocation(c.Context(), locationID, app_utils.StructToUpdateMap(request))

	if err != nil {
		if err == error_app.ErrDocumentNotFound {
			return c.Status(fiber.StatusNotFound).JSON(error_app.NotFoundErrorResponse("Không tìm thấy asset vị trí này"))
		}
		if err == error_app.ErrPermissionDenied {
			return c.Status(fiber.StatusForbidden).JSON(error_app.NotFoundErrorResponse("Không có quyền"))

		}
		return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse(err.Error()))
	}
	return c.SendString("Update success")
}

func (aq *LocationController) CreateHandler(c *fiber.Ctx) error {

	type createLocationRequest struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description" validate:"required"`
	}
	var request *createLocationRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(error_app.BadRequestErrorResponse(err.Error()))
	}
	if err := validator_app.ValidateStruct(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	location, err := aq.categoryUsecase.CreateLocation(c.Context(), &entity.LocationEntity{
		Name:        request.Name,
		Description: request.Description,
	})

	if err != nil {
		if err == error_app.ErrDocumentNotFound {
			return c.Status(fiber.StatusNotFound).JSON(error_app.NotFoundErrorResponse("Không tìm thấy asset vị trí này"))
		}
		if err == error_app.ErrPermissionDenied {
			return c.Status(fiber.StatusForbidden).JSON(error_app.NotFoundErrorResponse("Không có quyền"))

		}
		return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse(err.Error()))
	}
	return c.JSON(location)
}
